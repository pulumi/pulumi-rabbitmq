// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class FederationUpstream extends pulumi.CustomResource {
    /**
     * Get an existing FederationUpstream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FederationUpstreamState, opts?: pulumi.CustomResourceOptions): FederationUpstream {
        return new FederationUpstream(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/federationUpstream:FederationUpstream';

    /**
     * Returns true if the given object is an instance of FederationUpstream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FederationUpstream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FederationUpstream.__pulumiType;
    }

    public /*out*/ readonly component!: pulumi.Output<string>;
    public readonly definition!: pulumi.Output<outputs.FederationUpstreamDefinition>;
    public readonly name!: pulumi.Output<string>;
    public readonly vhost!: pulumi.Output<string>;

    /**
     * Create a FederationUpstream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FederationUpstreamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FederationUpstreamArgs | FederationUpstreamState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FederationUpstreamState | undefined;
            inputs["component"] = state ? state.component : undefined;
            inputs["definition"] = state ? state.definition : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as FederationUpstreamArgs | undefined;
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            if ((!args || args.vhost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vhost'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vhost"] = args ? args.vhost : undefined;
            inputs["component"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FederationUpstream.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FederationUpstream resources.
 */
export interface FederationUpstreamState {
    readonly component?: pulumi.Input<string>;
    readonly definition?: pulumi.Input<inputs.FederationUpstreamDefinition>;
    readonly name?: pulumi.Input<string>;
    readonly vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FederationUpstream resource.
 */
export interface FederationUpstreamArgs {
    readonly definition: pulumi.Input<inputs.FederationUpstreamDefinition>;
    readonly name?: pulumi.Input<string>;
    readonly vhost: pulumi.Input<string>;
}
