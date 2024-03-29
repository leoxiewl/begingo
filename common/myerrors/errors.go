package myerrors

type fundamental struct {
	msg string
}

func (f *fundamental) Error() string {
	return f.msg
}

func New(message string) error {
	return &fundamental{
		msg: message,
	}
}
