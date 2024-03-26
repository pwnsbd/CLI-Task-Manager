package cmd

import (
	"fmt"
	"strconv"

	"github.com/pwnsbd/Task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong", err.Error())
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invaid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println("Failed to mark \"%d\" as completed. Error : %s\n", id, err)
			} else {
				fmt.Println("Marked \"%d\" as completed. Error : %s\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
