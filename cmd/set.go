package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jedipunkz/awscreds/pkg/sts"
	"github.com/riywo/loginshell"
	"github.com/spf13/cobra"
)

var setFlags struct {
	mfa     string
	region  string
	profile string
	shell   string
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "setup aws creds",
	Long:  "setup aws credentials via MFA Number",
	Run: func(cmd *cobra.Command, args []string) {
		i := sts.NewIdentity(setFlags.profile)
		err := i.GetCallerIdentity()
		if err != nil {
			log.Fatal(err)
		}
		err = i.GetSessionToken(setFlags.mfa)
		if err != nil {
			log.Fatal(err)
		}

		var shell string
		if setFlags.shell == "" {
			loginshell, err := loginshell.Shell()
			if err != nil {
				log.Fatal(err)
			}
			s := strings.Split(loginshell, "/")
			shell = s[len(s)-1] // last word is shell name
		} else {
			shell = setFlags.shell
		}

		switch shell {
		case "fish":
			fmt.Println("set -x AWS_PROFILE " + setFlags.profile)
			fmt.Println("set -x AWS_REGION " + setFlags.region)
			fmt.Println("set -x AWS_ACCESS_KEY_ID " + i.AccessKeyID)
			fmt.Println("set -x AWS_SECRET_ACCESS_KEY " + i.SecretAccessKeyID)
			fmt.Println("set -x AWS_SESSION_TOKEN " + i.SessionToken)
		case "zsh", "bash", "sh":
			fmt.Println("export AWS_PROFILE=" + setFlags.profile)
			fmt.Println("export AWS_REGION=" + setFlags.region)
			fmt.Println("export AWS_ACCESS_KEY_ID=" + i.AccessKeyID)
			fmt.Println("export AWS_SECRET_ACCESS_KEY=" + i.SecretAccessKeyID)
			fmt.Println("export AWS_SESSION_TOKEN=" + i.SessionToken)
		default:
			fmt.Printf("Your shell: %s is no valid.", shell)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&setFlags.mfa, "mfanum", "m", "", "mfa number")
	err := setCmd.MarkFlagRequired("mfanum")
	if err != nil {
		log.Fatal(err)
	}
	setCmd.Flags().StringVarP(&setFlags.region, "region", "r", "", "region name")
	err = setCmd.MarkFlagRequired("region")
	if err != nil {
		log.Fatal(err)
	}
	setCmd.Flags().StringVarP(&setFlags.profile, "profile", "p", "", "profile name")
	err = setCmd.MarkFlagRequired("profile")
	if err != nil {
		log.Fatal(err)
	}
	setCmd.Flags().StringVarP(&setFlags.shell, "shell", "s", "", "current shell")
}
