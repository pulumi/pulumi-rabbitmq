// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ.Inputs
{

    public sealed class ShovelInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
        /// Defaults to `on-confirm`.
        /// </summary>
        [Input("ackMode")]
        public Input<string>? AckMode { get; set; }

        /// <summary>
        /// Whether to add `x-shovelled` headers to shovelled messages.
        /// </summary>
        [Input("addForwardHeaders")]
        public Input<bool>? AddForwardHeaders { get; set; }

        /// <summary>
        /// Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
        /// </summary>
        [Input("deleteAfter")]
        public Input<string>? DeleteAfter { get; set; }

        /// <summary>
        /// Whether to add `x-shovelled` headers to shovelled messages.
        /// </summary>
        [Input("destinationAddForwardHeaders")]
        public Input<bool>? DestinationAddForwardHeaders { get; set; }

        [Input("destinationAddTimestampHeader")]
        public Input<bool>? DestinationAddTimestampHeader { get; set; }

        /// <summary>
        /// The AMQP 1.0 destination link address.
        /// </summary>
        [Input("destinationAddress")]
        public Input<string>? DestinationAddress { get; set; }

        /// <summary>
        /// Application properties to set when shovelling messages.
        /// </summary>
        [Input("destinationApplicationProperties")]
        public Input<string>? DestinationApplicationProperties { get; set; }

        /// <summary>
        /// The exchange to which messages should be published.
        /// Either this or `destination_queue` must be specified but not both.
        /// </summary>
        [Input("destinationExchange")]
        public Input<string>? DestinationExchange { get; set; }

        /// <summary>
        /// The routing key when using `destination_exchange`.
        /// </summary>
        [Input("destinationExchangeKey")]
        public Input<string>? DestinationExchangeKey { get; set; }

        /// <summary>
        /// Properties to overwrite when shovelling messages.
        /// 
        /// For more details regarding dynamic shovel parameters please have a look at the official reference documentaion at [RabbitMQ: Configuring Dynamic Shovels](https://www.rabbitmq.com/shovel-dynamic.html).
        /// </summary>
        [Input("destinationProperties")]
        public Input<string>? DestinationProperties { get; set; }

        /// <summary>
        /// The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
        /// Defaults to `amqp091`.
        /// </summary>
        [Input("destinationProtocol")]
        public Input<string>? DestinationProtocol { get; set; }

        /// <summary>
        /// A map of properties to overwrite when shovelling messages.
        /// </summary>
        [Input("destinationPublishProperties")]
        public Input<string>? DestinationPublishProperties { get; set; }

        /// <summary>
        /// The queue to which messages should be published.
        /// Either this or `destination_exchange` must be specified but not both.
        /// </summary>
        [Input("destinationQueue")]
        public Input<string>? DestinationQueue { get; set; }

        /// <summary>
        /// The amqp uri for the destination .
        /// </summary>
        [Input("destinationUri", required: true)]
        public Input<string> DestinationUri { get; set; } = null!;

        /// <summary>
        /// The maximum number of unacknowledged messages copied over a shovel at any one time.
        /// </summary>
        [Input("prefetchCount")]
        public Input<int>? PrefetchCount { get; set; }

        /// <summary>
        /// The duration in seconds to reconnect to a broker after disconnected.
        /// Defaults to `1`.
        /// </summary>
        [Input("reconnectDelay")]
        public Input<int>? ReconnectDelay { get; set; }

        /// <summary>
        /// The AMQP 1.0 source link address.
        /// </summary>
        [Input("sourceAddress")]
        public Input<string>? SourceAddress { get; set; }

        /// <summary>
        /// Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
        /// </summary>
        [Input("sourceDeleteAfter")]
        public Input<string>? SourceDeleteAfter { get; set; }

        /// <summary>
        /// The exchange from which to consume.
        /// Either this or `source_queue` must be specified but not both.
        /// </summary>
        [Input("sourceExchange")]
        public Input<string>? SourceExchange { get; set; }

        /// <summary>
        /// The routing key when using `source_exchange`.
        /// </summary>
        [Input("sourceExchangeKey")]
        public Input<string>? SourceExchangeKey { get; set; }

        /// <summary>
        /// The maximum number of unacknowledged messages copied over a shovel at any one time.
        /// </summary>
        [Input("sourcePrefetchCount")]
        public Input<int>? SourcePrefetchCount { get; set; }

        /// <summary>
        /// The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
        /// Defaults to `amqp091`.
        /// </summary>
        [Input("sourceProtocol")]
        public Input<string>? SourceProtocol { get; set; }

        /// <summary>
        /// The queue from which to consume.
        /// Either this or `source_exchange` must be specified but not both.
        /// </summary>
        [Input("sourceQueue")]
        public Input<string>? SourceQueue { get; set; }

        /// <summary>
        /// The amqp uri for the source.
        /// </summary>
        [Input("sourceUri", required: true)]
        public Input<string> SourceUri { get; set; } = null!;

        public ShovelInfoGetArgs()
        {
        }
        public static new ShovelInfoGetArgs Empty => new ShovelInfoGetArgs();
    }
}
