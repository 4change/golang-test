package case_switch

import "fmt"

type Device interface {
	Open()        //打开
	Close()       //关闭
	VolumeUp()    //音量+
	VolumeDown()  //音量-
	ChannelUp()   //频道+
	ChannelDown() //频道-
}

type TV struct {
}

func (t TV) Open() {
	fmt.Println("打开TV")
}

func (t TV) Close() {
	fmt.Println("关闭TV")
}

func (t TV) VolumeUp() {
	fmt.Println("TV音量+")
}

func (t TV) VolumeDown() {
	fmt.Println("TV音量-")
}

func (t TV) ChannelUp() {
	fmt.Println("TV频道+")
}

func (t TV) ChannelDown() {
	fmt.Println("TV频道-")
}

type Radio struct {
}

func (t Radio) Open() {
	fmt.Println("打开Radio")
}

func (t Radio) Close() {
	fmt.Println("关闭Radio")
}

func (t Radio) VolumeUp() {
	fmt.Println("Radio音量+")
}

func (t Radio) VolumeDown() {
	fmt.Println("Radio音量-")
}

func (t Radio) ChannelUp() {
	fmt.Println("Radio频道+")
}

func (t Radio) ChannelDown() {
	fmt.Println("Radio频道-")
}
