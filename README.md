# GitEdit

The purpose of these programs is to list and edit the go projects that consists of numerous files which all import files from a respositary.  
This feature is useful if one wants to import a go project from github and compile it locally with the local imported files.  

In general the programs search all go files in a directory and all of its subdirectories for 

## listGoFiles

progam that lists all go files in the root directory and its subdirectories (the filetree).  
usage: ./ListGoFiles  /root=rootdir [/dbg]  

## copyGoFiles

progam that copies the filetree, but excludes all git files.  
usage: ./copyGoFiles  /root=rootDir /new=nrootdir /search=<search> /replace  [/dbg]  
note: search is the file string in the import section.  

## editGoFiles

progam that modifies all go source files of the filetree excluding git files.  
usage: ./editGoFiles  /root=rootDir /new=nrootdir /search=<search> /replace=<replace>  [/dbg]
note: replace is the string that will be substituted for each occurrence of the search string.  

