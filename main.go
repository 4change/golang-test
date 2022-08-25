package main

import "fmt"

/*
有限状态机是一种算法思想，简单而言，有限状态机由一组状态、一个初始状态、输入和根据输入及现有状态转换为下一个状态的转换函数组成。

请使用Python/Go语言实现一个灵活的状态机，使用者可以创建一个状态机，添加一组状态和状态转换表，并设置好初始状态，然后根据触发不同的事件进行状态的转移。

可以参考下列示例编写相应的测试用例：
示例1：https://img-blog.csdn.net/20140516223146296
*/

//// 接口
//type IFSMState interface {
//	Enter()
//	Exit()
//	CheckTransition(hour int) bool
//	Hour() int
//}
//
//// State父struct
//type FSMState struct{}
//
//// 进入状态
//func (this *FSMState) Enter() {
//	//
//}
//
//// 退出状态
//func (this *FSMState) Exit() {
//	//
//}
//
//// 状态转移检测
//func (this *FSMState) CheckTransition(hour int) {
//	//
//}
//
//////////// 打坐
//type ZazenState struct {
//	// TODO 此处应当修改为字符串状态
//	hour int
//	FSMState
//}
//
//func NewZazenState() *ZazenState {
//	return &ZazenState{hour: 8}
//}
//func (this *ZazenState) Enter() {
//	fmt.Println("ZazenState: 开门")
//}
//func (this *ZazenState) Exit() {
//	fmt.Println("ZazenState: 关门")
//}
//func (this *ZazenState) Hour() int {
//	return this.hour
//}
//// 状态转移检测
//func (this *ZazenState) CheckTransition(hour int) bool {
//	if hour == this.hour {
//		return true
//	}
//	return false
//}
//
//////////// 工作
//type WorkerState struct {
//	hour int    // 状态转移的条件，即输入参数
//	FSMState
//}
//func NewWorkerState() *WorkerState {
//	return &WorkerState{hour: 12}
//}
//func (this *WorkerState) Enter() {
//	fmt.Println("WorkerState: 锁门")
//}
//func (this *WorkerState) Exit() {
//	fmt.Println("WorkerState: 解锁")
//}
//func (this *WorkerState) Hour() int {
//	return this.hour
//}
//
//// 状态转移检测
//func (this *WorkerState) CheckTransition(hour int) bool {
//	if hour == this.hour {
//		return true
//	}
//	return false
//}
//
//////////// 状态机
//type FSM struct {
//	states map[string]IFSMState   // 持有状态集合
//	current_state IFSMState       // 当前状态
//	default_state IFSMState       // 默认状态
//	input_data int                // 外部输入数据
//	inited bool                   // 是否初始化
//}
//// 初始化FSM
//func (this *FSM) Init() {
//	this.Reset()
//}
//// 添加状态到FSM
//func (this *FSM) AddState(key string, state IFSMState) {
//	// 状态机的大小是否应当调整为 4
//	// 创建大小为2的状态机容器
//	if this.states == nil {
//		this.states = make(map[string]IFSMState, 2)
//	}
//	this.states[key] = state
//}
//// 设置默认的State
//func (this *FSM) SetDefaultState(state IFSMState) {
//	this.default_state = state
//}
//
//// 转移状态
//func (this *FSM) TransitionState() {
//	nextState := this.default_state
//	input_data := this.input_data
//
//	// 初始化
//	if this.inited {
//		for _, v := range this.states {
//			if input_data == v.Hour() {
//				nextState = v
//				break
//			}
//		}
//	}
//	if ok := nextState.CheckTransition(this.input_data); ok {
//		if this.current_state != nil {
//			// 退出前一个状态
//			this.current_state.Exit()
//		}
//		this.current_state = nextState
//		this.inited = true
//		// 进入下一个状态
//		nextState.Enter()
//	}
//}
//
//// 设置输入数据
//func (this *FSM) SetInputData(inputData int) {
//	this.input_data = inputData
//	this.TransitionState()
//}
//
//// 重置
//func (this *FSM) Reset() {
//	this.inited = false
//}
//
//func main() {
//	zazenState := NewZazenState()
//	workerState := NewWorkerState()
//	fsm := new(FSM)
//	fsm.AddState("ZazenState", zazenState)
//	fsm.AddState("WorkerState", workerState)
//	fsm.SetDefaultState(zazenState)
//	fsm.Init()
//	fsm.SetInputData(8)
//	fsm.SetInputData(12)
//	fsm.SetInputData(12)
//	fsm.SetInputData(8)
//	fsm.SetInputData(12)
//}














