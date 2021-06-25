package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type taskService interface {
	AddTask(input []string) error
	GetTasks() ([]string, error)
	ChangeTaskStatus(i int, active bool) error
	CleanDoneTasks() error
}

type Cli struct {
	taskService taskService
}

func New(taskService taskService) *Cli {
	return &Cli{
		taskService: taskService,
	}
}

func (c Cli) AddTask() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a task to the list.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("task name is required")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.taskService.AddTask([]string{args[0], "true", "false"})
		},
	}
}

func (c Cli) GetTasks() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all todolist tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			results, err := c.taskService.GetTasks()
			if err != nil {
				return err
			}

			fmt.Println(strings.Join(results, "\n"))

			return nil
		},
	}
}

func (c Cli) DoneTask() *cobra.Command {
	return &cobra.Command{
		Use:   "done",
		Short: "Mark Task as done",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("task id needed to mark task done")
			}
			_, err := strconv.ParseInt(args[0], 0, 32)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			i, _ := strconv.ParseInt(args[0], 0, 32)
			return c.taskService.ChangeTaskStatus(int(i), false)
		},
	}
}

func (c Cli) UnDoneTask() *cobra.Command {
	return &cobra.Command{
		Use:   "undone",
		Short: "Mark Task as undone",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("task id needed to mark task undone")
			}

			_, err := strconv.ParseInt(args[0], 0, 32)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			i, _ := strconv.ParseInt(args[0], 0, 32)
			return c.taskService.ChangeTaskStatus(int(i), true)
		},
	}
}

func (c Cli) Cleanup() *cobra.Command {
	return &cobra.Command{
		Use:   "cleanup",
		Short: "Cleanup done task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.taskService.CleanDoneTasks()
		},
	}
}
