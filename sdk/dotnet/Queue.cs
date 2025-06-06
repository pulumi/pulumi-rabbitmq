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
    /// The ``rabbitmq.Queue`` resource creates and manages a queue.
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic Example
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
    ///     var testQueue = new RabbitMQ.Queue("test", new()
    ///     {
    ///         Name = "test",
    ///         Vhost = guest.Vhost,
    ///         Settings = new RabbitMQ.Inputs.QueueSettingsArgs
    ///         {
    ///             Durable = false,
    ///             AutoDelete = true,
    ///             Arguments = 
    ///             {
    ///                 { "x-queue-type", "quorum" },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example With JSON Arguments
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using RabbitMQ = Pulumi.RabbitMQ;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var arguments = config.Get("arguments") ?? @"{
    ///   ""x-message-ttl"": 5000
    /// }
    /// ";
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
    ///     var testQueue = new RabbitMQ.Queue("test", new()
    ///     {
    ///         Name = "test",
    ///         Vhost = guest.Vhost,
    ///         Settings = new RabbitMQ.Inputs.QueueSettingsArgs
    ///         {
    ///             Durable = false,
    ///             AutoDelete = true,
    ///             ArgumentsJson = arguments,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Queues can be imported using the `id` which is composed of `name@vhost`. E.g.
    /// 
    /// ```sh
    /// $ pulumi import rabbitmq:index/queue:Queue test name@vhost
    /// ```
    /// </summary>
    [RabbitMQResourceType("rabbitmq:index/queue:Queue")]
    public partial class Queue : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the queue.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The settings of the queue. The structure is
        /// described below.
        /// </summary>
        [Output("settings")]
        public Output<Outputs.QueueSettings> Settings { get; private set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Output("vhost")]
        public Output<string?> Vhost { get; private set; } = null!;


        /// <summary>
        /// Create a Queue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Queue(string name, QueueArgs args, CustomResourceOptions? options = null)
            : base("rabbitmq:index/queue:Queue", name, args ?? new QueueArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Queue(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
            : base("rabbitmq:index/queue:Queue", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Queue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Queue Get(string name, Input<string> id, QueueState? state = null, CustomResourceOptions? options = null)
        {
            return new Queue(name, id, state, options);
        }
    }

    public sealed class QueueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the queue.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The settings of the queue. The structure is
        /// described below.
        /// </summary>
        [Input("settings", required: true)]
        public Input<Inputs.QueueSettingsArgs> Settings { get; set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public QueueArgs()
        {
        }
        public static new QueueArgs Empty => new QueueArgs();
    }

    public sealed class QueueState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the queue.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The settings of the queue. The structure is
        /// described below.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.QueueSettingsGetArgs>? Settings { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public QueueState()
        {
        }
        public static new QueueState Empty => new QueueState();
    }
}
