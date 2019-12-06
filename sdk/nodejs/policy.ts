// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq..Policy`` resource creates and manages policies for exchanges
 * and queues.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 * 
 * const testVHost = new rabbitmq.VHost("test", {});
 * const guest = new rabbitmq.Permissions("guest", {
 *     permissions: {
 *         configure: ".*",
 *         read: ".*",
 *         write: ".*",
 *     },
 *     user: "guest",
 *     vhost: testVHost.name,
 * });
 * const testPolicy = new rabbitmq.Policy("test", {
 *     policy: {
 *         applyTo: "all",
 *         definition: {
 *             "ha-mode": "all",
 *         },
 *         pattern: ".*",
 *         priority: 0,
 *     },
 *     vhost: guest.vhost,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/policy.html.markdown.
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * The name of the policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The settings of the policy. The structure is
     * described below.
     */
    public readonly policy!: pulumi.Output<outputs.PolicyPolicy>;
    /**
     * The vhost to create the resource in.
     */
    public readonly vhost!: pulumi.Output<string>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            if (!args || args.vhost === undefined) {
                throw new Error("Missing required property 'vhost'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["vhost"] = args ? args.vhost : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The settings of the policy. The structure is
     * described below.
     */
    readonly policy?: pulumi.Input<inputs.PolicyPolicy>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The settings of the policy. The structure is
     * described below.
     */
    readonly policy: pulumi.Input<inputs.PolicyPolicy>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost: pulumi.Input<string>;
}
