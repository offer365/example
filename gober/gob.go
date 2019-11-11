package gober

import (
	"encoding/gob"
	"log"
	"os"
)



type gobSend struct {
	A int
	B float64
}

type gobGet struct {
	A int
	D float64
}

func init() {
	gob.Register(&gobSend{})
	gob.Register(&gobGet{})
}

func main2() {
	gob.Encoder{}
	fmt.Println(gobEncode("1.gob"))
	fmt.Println(gobDecode("1.gob"))
}

func gobEncode(fileName string) error {
	fi, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fi.Close()

	encoder := gob.NewEncoder(fi)
	return encoder.Encode(gobSend{1,12})
}

func gobDecode(fileName string) (error) {
	fi, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fi.Close()

	decoder := gob.NewDecoder(fi)
	g := gobGet{}
	err = decoder.Decode(&g)
	fmt.Println(g)
	return err
}

//定义一个结构体
type Student struct {
	Name string
	Age uint8
	Address string
}

func main2(){
	//序列化
	s1:=Student{"张三",18,"江苏省"}
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)//创建编码器
	err1 := encoder.Encode(&s1)//编码
	if err1!=nil{
		log.Panic(err1)
	}
	fmt.Printf("序列化后：%x\n",buffer.Bytes())

	//反序列化
	byteEn:=buffer.Bytes()
	decoder := gob.NewDecoder(bytes.NewReader(byteEn)) //创建解密器
	var s2 Student
	err2 := decoder.Decode(&s2)//解密
	if err2!=nil{
		log.Panic(err2)
	}
	fmt.Println("反序列化后：",s2)
}

