package main

import (
	"flag"
	. "fmt"
	tt "time"
)

const (
	WHITE = iota
	RED
	BLUE
	GREEN
	YELLOW
)

var (
	msg = flag.String("msg", "Neko", "say something")
)

type Color byte

type Ghost struct {
	Name string
	Age  int
}

func (g *Ghost) Fly() {
	Println("i am", g.Name, "i fly over", g.Age, "years")
}

type Yuyuko struct {
	Ghost
	PetName string
}

func (y *Yuyuko) Fly() {
	y.Ghost.Fly()
	Println("my pet is", y.PetName)
}

func main() {
	MainInThird()

	return
	FuncCanExportFromSecFree()
	TimeAdd()

	flag.Parse()
	Println(*msg)

	Println(tt.Now().Hour(), ":", tt.Now().Minute())

	birthday, _ := tt.Parse("1.2006.2", "11.2009.10") // time.Time
	age := tt.Since(birthday)                         // time.Duration
	Printf("Go is %d years old\n", age/(tt.Hour*24*365))

	color := Color(WHITE)
	Println(color)

	color2 := RED
	Println(color2)

	yuyuko := Yuyuko{Ghost{"Yuyuko", 99}, "Yuumu"}
	yuyuko.Fly()

	Println("------")
	list := make([]interface{}, 3, 3)
	list[0] = "i am gopher"
	list[1] = yuyuko
	list[2] = 123

	for _, element := range list {
		switch t := element.(type) {
		case int:
			Println("You are INT", t)

		case string:
			Println("You are String", t)

		case Yuyuko:
			Println("Welocome home, Yuyuko Sama")

		default:
			Println("You are nothing")
		}
	}
}
