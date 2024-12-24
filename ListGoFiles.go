package main


import (
	"os"
	"fmt"
	"log"
	"path"

	tree "goDev/gitEdit/gitEditLib"
    cliutil "github.com/prr123/utility/utilLib"
)


func main() {

    numarg := len(os.Args)
    flags:=[]string{"dbg", "root"}

    useStr := " /root=rootdir [/dbg]"
    helpStr := "program file tree"

    if numarg > len(flags) +1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s %s\n", os.Args[0], useStr)
        os.Exit(-1)
    }

    if numarg == 1 || (numarg > 1 && os.Args[1] == "help") {
        fmt.Printf("help: %s\n", helpStr)
        fmt.Printf("usage is: %s %s\n", os.Args[0], useStr)
        os.Exit(1)
    }

    flagMap, err := cliutil.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}


    dbg:= false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}


    inFil := ""
    inval, ok := flagMap["root"]
    if !ok {
        log.Fatalf("error -- no in flag provided!\n")
    } else {
        if inval.(string) == "none" {log.Fatalf("error -- no input file name provided!\n")}
        inFil = inval.(string)
    }

    rootPath, err := os.Getwd()
    if err != nil {log.Fatalf("error Getwd: %v\n", err)}
    rootDir := path.Join(rootPath,inFil)

    if dbg {
        fmt.Printf("root:  %s\n", rootDir)
//        fmt.Printf("output: %s\n", newRoot)
    }

	err = tree.ListDirs(rootDir)
	if err != nil {log.Fatal("listDir: %v\n", err)}
	fmt.Println("success")
}

