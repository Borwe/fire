package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Program struct {}

func initProgram()*Program{
	return &Program{}
}

func (p Program) Init() tea.Cmd{
	return nil
}

func (p Program) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	return p, nil
}


func (p Program) View() string{
	return "YO"
}

func main(){
	prg := tea.NewProgram(initProgram())
	if _, err := prg.Run(); err!=nil {
		log.Fatalln("Error after run: ",err)
	}
}
