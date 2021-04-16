// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Synthetics
{
    /// <summary>
    /// Provides a Synthetics Canary resource.
    /// 
    /// &gt; **NOTE:** When you create a canary, AWS creates supporting implicit resources. See the Amazon CloudWatch Synthetics documentation on [DeleteCanary](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DeleteCanary.html) for a full list. Neither AWS nor this provider deletes these implicit resources automatically when the canary is deleted. Before deleting a canary, ensure you have all the information about the canary that you need to delete the implicit resources using the AWS Console, or AWS CLI.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var some = new Aws.Synthetics.Canary("some", new Aws.Synthetics.CanaryArgs
    ///         {
    ///             ArtifactS3Location = "s3://some-bucket/",
    ///             ExecutionRoleArn = "some-role",
    ///             Handler = "exports.handler",
    ///             RuntimeVersion = "syn-1.0",
    ///             Schedule = new Aws.Synthetics.Inputs.CanaryScheduleArgs
    ///             {
    ///                 Expression = "rate(0 minute)",
    ///             },
    ///             ZipFile = "test-fixtures/lambdatest.zip",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Synthetics Canaries can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:synthetics/canary:Canary some some-canary
    /// ```
    /// </summary>
    [AwsResourceType("aws:synthetics/canary:Canary")]
    public partial class Canary : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Canary.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
        /// </summary>
        [Output("artifactS3Location")]
        public Output<string> ArtifactS3Location { get; private set; } = null!;

        /// <summary>
        /// ARN of the Lambda function that is used as your canary's engine.
        /// </summary>
        [Output("engineArn")]
        public Output<string> EngineArn { get; private set; } = null!;

        /// <summary>
        /// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
        /// </summary>
        [Output("executionRoleArn")]
        public Output<string> ExecutionRoleArn { get; private set; } = null!;

        /// <summary>
        /// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
        /// </summary>
        [Output("failureRetentionPeriod")]
        public Output<int?> FailureRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
        /// </summary>
        [Output("handler")]
        public Output<string> Handler { get; private set; } = null!;

        /// <summary>
        /// Name for this canary.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration block for individual canary runs. Detailed below.
        /// </summary>
        [Output("runConfig")]
        public Output<Outputs.CanaryRunConfig> RunConfig { get; private set; } = null!;

        /// <summary>
        /// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
        /// </summary>
        [Output("runtimeVersion")]
        public Output<string> RuntimeVersion { get; private set; } = null!;

        /// <summary>
        /// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zip_file`.**
        /// </summary>
        [Output("s3Bucket")]
        public Output<string?> S3Bucket { get; private set; } = null!;

        /// <summary>
        /// S3 key of your script. **Conflicts with `zip_file`.**
        /// </summary>
        [Output("s3Key")]
        public Output<string?> S3Key { get; private set; } = null!;

        /// <summary>
        /// S3 version ID of your script. **Conflicts with `zip_file`.**
        /// </summary>
        [Output("s3Version")]
        public Output<string?> S3Version { get; private set; } = null!;

        /// <summary>
        /// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.CanarySchedule> Schedule { get; private set; } = null!;

        /// <summary>
        /// ARN of the Lambda layer where Synthetics stores the canary script code.
        /// </summary>
        [Output("sourceLocationArn")]
        public Output<string> SourceLocationArn { get; private set; } = null!;

        /// <summary>
        /// Whether to run or stop the canary.
        /// </summary>
        [Output("startCanary")]
        public Output<bool?> StartCanary { get; private set; } = null!;

        /// <summary>
        /// Canary status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
        /// </summary>
        [Output("successRetentionPeriod")]
        public Output<int?> SuccessRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Structure that contains information about when the canary was created, modified, and most recently run. see Timeline.
        /// </summary>
        [Output("timelines")]
        public Output<ImmutableArray<Outputs.CanaryTimeline>> Timelines { get; private set; } = null!;

        /// <summary>
        /// Configuration block. Detailed below.
        /// </summary>
        [Output("vpcConfig")]
        public Output<Outputs.CanaryVpcConfig?> VpcConfig { get; private set; } = null!;

        /// <summary>
        /// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3_bucket`, `s3_key`, and `s3_version`.**
        /// </summary>
        [Output("zipFile")]
        public Output<string?> ZipFile { get; private set; } = null!;


        /// <summary>
        /// Create a Canary resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Canary(string name, CanaryArgs args, CustomResourceOptions? options = null)
            : base("aws:synthetics/canary:Canary", name, args ?? new CanaryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Canary(string name, Input<string> id, CanaryState? state = null, CustomResourceOptions? options = null)
            : base("aws:synthetics/canary:Canary", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Canary resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Canary Get(string name, Input<string> id, CanaryState? state = null, CustomResourceOptions? options = null)
        {
            return new Canary(name, id, state, options);
        }
    }

    public sealed class CanaryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
        /// </summary>
        [Input("artifactS3Location", required: true)]
        public Input<string> ArtifactS3Location { get; set; } = null!;

        /// <summary>
        /// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
        /// </summary>
        [Input("executionRoleArn", required: true)]
        public Input<string> ExecutionRoleArn { get; set; } = null!;

        /// <summary>
        /// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
        /// </summary>
        [Input("failureRetentionPeriod")]
        public Input<int>? FailureRetentionPeriod { get; set; }

        /// <summary>
        /// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
        /// </summary>
        [Input("handler", required: true)]
        public Input<string> Handler { get; set; } = null!;

        /// <summary>
        /// Name for this canary.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration block for individual canary runs. Detailed below.
        /// </summary>
        [Input("runConfig")]
        public Input<Inputs.CanaryRunConfigArgs>? RunConfig { get; set; }

        /// <summary>
        /// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
        /// </summary>
        [Input("runtimeVersion", required: true)]
        public Input<string> RuntimeVersion { get; set; } = null!;

        /// <summary>
        /// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zip_file`.**
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// S3 key of your script. **Conflicts with `zip_file`.**
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// S3 version ID of your script. **Conflicts with `zip_file`.**
        /// </summary>
        [Input("s3Version")]
        public Input<string>? S3Version { get; set; }

        /// <summary>
        /// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<Inputs.CanaryScheduleArgs> Schedule { get; set; } = null!;

        /// <summary>
        /// Whether to run or stop the canary.
        /// </summary>
        [Input("startCanary")]
        public Input<bool>? StartCanary { get; set; }

        /// <summary>
        /// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
        /// </summary>
        [Input("successRetentionPeriod")]
        public Input<int>? SuccessRetentionPeriod { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Configuration block. Detailed below.
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.CanaryVpcConfigArgs>? VpcConfig { get; set; }

        /// <summary>
        /// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3_bucket`, `s3_key`, and `s3_version`.**
        /// </summary>
        [Input("zipFile")]
        public Input<string>? ZipFile { get; set; }

        public CanaryArgs()
        {
        }
    }

    public sealed class CanaryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Canary.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
        /// </summary>
        [Input("artifactS3Location")]
        public Input<string>? ArtifactS3Location { get; set; }

        /// <summary>
        /// ARN of the Lambda function that is used as your canary's engine.
        /// </summary>
        [Input("engineArn")]
        public Input<string>? EngineArn { get; set; }

        /// <summary>
        /// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
        /// </summary>
        [Input("executionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        /// <summary>
        /// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
        /// </summary>
        [Input("failureRetentionPeriod")]
        public Input<int>? FailureRetentionPeriod { get; set; }

        /// <summary>
        /// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// Name for this canary.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration block for individual canary runs. Detailed below.
        /// </summary>
        [Input("runConfig")]
        public Input<Inputs.CanaryRunConfigGetArgs>? RunConfig { get; set; }

        /// <summary>
        /// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
        /// </summary>
        [Input("runtimeVersion")]
        public Input<string>? RuntimeVersion { get; set; }

        /// <summary>
        /// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zip_file`.**
        /// </summary>
        [Input("s3Bucket")]
        public Input<string>? S3Bucket { get; set; }

        /// <summary>
        /// S3 key of your script. **Conflicts with `zip_file`.**
        /// </summary>
        [Input("s3Key")]
        public Input<string>? S3Key { get; set; }

        /// <summary>
        /// S3 version ID of your script. **Conflicts with `zip_file`.**
        /// </summary>
        [Input("s3Version")]
        public Input<string>? S3Version { get; set; }

        /// <summary>
        /// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.CanaryScheduleGetArgs>? Schedule { get; set; }

        /// <summary>
        /// ARN of the Lambda layer where Synthetics stores the canary script code.
        /// </summary>
        [Input("sourceLocationArn")]
        public Input<string>? SourceLocationArn { get; set; }

        /// <summary>
        /// Whether to run or stop the canary.
        /// </summary>
        [Input("startCanary")]
        public Input<bool>? StartCanary { get; set; }

        /// <summary>
        /// Canary status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
        /// </summary>
        [Input("successRetentionPeriod")]
        public Input<int>? SuccessRetentionPeriod { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("timelines")]
        private InputList<Inputs.CanaryTimelineGetArgs>? _timelines;

        /// <summary>
        /// Structure that contains information about when the canary was created, modified, and most recently run. see Timeline.
        /// </summary>
        public InputList<Inputs.CanaryTimelineGetArgs> Timelines
        {
            get => _timelines ?? (_timelines = new InputList<Inputs.CanaryTimelineGetArgs>());
            set => _timelines = value;
        }

        /// <summary>
        /// Configuration block. Detailed below.
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.CanaryVpcConfigGetArgs>? VpcConfig { get; set; }

        /// <summary>
        /// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3_bucket`, `s3_key`, and `s3_version`.**
        /// </summary>
        [Input("zipFile")]
        public Input<string>? ZipFile { get; set; }

        public CanaryState()
        {
        }
    }
}