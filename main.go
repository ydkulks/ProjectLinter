package main

import (
	"github.com/ydkulks/ProjectLinter/prg_struct"
	"github.com/AlecAivazis/survey/v2"
)

func main(){
	options := []string{"Javascript", "PHP", "Go", "Lua (Plugin)","Other"}

  var selectedOption string

  // Create a survey prompt for selecting an option
  prompt := &survey.Select{
    Message: "Select project type:",
    Options: options,
  }

  // Ask the user to select an option
  survey.AskOne(prompt, &selectedOption)

	if selectedOption == "Javascript"{
		prg_struct.DirWalk()
	}
	if selectedOption == "PHP"{
		return
	}
	if selectedOption == "Go"{
		return
	}
	if selectedOption == "Lua (Plugin)"{
		return
	}
	if selectedOption == "Other"{
		return
	}
}
