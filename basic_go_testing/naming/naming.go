package naming

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Bark(muzzled bool) {
	if muzzled {
		println("Woof")
	} else {
		println("WOOF !!")
	}
}
func Speak(lang string) {
	switch lang {
	case "spanish":
		println("Hola")
	default:
		println("Hello")
	}
}

func Color(name string) string {
	switch name {
	case "blue":
		return "#0000FF"
	case "red":
		return "#FFFFF"
	default:
		return "#00000"

	}
}
