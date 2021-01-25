package connectsql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)
type DBConnectionInfo struct {
	User	string//数据库用户名
	Host	string//数据库地址
	password	string//数据库用户名密码
	port string//数据库接口
	DbName	string//数据库名
}//定义了一个数据库信息结构
var DB *sql.DB//DB数据库连接池
func Conectsql(){
	var dbc DBConnectionInfo//定义了一个数据库数据结构
	dbc.DbName="library"
	dbc.User="root"
	dbc.password="123456"
	dbc.Host= "127.0.0.1"
	dbc.port="3306"//初始化连接数据库信息
	InitDB(&dbc)//调用连接数据库函数
}//初始化数据库信息并调用函数
func InitDB(dbc *DBConnectionInfo)  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{dbc.User, ":", dbc.password, "@tcp(",dbc.Host, ":", dbc.port, ")/", dbc.DbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接

	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")

}