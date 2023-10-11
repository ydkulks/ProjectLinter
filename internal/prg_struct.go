package prg_struct

import "fmt"

// Global custom type
type DirectoryStructure struct{
	IgnoreDir      []string
	IgnoreFileType string
	RootFiles      []string
	RootDirs       []string
	Public         string
}

// Global variables
var( 
	Pgr_dir string
  // ANSI colors
  green = "\033[32m"
  blue = "\033[34m"
  yellow = "\033[33m"
  red = "\033[31m"
  reset = "\033[0m"
  
  // Symbols and its color
  status_ok = green + "✓" + reset
  status_bad = red + "✗" + reset
  status_q = blue + "?" + reset
  status_ignore = yellow + "Ø" + reset
)

func ProjectStructure(prg_type string){

	fmt.Printf("%s Project Directory: ",status_q)
	_,err := fmt.Scan(&Pgr_dir)
	if err != nil {
		fmt.Printf("\n%s Error: %v\n",status_bad,err)
	}

	if prg_type == "Javascript"{
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
		DirWalk(data)
		return
	}
	if prg_type == "PHP"{
		data := DirectoryStructure{
			IgnoreDir: []string{
				".git",
			},
			RootFiles: []string{
				"LICENSE",
				"README.md",
				"CODE_OF_CONDUCT.md",
				".env",
				".env.local",
				".gitignore",
			},
			RootDirs: []string{
			},
		}
		DirWalk(data)
		return
	}
	if prg_type == "Go (Web app)"{
		data := DirectoryStructure{
			IgnoreDir: []string{
				".git",
			},
			RootFiles: []string{
				"go.mod",
				"go.sum",
				"LICENSE",
				"README.md",
				"CODE_OF_CONDUCT.md",
				".env",
				".env.local",
				".gitignore",
			},
			RootDirs: []string{
			},
		}
		DirWalk(data)
		return
	}
	if prg_type == "Go (CLI)"{
		data := DirectoryStructure{
			IgnoreDir: []string{
				".git",
			},
			RootFiles: []string{
				"go.mod",
				"go.sum",
				"LICENSE",
				"README.md",
				"CODE_OF_CONDUCT.md",
				".env",
				".env.local",
				".gitignore",
			},
			RootDirs: []string{
				"cmd",
				"internal",
				"pkg",
				"vendor",
			},
		}
		DirWalk(data)
		return
	}
	if prg_type == "Lua (Plugin)"{
		data := DirectoryStructure{
			IgnoreDir: []string{
				".git",
			},
			RootFiles: []string{
				"LICENSE",
				"README.md",
				"CODE_OF_CONDUCT.md",
				".gitignore",
			},
			RootDirs: []string{
				"plugin",
				"lua",
				"doc",
				"tests",
			},
		}
		DirWalk(data)
		return
	}
	if prg_type == "Other"{
		return
	}
}
