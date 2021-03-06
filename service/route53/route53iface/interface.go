// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package route53iface provides an interface for the Amazon Route 53.
package route53iface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/route53"
)

// Route53API is the interface type for route53.Route53.
type Route53API interface {
	AssociateVPCWithHostedZoneRequest(*route53.AssociateVPCWithHostedZoneInput) (*service.Request, *route53.AssociateVPCWithHostedZoneOutput)

	AssociateVPCWithHostedZone(*route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error)

	ChangeResourceRecordSetsRequest(*route53.ChangeResourceRecordSetsInput) (*service.Request, *route53.ChangeResourceRecordSetsOutput)

	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)

	ChangeTagsForResourceRequest(*route53.ChangeTagsForResourceInput) (*service.Request, *route53.ChangeTagsForResourceOutput)

	ChangeTagsForResource(*route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error)

	CreateHealthCheckRequest(*route53.CreateHealthCheckInput) (*service.Request, *route53.CreateHealthCheckOutput)

	CreateHealthCheck(*route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error)

	CreateHostedZoneRequest(*route53.CreateHostedZoneInput) (*service.Request, *route53.CreateHostedZoneOutput)

	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)

	CreateReusableDelegationSetRequest(*route53.CreateReusableDelegationSetInput) (*service.Request, *route53.CreateReusableDelegationSetOutput)

	CreateReusableDelegationSet(*route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error)

	DeleteHealthCheckRequest(*route53.DeleteHealthCheckInput) (*service.Request, *route53.DeleteHealthCheckOutput)

	DeleteHealthCheck(*route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error)

	DeleteHostedZoneRequest(*route53.DeleteHostedZoneInput) (*service.Request, *route53.DeleteHostedZoneOutput)

	DeleteHostedZone(*route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)

	DeleteReusableDelegationSetRequest(*route53.DeleteReusableDelegationSetInput) (*service.Request, *route53.DeleteReusableDelegationSetOutput)

	DeleteReusableDelegationSet(*route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error)

	DisassociateVPCFromHostedZoneRequest(*route53.DisassociateVPCFromHostedZoneInput) (*service.Request, *route53.DisassociateVPCFromHostedZoneOutput)

	DisassociateVPCFromHostedZone(*route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error)

	GetChangeRequest(*route53.GetChangeInput) (*service.Request, *route53.GetChangeOutput)

	GetChange(*route53.GetChangeInput) (*route53.GetChangeOutput, error)

	GetCheckerIPRangesRequest(*route53.GetCheckerIPRangesInput) (*service.Request, *route53.GetCheckerIPRangesOutput)

	GetCheckerIPRanges(*route53.GetCheckerIPRangesInput) (*route53.GetCheckerIPRangesOutput, error)

	GetGeoLocationRequest(*route53.GetGeoLocationInput) (*service.Request, *route53.GetGeoLocationOutput)

	GetGeoLocation(*route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error)

	GetHealthCheckRequest(*route53.GetHealthCheckInput) (*service.Request, *route53.GetHealthCheckOutput)

	GetHealthCheck(*route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error)

	GetHealthCheckCountRequest(*route53.GetHealthCheckCountInput) (*service.Request, *route53.GetHealthCheckCountOutput)

	GetHealthCheckCount(*route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error)

	GetHealthCheckLastFailureReasonRequest(*route53.GetHealthCheckLastFailureReasonInput) (*service.Request, *route53.GetHealthCheckLastFailureReasonOutput)

	GetHealthCheckLastFailureReason(*route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error)

	GetHealthCheckStatusRequest(*route53.GetHealthCheckStatusInput) (*service.Request, *route53.GetHealthCheckStatusOutput)

	GetHealthCheckStatus(*route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error)

	GetHostedZoneRequest(*route53.GetHostedZoneInput) (*service.Request, *route53.GetHostedZoneOutput)

	GetHostedZone(*route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error)

	GetHostedZoneCountRequest(*route53.GetHostedZoneCountInput) (*service.Request, *route53.GetHostedZoneCountOutput)

	GetHostedZoneCount(*route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error)

	GetReusableDelegationSetRequest(*route53.GetReusableDelegationSetInput) (*service.Request, *route53.GetReusableDelegationSetOutput)

	GetReusableDelegationSet(*route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error)

	ListGeoLocationsRequest(*route53.ListGeoLocationsInput) (*service.Request, *route53.ListGeoLocationsOutput)

	ListGeoLocations(*route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error)

	ListHealthChecksRequest(*route53.ListHealthChecksInput) (*service.Request, *route53.ListHealthChecksOutput)

	ListHealthChecks(*route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error)

	ListHealthChecksPages(*route53.ListHealthChecksInput, func(*route53.ListHealthChecksOutput, bool) bool) error

	ListHostedZonesRequest(*route53.ListHostedZonesInput) (*service.Request, *route53.ListHostedZonesOutput)

	ListHostedZones(*route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error)

	ListHostedZonesPages(*route53.ListHostedZonesInput, func(*route53.ListHostedZonesOutput, bool) bool) error

	ListHostedZonesByNameRequest(*route53.ListHostedZonesByNameInput) (*service.Request, *route53.ListHostedZonesByNameOutput)

	ListHostedZonesByName(*route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)

	ListResourceRecordSetsRequest(*route53.ListResourceRecordSetsInput) (*service.Request, *route53.ListResourceRecordSetsOutput)

	ListResourceRecordSets(*route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error)

	ListResourceRecordSetsPages(*route53.ListResourceRecordSetsInput, func(*route53.ListResourceRecordSetsOutput, bool) bool) error

	ListReusableDelegationSetsRequest(*route53.ListReusableDelegationSetsInput) (*service.Request, *route53.ListReusableDelegationSetsOutput)

	ListReusableDelegationSets(*route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error)

	ListTagsForResourceRequest(*route53.ListTagsForResourceInput) (*service.Request, *route53.ListTagsForResourceOutput)

	ListTagsForResource(*route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)

	ListTagsForResourcesRequest(*route53.ListTagsForResourcesInput) (*service.Request, *route53.ListTagsForResourcesOutput)

	ListTagsForResources(*route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error)

	UpdateHealthCheckRequest(*route53.UpdateHealthCheckInput) (*service.Request, *route53.UpdateHealthCheckOutput)

	UpdateHealthCheck(*route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error)

	UpdateHostedZoneCommentRequest(*route53.UpdateHostedZoneCommentInput) (*service.Request, *route53.UpdateHostedZoneCommentOutput)

	UpdateHostedZoneComment(*route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error)
}
