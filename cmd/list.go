package cmd

func init() {
	rootCmd.AddCommand(Cli.GetTasks())
}
