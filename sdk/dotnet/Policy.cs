// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ
{
    /// <summary>
    /// The ``rabbitmq.Policy`` resource creates and manages policies for exchanges
    /// and queues.
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
    ///     var guest = new RabbitMQ.Permissions("guest", new()
    ///     {
    ///         User = "guest",
    ///         Vhost = test.Name,
    ///         PermissionDetails = new RabbitMQ.Inputs.PermissionsPermissionsArgs
    ///         {
    ///             Configure = ".*",
    ///             Write = ".*",
    ///             Read = ".*",
    ///         },
    ///     });
    /// 
    ///     var testPolicy = new RabbitMQ.Policy("test", new()
    ///     {
    ///         Name = "test",
    ///         Vhost = guest.Vhost,
    ///         PolicyBlock = new RabbitMQ.Inputs.PolicyPolicyArgs
    ///         {
    ///             Pattern = ".*",
    ///             Priority = 0,
    ///             ApplyTo = "all",
    ///             Definition = 
    ///             {
    ///                 { "ha-mode", "all" },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Policies can be imported using the `id` which is composed of `name@vhost`.
    /// 
    /// E.g.
    /// 
    /// ```sh
    /// $ pulumi import rabbitmq:index/policy:Policy test name@vhost
    /// ```
    /// </summary>
    [RabbitMQResourceType("rabbitmq:index/policy:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The settings of the policy. The structure is
        /// described below.
        /// </summary>
        [Output("policy")]
        public Output<Outputs.PolicyPolicy> PolicyBlock { get; private set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Output("vhost")]
        public Output<string> Vhost { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("rabbitmq:index/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("rabbitmq:index/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The settings of the policy. The structure is
        /// described below.
        /// </summary>
        [Input("policy", required: true)]
        public Input<Inputs.PolicyPolicyArgs> PolicyBlock { get; set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost", required: true)]
        public Input<string> Vhost { get; set; } = null!;

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }

    public sealed class PolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The settings of the policy. The structure is
        /// described below.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.PolicyPolicyGetArgs>? PolicyBlock { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public PolicyState()
        {
        }
        public static new PolicyState Empty => new PolicyState();
    }
}
