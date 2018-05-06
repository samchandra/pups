package pupstemplate

func ProjectDomain() string {
	t :=
		`package {{.Name}}

type User struct {
    name     string
    password string
}

type Token struct {
    userID  string
    Token   string
}
`
	return t
}
