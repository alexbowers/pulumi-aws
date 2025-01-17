// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getDefaultTags(args?: GetDefaultTagsArgs, opts?: pulumi.InvokeOptions): Promise<GetDefaultTagsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getDefaultTags:getDefaultTags", {
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefaultTags.
 */
export interface GetDefaultTagsArgs {
    /**
     * Blocks of default tags set on the provider. See details below.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getDefaultTags.
 */
export interface GetDefaultTagsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Blocks of default tags set on the provider. See details below.
     */
    readonly tags: {[key: string]: string};
}
