package prg_struct

import (
	"github.com/AlecAivazis/survey/v2"
	"encoding/json"
	"fmt"
	"os"
)

func addFilesAndDirs(data DirectoryStructure){
	// Store the data in a file for later use
	// Check if save file exists or not
	if _,err := os.Stat("structure.json"); os.IsNotExist(err){
		fmt.Println(red + "save file does not exist" + reset)
	}else{
		// fmt.Println("save file exist")
		file, err := os.Open("structure.json")
		if err != nil {fmt.Printf(red + "%s\n" + reset, err)}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&data); err!=nil {fmt.Printf(red + "%s\n" + reset,err)}
	}

	// Add survey options to select the arrays under data
			// IgnoreDir: []string{},
			// RootFiles: []string{},
			// RootDirs: []string{},
			// NonRootFiles: []string{},
			// NonRootDirs: []string{},
  var selectedOption string
	options := []string{
		"Directories to ignore",
		"Files in Root",
		"Directories in Root",
		"NonRoot files",
		"NonRoot Directories",
	}
	arrayMap := map[string]*[]string{
		"Directories to ignore": &data.IgnoreDir,
		"Files in Root": &data.RootFiles,
		"Directories in Root": &data.RootFiles,
		"NonRoot files": &data.NonRootFiles,
		"NonRoot Directories": &data.NonRootDirs,
	}
  prompt := &survey.Select{
		Message: "Define you're project structure: ",
    Options: options,
  }
  survey.AskOne(prompt, &selectedOption)
	fmt.Println(arrayMap[selectedOption])
	selected := arrayMap[selectedOption]

	for{
		fmt.Printf("%s Add file or folder (q to quit): ",status_q)
		var userData string
		_,err := fmt.Scan(&userData)
		if err != nil {fmt.Println(red + "%s\n" + reset,err)}
		if userData == "q"{break}else{
			// data.IgnoreDir = append(data.IgnoreDir, userData)
			*selected = append(*selected, userData)
			fmt.Println(*selected)
		}
	}
	// If data != structure.json > update json
	// Display, add and modify json using CLI
}
