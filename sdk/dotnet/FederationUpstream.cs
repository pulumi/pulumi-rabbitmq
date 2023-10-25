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
    /// The ``rabbitmq.FederationUpstream`` resource creates and manages a federation upstream parameter.
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
    ///     var test = new RabbitMQ.VHost("test");
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
    ///     // downstream exchange
    ///     var fooExchange = new RabbitMQ.Exchange("fooExchange", new()
    ///     {
    ///         Vhost = guest.Vhost,
    ///         Settings = new RabbitMQ.Inputs.ExchangeSettingsArgs
    ///         {
    ///             Type = "topic",
    ///             Durable = true,
    ///         },
    ///     });
    /// 
    ///     // upstream broker
    ///     var fooFederationUpstream = new RabbitMQ.FederationUpstream("fooFederationUpstream", new()
    ///     {
    ///         Vhost = guest.Vhost,
    ///         Definition = new RabbitMQ.Inputs.FederationUpstreamDefinitionArgs
    ///         {
    ///             Uri = "amqp://guest:guest@upstream-server-name:5672/%2f",
    ///             PrefetchCount = 1000,
    ///             ReconnectDelay = 5,
    ///             AckMode = "on-confirm",
    ///             TrustUserId = false,
    ///             MaxHops = 1,
    ///         },
    ///     });
    /// 
    ///     var fooPolicy = new RabbitMQ.Policy("fooPolicy", new()
    ///     {
    ///         Vhost = guest.Vhost,
    ///         PolicyBlock = new RabbitMQ.Inputs.PolicyPolicyArgs
    ///         {
    ///             Pattern = fooExchange.Name.Apply(name =&gt; $"(^{name}$)"),
    ///             Priority = 1,
    ///             ApplyTo = "exchanges",
    ///             Definition = 
    ///             {
    ///                 { "federation-upstream", fooFederationUpstream.Name },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A Federation upstream can be imported using the resource `id` which is composed of `name@vhost`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import rabbitmq:index/federationUpstream:FederationUpstream foo foo@test
    /// ```
    /// </summary>
    [RabbitMQResourceType("rabbitmq:index/federationUpstream:FederationUpstream")]
    public partial class FederationUpstream : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
        /// </summary>
        [Output("component")]
        public Output<string> Component { get; private set; } = null!;

        /// <summary>
        /// The configuration of the federation upstream. The structure is described below.
        /// </summary>
        [Output("definition")]
        public Output<Outputs.FederationUpstreamDefinition> Definition { get; private set; } = null!;

        /// <summary>
        /// The name of the federation upstream.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Output("vhost")]
        public Output<string> Vhost { get; private set; } = null!;


        /// <summary>
        /// Create a FederationUpstream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FederationUpstream(string name, FederationUpstreamArgs args, CustomResourceOptions? options = null)
            : base("rabbitmq:index/federationUpstream:FederationUpstream", name, args ?? new FederationUpstreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FederationUpstream(string name, Input<string> id, FederationUpstreamState? state = null, CustomResourceOptions? options = null)
            : base("rabbitmq:index/federationUpstream:FederationUpstream", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FederationUpstream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FederationUpstream Get(string name, Input<string> id, FederationUpstreamState? state = null, CustomResourceOptions? options = null)
        {
            return new FederationUpstream(name, id, state, options);
        }
    }

    public sealed class FederationUpstreamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration of the federation upstream. The structure is described below.
        /// </summary>
        [Input("definition", required: true)]
        public Input<Inputs.FederationUpstreamDefinitionArgs> Definition { get; set; } = null!;

        /// <summary>
        /// The name of the federation upstream.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost", required: true)]
        public Input<string> Vhost { get; set; } = null!;

        public FederationUpstreamArgs()
        {
        }
        public static new FederationUpstreamArgs Empty => new FederationUpstreamArgs();
    }

    public sealed class FederationUpstreamState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
        /// </summary>
        [Input("component")]
        public Input<string>? Component { get; set; }

        /// <summary>
        /// The configuration of the federation upstream. The structure is described below.
        /// </summary>
        [Input("definition")]
        public Input<Inputs.FederationUpstreamDefinitionGetArgs>? Definition { get; set; }

        /// <summary>
        /// The name of the federation upstream.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public FederationUpstreamState()
        {
        }
        public static new FederationUpstreamState Empty => new FederationUpstreamState();
    }
}
