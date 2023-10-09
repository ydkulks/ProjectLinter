package prg_struct

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func DirWalk(){
	// ANSI colors
	green := "\033[32m"
	blue := "\033[34m"
	// yellow := "\033[33m"
	red := "\033[31m"
	reset := "\033[0m"

	// Symbols and its color
	status_ok := green + "✓" + reset
	status_bad := red + "✗" + reset
	status_q := blue + "?" + reset

	fmt.Printf("%s Project Directory: ",status_q)
	var pgr_dir string
	_,err := fmt.Scan(&pgr_dir)
	if err != nil {
		fmt.Printf("\n%s Error: %v\n",status_bad,err)
	}

	error := filepath.Walk(pgr_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			return err
		}

		// Ignore folders
		ignoreDir := [20]string{
			".git",
			".next",
			"node_modules",
		}
		ignoreFileType := `\.md$|\.env$|\.local$`
		ignoreFiles := [20]string{
			"package.json",
			"package-lock.json",
			"LICENSE",
			"vercel.json",
			".gitignore",
			"tsconfig.json",
			".eslintrc.json",
		}

		regex, err1 := regexp.Compile(ignoreFileType)
		if err1 != nil {return err1}
		if regex.MatchString(path) && !info.IsDir(){
			fmt.Printf("%s Ignored file extensions: %q\n",status_bad,path)
			return nil
		}

		for i:=0; i <=len(ignoreFiles)-1; i++{
			if !info.IsDir() && info.Name() == ignoreFiles[i]{
				fmt.Printf("%s Skipped file: %q\n",status_bad,path)
				return nil
			}
		}

		for i:=0; i <= len(ignoreDir)-1; i++ {
			if info.IsDir() && info.Name() == ignoreDir[i] {
				fmt.Printf("%s Skipped directory: %q\n",status_bad,path)
				return filepath.SkipDir
			}
		}

		fmt.Printf("%s Other file: %q\n",status_ok,path)
		return nil
	})

	if error != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}

}
