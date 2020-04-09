// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ
{
    /// <summary>
    /// The ``rabbitmq..Permissions`` resource creates and manages a user's set of
    /// permissions.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/permissions.html.markdown.
    /// </summary>
    public partial class Permissions : Pulumi.CustomResource
    {
        /// <summary>
        /// The settings of the permissions. The structure is
        /// described below.
        /// </summary>
        [Output("permissions")]
        public Output<Outputs.PermissionsPermissions> PermissionDetails { get; private set; } = null!;

        /// <summary>
        /// The user to apply the permissions to.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Output("vhost")]
        public Output<string?> Vhost { get; private set; } = null!;


        /// <summary>
        /// Create a Permissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Permissions(string name, PermissionsArgs args, CustomResourceOptions? options = null)
            : base("rabbitmq:index/permissions:Permissions", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Permissions(string name, Input<string> id, PermissionsState? state = null, CustomResourceOptions? options = null)
            : base("rabbitmq:index/permissions:Permissions", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Permissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Permissions Get(string name, Input<string> id, PermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new Permissions(name, id, state, options);
        }
    }

    public sealed class PermissionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The settings of the permissions. The structure is
        /// described below.
        /// </summary>
        [Input("permissions", required: true)]
        public Input<Inputs.PermissionsPermissionsArgs> PermissionDetails { get; set; } = null!;

        /// <summary>
        /// The user to apply the permissions to.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public PermissionsArgs()
        {
        }
    }

    public sealed class PermissionsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The settings of the permissions. The structure is
        /// described below.
        /// </summary>
        [Input("permissions")]
        public Input<Inputs.PermissionsPermissionsGetArgs>? PermissionDetails { get; set; }

        /// <summary>
        /// The user to apply the permissions to.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public PermissionsState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class PermissionsPermissionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The "configure" ACL.
        /// </summary>
        [Input("configure", required: true)]
        public Input<string> Configure { get; set; } = null!;

        /// <summary>
        /// The "read" ACL.
        /// </summary>
        [Input("read", required: true)]
        public Input<string> Read { get; set; } = null!;

        /// <summary>
        /// The "write" ACL.
        /// </summary>
        [Input("write", required: true)]
        public Input<string> Write { get; set; } = null!;

        public PermissionsPermissionsArgs()
        {
        }
    }

    public sealed class PermissionsPermissionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The "configure" ACL.
        /// </summary>
        [Input("configure", required: true)]
        public Input<string> Configure { get; set; } = null!;

        /// <summary>
        /// The "read" ACL.
        /// </summary>
        [Input("read", required: true)]
        public Input<string> Read { get; set; } = null!;

        /// <summary>
        /// The "write" ACL.
        /// </summary>
        [Input("write", required: true)]
        public Input<string> Write { get; set; } = null!;

        public PermissionsPermissionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class PermissionsPermissions
    {
        /// <summary>
        /// The "configure" ACL.
        /// </summary>
        public readonly string Configure;
        /// <summary>
        /// The "read" ACL.
        /// </summary>
        public readonly string Read;
        /// <summary>
        /// The "write" ACL.
        /// </summary>
        public readonly string Write;

        [OutputConstructor]
        private PermissionsPermissions(
            string configure,
            string read,
            string write)
        {
            Configure = configure;
            Read = read;
            Write = write;
        }
    }
    }
}
