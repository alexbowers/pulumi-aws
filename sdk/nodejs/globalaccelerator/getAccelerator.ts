// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about a Global Accelerator accelerator.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const acceleratorArn = config.get("acceleratorArn") || "";
 * const acceleratorName = config.get("acceleratorName") || "";
 * const example = aws.globalaccelerator.getAccelerator({
 *     arn: acceleratorArn,
 *     name: acceleratorName,
 * });
 * ```
 */
export function getAccelerator(args?: GetAcceleratorArgs, opts?: pulumi.InvokeOptions): Promise<GetAcceleratorResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:globalaccelerator/getAccelerator:getAccelerator", {
        "arn": args.arn,
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccelerator.
 */
export interface GetAcceleratorArgs {
    /**
     * The full ARN of the Global Accelerator.
     */
    arn?: string;
    /**
     * The unique name of the Global Accelerator.
     */
    name?: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getAccelerator.
 */
export interface GetAcceleratorResult {
    readonly arn: string;
    readonly attributes: outputs.globalaccelerator.GetAcceleratorAttribute[];
    readonly dnsName: string;
    readonly enabled: boolean;
    readonly hostedZoneId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddressType: string;
    readonly ipSets: outputs.globalaccelerator.GetAcceleratorIpSet[];
    readonly name: string;
    readonly tags: {[key: string]: string};
}
