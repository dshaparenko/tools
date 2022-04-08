package cmd

import (
	"path/filepath"
	"strings"

	"github.com/devopsext/tools/common"
	"github.com/devopsext/tools/vendors"
	"github.com/devopsext/utils"
	"github.com/spf13/cobra"
)

var slackOptions = vendors.SlackOptions{
	URL:      envGet("SLACK_URL", "").(string),
	Timeout:  envGet("SLACK_TIMEOUT", 30).(int),
	Insecure: envGet("SLACK_INSECURE", false).(bool),
	Message:  envGet("SLACK_MESSAGE", "").(string),
	FileName: envGet("SLACK_FILENAME", "").(string),
	Title:    envGet("SLACK_TITLE", "").(string),
	Content:  envGet("SLACK_CONTENT", "").(string),
	Output:   envGet("SLACK_OUTPUT", "").(string),
	Query:    envGet("SLACK_QUERY", "").(string),
}

func slackNew(stdout *common.Stdout) common.Messenger {

	messageBytes, err := utils.Content(slackOptions.Message)
	if err != nil {
		stdout.Panic(err)
	}
	slackOptions.Message = string(messageBytes)

	contentBytes, err := utils.Content(slackOptions.Content)
	if err != nil {
		stdout.Panic(err)
	}
	slackOptions.Content = string(contentBytes)

	if utils.IsEmpty(slackOptions.FileName) && utils.FileExists(slackOptions.Content) {
		slackOptions.FileName = strings.TrimSuffix(slackOptions.Content, filepath.Ext(slackOptions.Content))
	}

	slack := vendors.NewSlack(slackOptions)
	if slack == nil {
		stdout.Panic("No slack")
	}
	return slack
}

func NewSlackCommand() *cobra.Command {

	slackCmd := &cobra.Command{
		Use:   "slack",
		Short: "Slack tools",
	}

	flags := slackCmd.PersistentFlags()
	flags.StringVar(&slackOptions.URL, "slack-url", slackOptions.URL, "Slack URL")
	flags.IntVar(&slackOptions.Timeout, "slack-timeout", slackOptions.Timeout, "Slack timeout")
	flags.BoolVar(&slackOptions.Insecure, "slack-insecure", slackOptions.Insecure, "Slack insecure")
	flags.StringVar(&slackOptions.Message, "slack-message", slackOptions.Message, "Slack message")
	flags.StringVar(&slackOptions.FileName, "slack-filename", slackOptions.FileName, "Slack file name")
	flags.StringVar(&slackOptions.Title, "slack-title", slackOptions.Title, "Slack title")
	flags.StringVar(&slackOptions.Content, "slack-content", slackOptions.Content, "Slack content")
	flags.StringVar(&slackOptions.Output, "slack-output", slackOptions.Output, "Slack output")
	flags.StringVar(&slackOptions.Query, "slack-query", slackOptions.Query, "Slack query")

	slackCmd.AddCommand(&cobra.Command{
		Use:   "send",
		Short: "Send text message",
		Run: func(cmd *cobra.Command, args []string) {

			stdout.Debug("Slack sending message...")
			bytes, err := slackNew(stdout).Send()
			if err != nil {
				stdout.Error(err)
				return
			}
			common.Output(slackOptions.Query, slackOptions.Output, bytes, stdout)
		},
	})

	slackCmd.AddCommand(&cobra.Command{
		Use:   "send-file",
		Short: "Send file",
		Run: func(cmd *cobra.Command, args []string) {

			stdout.Debug("Slack sending file...")
			bytes, err := slackNew(stdout).SendFile()
			if err != nil {
				stdout.Error(err)
				return
			}
			common.Output(slackOptions.Query, slackOptions.Output, bytes, stdout)
		},
	})
	return slackCmd
}