package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/d4v1d0v1c/gshell/internal/repl"
)

/*
VAR int ret,freq, int, vpid, ppid, apid
VAR string para, amode
MAIN
    para <- getenv "#1"
    if para=""
      print "Frequency missing"
      freq <- 626000
    else
      freq <- integer para
    endif
END
*/

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the gshell!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
