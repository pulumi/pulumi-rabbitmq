// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ.Inputs
{

    public sealed class FederationUpstreamDefinitionGetArgs : Pulumi.ResourceArgs
    {
        [Input("ackMode")]
        public Input<string>? AckMode { get; set; }

        [Input("exchange")]
        public Input<string>? Exchange { get; set; }

        [Input("expires")]
        public Input<int>? Expires { get; set; }

        [Input("maxHops")]
        public Input<int>? MaxHops { get; set; }

        [Input("messageTtl")]
        public Input<int>? MessageTtl { get; set; }

        [Input("prefetchCount")]
        public Input<int>? PrefetchCount { get; set; }

        [Input("queue")]
        public Input<string>? Queue { get; set; }

        [Input("reconnectDelay")]
        public Input<int>? ReconnectDelay { get; set; }

        [Input("trustUserId")]
        public Input<bool>? TrustUserId { get; set; }

        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public FederationUpstreamDefinitionGetArgs()
        {
        }
    }
}
