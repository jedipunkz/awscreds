package cmd

import (
	"awscreds/pkg/sts"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var setFlags struct {
	mfa     string
	region  string
	profile string
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "setup aws creds",
	Long:  "setup aws credentials via MFA Number",
	Run: func(cmd *cobra.Command, args []string) {
		i := sts.NewIdentity()
		err := i.GetCallerIdentity()
		if err != nil {
			log.Fatal(err)
		}
		err = i.GetSessionToken(setFlags.mfa)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("set -x AWS_PROFILE " + setFlags.profile)
		fmt.Println("set -x AWS_REGION " + setFlags.region)
		fmt.Println("set -x AWS_ACCESS_KEY_ID " + i.AccessKeyID)
		fmt.Println("set -x AWS_SECRET_ACCESS_KEY " + i.SecretAccessKeyID)
		fmt.Println("set -x AWS_SESSION_TOKEN " + i.SessionToken)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&setFlags.mfa, "mfanum", "m", "", "mfanum: <MFA_NUM>")
	setCmd.MarkFlagRequired("mfanum")
	setCmd.Flags().StringVarP(&setFlags.region, "region", "r", "", "region: <region name>")
	setCmd.MarkFlagRequired("region")
	setCmd.Flags().StringVarP(&setFlags.profile, "profile", "p", "", "profile: <profile name>")
	setCmd.MarkFlagRequired("profile")
}
