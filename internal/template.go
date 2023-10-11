package prg_struct

import (
	"fmt"
	"os"
	"path/filepath"
	// "regexp"
)

func DirWalk(data DirectoryStructure){
	error := filepath.Walk(Pgr_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			return err
		}

		// Root files
		for i:=0; i <=len(data.RootFiles)-1; i++{
			if !info.IsDir() && info.Name() == data.RootFiles[i]{
				if filepath.Dir(path) == Pgr_dir{
					fmt.Printf("%s Root: %q\n",status_ok,path)
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
					fmt.Printf("%s Root directory: %q\n",status_ok,path)
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
				fmt.Printf("%s Skipped directory: %q\n",status_ignore,path)
				return filepath.SkipDir
			}
		}

		// Other files and directories
		if info.IsDir() {
			fmt.Printf("%s Other directory: %q\n",status_q,path)
		}else{
			fmt.Printf("%s Other file: %q\n",status_q,path)
		}
		return nil
	})

	if error != nil {
		fmt.Printf("Error walking directory: %v\n", error)
	}

}
