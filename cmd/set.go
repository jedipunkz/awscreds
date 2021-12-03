package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	mysts "github.com/jedipunkz/awscreds/internal/sts"
	"github.com/riywo/loginshell"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var setFlags struct {
	mfa     string
	region  string
	profile string
	shell   string
}

func getShell() (shell string) {
	if setFlags.shell == "" {
		loginshell, err := loginshell.Shell()
		if err != nil {
			log.Fatalln(err)
		}
		s := strings.Split(loginshell, "/")
		shell = s[len(s)-1] // last word is shell name
	} else {
		shell = setFlags.shell
	}
	return shell
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "setup aws creds",
	Long:  "setup aws credentials via MFA Number",
	Run: func(cmd *cobra.Command, args []string) {
		sess := session.Must(session.NewSessionWithOptions(session.Options{Profile: setFlags.profile}))
		svc := sts.New(sess)
		i := mysts.Identity{}

		// i := sts.NewIdentity(setFlags.profile)
		if err := i.GetIdentity(svc); err != nil {
			log.Fatalln(err)
		}
		if err := i.GetToken(setFlags.mfa, svc); err != nil {
			log.Fatalln(err)
		}

		shell := getShell()
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
			log.Fatalln("Your shell is invalid.:", shell)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&setFlags.mfa, "mfanum", "m", "", "mfa number")
	if err := setCmd.MarkFlagRequired("mfanum"); err != nil {
		log.Fatalln(err)
	}
	setCmd.Flags().StringVarP(&setFlags.region, "region", "r", "", "region name")
	if err := setCmd.MarkFlagRequired("region"); err != nil {
		log.Fatalln(err)
	}
	setCmd.Flags().StringVarP(&setFlags.profile, "profile", "p", "", "profile name")
	if err := setCmd.MarkFlagRequired("profile"); err != nil {
		log.Fatalln(err)
	}
	setCmd.Flags().StringVarP(&setFlags.shell, "shell", "s", "", "current shell")
}