// 接口
type IFSMState interface {
	Enter()
	Exit()
	CheckTransition(state string) bool
	State() string
}

// State父struct
type FSMState struct{}

// 进入状态
func (this *FSMState) Enter() {
	//
}

// 退出状态
func (this *FSMState) Exit() {
	//
}

// 状态转移检测
func (this *FSMState) CheckTransition(state string) {
	//
}

type OpenState struct {
	state string
	FSMState
}

func NewOpenState() *OpenState {
	return &OpenState{state: "closed"}
}
func (this *OpenState) Enter() {
	fmt.Println("OpenState: 开门")
}
func (this *OpenState) Exit() {
	fmt.Println("OpenState: 关门")
}
func (this *OpenState) State() string {
	return this.state
}
// 状态转移检测
func (this *OpenState) CheckTransition(state string) bool {
	fmt.Println("状态判定-------------------")
	if state == this.state {
		return true
	}
	return false
}

type LockState struct {
	state string    // 状态转移的条件，即输入参数
	FSMState
}
func NewLockState() *LockState {
	return &LockState{state: "locked"}
}
func (this *LockState) Enter() {
	fmt.Println("LockState: 锁门")
}
func (this *LockState) Exit() {
	fmt.Println("LockState: 解锁")
}
func (this *LockState) State() string {
	return this.state
}

// 状态转移检测
func (this *LockState) CheckTransition(state string) bool {
	if state == this.state {
		return true
	}
	return false
}

////////// 状态机
type FSM struct {
	states map[string]IFSMState   // 持有状态集合
	current_state IFSMState       // 当前状态
	default_state IFSMState       // 默认状态
	input_data string             // 外部输入数据
	inited bool                   // 是否初始化
}
// 初始化FSM
func (this *FSM) Init() {
	this.Reset()
}
// 添加状态到FSM
func (this *FSM) AddState(key string, state IFSMState) {
	// 状态机的大小是否应当调整为 4
	// 创建大小为2的状态机容器
	if this.states == nil {
		this.states = make(map[string]IFSMState, 2)
	}
	this.states[key] = state
}
// 设置默认的State
func (this *FSM) SetDefaultState(state IFSMState) {
	fmt.Println("设置状态--------------------")
	this.default_state = state
}

// 转移状态
func (this *FSM) TransitionState() {
	nextState := this.default_state
	input_data := this.input_data

	// 初始化
	if this.inited {
		fmt.Println("进入初始化------------")
		for _, v := range this.states {
			if input_data == v.State() {
				nextState = v
				break
			}
		}
	}

	if ok := nextState.CheckTransition(this.input_data); ok {
		if this.current_state != nil {
			// 退出前一个状态
			this.current_state.Exit()
		}
		this.current_state = nextState
		this.inited = true
		// 进入下一个状态
		nextState.Enter()
	}
}

// 设置输入数据
func (this *FSM) SetInputData(inputData string) {
	fmt.Println("设置输入数据----------------")
	this.input_data = inputData
	this.TransitionState()
}

// 重置
func (this *FSM) Reset() {
	this.inited = false
}

func main() {
	openState := NewOpenState()
	workerState := NewLockState()
	fsm := new(FSM)
	fsm.AddState("openState", openState)
	fsm.AddState("lockState", workerState)
	fsm.SetDefaultState(openState)
	fsm.Init()
	fsm.SetInputData("closed")
	fsm.SetInputData("locked")
	fsm.SetInputData("closed")
	fsm.SetInputData("locked")
}
