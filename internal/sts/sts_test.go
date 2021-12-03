package mysts2

import (
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

type mockSTSSvc struct {
	stsiface.STSAPI
}

func (m *mockSTSSvc) GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return &sts.GetCallerIdentityOutput{
		Account: aws.String("555555555555"),
		Arn:     aws.String("arn:aws:sts::555555555555:user/foo"),
		UserId:  aws.String("foo"),
	}, nil
}

func TestGetIdentity(t *testing.T) {
	svc := &mockSTSSvc{}
	i := Identity{}
	err := i.GetIdentity(svc)
	if err != nil {
		t.Error(err)
	}

	expectedAccount := "555555555555"
	expectedArn := "arn:aws:sts::555555555555:user/foo"
	expectedUserID := "foo"

	if !reflect.DeepEqual(expectedAccount, i.Account) {
		t.Errorf("expected %q to eq %q", expectedAccount, i.Account)
	}
	if !reflect.DeepEqual(expectedArn, i.Arn) {
		t.Errorf("expected %q to eq %q", expectedArn, i.Arn)
	}
	if !reflect.DeepEqual(expectedUserID, i.UserID) {
		t.Errorf("expected %q to eq %q", expectedUserID, i.UserID)
	}
}

func (m *mockSTSSvc) GetSessionToken(input *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
	return &sts.GetSessionTokenOutput{
		Credentials: &sts.Credentials{
			AccessKeyId:     aws.String("foo"),
			Expiration:      aws.Time(time.Now()),
			SecretAccessKey: aws.String("bar"),
			SessionToken:    aws.String("buzz"),
		},
	}, nil
}

func TestGetToken(t *testing.T) {
	svc := &mockSTSSvc{}
	i := Identity{}
	err := i.GetToken("111111", svc)
	if err != nil {
		t.Error(err)
	}

	expectedAccessKeyID := "foo"
	expectedSecretAccessKeyID := "bar"
	expectedSessionToken := "buzz"

	if !reflect.DeepEqual(expectedAccessKeyID, i.AccessKeyID) {
		t.Errorf("expected %q to eq %q", expectedAccessKeyID, i.AccessKeyID)
	}
	if !reflect.DeepEqual(expectedSecretAccessKeyID, i.SecretAccessKeyID) {
		t.Errorf("expected %q to eq %q", expectedSecretAccessKeyID, i.SecretAccessKeyID)
	}
	if !reflect.DeepEqual(expectedSessionToken, i.SessionToken) {
		t.Errorf("expected %q to eq %q", expectedSessionToken, i.SessionToken)
	}
}
