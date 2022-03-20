package dubbo

type User struct {
	ID   int32
	Name string
	Pass string
	Del  int32
}

func (user *User) JavaClassName() string {
	return "com.shy.provider.model.User"
}
