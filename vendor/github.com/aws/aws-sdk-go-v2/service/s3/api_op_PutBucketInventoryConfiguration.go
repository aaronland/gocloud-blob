// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported for directory buckets.
//
// This implementation of the PUT action adds an S3 Inventory configuration
// (identified by the inventory ID) to the bucket. You can have up to 1,000
// inventory configurations per bucket.
//
// Amazon S3 inventory generates inventories of the objects in the bucket on a
// daily or weekly basis, and the results are published to a flat file. The bucket
// that is inventoried is called the source bucket, and the bucket where the
// inventory flat file is stored is called the destination bucket. The destination
// bucket must be in the same Amazon Web Services Region as the source bucket.
//
// When you configure an inventory for a source bucket, you specify the
// destination bucket where you want the inventory to be stored, and whether to
// generate the inventory daily or weekly. You can also configure what object
// metadata to include and whether to inventory all object versions or only current
// versions. For more information, see [Amazon S3 Inventory]in the Amazon S3 User Guide.
//
// You must create a bucket policy on the destination bucket to grant permissions
// to Amazon S3 to write objects to the bucket in the defined location. For an
// example policy, see [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis].
//
// Permissions To use this operation, you must have permission to perform the
// s3:PutInventoryConfiguration action. The bucket owner has this permission by
// default and can grant this permission to others.
//
// The s3:PutInventoryConfiguration permission allows a user to create an [S3 Inventory] report
// that includes all object metadata fields available and to specify the
// destination bucket to store the inventory. A user with read access to objects in
// the destination bucket can also access all object metadata fields that are
// available in the inventory report.
//
// To restrict access to an inventory report, see [Restricting access to an Amazon S3 Inventory report] in the Amazon S3 User Guide.
// For more information about the metadata fields available in S3 Inventory, see [Amazon S3 Inventory lists]
// in the Amazon S3 User Guide. For more information about permissions, see [Permissions related to bucket subresource operations]and [Identity and access management in Amazon S3]
// in the Amazon S3 User Guide.
//
// PutBucketInventoryConfiguration has the following special errors:
//
// HTTP 400 Bad Request Error  Code: InvalidArgument
//
// Cause: Invalid Argument
//
// HTTP 400 Bad Request Error  Code: TooManyConfigurations
//
// Cause: You are attempting to create a new configuration but have already
// reached the 1,000-configuration limit.
//
// HTTP 403 Forbidden Error  Cause: You are not the owner of the specified bucket,
// or you do not have the s3:PutInventoryConfiguration bucket permission to set
// the configuration on the bucket.
//
// The following operations are related to PutBucketInventoryConfiguration :
//
// [GetBucketInventoryConfiguration]
//
// [DeleteBucketInventoryConfiguration]
//
// [ListBucketInventoryConfigurations]
//
// [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html#example-bucket-policies-use-case-9
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [ListBucketInventoryConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html
// [S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-inventory.html
// [Permissions related to bucket subresource operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html
// [Identity and access management in Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Restricting access to an Amazon S3 Inventory report]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html#example-bucket-policies-use-case-10
// [Amazon S3 Inventory lists]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-inventory.html#storage-inventory-contents
// [GetBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html
func (c *Client) PutBucketInventoryConfiguration(ctx context.Context, params *PutBucketInventoryConfigurationInput, optFns ...func(*Options)) (*PutBucketInventoryConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketInventoryConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketInventoryConfiguration", params, optFns, c.addOperationPutBucketInventoryConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketInventoryConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketInventoryConfigurationInput struct {

	// The name of the bucket where the inventory configuration will be stored.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the inventory configuration.
	//
	// This member is required.
	Id *string

	// Specifies the inventory configuration.
	//
	// This member is required.
	InventoryConfiguration *types.InventoryConfiguration

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *PutBucketInventoryConfigurationInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type PutBucketInventoryConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketInventoryConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketInventoryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketInventoryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketInventoryConfiguration"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIsExpressUserAgent(stack); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpPutBucketInventoryConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketInventoryConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketInventoryConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func (v *PutBucketInventoryConfigurationInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketInventoryConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketInventoryConfiguration",
	}
}

// getPutBucketInventoryConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getPutBucketInventoryConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketInventoryConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketInventoryConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketInventoryConfigurationBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
