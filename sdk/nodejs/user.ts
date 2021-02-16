// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq.User`` resource creates and manages a user.
 *
 * > **Note:** All arguments including username and password will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 *
 * const test = new rabbitmq.User("test", {
 *     password: "foobar",
 *     tags: [
 *         "administrator",
 *         "management",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Users can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import rabbitmq:index/user:User test mctest
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The name of the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password of the user. The value of this argument
     * is plain-text so make sure to secure where this is defined.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Which permission model to apply to the user. Valid
     * options are: management, policymaker, monitoring, and administrator.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The name of the user.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password of the user. The value of this argument
     * is plain-text so make sure to secure where this is defined.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Which permission model to apply to the user. Valid
     * options are: management, policymaker, monitoring, and administrator.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The name of the user.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password of the user. The value of this argument
     * is plain-text so make sure to secure where this is defined.
     */
    readonly password: pulumi.Input<string>;
    /**
     * Which permission model to apply to the user. Valid
     * options are: management, policymaker, monitoring, and administrator.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
