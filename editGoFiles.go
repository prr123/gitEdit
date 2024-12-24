// program that copies directory tree


package main

import (
	"fmt"
	"log"
	"os"
	"path"

    tree "goDev/gitEdit/gitEditLib"
    util "github.com/prr123/utility/utilLib"
)

func main() {

    numarg := len(os.Args)
    flags:=[]string{"dbg", "root", "new", "search", "replace"}

    useStr := " /root=rootDir /new=nrootdir /search=<search> /replace  [/dbg]"
    helpStr := "gitEdit program"

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

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    dbg:= false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}

    rootDir := ""
    inval, ok := flagMap["root"]
    if !ok {
        log.Fatalf("error -- no in flag provided!\n")
    } else {
        if inval.(string) == "none" {log.Fatalf("error -- no input file name provided!\n")}
        rootDir = inval.(string)
    }

    newRootDir := ""
    nrval, ok := flagMap["new"]
    if !ok {
        log.Fatalf("error -- no new flag provided!\n")
    } else {
        if nrval.(string) == "none" {log.Fatalf("error -- no new root dir name provided!\n")}
        newRootDir = nrval.(string)
    }

    searchStr := ""
    sval, ok := flagMap["search"]
    if ok {
        if sval.(string) != "none" {searchStr = sval.(string)}
    }

    replStr := ""
    rval, ok := flagMap["replace"]
    if ok {
//        log.Fatalf("error -- no search flag provided!\n")
//    } else {
        if rval.(string) == "none" {log.Fatalf("error -- no replace string provided!\n")}
        replStr = rval.(string)
    }

	wdPath, err := os.Getwd()
	if err != nil {log.Fatalf("could not get working directory!")}

    rootDirnam := path.Join(wdPath, rootDir)
    newRootDirnam := path.Join(wdPath, newRootDir)

    if dbg {
        fmt.Printf("root:     %s\n", rootDirnam)
        fmt.Printf("new root: %s\n", newRootDirnam)
		fmt.Printf("search:   %s\n", searchStr)
		fmt.Printf("replace:  %s\n", replStr)
    }


	replace := true
	if len(searchStr) == 0 {replace = false}
	if len(replStr) == 0 {replace = false}
	if dbg {fmt.Printf("replace: %t\n", replace)}

	if !replace {
		log.Printf("no search or replace string!\n")
		os.Exit(0)
	}

	searchStr = "\"github.com/" + searchStr
	if dbg {fmt.Printf("actual search: '%s'\n", searchStr)}

    err = tree.EditFiles(rootDirnam, newRootDirnam, searchStr, replStr)
    if err != nil {log.Fatalf("modDirs: %v\n", err)}

	fmt.Println("*** success ***")
}

