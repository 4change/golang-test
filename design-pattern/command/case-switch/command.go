package case_switch

// 抽象命令模块：这里的开和关，音量的大和小，频道调节上和下，屏幕的亮和暗等等，其实就是两个反向的操作指令。
type Command interface {
	Do()
	UnDo()
}

// 开关类,负责设备的开和关,d负责持有设备对象，对设备进行控制
//SwitchCommand　开关控制
type SwitchCommand struct {
	d Device
}

func (t SwitchCommand) Do() {
	t.d.Open()
}
func (t SwitchCommand) UnDo() {
	t.d.Close()
}

// 音量类，负责控制设备音量大小
//VolumeCommand 音量控制
type VolumeCommand struct {
	d Device
}

func (t VolumeCommand) Do() {
	t.d.VolumeUp()
}

func (t VolumeCommand) UnDo() {
	t.d.VolumeDown()
}

// 频道类，负责控制设备频道调节
//ChannelCommand 频道控制
type ChannelCommand struct {
	d Device
}

func (t ChannelCommand) Do() {
	t.d.ChannelUp()
}

func (t ChannelCommand) UnDo() {
	t.d.ChannelDown()
}
