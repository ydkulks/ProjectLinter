package prg_struct

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirWalk(){
	green := "\033[32m"
	blue := "\033[34m"
	// yellow := "\033[33m"
	red := "\033[31m"
	reset := "\033[0m"

	status_ok := green + "✓" + reset
	status_bad := red + "✗" + reset
	status_q := blue + "?" + reset

	fmt.Printf("%s Project Directory: ",status_q)
	var pgr_dir string
	_,err := fmt.Scan(&pgr_dir)
	if err != nil {
		fmt.Printf("\n%s Error: %v\n",status_bad,err)
	}

	// fmt.Printf("%s Entered: %s\n",status_ok,pgr_dir)
	error := filepath.Walk(pgr_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			return err
		}
		// Example: Check for the presence of a "src" directory.
		if info.IsDir() && (info.Name() == ".git") || (info.Name() == "node_modules") {
			fmt.Printf("%s Skipped directory at %q\n",status_bad,path)
			return filepath.SkipDir
		}
		fmt.Printf("%s Visited file or directory: %q\n",status_ok,path)
		return nil
	})

	if error != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}

}
