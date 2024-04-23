// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq.Queue`` resource creates and manages a queue.
 *
 * ## Example Usage
 *
 * ### Basic Example
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
 * const testQueue = new rabbitmq.Queue("test", {
 *     name: "test",
 *     vhost: guest.vhost,
 *     settings: {
 *         durable: false,
 *         autoDelete: true,
 *         arguments: {
 *             "x-queue-type": "quorum",
 *         },
 *     },
 * });
 * ```
 *
 * ### Example With JSON Arguments
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 *
 * const config = new pulumi.Config();
 * const arguments = config.get("arguments") || `{
 *   "x-message-ttl": 5000
 * }
 * `;
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
 * const testQueue = new rabbitmq.Queue("test", {
 *     name: "test",
 *     vhost: guest.vhost,
 *     settings: {
 *         durable: false,
 *         autoDelete: true,
 *         argumentsJson: arguments,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Queues can be imported using the `id` which is composed of `name@vhost`. E.g.
 *
 * ```sh
 * $ pulumi import rabbitmq:index/queue:Queue test name@vhost
 * ```
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/queue:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * The name of the queue.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The settings of the queue. The structure is
     * described below.
     */
    public readonly settings!: pulumi.Output<outputs.QueueSettings>;
    /**
     * The vhost to create the resource in.
     */
    public readonly vhost!: pulumi.Output<string | undefined>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueArgs | QueueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QueueState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            if ((!args || args.settings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'settings'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["vhost"] = args ? args.vhost : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Queue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * The name of the queue.
     */
    name?: pulumi.Input<string>;
    /**
     * The settings of the queue. The structure is
     * described below.
     */
    settings?: pulumi.Input<inputs.QueueSettings>;
    /**
     * The vhost to create the resource in.
     */
    vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * The name of the queue.
     */
    name?: pulumi.Input<string>;
    /**
     * The settings of the queue. The structure is
     * described below.
     */
    settings: pulumi.Input<inputs.QueueSettings>;
    /**
     * The vhost to create the resource in.
     */
    vhost?: pulumi.Input<string>;
}
