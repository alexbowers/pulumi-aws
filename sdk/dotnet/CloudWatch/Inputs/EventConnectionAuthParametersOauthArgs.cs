// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch.Inputs
{

    public sealed class EventConnectionAuthParametersOauthArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A username for the authorization.
        /// </summary>
        [Input("authorizationEndpoint", required: true)]
        public Input<string> AuthorizationEndpoint { get; set; } = null!;

        /// <summary>
        /// Contains the client parameters for OAuth authorization. Contains the following two parameters.
        /// </summary>
        [Input("clientParameters")]
        public Input<Inputs.EventConnectionAuthParametersOauthClientParametersArgs>? ClientParameters { get; set; }

        /// <summary>
        /// A password for the authorization. Created and stored in AWS Secrets Manager.
        /// </summary>
        [Input("httpMethod", required: true)]
        public Input<string> HttpMethod { get; set; } = null!;

        /// <summary>
        /// OAuth Http Parameters are additional credentials used to sign the request to the authorization endpoint to exchange the OAuth Client information for an access token. Secret values are stored and managed by AWS Secrets Manager. A maximum of 1 are allowed. Documented below.
        /// </summary>
        [Input("oauthHttpParameters", required: true)]
        public Input<Inputs.EventConnectionAuthParametersOauthOauthHttpParametersArgs> OauthHttpParameters { get; set; } = null!;

        public EventConnectionAuthParametersOauthArgs()
        {
        }
    }
}
