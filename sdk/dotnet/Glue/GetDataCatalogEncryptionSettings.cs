// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    public static class GetDataCatalogEncryptionSettings
    {
        /// <summary>
        /// This data source can be used to fetch information about AWS Glue Data Catalog Encryption Settings.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Glue.GetDataCatalogEncryptionSettings.InvokeAsync(new Aws.Glue.GetDataCatalogEncryptionSettingsArgs
        ///         {
        ///             Id = "123456789123",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDataCatalogEncryptionSettingsResult> InvokeAsync(GetDataCatalogEncryptionSettingsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataCatalogEncryptionSettingsResult>("aws:glue/getDataCatalogEncryptionSettings:getDataCatalogEncryptionSettings", args ?? new GetDataCatalogEncryptionSettingsArgs(), options.WithVersion());
    }


    public sealed class GetDataCatalogEncryptionSettingsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Data Catalog. This is typically the AWS account ID.
        /// </summary>
        [Input("catalogId", required: true)]
        public string CatalogId { get; set; } = null!;

        public GetDataCatalogEncryptionSettingsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataCatalogEncryptionSettingsResult
    {
        public readonly string CatalogId;
        /// <summary>
        /// The security configuration to set. see Data Catalog Encryption Settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataCatalogEncryptionSettingsDataCatalogEncryptionSettingResult> DataCatalogEncryptionSettings;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDataCatalogEncryptionSettingsResult(
            string catalogId,

            ImmutableArray<Outputs.GetDataCatalogEncryptionSettingsDataCatalogEncryptionSettingResult> dataCatalogEncryptionSettings,

            string id)
        {
            CatalogId = catalogId;
            DataCatalogEncryptionSettings = dataCatalogEncryptionSettings;
            Id = id;
        }
    }
}
