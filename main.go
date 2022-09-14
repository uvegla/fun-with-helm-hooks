package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	ch := make(chan error)

	go func() {
		//coreBasic()

		coreHelm()

		close(ch)
	}()

	select {
	case <-ch:
		// Fall through.
	case <-time.After(5 * time.Second):
		fmt.Print("Timeout reached...")
	}
}

func coreHelm() {
	cmd := exec.Command("helm", "install", "--set", "pi.digits=6000,pi.backoffLimit=10", "dummy", "/Users/laszlouveges/sandbox/investigations/dummy/helm/dummy")
	stdout, err := cmd.Output()

	if err != nil {
		log.Println(err)
	}

	fmt.Print(stdout)
}

func coreBasic() {
	//cmd := exec.Command("echo", "Hello world!")
	cmd := exec.Command("pwd")
	pwd, _ := cmd.Output()

	fmt.Print("PWD: " + string(pwd))

	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	n := 1
	for n < 1000 {
		if _, err := f.WriteString(strconv.Itoa(n) + "\n"); err != nil {
			log.Println(err)
		}

		n += 1

		cmd = exec.Command("sleep", "1")
		stdout, _ := cmd.Output()
		fmt.Print(string(stdout))
	}

	if err != nil {
		fmt.Printf("An error occured: %#v", err)
	}

	exec.Command("touch", "/Users/laszlouveges/sandbox/investigations/dummy/alma.txt")

	fmt.Print("Done here")
}
