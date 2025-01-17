// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about a MQ Broker.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const brokerId = config.get("brokerId") || "";
 * const brokerName = config.get("brokerName") || "";
 * const byId = aws.mq.getBroker({
 *     brokerId: brokerId,
 * });
 * const byName = aws.mq.getBroker({
 *     brokerName: brokerName,
 * });
 * ```
 */
export function getBroker(args?: GetBrokerArgs, opts?: pulumi.InvokeOptions): Promise<GetBrokerResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:mq/getBroker:getBroker", {
        "brokerId": args.brokerId,
        "brokerName": args.brokerName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getBroker.
 */
export interface GetBrokerArgs {
    /**
     * The unique id of the mq broker.
     */
    brokerId?: string;
    /**
     * The unique name of the mq broker.
     */
    brokerName?: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getBroker.
 */
export interface GetBrokerResult {
    readonly arn: string;
    readonly authenticationStrategy: string;
    readonly autoMinorVersionUpgrade: boolean;
    readonly brokerId: string;
    readonly brokerName: string;
    readonly configuration: outputs.mq.GetBrokerConfiguration;
    readonly deploymentMode: string;
    readonly encryptionOptions: outputs.mq.GetBrokerEncryptionOption[];
    readonly engineType: string;
    readonly engineVersion: string;
    readonly hostInstanceType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instances: outputs.mq.GetBrokerInstance[];
    readonly ldapServerMetadatas: outputs.mq.GetBrokerLdapServerMetadata[];
    readonly logs: outputs.mq.GetBrokerLogs;
    readonly maintenanceWindowStartTime: outputs.mq.GetBrokerMaintenanceWindowStartTime;
    readonly publiclyAccessible: boolean;
    readonly securityGroups: string[];
    readonly storageType: string;
    readonly subnetIds: string[];
    readonly tags: {[key: string]: string};
    readonly users: outputs.mq.GetBrokerUser[];
}
