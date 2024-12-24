package editFileTree


import (
	"os"
	"fmt"
	"bytes"
	"io/fs"
	"io"
//	"path"
)

type lintyp []byte

func ListDirs(rootDirnam string) (err error) {

//fmt.Printf("dbg -- dir name: %s\n", rootDirnam)
    entryList, err := os.ReadDir(rootDirnam)
   	if err != nil {return fmt.Errorf("ReadDir: %v", err)}
	subDirList := make([]fs.DirEntry,0, 100)
	goList := make([]fs.DirEntry,0, 100)
	dcount :=0
	for i:=0; i<len(entryList); i++ {
       	if entryList[i].IsDir() {
            nam := []byte(entryList[i].Name())
//fmt.Printf("dbg -- nam: %q %s\n", nam[0], nam)
            if nam[0] != '.' {
	           	subDirList = append(subDirList,entryList[i])
//fmt.Printf("dbg -- chDir[%d]: %s\n", dcount, subDirList[dcount].Name())
    	       	dcount++
			}
		} else {
			nam := []byte(entryList[i].Name())
			lnam := len(nam)
			if lnam > 3 {
				ext:= nam[lnam-3:]
				match := bytes.Equal([]byte(".go"),ext)
				if match {
					goList = append(goList,entryList[i])
				}
			}
		}
	}
	if len(goList) > 0 {
		fmt.Printf("***** %s: go files *****\n", rootDirnam)
		for i,file := range goList {
			fmt.Printf("file [%d]: %s\n",i, file.Name())
		}
	}

// replace with path
	for i:=0; i< len(subDirList); i++ {
		subNam := rootDirnam + "/" + subDirList[i].Name()
//fmt.Printf("dbg -- child[%d] name: %s\n", i, subNam)
		err = ListDirs(subNam)
	   	if err != nil {fmt.Errorf("listDir: %v", err)}
	}
	return nil
}

func CopyDirs(rootDirnam, newrootDir string) (err error) {

//fmt.Printf("ModDirs: %s %s\n", rootDirnam, newrootDir)
    entryList, err := os.ReadDir(rootDirnam)
    if err != nil {return fmt.Errorf("ReadDir root: %v", err)}
//fmt.Printf("entrylist: %d\n", len(entryList))

	err = os.Mkdir(newrootDir, 0777)
    if err != nil {return fmt.Errorf("mkdir: %v", err)}
//fmt.Printf("created new rootdir: %s\n", newrootDir)

    subDirList := make([]fs.DirEntry,0, 100)
    goList := make([]fs.DirEntry,0, 100)
    dcount :=0
    for i:=0; i<len(entryList); i++ {
        if entryList[i].IsDir() {
            nam := []byte(entryList[i].Name())
//fmt.Printf("dbg -- nam: %q %s\n", nam[0], nam)
            if nam[0] != '.' {
                subDirList = append(subDirList,entryList[i])
//				childDirnam := path.Join(newrootDir,entryList[i])
//fmt.Printf("dbg -- chDir[%d]: %s\n", dcount, subDirList[dcount])
                dcount++
            }
        } else {
            nam := []byte(entryList[i].Name())
            lnam := len(nam)
            if lnam > 3 {
                ext:= nam[lnam-3:]
                match := bytes.Equal([]byte(".go"),ext)
                if match {
                    goList = append(goList,entryList[i])
                }
            }
        }
    }

	fmt.Printf("***** %s: sub dir list *****\n", rootDirnam)
// replace with path
    for i:=0; i< len(subDirList); i++ {
        subNam := rootDirnam + "/" + subDirList[i].Name()
		childDirnam := newrootDir +"/" +subDirList[i].Name()
fmt.Printf("dbg -- child[%d] name: %s\n", i, childDirnam)
        err = CopyDirs(subNam, childDirnam)
        if err != nil {fmt.Errorf("modDir <%s:%s>: %v",subNam, childDirnam, err)}
    }

	fmt.Printf("***** %s: go files *****\n", rootDirnam)
    if len(goList) > 0 {
    	for i,file := range goList {
        	fmt.Printf("file [%d]: %s\n",i, file.Name())
			srcFilnam := rootDirnam + "/" + file.Name()
			destFilnam := newrootDir +"/" + file.Name()

    		srcFile, err := os.Open(srcFilnam)
    		if err != nil {return fmt.Errorf("src file open: %v", err)}
		    defer srcFile.Close()

    		destFile, err := os.Create(destFilnam) // creates if file doesn't exist
    		if err != nil {return fmt.Errorf("dest file create: %v", err)}
		    defer destFile.Close()

    		_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
    		if err != nil {return fmt.Errorf("copy file: %v", err)}
//    		err = destFile.Sync()
//    check(err)

    	}
	}

    return nil
}

