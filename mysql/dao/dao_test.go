package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

var Db *sqlx.DB

type User struct {
	Uid  int64  `db:"uid"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 建库sql: CREATE DATABASE IF NOT EXISTS golang DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
func TestNewDB(t *testing.T) {
	driver := NewDB("mysql")
	var err error
	Db, err = driver.Init(
		WithHost("10.0.0.110"),
		WithDatabase("golang"),
		WithUsername("root"),
		WithPwd("Zrz0123456789"))
	fmt.Println(err)

	//ExampleCreateTable()
	//ExampleExecInsert()
	//ExampleExecUpdate()
	//ExamplePrepareSelect()
	//ExamplePrepareInsert()
	//ExamplePreparexSelect()
	//ExampleQuery()
	//ExampleQueryx()
	//ExampleQueryRow()
	//ExampleSelectGet()
	ExampleTx()
}

func ExampleCreateTable() {
	user := `CREATE TABLE user (
    uid INT(10) NOT NULL AUTO_INCREMENT,
    name VARCHAR(64)  DEFAULT NULL,
    age int(3)  DEFAULT NULL,
    PRIMARY KEY (uid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8`
	Db.MustExec(user)

}

//Exec和MustExec从连接池中获取一个连接然后只想对应的query操作。
// 对于不支持ad-hoc query execution的驱动，在操作执行的背后会创建一个prepared statement。
// 在结果返回前这个connection会返回到连接池中。
//
//需要注意的是不同的数据库类型使用的占位符不同，mysql采用？作为占位符号。
//MySQL 使用？
//PostgreSQL 使用1,1,2等等
//SQLite 使用？或$1
//Oracle 使用:name

// 增
func ExampleExecInsert() {
	sql := `INSERT INTO user (name, age) VALUES (?, ?)`
	result, err := Db.Exec(sql, "jack", 18)
	if err != nil {
		fmt.Println("insert failed,error： ", err)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("insert id is :", id)
}

// 改
func ExampleExecUpdate() {
	sql := `update user set age = ? where uid = ?`
	_, err1 := Db.Exec(sql, 19, 1)
	if err1 != nil {
		fmt.Println("update failed error:", err1)
	} else {
		fmt.Println("update success!")
	}
}

// 删
func ExampleExecDelete() {
	sql := `delete from user where uid = ? `
	_, err2 := Db.Exec(sql, 1)
	if err2 != nil {
		fmt.Println("delete error:", err2)
	} else {
		fmt.Println("delete success")
	}
}

//sql预处理（Prepared Statements）
// 对于大部分的数据库来说，当一个query执行的时候，在sql语句数据库内部声明已经声明过了，其声明是在数据库中，我们可以提前进行声明，以便在其他地方重用。
func ExamplePrepareSelect() {
	sql := `SELECT * FROM user WHERE uid>?`
	stmt, _ := Db.Prepare(sql)

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	rows, err := stmt.Query(0)
	defer rows.Close()
	if err != nil {
		fmt.Println("perpare error,", err)
	}
	var user User
	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Name, &user.Age)
		if err != nil {
			fmt.Println("scan error,", err)
		}
		fmt.Println(user)
	}
}

// sqlx还提供了Preparex()进行扩展，可直接用于结构体转换
func ExamplePreparexSelect() {
	sql := `SELECT * FROM user WHERE uid=?`
	stmt, _ := Db.Preparex(sql)

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	var u User
	stmt.Get(&u, 1)
	fmt.Println(u.Age, u.Uid, u.Name)

	var us []User
	sql = `SELECT * FROM user WHERE uid>?`
	stmt, _ = Db.Preparex(sql)
	stmt.Select(&us, 0)
	for _, u := range us {
		fmt.Println(u.Age, u.Uid, u.Name)
	}
}

// 预处理 插入数据
func ExamplePrepareInsert() {
	sql := "insert into user(name,age)values(?,?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		fmt.Println("perpare error,", err)
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Exec("mali", 55)
	result, err = stmt.Exec("tom", 35)
	result, err = stmt.Exec("lisi", 25)
	result, err = stmt.Exec("lili", 15)
	if err != nil {
		fmt.Println("insrt error,", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("id error,", err)
	}
	fmt.Println("id,", id)
}

//Query是database/sql中执行查询主要使用的方法，该方法返回row结果。
// Query返回一个sql.Rows对象和一个error对象。
//在使用的时候应该把Rows当成一个游标而不是一系列的结果。
// 尽管数据库驱动缓存的方法不一样，通过Next()迭代每次获取一列结果，对于查询结果非常巨大的情况下，可以有效的限制内存的使用，
// Scan()利用reflect把sql每一列结果映射到go语言的数据类型如string，[]byte等。
// 如果你没有遍历完全部的rows结果，一定要记得在把connection返回到连接池之前调用rows.Close()。
//
//Query返回的error有可能是在server准备查询的时候发生的，
// 也有可能是在执行查询语句的时候发生的。
// 例如可能从连接池中获取一个坏的连级（尽管数据库会尝试10次去发现或创建一个工作连接）。
// 一般来说，错误主要由错误的sql语句，错误的类似匹配，错误的域名或表名等。
//
//在大部分情况下，Rows.Scan()会把从驱动获取的数据进行拷贝，无论驱动如何使用缓存。
// 特殊类型sql.RawBytes可以用来从驱动返回的数据总获取一个zero-copy的slice byte。
// 当下一次调用Next的时候，这个值就不在有效了，因为它指向的内存已经被驱动重写了别的数据。
//
//Query使用的connection在所有的rows通过Next()遍历完后或者调用rows.Close()后释放。
func ExampleQuery() {
	var (
		name     string
		uid, age int
	)
	sql := "select * from user"
	rows, _ := Db.Query(sql)
	for rows.Next() {
		rows.Scan(&uid, &name, &age)
		fmt.Println(uid, name, age)
	}
}

// Queryx和Query行为很相似，不过返回一个sqlx.Rows对象，支持扩展的scan行为,同时可将对数据进行结构体转换。
func ExampleQueryx() {
	var u User
	sql := "select * from user"
	rows, err := Db.Queryx(sql)
	fmt.Println(err)
	for rows.Next() {
		rows.StructScan(&u) // 转换为结构体
		fmt.Println(u.Uid, u.Name, u.Age)
	}
}

// QueryRow和QueryRowx都是从数据库中获取一条数据，
// 但是QueryRowx提供scan扩展，可直接将结果转换为结构体。
func ExampleQueryRow() {
	var (
		name string
		age  int
	)
	sql := "select name,age from user where uid=?"
	row := Db.QueryRow(sql, 2)
	row.Scan(&name, &age)
	fmt.Println(name, age)

	var u User
	err := Db.QueryRowx(sql, 3).StructScan(&u)
	fmt.Println(err)
	fmt.Println(u.Name, u.Age)

}

// Get和Select是一个非常省时的扩展，可直接将结果赋值给结构体，其内部封装了StructScan进行转化。
// Get用于获取单个结果然后Scan，Select用来获取结果切片。
func ExampleSelectGet() {
	var (
		u  User
		us []User
	)
	sql := `select * from user where uid>?`
	Db.Select(&us, sql, 0)
	for _, u := range us {
		fmt.Println(u.Uid, u.Name, u.Age)
	}
	Db.Get(&u, "select * from user where uid=?", 1)
	fmt.Println(u.Uid, u.Name, u.Age)

}

//事务操作是通过三个方法实现：
//Begin()：开启事务
//Commit()：提交事务（执行sql)
//Rollback()：回滚
func ExampleTx() {
	sql := `insert into user(name,age) values(?,?)`
	tx, err := Db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println(" tx error,", err)
		return
	}
	_, err = tx.Exec(sql, "zhangsan", 34)
	if err != nil {
		tx.Rollback()
	}
	_, err = tx.Exec(sql, "wangwu", 32)
	if err != nil {
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("commit error")
	}
}
