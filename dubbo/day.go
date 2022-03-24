package dubbo

type Day struct {
	Day  string
	Msg1 string
	Msg2 string
	Msg3 string
	Msg4 string
	Msg5 string
	Msg6 string
}

func (day *Day) JavaClassName() string {
	return "com.shy.provider.model.Day"
}
