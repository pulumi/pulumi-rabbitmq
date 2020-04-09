// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq..VHost`` resource creates and manages a vhost.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 * 
 * const myVhost = new rabbitmq.VHost("myVhost", {});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/vhost.html.markdown.
 */
export class VHost extends pulumi.CustomResource {
    /**
     * Get an existing VHost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VHostState, opts?: pulumi.CustomResourceOptions): VHost {
        return new VHost(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/vHost:VHost';

    /**
     * Returns true if the given object is an instance of VHost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VHost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VHost.__pulumiType;
    }

    /**
     * The name of the vhost.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a VHost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VHostArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VHostArgs | VHostState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VHostState | undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as VHostArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VHost.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VHost resources.
 */
export interface VHostState {
    /**
     * The name of the vhost.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VHost resource.
 */
export interface VHostArgs {
    /**
     * The name of the vhost.
     */
    readonly name?: pulumi.Input<string>;
}
