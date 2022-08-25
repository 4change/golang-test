package case_switch

import (
	"fmt"
	"testing"
)

func TestCaseSwitch(t *testing.T) {

	tv := TV{}

	c := Controller{ //　c通过组合３种命令和TV{}关联，变成电视机控制器
		switchC:  SwitchCommand{d: tv},
		volumeC:  VolumeCommand{d: tv},
		channelC: ChannelCommand{d: tv},
	}
	c.buttonOkHold()
	c.buttonUpClick()
	c.buttonRightClick()
	c.buttonOkClick()

	fmt.Println("-----------")
	radio := Radio{}
	c = Controller{ //　c通过组合３种命令和Radio{}关联，变成电视机控制器
		switchC:  SwitchCommand{d: radio},
		channelC: VolumeCommand{d: radio},  //调整按键功能【左右】键负责音量
		volumeC:  ChannelCommand{d: radio}, //调整按键功能【上下】键负责频道
	}

	c.buttonOkHold()
	c.buttonUpClick()
	c.buttonRightClick()
	c.buttonOkClick()

	return
}
