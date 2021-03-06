// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sesiface provides an interface for the Amazon Simple Email Service.
package sesiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SESAPI is the interface type for ses.SES.
type SESAPI interface {
	DeleteIdentityRequest(*ses.DeleteIdentityInput) (*service.Request, *ses.DeleteIdentityOutput)

	DeleteIdentity(*ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error)

	DeleteIdentityPolicyRequest(*ses.DeleteIdentityPolicyInput) (*service.Request, *ses.DeleteIdentityPolicyOutput)

	DeleteIdentityPolicy(*ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error)

	DeleteVerifiedEmailAddressRequest(*ses.DeleteVerifiedEmailAddressInput) (*service.Request, *ses.DeleteVerifiedEmailAddressOutput)

	DeleteVerifiedEmailAddress(*ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error)

	GetIdentityDKIMAttributesRequest(*ses.GetIdentityDKIMAttributesInput) (*service.Request, *ses.GetIdentityDKIMAttributesOutput)

	GetIdentityDKIMAttributes(*ses.GetIdentityDKIMAttributesInput) (*ses.GetIdentityDKIMAttributesOutput, error)

	GetIdentityNotificationAttributesRequest(*ses.GetIdentityNotificationAttributesInput) (*service.Request, *ses.GetIdentityNotificationAttributesOutput)

	GetIdentityNotificationAttributes(*ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error)

	GetIdentityPoliciesRequest(*ses.GetIdentityPoliciesInput) (*service.Request, *ses.GetIdentityPoliciesOutput)

	GetIdentityPolicies(*ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error)

	GetIdentityVerificationAttributesRequest(*ses.GetIdentityVerificationAttributesInput) (*service.Request, *ses.GetIdentityVerificationAttributesOutput)

	GetIdentityVerificationAttributes(*ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error)

	GetSendQuotaRequest(*ses.GetSendQuotaInput) (*service.Request, *ses.GetSendQuotaOutput)

	GetSendQuota(*ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error)

	GetSendStatisticsRequest(*ses.GetSendStatisticsInput) (*service.Request, *ses.GetSendStatisticsOutput)

	GetSendStatistics(*ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error)

	ListIdentitiesRequest(*ses.ListIdentitiesInput) (*service.Request, *ses.ListIdentitiesOutput)

	ListIdentities(*ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error)

	ListIdentitiesPages(*ses.ListIdentitiesInput, func(*ses.ListIdentitiesOutput, bool) bool) error

	ListIdentityPoliciesRequest(*ses.ListIdentityPoliciesInput) (*service.Request, *ses.ListIdentityPoliciesOutput)

	ListIdentityPolicies(*ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error)

	ListVerifiedEmailAddressesRequest(*ses.ListVerifiedEmailAddressesInput) (*service.Request, *ses.ListVerifiedEmailAddressesOutput)

	ListVerifiedEmailAddresses(*ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error)

	PutIdentityPolicyRequest(*ses.PutIdentityPolicyInput) (*service.Request, *ses.PutIdentityPolicyOutput)

	PutIdentityPolicy(*ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error)

	SendEmailRequest(*ses.SendEmailInput) (*service.Request, *ses.SendEmailOutput)

	SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error)

	SendRawEmailRequest(*ses.SendRawEmailInput) (*service.Request, *ses.SendRawEmailOutput)

	SendRawEmail(*ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)

	SetIdentityDKIMEnabledRequest(*ses.SetIdentityDKIMEnabledInput) (*service.Request, *ses.SetIdentityDKIMEnabledOutput)

	SetIdentityDKIMEnabled(*ses.SetIdentityDKIMEnabledInput) (*ses.SetIdentityDKIMEnabledOutput, error)

	SetIdentityFeedbackForwardingEnabledRequest(*ses.SetIdentityFeedbackForwardingEnabledInput) (*service.Request, *ses.SetIdentityFeedbackForwardingEnabledOutput)

	SetIdentityFeedbackForwardingEnabled(*ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)

	SetIdentityNotificationTopicRequest(*ses.SetIdentityNotificationTopicInput) (*service.Request, *ses.SetIdentityNotificationTopicOutput)

	SetIdentityNotificationTopic(*ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error)

	VerifyDomainDKIMRequest(*ses.VerifyDomainDKIMInput) (*service.Request, *ses.VerifyDomainDKIMOutput)

	VerifyDomainDKIM(*ses.VerifyDomainDKIMInput) (*ses.VerifyDomainDKIMOutput, error)

	VerifyDomainIdentityRequest(*ses.VerifyDomainIdentityInput) (*service.Request, *ses.VerifyDomainIdentityOutput)

	VerifyDomainIdentity(*ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error)

	VerifyEmailAddressRequest(*ses.VerifyEmailAddressInput) (*service.Request, *ses.VerifyEmailAddressOutput)

	VerifyEmailAddress(*ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error)

	VerifyEmailIdentityRequest(*ses.VerifyEmailIdentityInput) (*service.Request, *ses.VerifyEmailIdentityOutput)

	VerifyEmailIdentity(*ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error)
}
