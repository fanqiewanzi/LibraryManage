package login

import (
	"connectsql"
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"register"
	"system"
)
type User struct {
	Userid   string
	Name     string
	PassWord string
	Phonenum string
}
func Login(){
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("登陆界面.glade") //读取glade文件

	// 获取窗口控件指针，注意"window1"要和glade里的标志名称匹配
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(500, 300)        //设置窗口大小
	window.SetTitle("登陆界面")            //设置标题 //设置icon
	window.SetResizable(true)             //设置不可伸缩
	window.SetPosition(gtk.WIN_POS_CENTER) //设置居中显示
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	button2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	Entry1 := gtk.EntryFromObject(builder.GetObject("entry1"))
	Entry2 := gtk.EntryFromObject(builder.GetObject("entry2"))
	//按窗口关闭按钮，自动触发"destroy"信号
	button2.Connect("clicked", func() {//按注册按钮进入注册部分
		register.Registe()
	})
	button1.Connect("clicked", func(){
		var users User
		var user1 User
		users.Name = Entry1.GetText()
		users.PassWord = Entry2.GetText()
		err := (connectsql.DB).QueryRow("SELECT * FROM users WHERE UserName=?", users.Name).Scan(&user1.Userid,&user1.Name,&user1.PassWord,&user1.Phonenum)
		if err != nil {
			fmt.Println("查询出错了")
		}
		if users.PassWord == user1.PassWord {
			window.SetVisible(false)
			system.System()
		}
	})
	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()

}