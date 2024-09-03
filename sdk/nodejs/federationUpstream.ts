// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `rabbitmq.FederationUpstream` resource creates and manages a federation upstream parameter.
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
 * // downstream exchange
 * const foo = new rabbitmq.Exchange("foo", {
 *     name: "foo",
 *     vhost: guest.vhost,
 *     settings: {
 *         type: "topic",
 *         durable: true,
 *     },
 * });
 * // upstream broker
 * const fooFederationUpstream = new rabbitmq.FederationUpstream("foo", {
 *     name: "foo",
 *     vhost: guest.vhost,
 *     definition: {
 *         uri: "amqp://guest:guest@upstream-server-name:5672/%2f",
 *         prefetchCount: 1000,
 *         reconnectDelay: 5,
 *         ackMode: "on-confirm",
 *         trustUserId: false,
 *         maxHops: 1,
 *     },
 * });
 * const fooPolicy = new rabbitmq.Policy("foo", {
 *     name: "foo",
 *     vhost: guest.vhost,
 *     policy: {
 *         pattern: pulumi.interpolate`(^${foo.name}$)`,
 *         priority: 1,
 *         applyTo: "exchanges",
 *         definition: {
 *             "federation-upstream": fooFederationUpstream.name,
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * A Federation upstream can be imported using the resource `id` which is composed of `name@vhost`, e.g.
 *
 * ```sh
 * $ pulumi import rabbitmq:index/federationUpstream:FederationUpstream foo foo@test
 * ```
 */
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

    /**
     * Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
     */
    public /*out*/ readonly component!: pulumi.Output<string>;
    /**
     * The configuration of the federation upstream. The structure is described below.
     */
    public readonly definition!: pulumi.Output<outputs.FederationUpstreamDefinition>;
    /**
     * The name of the federation upstream.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The vhost to create the resource in.
     */
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FederationUpstreamState | undefined;
            resourceInputs["component"] = state ? state.component : undefined;
            resourceInputs["definition"] = state ? state.definition : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as FederationUpstreamArgs | undefined;
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            if ((!args || args.vhost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vhost'");
            }
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vhost"] = args ? args.vhost : undefined;
            resourceInputs["component"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FederationUpstream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FederationUpstream resources.
 */
export interface FederationUpstreamState {
    /**
     * Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
     */
    component?: pulumi.Input<string>;
    /**
     * The configuration of the federation upstream. The structure is described below.
     */
    definition?: pulumi.Input<inputs.FederationUpstreamDefinition>;
    /**
     * The name of the federation upstream.
     */
    name?: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FederationUpstream resource.
 */
export interface FederationUpstreamArgs {
    /**
     * The configuration of the federation upstream. The structure is described below.
     */
    definition: pulumi.Input<inputs.FederationUpstreamDefinition>;
    /**
     * The name of the federation upstream.
     */
    name?: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    vhost: pulumi.Input<string>;
}
