package cmd

// var (
// 	cfgFile   string
// 	logLevel  string
// 	logFormat string
// 	rootCmd   = &cobra.Command{
// 		Use:   "mycli",
// 		Short: "mycli is a fast, friendly CLI starter for Arch & Linux",
// 		Long:  "mycli is a batteries-included CLI scaffold using Cobra, Viper, slog, and XDG paths.",
// 		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
// 			// Bind CLI flags → viper (so env/config can override as well)
// 			_ = viper.BindPFlag("log.level", cmd.Flags().Lookup("log-level"))
// 			_ = viper.BindPFlag("log.format", cmd.Flags().Lookup("log-format"))

// 			// Initialize config (loads config file + env)
// 			if err := config.Init(cfgFile); err != nil {
// 				return err
// 			}
// 			// Initialize logger
// 			return logger.Init(viper.GetString("log.level"), viper.GetString("log.format"))
// 		},
// 	}
// )

// Execute runs the root command.
func Execute() error {
	return nil
	// return rootCmd.Execute()
}

// func init() {
// 	// Global flags
// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path (default: XDG config dir)")
// 	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level (debug, info, warn, error)")
// 	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "text", "log format (text, json)")

// 	// Env prefix: MYCLI_ (e.g., MYCLI_LOG_LEVEL)
// 	viper.SetEnvPrefix("MYCLI")
// 	viper.AutomaticEnv()

// 	// Add subcommands
// 	rootCmd.AddCommand(versionCmd)
// 	rootCmd.AddCommand(helloCmd)
// }
