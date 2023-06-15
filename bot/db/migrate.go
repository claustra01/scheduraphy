package db

func Migrate() {
	Connect()
	Psql.Exec("DROP TABLE IF EXISTS users")
	Psql.AutoMigrate(&User{})
}
