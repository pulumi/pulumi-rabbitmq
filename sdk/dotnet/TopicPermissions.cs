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
    /// The ``rabbitmq..TopicPermissions`` resource creates and manages a user's set of
    /// topic permissions.
    /// </summary>
    public partial class TopicPermissions : Pulumi.CustomResource
    {
        /// <summary>
        /// The settings of the permissions. The structure is
        /// described below.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.TopicPermissionsPermission>> Permissions { get; private set; } = null!;

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
        /// Create a TopicPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TopicPermissions(string name, TopicPermissionsArgs args, CustomResourceOptions? options = null)
            : base("rabbitmq:index/topicPermissions:TopicPermissions", name, args ?? new TopicPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TopicPermissions(string name, Input<string> id, TopicPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("rabbitmq:index/topicPermissions:TopicPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TopicPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TopicPermissions Get(string name, Input<string> id, TopicPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new TopicPermissions(name, id, state, options);
        }
    }

    public sealed class TopicPermissionsArgs : Pulumi.ResourceArgs
    {
        [Input("permissions", required: true)]
        private InputList<Inputs.TopicPermissionsPermissionArgs>? _permissions;

        /// <summary>
        /// The settings of the permissions. The structure is
        /// described below.
        /// </summary>
        public InputList<Inputs.TopicPermissionsPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.TopicPermissionsPermissionArgs>());
            set => _permissions = value;
        }

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

        public TopicPermissionsArgs()
        {
        }
    }

    public sealed class TopicPermissionsState : Pulumi.ResourceArgs
    {
        [Input("permissions")]
        private InputList<Inputs.TopicPermissionsPermissionGetArgs>? _permissions;

        /// <summary>
        /// The settings of the permissions. The structure is
        /// described below.
        /// </summary>
        public InputList<Inputs.TopicPermissionsPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.TopicPermissionsPermissionGetArgs>());
            set => _permissions = value;
        }

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

        public TopicPermissionsState()
        {
        }
    }
}
