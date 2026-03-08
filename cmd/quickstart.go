package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var quickstartCmd = &cobra.Command{
	Use:   "quickstart",
	Short: "Show usage modes and getting started guide",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
Usage modes:

  1. Gateway (recommended — web-based setup at bot.lingti.com):
     lingti-bot gateway       # Opens bot.lingti.com/bots/xxxx to connect platforms

  2. MCP Server (for Claude Desktop / Cursor / Windsurf):
     Add to your MCP config (claude_desktop_config.json):

     {
       "mcpServers": {
         "lingti-bot": {
           "command": "/usr/local/bin/lingti-bot",
           "args": ["serve"]
         }
       }
     }


  3. Cloud Relay (connect to Lingti cloud for Feishu/Slack bots):
     lingti-bot relay         # Connect to cloud relay service

For more information:
  lingti-bot help                       # Show all commands
  lingti-bot <command> --help           # Help for specific command
  https://bot.lingti.com            # Documentation
  https://github.com/ruilisi/lingti-bot # Source code

`)
	},
}

func init() {
	rootCmd.AddCommand(quickstartCmd)
}
