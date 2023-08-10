package main

func main() {

}

// TODO
type Config struct {
	DB string
}
type MyDB interface {
	GetConfig() *Config
}

type User struct {
}

func (u *User) GetConfig() *Config {
	return &Config{DB: "121313"}
}

// func getConfig() {
// 	once := sync.Once{}
// 	once.Do(func() {
// 		var config Config
// 		config.DB = "DB"
// 	})
// }