func EditFiles(rootDirnam, newrootDir, searchStr, replStr string) (err error) {

//fmt.Printf("ModDirs: %s %s\n", rootDirnam, newrootDir)
    entryList, err := os.ReadDir(rootDirnam)
    if err != nil {return fmt.Errorf("ReadDir root: %v", err)}
//fmt.Printf("entrylist: %d\n", len(entryList))

	err = os.Mkdir(newrootDir, 0777)
    if err != nil {return fmt.Errorf("mkdir: %v", err)}
//fmt.Printf("created new rootdir: %s\n", newrootDir)

    subDirList := make([]fs.DirEntry,0, 100)
    goList := make([]fs.DirEntry,0, 100)
    dcount :=0
    for i:=0; i<len(entryList); i++ {
        if entryList[i].IsDir() {
            nam := []byte(entryList[i].Name())
//fmt.Printf("dbg -- nam: %q %s\n", nam[0], nam)
            if nam[0] != '.' {
                subDirList = append(subDirList,entryList[i])
//				childDirnam := path.Join(newrootDir,entryList[i])
//fmt.Printf("dbg -- chDir[%d]: %s\n", dcount, subDirList[dcount])
                dcount++
            }
        } else {
            nam := []byte(entryList[i].Name())
            lnam := len(nam)
            if lnam > 3 {
                ext:= nam[lnam-3:]
                match := bytes.Equal([]byte(".go"),ext)
                if match {
                    goList = append(goList,entryList[i])
                }
            }
        }
    }

	fmt.Printf("***** %s: sub dir list *****\n", rootDirnam)
// replace with path
    for i:=0; i< len(subDirList); i++ {
        subNam := rootDirnam + "/" + subDirList[i].Name()
		childDirnam := newrootDir +"/" +subDirList[i].Name()
//fmt.Printf("dbg -- child[%d] name: %s\n", i, childDirnam)
        err = EditFiles(subNam, childDirnam, searchStr, replStr)
        if err != nil {fmt.Errorf("modDir <%s:%s>: %v",subNam, childDirnam, err)}
    }

	fmt.Printf("***** %s: go files *****\n", rootDirnam)
    if len(goList) > 0 {
    	for i,file := range goList {
        	fmt.Printf("file [%d]: %s\n",i, file.Name())
			srcFilnam := rootDirnam + "/" + file.Name()
			destFilnam := newrootDir +"/" + file.Name()

    		source, err := os.ReadFile(srcFilnam)
    		if err != nil {return fmt.Errorf("src file open: %v", err)}

			dest, err := EditImportContent(source, searchStr, replStr)
    		if err != nil {return fmt.Errorf("EditFile: %v", err)}

			// creates if file doesn't exist
    		err = os.WriteFile(destFilnam, dest, 0666)
    		if err != nil {return fmt.Errorf("dest file write: %v", err)}
    	}
	}

    return nil
}

func EditImportContent(inData []byte, searchStr, replStr string) (out []byte, err error) {

	dbg := false
    idx := bytes.Index(inData, []byte("import ("))
    if idx == -1 {return nil, fmt.Errorf("no import statement found!")}
    if dbg {fmt.Printf("found import statement\n")}
    idxend := bytes.Index(inData[idx:], []byte(")"))
    if idxend == -1 {return nil, fmt.Errorf("no end to import statement found!")}

    searchst := idx + len("import (")
    searchend:= idx +idxend
    if dbg {fmt.Printf("dbg -- search [%d:%d] %s\n", searchst, searchend, inData[searchst:searchend])}

    fmt.Printf("***** import: ***\n")
    fmt.Printf("%s", inData[searchst:searchend])

    outMod, err := EditFileContent(inData[searchst:searchend], searchStr, replStr)
    if err != nil {return nil, fmt.Errorf("EditFileContent: %v", err)}

    outData := make([]byte, 0, 1024*1024)
    outData = append(outData, inData[:idx]...)
    outData = append(outData, []byte("import (\n")...)
    outData = append(outData, outMod...)
    outData = append(outData, []byte(")")...)
    outData = append(outData, inData[searchend:]...)

	return outData, nil
}


func EditFileContent(in []byte, searchStr, newStr string) (out []byte, err error) {
//	linCount :=0
	if len(in) == 0 {return nil, fmt.Errorf("no input\n")}
	nline := make([]byte,128)
	lines := make ([]lintyp, 0, 64)
	linst :=0
	for i:=0; i<len(in); i++ {
		if in[i] == '\n' {
			lines = append(lines,in[linst:i+1])
			linst = i+1
		}
	}
//	PrintLines(lines)

	// parse each line
	out = make([]byte,0, 1024*1024)

	newB := []byte(newStr)
	for i:=0; i< len(lines); i++ {
		line := lines[i]
//		fmt.Printf("line [%d]:%s\n", i, line)
		if len(line) > 1 && bytes.Equal(line[0:1],[]byte("//")) {continue}
		sidx := bytes.Index(line, []byte(searchStr))
		if sidx > -1 {
			tailst := sidx+len(searchStr)
			tailend := len(line)
//			fmt.Printf("match: %s\n", string(line))
			for m:=sidx;m<len(line); m++ {
				if line[m] == '"' {tailend=m}
			}

			for j:=0;j<4; j++ {nline[j] = ' '}
			nline[4]='"'
			for k:=0; k<len(newB); k++ {nline[k+5] = newB[k]}
			nlinSt := 5 + len(newB)
			for j:= 0; j<tailend-tailst+1; j++ {nline[nlinSt + j] = line[tailst+j]}
			nlinEnd := nlinSt+tailend-tailst
			nline[nlinEnd+1] = '\n'
//		fmt.Printf("new [%d]:%s\n",i, nline[:nlinEnd+2])
			lines[i] = nline[:nlinEnd+2]
		}
		out = append(out, lines[i]...)
	}

	return out, nil
}

func PrintLines(lines []lintyp) {

	fmt.Println("************ lines ***********")
	for i:=0; i< len(lines); i++ {
		fmt.Printf(" --%d:%s\n", i, lines[i])
	}
	fmt.Println("********** end lines *********")
}
