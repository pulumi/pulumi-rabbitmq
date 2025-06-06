// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq.OperatorPolicy`` resource creates and manages operator policies for queues.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 *
 * const test = new rabbitmq.VHost("test", {name: "test"});
 * const guest = new rabbitmq.Permissions("guest", {
 *     user: "guest",
 *     vhost: test.name,
 *     permissions: {
 *         configure: ".*",
 *         write: ".*",
 *         read: ".*",
 *     },
 * });
 * const testOperatorPolicy = new rabbitmq.OperatorPolicy("test", {
 *     name: "test",
 *     vhost: guest.vhost,
 *     policy: {
 *         pattern: ".*",
 *         priority: 0,
 *         applyTo: "queues",
 *         definition: {
 *             "message-ttl": "3600000",
 *             expires: "1800000",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Operator policies can be imported using the `id` which is composed of `name@vhost`.
 *
 * E.g.
 *
 * ```sh
 * $ pulumi import rabbitmq:index/operatorPolicy:OperatorPolicy test name@vhost
 * ```
 */
export class OperatorPolicy extends pulumi.CustomResource {
    /**
     * Get an existing OperatorPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OperatorPolicyState, opts?: pulumi.CustomResourceOptions): OperatorPolicy {
        return new OperatorPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/operatorPolicy:OperatorPolicy';

    /**
     * Returns true if the given object is an instance of OperatorPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OperatorPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OperatorPolicy.__pulumiType;
    }

    /**
     * The name of the operator policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The settings of the operator policy. The structure is
     * described below.
     */
    public readonly policy!: pulumi.Output<outputs.OperatorPolicyPolicy>;
    /**
     * The vhost to create the resource in.
     */
    public readonly vhost!: pulumi.Output<string>;

    /**
     * Create a OperatorPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OperatorPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OperatorPolicyArgs | OperatorPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OperatorPolicyState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as OperatorPolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.vhost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vhost'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["vhost"] = args ? args.vhost : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OperatorPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OperatorPolicy resources.
 */
export interface OperatorPolicyState {
    /**
     * The name of the operator policy.
     */
    name?: pulumi.Input<string>;
    /**
     * The settings of the operator policy. The structure is
     * described below.
     */
    policy?: pulumi.Input<inputs.OperatorPolicyPolicy>;
    /**
     * The vhost to create the resource in.
     */
    vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OperatorPolicy resource.
 */
export interface OperatorPolicyArgs {
    /**
     * The name of the operator policy.
     */
    name?: pulumi.Input<string>;
    /**
     * The settings of the operator policy. The structure is
     * described below.
     */
    policy: pulumi.Input<inputs.OperatorPolicyPolicy>;
    /**
     * The vhost to create the resource in.
     */
    vhost: pulumi.Input<string>;
}
