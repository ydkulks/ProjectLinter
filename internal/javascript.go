package prg_struct

import (
	"fmt"
	"os"
	"path/filepath"
	// "regexp"
)

type DirectoryStructure struct{
	IgnoreDir      []string
	IgnoreFileType string
	RootFiles      []string
	RootDirs       []string
	Public         string
}

func JsDirWalk(){
	// ANSI colors
	green := "\033[32m"
	blue := "\033[34m"
	yellow := "\033[33m"
	red := "\033[31m"
	reset := "\033[0m"

	// Symbols and its color
	status_ok := green + "✓" + reset
	status_bad := red + "✗" + reset
	status_q := blue + "?" + reset
	status_ignore := yellow + "Ø" + reset

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

		data := DirectoryStructure{
			IgnoreDir: []string{
				".git",
				"node_modules",
				".next",
			},
			RootFiles: []string{
				"package.json",
				"package-lock.json",
				"tailwind.config.js",
				"tailwind.config.ts",
				"tsconfig.json",
				"tslint.json",
				"webpack.config.js",
				"yarn.lock",
				"postcss.config.js",
				".eslintrc.json",
				"LICENSE",
				"README.md",
				"CODE_OF_CONDUCT.md",
				".env",
				".env.local",
				".gitignore",
				"app.js",
			},
			RootDirs: []string{
				"public",
				"src",
				"tests",
				"config",
				"typings",
			},
		}

		// Root files
		for i:=0; i <=len(data.RootFiles)-1; i++{
			if !info.IsDir() && info.Name() == data.RootFiles[i]{
				if filepath.Dir(path) == pgr_dir{
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
				if filepath.Dir(path) == pgr_dir{
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
		fmt.Printf("Error walking directory: %v\n", err)
	}

}
