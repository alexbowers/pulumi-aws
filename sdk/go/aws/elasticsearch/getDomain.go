// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an Elasticsearch Domain
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticsearch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticsearch.LookupDomain(ctx, &elasticsearch.LookupDomainArgs{
// 			DomainName: "my-domain-name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("aws:elasticsearch/getDomain:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomain.
type LookupDomainArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// The tags assigned to the domain.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDomain.
type LookupDomainResult struct {
	// The policy document attached to the domain.
	AccessPolicies string `pulumi:"accessPolicies"`
	// Key-value string pairs to specify advanced configuration options.
	AdvancedOptions map[string]string `pulumi:"advancedOptions"`
	// Status of the Elasticsearch domain's advanced security options. The block consists of the following attributes:
	AdvancedSecurityOptions []GetDomainAdvancedSecurityOption `pulumi:"advancedSecurityOptions"`
	// The Amazon Resource Name (ARN) of the domain.
	Arn string `pulumi:"arn"`
	// Cluster configuration of the domain.
	ClusterConfigs []GetDomainClusterConfig `pulumi:"clusterConfigs"`
	// Domain Amazon Cognito Authentication options for Kibana.
	CognitoOptions []GetDomainCognitoOption `pulumi:"cognitoOptions"`
	// Status of the creation of the domain.
	Created bool `pulumi:"created"`
	// Status of the deletion of the domain.
	Deleted bool `pulumi:"deleted"`
	// Unique identifier for the domain.
	DomainId   string `pulumi:"domainId"`
	DomainName string `pulumi:"domainName"`
	// EBS Options for the instances in the domain.
	EbsOptions []GetDomainEbsOption `pulumi:"ebsOptions"`
	// ElasticSearch version for the domain.
	ElasticsearchVersion string `pulumi:"elasticsearchVersion"`
	// Domain encryption at rest related options.
	EncryptionAtRests []GetDomainEncryptionAtRest `pulumi:"encryptionAtRests"`
	// Domain-specific endpoint used to submit index, search, and data upload requests.
	Endpoint string `pulumi:"endpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Domain-specific endpoint used to access the Kibana application.
	KibanaEndpoint string `pulumi:"kibanaEndpoint"`
	// Domain log publishing related options.
	LogPublishingOptions []GetDomainLogPublishingOption `pulumi:"logPublishingOptions"`
	// Domain in transit encryption related options.
	NodeToNodeEncryptions []GetDomainNodeToNodeEncryption `pulumi:"nodeToNodeEncryptions"`
	// Status of a configuration change in the domain.
	// * `snapshotOptions` – Domain snapshot related options.
	Processing      bool                      `pulumi:"processing"`
	SnapshotOptions []GetDomainSnapshotOption `pulumi:"snapshotOptions"`
	// The tags assigned to the domain.
	Tags map[string]string `pulumi:"tags"`
	// VPC Options for private Elasticsearch domains.
	VpcOptions []GetDomainVpcOption `pulumi:"vpcOptions"`
}
