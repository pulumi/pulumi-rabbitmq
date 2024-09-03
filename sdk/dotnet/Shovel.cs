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
    /// The `rabbitmq.Shovel` resource creates and manages a dynamic shovel.
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
    ///     var testExchange = new RabbitMQ.Exchange("test", new()
    ///     {
    ///         Name = "test_exchange",
    ///         Vhost = test.Name,
    ///         Settings = new RabbitMQ.Inputs.ExchangeSettingsArgs
    ///         {
    ///             Type = "fanout",
    ///             Durable = false,
    ///             AutoDelete = true,
    ///         },
    ///     });
    /// 
    ///     var testQueue = new RabbitMQ.Queue("test", new()
    ///     {
    ///         Name = "test_queue",
    ///         Vhost = test.Name,
    ///         Settings = new RabbitMQ.Inputs.QueueSettingsArgs
    ///         {
    ///             Durable = false,
    ///             AutoDelete = true,
    ///         },
    ///     });
    /// 
    ///     var shovelTest = new RabbitMQ.Shovel("shovelTest", new()
    ///     {
    ///         Name = "shovelTest",
    ///         Vhost = test.Name,
    ///         Info = new RabbitMQ.Inputs.ShovelInfoArgs
    ///         {
    ///             SourceUri = "amqp:///test",
    ///             SourceExchange = testExchange.Name,
    ///             SourceExchangeKey = "test",
    ///             DestinationUri = "amqp:///test",
    ///             DestinationQueue = testQueue.Name,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Shovels can be imported using the `name` and `vhost`
    /// 
    /// E.g.
    /// 
    /// ```sh
    /// $ pulumi import rabbitmq:index/shovel:Shovel test shovelTest@test
    /// ```
    /// </summary>
    [RabbitMQResourceType("rabbitmq:index/shovel:Shovel")]
    public partial class Shovel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The settings of the dynamic shovel. The structure is
        /// described below.
        /// </summary>
        [Output("info")]
        public Output<Outputs.ShovelInfo> Info { get; private set; } = null!;

        /// <summary>
        /// The shovel name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Output("vhost")]
        public Output<string> Vhost { get; private set; } = null!;


        /// <summary>
        /// Create a Shovel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Shovel(string name, ShovelArgs args, CustomResourceOptions? options = null)
            : base("rabbitmq:index/shovel:Shovel", name, args ?? new ShovelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Shovel(string name, Input<string> id, ShovelState? state = null, CustomResourceOptions? options = null)
            : base("rabbitmq:index/shovel:Shovel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Shovel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Shovel Get(string name, Input<string> id, ShovelState? state = null, CustomResourceOptions? options = null)
        {
            return new Shovel(name, id, state, options);
        }
    }

    public sealed class ShovelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The settings of the dynamic shovel. The structure is
        /// described below.
        /// </summary>
        [Input("info", required: true)]
        public Input<Inputs.ShovelInfoArgs> Info { get; set; } = null!;

        /// <summary>
        /// The shovel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost", required: true)]
        public Input<string> Vhost { get; set; } = null!;

        public ShovelArgs()
        {
        }
        public static new ShovelArgs Empty => new ShovelArgs();
    }

    public sealed class ShovelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The settings of the dynamic shovel. The structure is
        /// described below.
        /// </summary>
        [Input("info")]
        public Input<Inputs.ShovelInfoGetArgs>? Info { get; set; }

        /// <summary>
        /// The shovel name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public ShovelState()
        {
        }
        public static new ShovelState Empty => new ShovelState();
    }
}
