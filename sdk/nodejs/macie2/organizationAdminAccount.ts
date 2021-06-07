// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an [Amazon Macie Organization Admin Account](https://docs.aws.amazon.com/macie/latest/APIReference/admin.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.macie2.Account("example", {});
 * const test = new aws.macie2.OrganizationAdminAccount("test", {adminAccountId: "ID OF THE ADMIN ACCOUNT"}, {
 *     dependsOn: [aws_macie2_account.test],
 * });
 * ```
 *
 * ## Import
 *
 * `aws_macie2_organization_admin_account` can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import aws:macie2/organizationAdminAccount:OrganizationAdminAccount example abcd1
 * ```
 */
export class OrganizationAdminAccount extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationAdminAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationAdminAccountState, opts?: pulumi.CustomResourceOptions): OrganizationAdminAccount {
        return new OrganizationAdminAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:macie2/organizationAdminAccount:OrganizationAdminAccount';

    /**
     * Returns true if the given object is an instance of OrganizationAdminAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationAdminAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationAdminAccount.__pulumiType;
    }

    /**
     * The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
     */
    public readonly adminAccountId!: pulumi.Output<string>;

    /**
     * Create a OrganizationAdminAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationAdminAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationAdminAccountArgs | OrganizationAdminAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationAdminAccountState | undefined;
            inputs["adminAccountId"] = state ? state.adminAccountId : undefined;
        } else {
            const args = argsOrState as OrganizationAdminAccountArgs | undefined;
            if ((!args || args.adminAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminAccountId'");
            }
            inputs["adminAccountId"] = args ? args.adminAccountId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationAdminAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationAdminAccount resources.
 */
export interface OrganizationAdminAccountState {
    /**
     * The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
     */
    adminAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationAdminAccount resource.
 */
export interface OrganizationAdminAccountArgs {
    /**
     * The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
     */
    adminAccountId: pulumi.Input<string>;
}