package find

import (
	"connectsql"
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)
type User struct {
	Num   string
	Name_Ch    string
	Name_En string
	Borrow string
}
	func Find(){
		gtk.Init(&os.Args)
		builder := gtk.NewBuilder()       //新建builder
		builder.AddFromFile("find.glade") //读取glade文件
		window := gtk.WindowFromObject(builder.GetObject("window1"))
		window.SetSizeRequest(600, 400)        //设置窗口大小
		window.SetTitle("管理员")            //设置标题 //设置icon
		window.SetResizable(true)             //设置不可伸缩
		window.SetPosition(gtk.WIN_POS_CENTER) //设置居中显示
		button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
		entry1 := gtk.EntryFromObject(builder.GetObject("entry1"))
		entry2 := gtk.EntryFromObject(builder.GetObject("entry2"))
		entry3 := gtk.EntryFromObject(builder.GetObject("entry3"))
		entry4 := gtk.EntryFromObject(builder.GetObject("entry4"))
		entry5 := gtk.EntryFromObject(builder.GetObject("entry5"))
		button1.Connect("clicked", func() {
			var user1 User
			err := (connectsql.DB).QueryRow("SELECT * FROM books WHERE Name_Ch=?", entry1.GetText()).Scan(&user1.Num,&user1.Name_Ch,&user1.Name_En,&user1.Borrow)
			if err != nil {
				fmt.Println("查询出错了")
			}
			entry2.SetText(user1.Num)
			entry3.SetText(user1.Name_Ch)
			entry4.SetText(user1.Name_En)
			entry5.SetText(user1.Borrow)
		})
		window.Connect("destroy",gtk.MainQuit)
		window.ShowAll()
		gtk.Main()
	}
	func Find1(){
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("图书信息.glade") //读取glade文件
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(600, 400)        //设置窗口大小
	window.SetTitle("管理员")            //设置标题 //设置icon
	window.SetResizable(true)             //设置不可伸缩
	window.SetPosition(gtk.WIN_POS_CENTER) //设置居中显示

	
	window.Connect("destroy",gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}