// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2.Outputs
{

    [OutputType]
    public sealed class StageAccessLogSettings
    {
        /// <summary>
        /// The ARN of the CloudWatch Logs log group to receive access logs. Any trailing `:*` is trimmed from the ARN.
        /// </summary>
        public readonly string DestinationArn;
        /// <summary>
        /// A single line [format](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#apigateway-cloudwatch-log-formats) of the access logs of data, as specified by [selected $context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-logging.html).
        /// </summary>
        public readonly string Format;

        [OutputConstructor]
        private StageAccessLogSettings(
            string destinationArn,

            string format)
        {
            DestinationArn = destinationArn;
            Format = format;
        }
    }
}