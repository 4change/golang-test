package main

//Subject　目标对象抽象
type Subject interface {
	Add(o Observer)
	Remove(o Observer)
	Notify()
}

// =====================================================================================================================

type ConcreteSubject struct {
	obs []Observer // 负责保存所有观察者对象
}

func (t *ConcreteSubject) Add(o Observer) { //注册观察者
	t.obs = append(t.obs, o)
}

func (t *ConcreteSubject) Remove(o Observer) { //删除观察者
	for k, v := range t.obs {
		if v == o {
			t.obs = append(t.obs[0:k], t.obs[k+1:]...)
			return
		}
	}
	return
}

func (t *ConcreteSubject) Notify() { //通知观察者
	for _, v := range t.obs {
		v.Response("你有一条消息")
	}
}