package notetaker

import (
	"fmt"
	"os"

	fancyui "github.com/common-nighthawk/go-figure"
	writeTable "github.com/olekukonko/tablewriter"
)

func NoteTaker(){
	var startAction string

	data := [][]string{
		{"S/N", "Task", "Status"},
		{"1", "Short description", "DONE"},
		{"2", "Another lengthy description requiring truncation or wrapping", "ERROR"},
	}

	noteTaker := fancyui.NewColorFigure("Note Taker", "", "purple", true)
	noteTaker.Print()

	fmt.Println("Please choose an action:\n 1. To Add a note\n 2. To View notes:")	
	fmt.Scan(&startAction) 

	switch startAction {
		case "1":	
		AddNote();
		return
		case "2":
			fmt.Println("Viewing notes...")
			if len(data) <=1 {
				fmt.Println("No notes available.")
				return
			}
				table := writeTable.NewWriter(os.Stdout)
				table.SetHeader(data[0])
				table.AppendBulk(data[1:])
				table.Render()
				return
	}

}