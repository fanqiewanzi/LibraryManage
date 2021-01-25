package register

import (
	"connectsql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

type User struct {
	Userid   string
	Name     string
	PassWord string
	Phonenum string
}

func Registe() {
	gtk.Init(&os.Args)
	var Users User
	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("注册界面.glade") //读取glade文件

	// 获取窗口控件指针，注意"window1"要和glade里的标志名称匹配
	window := gtk.WindowFromObject(builder.GetObject("window1"))

	window.SetSizeRequest(500, 300)        //设置窗口大小
	window.SetTitle("注册界面")                //设置标题 //设置icon
	window.SetResizable(false)             //设置不可伸缩
	window.SetPosition(gtk.WIN_POS_CENTER) //设置居中显示
	//按窗口关闭按钮，自动触发"destroy"信号
	tx, err := connectsql.DB.Begin()
	if nil != err {
		fmt.Println("开启事务失败:", err)

	}
	button1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	button2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	Entry1 := gtk.EntryFromObject(builder.GetObject("entry1"))
	Entry2 := gtk.EntryFromObject(builder.GetObject("entry2"))
	Entry3 := gtk.EntryFromObject(builder.GetObject("entry3"))
	Entry4 := gtk.EntryFromObject(builder.GetObject("entry4"))
	//按窗口关闭按钮，自动触发"destroy"信号
	button1.Connect("clicked", func() (id int64, err error) {
		Users.Userid = Entry1.GetText()
		Users.Name = Entry2.GetText()
		Users.PassWord = Entry3.GetText()
		Users.Phonenum = Entry4.GetText()
		stmt, err := tx.Prepare("insert into users(`idUsers`, `UserName`,`PassWord`,`Phonenumer`) values(?,?,?,?)")
		if nil != err {
			fmt.Println("Prepare Failed!")

		}
		res, err := stmt.Exec(Users.Userid, Users.Name, Users.PassWord, Users.Phonenum)
		if nil != err {
			fmt.Println("Exec Failed!")
			tx.Commit()
			id, err = res.LastInsertId()
			return id, nil
		}
		tx.Commit()
		id, err = res.LastInsertId()
		gtk.MainQuit()
		return id, nil
	})
	button2.Connect("clicked", func() {
		gtk.MainQuit()
	})
	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll()
	gtk.Main()
}
