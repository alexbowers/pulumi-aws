// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static class GetDefaultTags
    {
        public static Task<GetDefaultTagsResult> InvokeAsync(GetDefaultTagsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultTagsResult>("aws:index/getDefaultTags:getDefaultTags", args ?? new GetDefaultTagsArgs(), options.WithVersion());
    }


    public sealed class GetDefaultTagsArgs : Pulumi.InvokeArgs
    {
        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Blocks of default tags set on the provider. See details below.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetDefaultTagsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDefaultTagsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Blocks of default tags set on the provider. See details below.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetDefaultTagsResult(
            string id,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Tags = tags;
        }
    }
}
