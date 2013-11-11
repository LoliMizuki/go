package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	os.Mkdir("nekoko", 0777)
	os.MkdirAll("nekoko2/nekoinneko/whyneko", 0777)

	os.Remove("nekoko")
	os.RemoveAll("nekoko2")

	f, _ := os.Create("myneko.txt")
	defer f.Close()

	f.WriteString("abcdefghijklmnopq")

	bytes := make([]byte, 6)
	ix, _ := f.ReadAt(bytes[:], 4)
	fmt.Println("read #bytes:", ix)
	fmt.Println(string(bytes))
}
