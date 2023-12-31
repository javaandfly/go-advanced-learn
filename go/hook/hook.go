package main

func main() {

	hook := MyHook{}
	DoSome(&hook)

}

func DoSome(hook Hook) {
	hook.Before()

	hook.After()
}

type Hook interface {
	// 钩子方法
	Before()
	After()
}

type MyHook struct {
	Hook Hook
}

// func (m *MyHook) Before() {
// 	fmt.Println("before this is a hook")
// }

// func (m *MyHook) After() {
// 	fmt.Println("after this is a hook")
// }
