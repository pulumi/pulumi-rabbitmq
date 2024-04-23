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
    /// The ``rabbitmq.Permissions`` resource creates and manages a user's set of
    /// permissions.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using RabbitMQ = Pulumi.RabbitMQ;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new RabbitMQ.VHost("test", new()
    ///     {
    ///         Name = "test",
    ///     });
    /// 
    ///     var testUser = new RabbitMQ.User("test", new()
    ///     {
    ///         Name = "mctest",
    ///         Password = "foobar",
    ///         Tags = new[]
    ///         {
    ///             "administrator",
    ///         },
    ///     });
    /// 
    ///     var testPermissions = new RabbitMQ.Permissions("test", new()
    ///     {
    ///         User = testUser.Name,
    ///         Vhost = test.Name,
    ///         PermissionDetails = new RabbitMQ.Inputs.PermissionsPermissionsArgs
    ///         {
    ///             Configure = ".*",
    ///             Write = ".*",
    ///             Read = ".*",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Permissions can be imported using the `id` which is composed of  `user@vhost`.
    /// 
    /// E.g.
    /// 
    /// ```sh
    /// $ pulumi import rabbitmq:index/permissions:Permissions test user@vhost
    /// ```
    /// </summary>
    [RabbitMQResourceType("rabbitmq:index/permissions:Permissions")]
    public partial class Permissions : global::Pulumi.CustomResource
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
            : base("rabbitmq:index/permissions:Permissions", name, args ?? new PermissionsArgs(), MakeResourceOptions(options, ""))
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

    public sealed class PermissionsArgs : global::Pulumi.ResourceArgs
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
        public static new PermissionsArgs Empty => new PermissionsArgs();
    }

    public sealed class PermissionsState : global::Pulumi.ResourceArgs
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
        public static new PermissionsState Empty => new PermissionsState();
    }
}
