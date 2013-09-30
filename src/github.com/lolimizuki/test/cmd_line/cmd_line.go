package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	useOS()
	// useExec()
}

func useOS() {
	env := os.Environ()
	// 輸出 os.Environ()
	// for _, s := range env {
	// 	fmt.Println(s)
	// }

	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
}

func useExec() {
	cmd := exec.Command("ls", "-l")
	out, err := cmd.Output()
	if err != nil {
		return
	}

	fmt.Println(string(out))
}
