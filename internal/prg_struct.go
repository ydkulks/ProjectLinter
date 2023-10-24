package prg_struct

import (
	"fmt"
)

// Global custom type
type DirectoryStructure struct{
	IgnoreDir      []string
	RootFiles      []string
	RootDirs       []string
	NonRootFiles   []string
	NonRootDirs    []string
	DirectoryFileExtensions    map[string]map[string]bool
}

// Global variables
var( 
	Pgr_dir string
  // ANSI colors
  green = "\033[32m"
  blue = "\033[34m"
  yellow = "\033[33m"
  red = "\033[31m"
  dark_gray = "\033[38;2;100;100;100m"
  medium_gray = "\033[38;2;150;150;150m"
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
		fmt.Printf(red + "%s\n" + reset,err)
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
			DirectoryFileExtensions: map[string]map[string]bool{
				"public": {
					".html":true,
					".js":true,
					".css":true,
				},
				"src": {
					".ts":true,
				},
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
			DirectoryFileExtensions: map[string]map[string]bool{
				"app/Controllers":{".php":true},
				"app/Models":{".php":true},
				"app/Views":{".blade.php":true,".html":true,".twig":true},
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
			DirectoryFileExtensions: map[string]map[string]bool{
				"assets":{".css":true, ".js":true, ".jpg":true, ".png":true},
				"templates":{".html":true},
				"internal":{".go":true},
				"pkg":{".go":true},
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
			DirectoryFileExtensions: map[string]map[string]bool{
				"cmd":{".go":true},
				"internal":{".go":true},
				"pkg":{".go":true},
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
				"test",
			},
			NonRootFiles: []string{
				"lua/commands.lua",
				"lua/mappings.lua",
				"lua/utils.lua",
				"doc/tags",
			},
			NonRootDirs: []string{},
			DirectoryFileExtensions: map[string]map[string]bool{
				"lua":{".lua":true},
				"plugin":{".lua":true},
				"test":{".lua":true},
				"doc":{".txt":true},
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
			DirectoryFileExtensions: map[string]map[string]bool{},
		}
		// Take input for file structure
		updatedData := addFilesAndDirs(data)
		DirWalk(updatedData)
		return
	}
}
