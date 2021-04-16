// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Get permissions for a principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, and tables. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
 *
 * > **NOTE:** This data source deals with explicitly granted permissions. Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).
 *
 * ## Example Usage
 * ### Permissions For A Glue Catalog Database
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.lakeformation.getPermissions({
 *     principal: aws_iam_role.workflow_role.arn,
 *     database: {
 *         name: aws_glue_catalog_database.test.name,
 *         catalogId: "110376042874",
 *     },
 * });
 * ```
 */
export function getPermissions(args: GetPermissionsArgs, opts?: pulumi.InvokeOptions): Promise<GetPermissionsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:lakeformation/getPermissions:getPermissions", {
        "catalogId": args.catalogId,
        "catalogResource": args.catalogResource,
        "dataLocation": args.dataLocation,
        "database": args.database,
        "principal": args.principal,
        "table": args.table,
        "tableWithColumns": args.tableWithColumns,
    }, opts);
}

/**
 * A collection of arguments for invoking getPermissions.
 */
export interface GetPermissionsArgs {
    /**
     * Identifier for the Data Catalog. By default, it is the account ID of the caller.
     */
    readonly catalogId?: string;
    /**
     * Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
     */
    readonly catalogResource?: boolean;
    /**
     * Configuration block for a data location resource. Detailed below.
     */
    readonly dataLocation?: inputs.lakeformation.GetPermissionsDataLocation;
    /**
     * Configuration block for a database resource. Detailed below.
     */
    readonly database?: inputs.lakeformation.GetPermissionsDatabase;
    /**
     * Principal to be granted the permissions on the resource. Supported principals are IAM users or IAM roles.
     */
    readonly principal: string;
    /**
     * Configuration block for a table resource. Detailed below.
     */
    readonly table?: inputs.lakeformation.GetPermissionsTable;
    /**
     * Configuration block for a table with columns resource. Detailed below.
     */
    readonly tableWithColumns?: inputs.lakeformation.GetPermissionsTableWithColumns;
}

/**
 * A collection of values returned by getPermissions.
 */
export interface GetPermissionsResult {
    readonly catalogId?: string;
    readonly catalogResource?: boolean;
    readonly dataLocation: outputs.lakeformation.GetPermissionsDataLocation;
    readonly database: outputs.lakeformation.GetPermissionsDatabase;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of permissions granted to the principal. For details on permissions, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     */
    readonly permissions: string[];
    /**
     * Subset of `permissions` which the principal can pass.
     */
    readonly permissionsWithGrantOptions: string[];
    readonly principal: string;
    readonly table: outputs.lakeformation.GetPermissionsTable;
    readonly tableWithColumns: outputs.lakeformation.GetPermissionsTableWithColumns;
}