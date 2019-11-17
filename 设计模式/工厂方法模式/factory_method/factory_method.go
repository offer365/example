package factory_method

// 工厂方法模式
// 工厂方法模式使用子类的方式延迟生成对象到子类中实现。
//
// Go中不存在继承 所以使用匿名组合来实现

// 被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type OptFactory interface {
	Create() Operator
}

type OptBase struct {
	a, b int
}
