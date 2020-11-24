// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq.Binding`` resource creates and manages a binding relationship
 * between a queue an exchange.
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
 * const testExchange = new rabbitmq.Exchange("test", {
 *     settings: {
 *         autoDelete: true,
 *         durable: false,
 *         type: "fanout",
 *     },
 *     vhost: guest.vhost,
 * });
 * const testQueue = new rabbitmq.Queue("test", {
 *     settings: {
 *         autoDelete: false,
 *         durable: true,
 *     },
 *     vhost: guest.vhost,
 * });
 * const testBinding = new rabbitmq.Binding("test", {
 *     destination: testQueue.name,
 *     destinationType: "queue",
 *     routingKey: "#",
 *     source: testExchange.name,
 *     vhost: testVHost.name,
 * });
 * ```
 *
 * ## Import
 *
 * Bindings can be imported using the `id` which is composed of
 *
 *  `vhost/source/destination/destination_type/properties_key`. E.g.
 *
 * ```sh
 *  $ pulumi import rabbitmq:index/binding:Binding test test/test/test/queue/%23
 * ```
 */
export class Binding extends pulumi.CustomResource {
    /**
     * Get an existing Binding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BindingState, opts?: pulumi.CustomResourceOptions): Binding {
        return new Binding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/binding:Binding';

    /**
     * Returns true if the given object is an instance of Binding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Binding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Binding.__pulumiType;
    }

    /**
     * Additional key/value arguments for the binding.
     */
    public readonly arguments!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly argumentsJson!: pulumi.Output<string | undefined>;
    /**
     * The destination queue or exchange.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * The type of destination (queue or exchange).
     */
    public readonly destinationType!: pulumi.Output<string>;
    /**
     * A unique key to refer to the binding.
     */
    public /*out*/ readonly propertiesKey!: pulumi.Output<string>;
    /**
     * A routing key for the binding.
     */
    public readonly routingKey!: pulumi.Output<string | undefined>;
    /**
     * The source exchange.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * The vhost to create the resource in.
     */
    public readonly vhost!: pulumi.Output<string>;

    /**
     * Create a Binding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BindingArgs | BindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BindingState | undefined;
            inputs["arguments"] = state ? state.arguments : undefined;
            inputs["argumentsJson"] = state ? state.argumentsJson : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["destinationType"] = state ? state.destinationType : undefined;
            inputs["propertiesKey"] = state ? state.propertiesKey : undefined;
            inputs["routingKey"] = state ? state.routingKey : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as BindingArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            if (!args || args.destinationType === undefined) {
                throw new Error("Missing required property 'destinationType'");
            }
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            if (!args || args.vhost === undefined) {
                throw new Error("Missing required property 'vhost'");
            }
            inputs["arguments"] = args ? args.arguments : undefined;
            inputs["argumentsJson"] = args ? args.argumentsJson : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["destinationType"] = args ? args.destinationType : undefined;
            inputs["routingKey"] = args ? args.routingKey : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["vhost"] = args ? args.vhost : undefined;
            inputs["propertiesKey"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Binding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Binding resources.
 */
export interface BindingState {
    /**
     * Additional key/value arguments for the binding.
     */
    readonly arguments?: pulumi.Input<{[key: string]: any}>;
    readonly argumentsJson?: pulumi.Input<string>;
    /**
     * The destination queue or exchange.
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * The type of destination (queue or exchange).
     */
    readonly destinationType?: pulumi.Input<string>;
    /**
     * A unique key to refer to the binding.
     */
    readonly propertiesKey?: pulumi.Input<string>;
    /**
     * A routing key for the binding.
     */
    readonly routingKey?: pulumi.Input<string>;
    /**
     * The source exchange.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Binding resource.
 */
export interface BindingArgs {
    /**
     * Additional key/value arguments for the binding.
     */
    readonly arguments?: pulumi.Input<{[key: string]: any}>;
    readonly argumentsJson?: pulumi.Input<string>;
    /**
     * The destination queue or exchange.
     */
    readonly destination: pulumi.Input<string>;
    /**
     * The type of destination (queue or exchange).
     */
    readonly destinationType: pulumi.Input<string>;
    /**
     * A routing key for the binding.
     */
    readonly routingKey?: pulumi.Input<string>;
    /**
     * The source exchange.
     */
    readonly source: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost: pulumi.Input<string>;
}
