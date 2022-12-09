package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	qs = []*survey.Question{
		{
			Name: "username",
			Prompt: &survey.Input{
				Message: "What is your username?",
				Default: "admin",
			},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "What is your password?"},
			Validate: survey.Required,
		},
	}

	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Store your T-Mobile router creds so that other commands can work",
		RunE: func(cmd *cobra.Command, args []string) error {

			answers := struct {
				Username string
				Password string
			}{}

			err := survey.Ask(qs, &answers)
			if err != nil {
				return err
			}

			viper.Set("username", answers.Username)
			viper.Set("password", answers.Password)

			viper.WriteConfig()

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)
}
