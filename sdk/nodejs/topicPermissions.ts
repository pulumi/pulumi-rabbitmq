// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq..TopicPermissions`` resource creates and manages a user's set of
 * topic permissions.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 * 
 * const testVHost = new rabbitmq.VHost("test", {});
 * const testUser = new rabbitmq.User("test", {
 *     password: "foobar",
 *     tags: ["administrator"],
 * });
 * const testTopicPermissions = new rabbitmq.TopicPermissions("test", {
 *     permissions: [{
 *         exchange: "amq.topic",
 *         read: ".*",
 *         write: ".*",
 *     }],
 *     user: testUser.name,
 *     vhost: testVHost.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/topic-permissions.html.markdown.
 */
export class TopicPermissions extends pulumi.CustomResource {
    /**
     * Get an existing TopicPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TopicPermissionsState | undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["user"] = state ? state.user : undefined;
            inputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as TopicPermissionsArgs | undefined;
            if (!args || args.permissions === undefined) {
                throw new Error("Missing required property 'permissions'");
            }
            if (!args || args.user === undefined) {
                throw new Error("Missing required property 'user'");
            }
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["user"] = args ? args.user : undefined;
            inputs["vhost"] = args ? args.vhost : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TopicPermissions.__pulumiType, name, inputs, opts);
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
    readonly permissions?: pulumi.Input<pulumi.Input<inputs.TopicPermissionsPermission>[]>;
    /**
     * The user to apply the permissions to.
     */
    readonly user?: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TopicPermissions resource.
 */
export interface TopicPermissionsArgs {
    /**
     * The settings of the permissions. The structure is
     * described below.
     */
    readonly permissions: pulumi.Input<pulumi.Input<inputs.TopicPermissionsPermission>[]>;
    /**
     * The user to apply the permissions to.
     */
    readonly user: pulumi.Input<string>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost?: pulumi.Input<string>;
}
