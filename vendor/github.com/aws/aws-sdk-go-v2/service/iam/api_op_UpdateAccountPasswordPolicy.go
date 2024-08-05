// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the password policy settings for the Amazon Web Services account.
//
// This operation does not support partial updates. No parameters are required,
// but if you do not specify a parameter, that parameter's value reverts to its
// default value. See the Request Parameters section for each parameter's default
// value. Also note that some parameters do not allow the default parameter to be
// explicitly set. Instead, to invoke the default value, do not include that
// parameter when you invoke the operation.
//
// For more information about using a password policy, see [Managing an IAM password policy] in the IAM User Guide.
//
// [Managing an IAM password policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingPasswordPolicies.html
func (c *Client) UpdateAccountPasswordPolicy(ctx context.Context, params *UpdateAccountPasswordPolicyInput, optFns ...func(*Options)) (*UpdateAccountPasswordPolicyOutput, error) {
	if params == nil {
		params = &UpdateAccountPasswordPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountPasswordPolicy", params, optFns, c.addOperationUpdateAccountPasswordPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountPasswordPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountPasswordPolicyInput struct {

	//  Allows all IAM users in your account to use the Amazon Web Services Management
	// Console to change their own passwords. For more information, see [Permitting IAM users to change their own passwords]in the IAM
	// User Guide.
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of false . The result is that IAM users in the account do not
	// automatically have permissions to change their own password.
	//
	// [Permitting IAM users to change their own passwords]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_enable-user-change.html
	AllowUsersToChangePassword bool

	//  Prevents IAM users who are accessing the account via the Amazon Web Services
	// Management Console from setting a new console password after their password has
	// expired. The IAM user cannot access the console until an administrator resets
	// the password.
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of false . The result is that IAM users can change their passwords
	// after they expire and continue to sign in as the user.
	//
	// In the Amazon Web Services Management Console, the custom password policy
	// option Allow users to change their own password gives IAM users permissions to
	// iam:ChangePassword for only their user and to the iam:GetAccountPasswordPolicy
	// action. This option does not attach a permissions policy to each user, rather
	// the permissions are applied at the account-level for all users by IAM. IAM users
	// with iam:ChangePassword permission and active access keys can reset their own
	// expired console password using the CLI or API.
	HardExpiry *bool

	// The number of days that an IAM user password is valid.
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of 0 . The result is that IAM user passwords never expire.
	MaxPasswordAge *int32

	// The minimum number of characters allowed in an IAM user password.
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of 6 .
	MinimumPasswordLength *int32

	// Specifies the number of previous passwords that IAM users are prevented from
	// reusing.
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of 0 . The result is that IAM users are not prevented from reusing
	// previous passwords.
	PasswordReusePrevention *int32

	// Specifies whether IAM user passwords must contain at least one lowercase
	// character from the ISO basic Latin alphabet (a to z).
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of false . The result is that passwords do not require at least
	// one lowercase character.
	RequireLowercaseCharacters bool

	// Specifies whether IAM user passwords must contain at least one numeric
	// character (0 to 9).
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of false . The result is that passwords do not require at least
	// one numeric character.
	RequireNumbers bool

	// Specifies whether IAM user passwords must contain at least one of the following
	// non-alphanumeric characters:
	//
	// ! @ # $ % ^ & * ( ) _ + - = [ ] { } | '
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of false . The result is that passwords do not require at least
	// one symbol character.
	RequireSymbols bool

	// Specifies whether IAM user passwords must contain at least one uppercase
	// character from the ISO basic Latin alphabet (A to Z).
	//
	// If you do not specify a value for this parameter, then the operation uses the
	// default value of false . The result is that passwords do not require at least
	// one uppercase character.
	RequireUppercaseCharacters bool

	noSmithyDocumentSerde
}

type UpdateAccountPasswordPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccountPasswordPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateAccountPasswordPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateAccountPasswordPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAccountPasswordPolicy"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountPasswordPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateAccountPasswordPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAccountPasswordPolicy",
	}
}
