package system

import (
	"find"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func System(){
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("system1.glade") //读取glade文件

	// 获取窗口控件指针，注意"window1"要和glade里的标志名称匹配
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(600, 400)        //设置窗口大小
	window.SetTitle("管理员")            //设置标题 //设置icon
	window.SetResizable(true)             //设置不可伸缩
	window.SetPosition(gtk.WIN_POS_CENTER) //设置居中显示
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	button2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	button5 := gtk.ButtonFromObject(builder.GetObject("button5"))
	button1.Connect("clicked", func() {
		find.Find()
	})
	button2.Connect("clicked", func() {
		find.Find1()
	})
	button5.Connect("clicked",func(){
		gtk.MainQuit()
	})
	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
