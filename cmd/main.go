package main

import (
	"github.com/ydkulks/ProjectLinter/internal"
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

	prg_struct.DirWalk(selectedOption)
}
