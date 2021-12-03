package mysts2

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

// Identity is struct for communicate sts identity
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

// NewIdentity is constractor for sts
// func NewIdentity(profilename string) *Identity {
// 	identity := new(Identity)
// 	identity.session = session.Must(session.NewSessionWithOptions(session.Options{Profile: profilename}))
// 	return identity
// }

// GetIdentity is function for getting identity
func (i *Identity) GetIdentity(svc stsiface.STSAPI) error {
	// svc := sts.New(i.session)
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

// GetToken is function for getting token
func (i *Identity) GetToken(mfanum string, svc stsiface.STSAPI) error {
	// svc := sts.New(i.session)
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(43200), // 12 hours
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
