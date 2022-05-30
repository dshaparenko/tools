package cmd

import (
	"github.com/devopsext/tools/common"
	"github.com/devopsext/tools/vendors"
	"github.com/spf13/cobra"
)

var gitlabOptions = vendors.GitlabOptions{
	Timeout:  envGet("GITLAB_TIMEOUT", 30).(int),
	Insecure: envGet("GITLAB_INSECURE", false).(bool),
	URL:      envGet("GITLAB_URL", "").(string),
	Token:    envGet("GITLAB_TOKEN", "").(string),
}

var gitlabOutput = common.OutputOptions{
	Output: envGet("GITLAB_OUTPUT", "").(string),
	Query:  envGet("GITLAB_OUTPUT_QUERY", "").(string),
}

var pipelineOptions = vendors.PipelineOptions{
	Project: envGet("GITLAB_PROJECT", "").(string),
	Ref:     envGet("GITLAB_PROJECT_REF", "").(string),
}

func NewGitlabCommand() *cobra.Command {
	gitlabCmd := &cobra.Command{
		Use:   "gitlab",
		Short: "Gitlab tools",
	}

	flags := gitlabCmd.PersistentFlags()
	flags.IntVar(&gitlabOptions.Timeout, "gitlab-timeout", gitlabOptions.Timeout, "Gitlab Timeout in seconds")
	flags.BoolVar(&gitlabOptions.Insecure, "gitlab-insecure", gitlabOptions.Insecure, "Gitlab Insecure")
	flags.StringVar(&gitlabOptions.URL, "gitlab-url", gitlabOptions.URL, "Gitlab URL")
	flags.StringVar(&gitlabOptions.Token, "gitlab-token", gitlabOptions.Token, "Gitlab Token")
	flags.StringVar(&gitlabOutput.Output, "gitlab-output", gitlabOutput.Output, "Gitlab Output")
	flags.StringVar(&gitlabOutput.Query, "gitlab-output-query", gitlabOutput.Query, "Gitlab Output Query")

	pipelineCmd := &cobra.Command{
		Use:   "pipeline",
		Short: "Get pipeline from Gitlab",
	}

	flags = pipelineCmd.PersistentFlags()

	flags.StringVar(&pipelineOptions.Project, "gitlab-project", pipelineOptions.Project, "Gitlab Project")
	flags.StringVar(&pipelineOptions.Ref, "gitlab-project-ref", pipelineOptions.Ref, "Gitlab Project Ref")

	pipelineCmd.AddCommand(&cobra.Command{
		Use:   "last",
		Short: "Get last successful gitlab pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			if pipelineOptions.Project == "" {
				stdout.Panic("No Gitlab project")
			}
			stdout.Debug("Getting pipelines…")
			bytes, err := vendors.NewGitlab(gitlabOptions).GetLastPipeline(pipelineOptions.Project, pipelineOptions.Ref)
			if err != nil {
				stdout.Error(err)
				return
			}
			common.OutputJson(gitlabOutput, "Gitlab", []interface{}{gitlabOptions}, bytes, stdout)
		},
	})

	pipelineCmd.AddCommand(&cobra.Command{
		Use:   "last-variables",
		Short: "Get last successful gitlab pipeline variables",
		Run: func(cmd *cobra.Command, args []string) {
			stdout.Debug("Getting pipelines…")
			bytes, err := vendors.NewGitlab(gitlabOptions).GetLastPipelineVariables(pipelineOptions.Project, pipelineOptions.Ref)
			if err != nil {
				stdout.Error(err)
				return
			}
			common.OutputJson(gitlabOutput, "Gitlab", []interface{}{gitlabOptions}, bytes, stdout)
		},
	})

	gitlabCmd.AddCommand(pipelineCmd)

	return gitlabCmd
}