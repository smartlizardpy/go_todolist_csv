package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func showtodo() {
	file, err := os.Open("tasks.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := records[1:]
	t := table.NewWriter()
	t.SetTitle("Todo")
	t.AppendHeader(table.Row{"#", "Task", "Status"})

	for i, c := range data {
		t.AppendRow(table.Row{i, c[0], c[1]})
	}
	fmt.Print(t.Render(), "\n")
}

func addtodo(data string) {
	filename := "tasks.csv"
	final_text := data + ",Not Done\n"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(final_text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully Added:", data, " and marked it as Not Done")
	var selected string
	fmt.Print("\nDo you want to see the list? (Y/n): ")
	fmt.Scanln(&selected)
	if selected == "y" || selected == "Y" || selected == "" {
		showtodo()
	}
}

func mark(id int) {
	filename := "tasks.csv"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv:", err)
		return
	}

	if id < 0 || id >= len(records)-1 {
		fmt.Println("Invalid task ID.")
		return
	}

	records[id+1][1] = "Done"

	file, err = os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	writer.Flush()
	fmt.Println("Successfully marked task", id, "as Done.")
	showtodo()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Help:\n--list : lists the content of the todo list\n--add : adds a todo to the list\n--mark: marks a todo done\n\n-----------------------------")
		return
	}

	args := os.Args[1:]
	command := args[0]

	if !strings.HasPrefix(command, "--") {
		fmt.Println("Invalid command. Please use '--' prefix for commands. Use --help for more info.")
		return
	}

	switch command {
	case "--list":
		showtodo()
	case "--add":
		if len(args) < 2 {
			fmt.Println("Please provide a task to add.")
			return
		}
		addtodo(strings.Join(args[1:], " "))
	case "--mark":
		if len(args) < 2 {
			fmt.Println("Please provide a task ID to mark as done.")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a number.")
			return	
		}
		mark(id)
	case "--help":
		fmt.Println("Help:\n--list : lists the content of the todo list\n--add : adds a todo to the list\n--mark: marks a todo done\n\n-----------------------------")
	default:
		fmt.Println("Unknown command. Use --help for a list of available commands.")
	}
}
