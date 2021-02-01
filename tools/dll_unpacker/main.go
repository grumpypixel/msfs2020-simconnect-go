package main

import (
	"fmt"
	"os"
	"path"

	"github.com/grumpypixel/msfs2020-simconnect-go/simconnect"
)

func main() {
	args := os.Args
	pathname := "."
	if len(args) > 1 {
		pathname = args[1]
	}

	filename := simconnect.SimConnectDLL
	fullpath := path.Join(pathname, filename)

	err := simconnect.UnpackDLL(fullpath)
	if err != nil {
		fmt.Println("Ugh. An error.", err)
	} else {
		fmt.Printf("Unpacked DLL to %s", fullpath)
	}
}
