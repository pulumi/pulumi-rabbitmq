// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ.Inputs
{

    public sealed class ExchangeSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("arguments")]
        private InputMap<object>? _arguments;

        /// <summary>
        /// Additional key/value settings for the exchange.
        /// </summary>
        public InputMap<object> Arguments
        {
            get => _arguments ?? (_arguments = new InputMap<object>());
            set => _arguments = value;
        }

        /// <summary>
        /// Whether the exchange will self-delete when all
        /// queues have finished using it.
        /// </summary>
        [Input("autoDelete")]
        public Input<bool>? AutoDelete { get; set; }

        /// <summary>
        /// Whether the exchange survives server restarts.
        /// Defaults to `false`.
        /// </summary>
        [Input("durable")]
        public Input<bool>? Durable { get; set; }

        /// <summary>
        /// The type of exchange.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ExchangeSettingsArgs()
        {
        }
    }
}
