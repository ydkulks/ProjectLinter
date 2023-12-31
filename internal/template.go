package prg_struct

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func DirWalk(data DirectoryStructure){
	// Regex compilation
	var fileRegexPatterns []*regexp.Regexp
	var dirRegexPatterns []*regexp.Regexp

	for _, pattern := range data.NonRootFiles {
		re, err := regexp.Compile(pattern)
		if err != nil{
			fmt.Println(red + "%s\n" + reset,err)
			continue
		}
		fileRegexPatterns = append(fileRegexPatterns, re)
	}
	for _, pattern := range data.NonRootDirs {
		re, err := regexp.Compile(pattern)
		if err != nil{
			fmt.Println(red + "%s\n" + reset,err)
			continue
		}
		dirRegexPatterns = append(dirRegexPatterns, re)
	}

	// Regex for file extension validation
	dirRegexMap := make(map[string]*regexp.Regexp)
	allowedExtensionsMap := make(map[string]map[string]bool)
	for dir := range data.DirectoryFileExtensions {
		dirRegex := regexp.MustCompile(`.*[/\\]` + dir + `[/\\].*`)
		dirRegexMap[dir] = dirRegex

		allowedExtensions := make(map[string]bool)
		for i := range data.DirectoryFileExtensions[dir] {
			allowedExtensions[i] = true
		}
		allowedExtensionsMap[dir] = allowedExtensions
	}

	// Title
	fmt.Print("\n")
	fmt.Print(yellow + "     PROJECT LINTER" + reset)
	fmt.Print("\n")
	fmt.Print("\n")

	// Path walk
	error := filepath.Walk(Pgr_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(red + "%s\n" + reset,err)
			return err
		}

		// Root files
		for i:=0; i <=len(data.RootFiles)-1; i++{
			if !info.IsDir() && info.Name() == data.RootFiles[i]{
				if filepath.Dir(path) == Pgr_dir{
					fmt.Printf("%s Root files: "+dark_gray+"%q\n",status_ok,path)
					return nil
				}else{
					fmt.Printf("%s FILE NOT IN ROOT: %q\n",status_bad,path)
					return nil
				}
			}
		}

		// Root directory folders
		for i:=0; i <=len(data.RootDirs)-1; i++{
			if info.IsDir() && info.Name() == data.RootDirs[i]{
				if filepath.Dir(path) == Pgr_dir{
					fmt.Printf("%s Root directory: "+dark_gray+"%q\n",status_ok,path)
					return nil
				}else{
					fmt.Printf("%s FOLDER NOT IN ROOT: %q\n",status_bad,path)
					return nil
				}
			}
		}

		// Ignore directory
		for i:=0; i <= len(data.IgnoreDir)-1; i++ {
			if info.IsDir() && info.Name() == data.IgnoreDir[i] {
				fmt.Printf("%s Skipped directory: "+dark_gray+"%q\n",status_ignore,path)
				return filepath.SkipDir
			}
		}

		// File validation
		for _, re := range fileRegexPatterns{
			if re != nil && re.MatchString(path){
				fmt.Printf("%s File: "+dark_gray+"%q\n",status_ok,path)
				return nil
			}
		}

		// Dir validation
		for _, re := range dirRegexPatterns{
			if re != nil && re.MatchString(path){
				fmt.Printf("%s Directory: "+dark_gray+"%q\n",status_ok,path)
				return nil
			}
		}

		// File extension validation
		if !info.IsDir(){
			for dir, dirRegex := range dirRegexMap {
				if dirRegex.MatchString(path) {
					ext := filepath.Ext(info.Name())
					if allowedExtensions, ok := allowedExtensionsMap[dir]; ok {
						if !allowedExtensions[ext] {
							fmt.Printf("%s File with invalid extension: "+red+"%q\n", status_bad, path)
							return nil
						} else {
							fmt.Printf("%s Allowed file: "+dark_gray+"%q\n", status_ok, path)
							return nil
						}
					}
				}
			}
		}

		// Other files and directories
		if info.IsDir() {
			fmt.Printf("%s Other directory: "+medium_gray+"%q\n",status_q,path)
		}else if !info.IsDir(){
			fmt.Printf("%s Other file: "+medium_gray+"%q\n",status_q,path)
		}
		return nil
	})

	if error != nil {
		fmt.Printf("Error walking directory: %v\n", error)
	}

}
