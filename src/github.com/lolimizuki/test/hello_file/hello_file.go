package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// UseIO()
	// UseIOUtil()
	// Column_File()
	// WriteFile()
	// CopyFile()
	// ReadWithSlice()
	put_3_til_5_char_to_new_file()
}

func UseIO() {
	inputFileName := "testFile.txt"

	file, err := os.Open(inputFileName)

	if err != nil {
		fmt.Println("open error")
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	lineNum := 0
	for {
		lineText, readErr := reader.ReadString('\n')
		if readErr == io.EOF {
			fmt.Println("... read end")
			break
		}

		fmt.Println(lineNum, ": ", lineText)
		lineNum++
	}
}

func UseIOUtil() {
	inputFileName := "testFile.txt"
	outputFileName := "testFile_copy.txt"

	buff, inErr := ioutil.ReadFile(inputFileName)

	if inErr != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", inErr)
	}

	fmt.Printf("%s", string(buff))

	outErr := ioutil.WriteFile(outputFileName, buff, 0x644)
	if outErr != nil {
		panic(outErr) // 同等 printf
	}

	fmt.Println(" ... end")
}

// 若檔案以 以空白作為 column 時, 使用 FScan 讀取
func Column_File() {
	fileName := "column_file.txt"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		var op, v1str, v2str string
		_, scanErr := fmt.Fscanln(file, &op, &v1str, &v2str)
		if scanErr != nil {
			break
		}

		v1, _ := strconv.Atoi(v1str)
		v2, _ := strconv.Atoi(v2str)
		result := func() int {
			switch strings.ToLower(op) {
			case "add":
				return v1 + v2
			case "min":
				return v1 - v2
			case "mul":
				return v1 * v2
			case "div":
				return v1 / v2
			}
			return 0
		}()

		fmt.Println(op, v1str, v2str, " = ", result)
	}

	fmt.Println(" ... end")
}

func WriteFile() {
	outputFile, fErr := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if fErr != nil {
		panic(fErr)
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	writeInString := "Write to file !!!!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(writeInString)
	}
	outputWriter.Flush()
}

func CopyFile() {
	srcFile, _ := os.Open("testFile.txt")
	dstFile, _ := os.OpenFile("testFile_COPY.txt", os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(dstFile, srcFile)
}

func ReadWithSlice() {
	file, err := os.Open("testFile.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("can not open file to read")
		return
	}

	ReadFileUseSlice(file)
}

func ReadFileUseSlice(file *os.File) {
	// 每一次只讀出 3 個 byte
	const NUMBER_OF_BYTE_PER_READ = 3

	for {
		var bytes [NUMBER_OF_BYTE_PER_READ]byte
		num, _ := file.Read(bytes[:])
		if num == 0 {
			break
		}

		fmt.Printf("[")
		for i := 0; i < num; i++ {
			fmt.Printf("%s", string(bytes[i]))
			if i < num-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("]\n")
	}
}

func put_3_til_5_char_to_new_file() {
	srcFile, _ := os.Open("testFile.txt")
	dstFile, _ := os.OpenFile("35File.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer srcFile.Close()
	defer dstFile.Close()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)

	for {
		lineString, _, isEof := reader.ReadLine()
		if isEof == io.EOF {
			fmt.Println("... ok")
			break
		}

		outputString := string(lineString[3:6]) + "\n"
		writer.WriteString(outputString)
	}

	writer.Flush()
}
