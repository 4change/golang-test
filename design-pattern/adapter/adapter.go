package adapter

/*
 Adapter 适配器模式：
        将一个类的接口转换成客户端希望的另一个接口。
		适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
 个人想法：代理和适配器：代理和代理的对象接口一致，客户端不知道代理对象，
        而适配器是客户端想要适配器的接口，适配器对象的接口和客户端想要的不一样，
		适配器将适配器对象的接口封装一下，改成客户端想要的接口
*/

// 适配器，实现了Player接口(source接口)，同时注入了ForeignCenter类(target类)
type Translator struct {
	f ForeignCenter
}

func (t *Translator) attack() {
	if t == nil {
		return
	}
	t.f.attack("进攻")
}

func (t *Translator) defense() {
	if t == nil {
		return
	}
	t.f.defense()
}

func NewTranslator(name string) Player {
	return &Translator{ForeignCenter{name}}
}
