package sts

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

const (
	profilename = "sandbox"
	region      = "ap-northeast-1"
)

// Identity is
type Identity struct {
	session           *session.Session
	Account           string
	Arn               string
	UserName          string
	UserID            string
	SerialNumber      string
	AccessKeyID       string
	SecretAccessKeyID string
	SessionToken      string
}

// NewIdentity is constractor
func NewIdentity() *Identity {
	identity := new(Identity)
	identity.session = session.Must(session.NewSessionWithOptions(session.Options{Profile: profilename}))
	return identity
}

// GetCallerIdentity is
func (i *Identity) GetCallerIdentity() error {
	svc := sts.New(i.session)
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		return err
	}

	i.Account = *result.Account
	i.Arn = *result.Arn
	username := strings.Split(i.Arn, "/")
	i.UserName = username[len(username)-1]
	i.UserID = *result.UserId
	i.SerialNumber = "arn:aws:iam::" + i.Account + ":mfa/" + i.UserName
	return nil
}

// GetSessionToken is
func (i *Identity) GetSessionToken(mfanum string) error {
	svc := sts.New(i.session)
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(3600),
		SerialNumber:    aws.String(i.SerialNumber),
		TokenCode:       aws.String(mfanum),
	}
	result, err := svc.GetSessionToken(input)
	if err != nil {
		return err
	}

	i.AccessKeyID = *result.Credentials.AccessKeyId
	i.SecretAccessKeyID = *result.Credentials.SecretAccessKey
	i.SessionToken = *result.Credentials.SessionToken
	return nil
}

// func main() {
// 	identity := NewIdentity()
// 	identity.getCallerIdentity()
// 	identity.getSessionToken()
//
// 	fmt.Println("set -x AWS_PROFILE " + profilename)
// 	fmt.Println("set -x AWS_REGION " + region)
// 	fmt.Println("set -x AWS_ACCESS_KEY_ID " + identity.AccessKeyID)
// 	fmt.Println("set -x AWS_SECRET_ACCESS_KEY " + identity.SecretAccessKeyID)
// 	fmt.Println("set -x AWS_SESSION_TOKEN " + identity.SessionToken)
// }
