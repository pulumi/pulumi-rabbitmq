// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq.TopicPermissions`` resource creates and manages a user's set of
 * topic permissions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 *
 * const test = new rabbitmq.VHost("test", {name: "test"});
 * const testUser = new rabbitmq.User("test", {
 *     name: "mctest",
 *     password: "foobar",
 *     tags: ["administrator"],
 * });
 * const testTopicPermissions = new rabbitmq.TopicPermissions("test", {
 *     user: testUser.name,
 *     vhost: test.name,
 *     permissions: [{
 *         exchange: "amq.topic",
 *         write: ".*",
 *         read: ".*",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Permissions can be imported using the `id` which is composed of  `user@vhost`.
 *
 * E.g.
 *
 * ```sh
 * $ pulumi import rabbitmq:index/topicPermissions:TopicPermissions test user@vhost
 * ```
 */
export class TopicPermissions extends pulumi.CustomResource {
    /**
     * Get an existing TopicPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicPermissionsState, opts?: pulumi.CustomResourceOptions): TopicPermissions {
        return new TopicPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/topicPermissions:TopicPermissions';

    /**
     * Returns true if the given object is an instance of TopicPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicPermissions.__pulumiType;
    }

    /**
     * The settings of the permissions. The structure is
     * described below.
     */
    public readonly permissions!: pulumi.Output<outputs.TopicPermissionsPermission[]>;
    /**
     * The user to apply the permissions to.
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * The vhost to create the resource in.
     */
    public readonly vhost!: pulumi.Output<string | undefined>;

    /**
     * Create a TopicPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicPermissionsArgs | TopicPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicPermissionsState | undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as TopicPermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["vhost"] = args ? args.vhost : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TopicPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TopicPermissions resources.
 */
export interface TopicPermissionsState {
    /**
     * The settings of the permissions. The structure is
     * described below.
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.TopicPermissionsPermission>[]>;
    /**
     * The user to apply the permissions to.
     */
    user?: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TopicPermissions resource.
 */
export interface TopicPermissionsArgs {
    /**
     * The settings of the permissions. The structure is
     * described below.
     */
    permissions: pulumi.Input<pulumi.Input<inputs.TopicPermissionsPermission>[]>;
    /**
     * The user to apply the permissions to.
     */
    user: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    vhost?: pulumi.Input<string>;
}
