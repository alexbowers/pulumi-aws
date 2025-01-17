// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The CodeArtifact Repository Endpoint data source returns the endpoint of a repository for a specific package format.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.codeartifact.getRepositoryEndpoint({
 *     domain: aws_codeartifact_domain.test.domain,
 *     repository: aws_codeartifact_repository.test.repository,
 *     format: "npm",
 * });
 * ```
 */
export function getRepositoryEndpoint(args: GetRepositoryEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryEndpointResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:codeartifact/getRepositoryEndpoint:getRepositoryEndpoint", {
        "domain": args.domain,
        "domainOwner": args.domainOwner,
        "format": args.format,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryEndpoint.
 */
export interface GetRepositoryEndpointArgs {
    /**
     * The name of the domain that contains the repository.
     */
    domain: string;
    /**
     * The account number of the AWS account that owns the domain.
     */
    domainOwner?: string;
    /**
     * Which endpoint of a repository to return. A repository has one endpoint for each package format: `npm`, `pypi`, `maven`, and `nuget`.
     */
    format: string;
    /**
     * The name of the repository.
     */
    repository: string;
}

/**
 * A collection of values returned by getRepositoryEndpoint.
 */
export interface GetRepositoryEndpointResult {
    readonly domain: string;
    readonly domainOwner: string;
    readonly format: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly repository: string;
    /**
     * The URL of the returned endpoint.
     */
    readonly repositoryEndpoint: string;
}
