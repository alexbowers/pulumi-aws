// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Workspaces
{
    /// <summary>
    /// Provides a workspace in [AWS Workspaces](https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces.html) Service
    /// 
    /// &gt; **NOTE:** During deletion of an `aws.workspaces.Workspace` resource, the service role `workspaces_DefaultRole` must be attached to the
    /// policy `arn:aws:iam::aws:policy/AmazonWorkSpacesServiceAccess`, or it will leak the ENI that the Workspaces service creates for the Workspace.
    /// </summary>
    public partial class Workspace : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the bundle for the WorkSpace.
        /// </summary>
        [Output("bundleId")]
        public Output<string> BundleId { get; private set; } = null!;

        /// <summary>
        /// The name of the WorkSpace, as seen by the operating system.
        /// </summary>
        [Output("computerName")]
        public Output<string> ComputerName { get; private set; } = null!;

        /// <summary>
        /// The ID of the directory for the WorkSpace.
        /// </summary>
        [Output("directoryId")]
        public Output<string> DirectoryId { get; private set; } = null!;

        /// <summary>
        /// The IP address of the WorkSpace.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the data stored on the root volume is encrypted.
        /// </summary>
        [Output("rootVolumeEncryptionEnabled")]
        public Output<bool?> RootVolumeEncryptionEnabled { get; private set; } = null!;

        /// <summary>
        /// The operational state of the WorkSpace.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The tags for the WorkSpace.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the data stored on the user volume is encrypted.
        /// </summary>
        [Output("userVolumeEncryptionEnabled")]
        public Output<bool?> UserVolumeEncryptionEnabled { get; private set; } = null!;

        /// <summary>
        /// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
        /// </summary>
        [Output("volumeEncryptionKey")]
        public Output<string?> VolumeEncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The WorkSpace properties.
        /// </summary>
        [Output("workspaceProperties")]
        public Output<Outputs.WorkspaceWorkspaceProperties> WorkspaceProperties { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs args, CustomResourceOptions? options = null)
            : base("aws:workspaces/workspace:Workspace", name, args ?? new WorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("aws:workspaces/workspace:Workspace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, state, options);
        }
    }

    public sealed class WorkspaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the bundle for the WorkSpace.
        /// </summary>
        [Input("bundleId", required: true)]
        public Input<string> BundleId { get; set; } = null!;

        /// <summary>
        /// The ID of the directory for the WorkSpace.
        /// </summary>
        [Input("directoryId", required: true)]
        public Input<string> DirectoryId { get; set; } = null!;

        /// <summary>
        /// Indicates whether the data stored on the root volume is encrypted.
        /// </summary>
        [Input("rootVolumeEncryptionEnabled")]
        public Input<bool>? RootVolumeEncryptionEnabled { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags for the WorkSpace.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        /// <summary>
        /// Indicates whether the data stored on the user volume is encrypted.
        /// </summary>
        [Input("userVolumeEncryptionEnabled")]
        public Input<bool>? UserVolumeEncryptionEnabled { get; set; }

        /// <summary>
        /// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
        /// </summary>
        [Input("volumeEncryptionKey")]
        public Input<string>? VolumeEncryptionKey { get; set; }

        /// <summary>
        /// The WorkSpace properties.
        /// </summary>
        [Input("workspaceProperties")]
        public Input<Inputs.WorkspaceWorkspacePropertiesArgs>? WorkspaceProperties { get; set; }

        public WorkspaceArgs()
        {
        }
    }

    public sealed class WorkspaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the bundle for the WorkSpace.
        /// </summary>
        [Input("bundleId")]
        public Input<string>? BundleId { get; set; }

        /// <summary>
        /// The name of the WorkSpace, as seen by the operating system.
        /// </summary>
        [Input("computerName")]
        public Input<string>? ComputerName { get; set; }

        /// <summary>
        /// The ID of the directory for the WorkSpace.
        /// </summary>
        [Input("directoryId")]
        public Input<string>? DirectoryId { get; set; }

        /// <summary>
        /// The IP address of the WorkSpace.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Indicates whether the data stored on the root volume is encrypted.
        /// </summary>
        [Input("rootVolumeEncryptionEnabled")]
        public Input<bool>? RootVolumeEncryptionEnabled { get; set; }

        /// <summary>
        /// The operational state of the WorkSpace.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags for the WorkSpace.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// Indicates whether the data stored on the user volume is encrypted.
        /// </summary>
        [Input("userVolumeEncryptionEnabled")]
        public Input<bool>? UserVolumeEncryptionEnabled { get; set; }

        /// <summary>
        /// The symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
        /// </summary>
        [Input("volumeEncryptionKey")]
        public Input<string>? VolumeEncryptionKey { get; set; }

        /// <summary>
        /// The WorkSpace properties.
        /// </summary>
        [Input("workspaceProperties")]
        public Input<Inputs.WorkspaceWorkspacePropertiesGetArgs>? WorkspaceProperties { get; set; }

        public WorkspaceState()
        {
        }
    }
}