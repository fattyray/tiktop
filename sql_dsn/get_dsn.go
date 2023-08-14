package sql_dsn

// 用于获取个人电脑数据库，原本各个文件获取的话换了电脑修改太麻烦，将来可以修改成自动获取 mysql
func GetDsn() string {
	return "root:123456@(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
}
