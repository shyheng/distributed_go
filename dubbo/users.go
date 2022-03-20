package dubbo

type Users struct {
	ID   int32
	Name string
	Pass string
	Del  int32
}

func (user *Users) JavaClassName() string {
	return "com.shy.provider.model.Users"
}
