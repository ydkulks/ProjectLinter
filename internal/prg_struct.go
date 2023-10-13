package prg_struct

import (
	"encoding/json"
	"fmt"
	"os"
)

// Global custom type
type DirectoryStructure struct{
	IgnoreDir      []string
	RootFiles      []string
	RootDirs       []string
	NonRootFiles   []string
	NonRootDirs    []string
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
			NonRootFiles: []string{
				// JS
				"config/database.js",
				"config/routes.js",
				"config/environment.js",
				// TS
				"config/constants.ts",
				"config/routes.ts",
			},
			NonRootDirs: []string{
				// JS
				"src/controllers",
				"src/models",
				"src/routes",
				"src/views",
				// TS
				"src/components",
				"src/containers",
				"src/services",
				"src/styles",
				"src/assets",
				"config/env",
				"config/",
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
				"composer.json",
				"composer.lock",
				"LICENSE",
				"README.md",
				"CODE_OF_CONDUCT.md",
				".env",
				".env.local",
				".gitignore",
			},
			RootDirs: []string{
				"app",
				"config",
				"public",
				"resources",
				"routes",
				"tests",
				"vendor",
			},
			NonRootFiles: []string{
				"config/database.php",
			},
			NonRootDirs: []string{
				"app/Controllers",
				"app/Models",
				"app/Views",
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
				"assets",
				"templates",
				"cmd",
				"internal",
				"pkg",
				"vendor",
			},
			NonRootFiles: []string{
				"cmd/main.go",
			},
			NonRootDirs: []string{
				"internal/handlers",
				"internal/models",
				"internal/middleware",
				"internal/config",
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
			NonRootFiles: []string{
				"cmd/mycli/main.go",
			},
			NonRootDirs: []string{
				"cmd/mycli",
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
			NonRootFiles: []string{
				"lua/commands.lua",
				"lua/mappings.lua",
				"lua/utils.lua",
			},
			NonRootDirs: []string{
			},
		}
		DirWalk(data)
		return
	}
	if prg_type == "Other"{
		data := DirectoryStructure{
			IgnoreDir: []string{},
			RootFiles: []string{},
			RootDirs: []string{},
			NonRootFiles: []string{},
			NonRootDirs: []string{},
		}
		// Take input for file structure
		addFilesAndDirs(data)
		return
	}
}

func addFilesAndDirs(data DirectoryStructure){
	// Store the data in a file for later use
	// If file is not available, then ask for user input

	// Check if save file exists or not
	if _,err := os.Stat("structure.json"); os.IsNotExist(err){
		fmt.Println("\033[31m" + "save file does not exist" + "\033[0m")
	}else{
		// fmt.Println("save file exist")
		file, err := os.Open("structure.json")
		if err != nil {fmt.Println(err)}
		defer file.Close()

		// var data DirectoryStructure
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&data); err!=nil {fmt.Println(err)}
	}

	// Add survey options to select the arrays under data
	for{
		fmt.Printf("%s Add file or folder (q to quit): ",status_q)
		var userData string
		_,err := fmt.Scan(&userData)
		if err != nil {fmt.Println(err)}
		if userData == "q"{break}else{
			data.IgnoreDir = append(data.IgnoreDir, userData)
			fmt.Println(data.IgnoreDir)
		}
	}
}
