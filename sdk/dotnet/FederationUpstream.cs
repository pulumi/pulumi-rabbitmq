// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ
{
    [RabbitMQResourceType("rabbitmq:index/federationUpstream:FederationUpstream")]
    public partial class FederationUpstream : Pulumi.CustomResource
    {
        [Output("component")]
        public Output<string> Component { get; private set; } = null!;

        [Output("definition")]
        public Output<Outputs.FederationUpstreamDefinition> Definition { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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

    public sealed class FederationUpstreamArgs : Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<Inputs.FederationUpstreamDefinitionArgs> Definition { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("vhost", required: true)]
        public Input<string> Vhost { get; set; } = null!;

        public FederationUpstreamArgs()
        {
        }
    }

    public sealed class FederationUpstreamState : Pulumi.ResourceArgs
    {
        [Input("component")]
        public Input<string>? Component { get; set; }

        [Input("definition")]
        public Input<Inputs.FederationUpstreamDefinitionGetArgs>? Definition { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public FederationUpstreamState()
        {
        }
    }
}
