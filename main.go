package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Archiker-715/Task-Tracker/constants"
	"github.com/Archiker-715/Task-Tracker/internal/task"
)

var tasks *task.Tasks = &task.Tasks{}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Hello! Insert your task request\n")

cliLoop:
	for {
		usersInput, _ := reader.ReadString('\n')
		formattedInput := strings.TrimSpace(usersInput)
		inp := strings.Split(formattedInput, " ")

		switch inp[0] {
		case constants.List:
			if err := listTasks(inp); err != nil {
				fmt.Printf("error when adding task: '%v'. Please try again", err)
			}
		case constants.Add:
			if err := addTask(formattedInput); err != nil {
				fmt.Printf("error when adding task: '%v'. Please try again", err)
			}
		case constants.MarkInProgress:
			if err := updateStatusTask(inp); err != nil {
				fmt.Printf("error when updating task's status: '%v'. Please try again", err)
			}
		case constants.MarkDone:
			if err := updateStatusTask(inp); err != nil {
				fmt.Printf("error when updating task's status: '%v'. Please try again", err)
			}
		case constants.Delete:
			if err := deleteTask(inp); err != nil {
				fmt.Printf("error when deleting task's status: '%v'. Please try again", err)
			}
		case constants.Update:
			if err := updateTask(inp, formattedInput); err != nil {
				fmt.Printf("error when updating task: '%v'. Please try again", err)
			}
		case constants.Exit:
			break cliLoop
		default:
			fmt.Printf("no one operation match with %q. Please try again", inp[1])
			continue
		}
	}
}

func addTask(rawTaskDescription string) error {
	taskDecription, err := task.Unquote(rawTaskDescription)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	newTaskId, err := tasks.AddTask(taskDecription)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("Task added successfully (ID: %d)\n", newTaskId)
	return nil
}

func listTasks(s []string) error {
	if len(s) == 1 {
		err := tasks.ListTasks()
		return task.CheckErr(err)
	}

	if s[1] == constants.Done || s[1] == constants.Todo || s[1] == constants.InProgress {
		err := tasks.FilteredListTasks(s[1])
		return task.CheckErr(err)
	} else {
		return fmt.Errorf("no one operation match with %q", s[1])
	}
}

func updateStatusTask(s []string) error {
	taskId, err := strconv.Atoi(s[1])
	if err != nil {
		return fmt.Errorf("failed get taskId")
	}
	err = tasks.UpdateTask(taskId, s[0])
	return task.CheckErr(err)
}

func updateTask(s []string, rawTaskDescription string) error {
	taskDecription, err := task.Unquote(rawTaskDescription)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	taskId, err := strconv.Atoi(s[1])
	if err != nil {
		return fmt.Errorf("failed get taskId")
	}
	err = tasks.UpdateTask(taskId, taskDecription)
	return task.CheckErr(err)
}

func deleteTask(s []string) error {
	taskId, err := strconv.Atoi(s[1])
	if err != nil {
		return fmt.Errorf("failed get taskId")
	}
	err = tasks.DeleteTask(taskId)
	return task.CheckErr(err)
}
