// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq.Shovel`` resource creates and manages a dynamic shovel.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 *
 * const testVHost = new rabbitmq.VHost("test", {});
 * const testExchange = new rabbitmq.Exchange("test", {
 *     settings: {
 *         autoDelete: true,
 *         durable: false,
 *         type: "fanout",
 *     },
 *     vhost: testVHost.name,
 * });
 * const testQueue = new rabbitmq.Queue("test", {
 *     settings: {
 *         autoDelete: true,
 *         durable: false,
 *     },
 *     vhost: testVHost.name,
 * });
 * const shovelTest = new rabbitmq.Shovel("shovelTest", {
 *     info: {
 *         destinationQueue: testQueue.name,
 *         destinationUri: "amqp:///test",
 *         sourceExchange: testExchange.name,
 *         sourceExchangeKey: "test",
 *         sourceUri: "amqp:///test",
 *     },
 *     vhost: testVHost.name,
 * });
 * ```
 *
 * ## Import
 *
 * Shovels can be imported using the `name` and `vhost` E.g.
 *
 * ```sh
 *  $ pulumi import rabbitmq:index/shovel:Shovel test shovelTest@test
 * ```
 */
export class Shovel extends pulumi.CustomResource {
    /**
     * Get an existing Shovel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShovelState, opts?: pulumi.CustomResourceOptions): Shovel {
        return new Shovel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/shovel:Shovel';

    /**
     * Returns true if the given object is an instance of Shovel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Shovel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Shovel.__pulumiType;
    }

    /**
     * The settings of the dynamic shovel. The structure is
     * described below.
     */
    public readonly info!: pulumi.Output<outputs.ShovelInfo>;
    /**
     * The shovel name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The vhost to create the resource in.
     */
    public readonly vhost!: pulumi.Output<string>;

    /**
     * Create a Shovel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShovelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShovelArgs | ShovelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShovelState | undefined;
            inputs["info"] = state ? state.info : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as ShovelArgs | undefined;
            if ((!args || args.info === undefined) && !opts.urn) {
                throw new Error("Missing required property 'info'");
            }
            if ((!args || args.vhost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vhost'");
            }
            inputs["info"] = args ? args.info : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vhost"] = args ? args.vhost : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Shovel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Shovel resources.
 */
export interface ShovelState {
    /**
     * The settings of the dynamic shovel. The structure is
     * described below.
     */
    readonly info?: pulumi.Input<inputs.ShovelInfo>;
    /**
     * The shovel name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Shovel resource.
 */
export interface ShovelArgs {
    /**
     * The settings of the dynamic shovel. The structure is
     * described below.
     */
    readonly info: pulumi.Input<inputs.ShovelInfo>;
    /**
     * The shovel name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost: pulumi.Input<string>;
}
