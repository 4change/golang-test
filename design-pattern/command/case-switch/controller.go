package case_switch

import "fmt"

//-----控制类------
type Controller struct { //组合不同的命令，构成控制端
	switchC  Command
	volumeC  Command
	channelC Command
}

func (t Controller) buttonOkHold() {
	fmt.Print("长按Ok键...")
	t.switchC.Do()
}
func (t Controller) buttonOkClick() {
	fmt.Print("单击Ok键...")
	t.switchC.UnDo()
}

func (t Controller) buttonUpClick() {
	fmt.Print("单击↑按键...")
	t.volumeC.Do()
}
func (t Controller) buttonDownClick() {
	fmt.Print("单击↓按键...")
	t.volumeC.UnDo()
}

func (t Controller) buttonRightClick() {
	fmt.Print("单击→按键...")
	t.channelC.Do()
}
func (t Controller) buttonLeftClick() {
	fmt.Print("单击←按键...")
	t.channelC.UnDo()
}
