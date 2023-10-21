package prg_struct

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func addFilesAndDirs(data DirectoryStructure) (DirectoryStructure){
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

	for{
		// Add survey options to select the arrays under data
		var selectedOption string
		options := []string{
			"Directories to ignore",
			"Files in Root",
			"Directories in Root",
			"NonRoot files",
			"NonRoot Directories",
			"File Extension Validation",
			"Quit",
		}
		arrayMap := map[string]*[]string{
			"Directories to ignore": &data.IgnoreDir,
			"Files in Root": &data.RootFiles,
			"Directories in Root": &data.RootDirs,
			"NonRoot files": &data.NonRootFiles,
			"NonRoot Directories": &data.NonRootDirs,
			"File Extension Validation": &data.NonRootDirs,
		}
		prompt := &survey.Select{
			Message: "Project structure: ",
			Options: options,
		}

		// Select array
		survey.AskOne(prompt, &selectedOption)
		fmt.Println(arrayMap[selectedOption])
		selected := arrayMap[selectedOption]
		if selectedOption == "Quit"{break}

		// Select operation
		operationOptions := []string{
			"Add",
			"Delete",
			"Display",
		}
		prompt = &survey.Select{
			Message: "Operation: ",
			Options: operationOptions,
		}
		survey.AskOne(prompt,&selectedOption)

		if selectedOption == "Add"{
			// Add
			for{
				fmt.Printf("%s Add file or folder "+green+"(m - menu): "+reset,status_q)
				var userData string
				_,err := fmt.Scan(&userData)
				if err != nil {fmt.Println(red + "%s\n" + reset,err)}
				if userData == "m"{break}else{
					// data.IgnoreDir = append(data.IgnoreDir, userData)
					*selected = append(*selected, userData)
					fmt.Println(*selected)
				}
			}
			updateJSON(data, selectedOption, *selected)
		}else if selectedOption == "Delete"{
			// Delete
			deleteOptions := *selected
			prompt = &survey.Select{
				Message: "Select File or Dir to delete: ",
				Options: deleteOptions,
			}
			survey.AskOne(prompt,&selectedOption)
			fmt.Println("Selected item to delete: ",selectedOption)

			var tempArray []string
			for _,value := range *selected{
				if value != selectedOption{
					tempArray = append(tempArray, value)
				}
			}
			*selected = tempArray
			fmt.Println("Current list:",*selected)
			updateJSON(data, selectedOption, *selected)
		}else if selectedOption == "Display"{
			// Display
			fmt.Println(*selected)
		}
	}
	// updateJSON(data)
	return data
}

func updateJSON(data DirectoryStructure, selectedOption string, selectedSlice []string){
	// Read JSON file
	jsonData, err := os.ReadFile("structure.json")
	if err != nil {
		fmt.Println(red + "%s\n" + reset, err)
	}

	// Unmarshal JSON into a map to update the appropriate slice
	var fileData map[string]interface{}
	err = json.Unmarshal(jsonData, &fileData)
	if err != nil {
		fmt.Println(red + "%s\n" + reset, err)
	}

	// Update the appropriate slice based on the selectedOption
	switch selectedOption {
	case "Directories to ignore":
		fileData["IgnoreDir"] = selectedSlice
	case "Files in Root":
		fileData["RootFiles"] = selectedSlice
	case "Directories in Root":
		fileData["RootDirs"] = selectedSlice
	case "NonRoot files":
		fileData["NonRootFiles"] = selectedSlice
	case "NonRoot Directories":
		fileData["NonRootDirs"] = selectedSlice
	}

	// Marshal the updated data back to JSON
	updatedJSON, err := json.MarshalIndent(fileData, "", " ")
	if err != nil {
		fmt.Println(red + "%s\n" + reset, err)
	}

	// Write JSON file
	err = os.WriteFile("structure.json", updatedJSON, 0644)
	if err != nil {
		fmt.Println(red + "%s\n" + reset, err)
	}

	fmt.Println(green + "Updated and saved" + reset)
	fmt.Println(data)
}
