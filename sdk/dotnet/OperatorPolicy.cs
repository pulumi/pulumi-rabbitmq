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
    /// The ``rabbitmq.OperatorPolicy`` resource creates and manages operator policies for queues.
    /// 
    /// ## Import
    /// 
    /// Operator policies can be imported using the `id` which is composed of `name@vhost`. E.g.
    /// 
    /// ```sh
    ///  $ pulumi import rabbitmq:index/operatorPolicy:OperatorPolicy test name@vhost
    /// ```
    /// </summary>
    [RabbitMQResourceType("rabbitmq:index/operatorPolicy:OperatorPolicy")]
    public partial class OperatorPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the operator policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The settings of the operator policy. The structure is
        /// described below.
        /// </summary>
        [Output("policy")]
        public Output<Outputs.OperatorPolicyPolicy> Policy { get; private set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Output("vhost")]
        public Output<string> Vhost { get; private set; } = null!;


        /// <summary>
        /// Create a OperatorPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OperatorPolicy(string name, OperatorPolicyArgs args, CustomResourceOptions? options = null)
            : base("rabbitmq:index/operatorPolicy:OperatorPolicy", name, args ?? new OperatorPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OperatorPolicy(string name, Input<string> id, OperatorPolicyState? state = null, CustomResourceOptions? options = null)
            : base("rabbitmq:index/operatorPolicy:OperatorPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OperatorPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OperatorPolicy Get(string name, Input<string> id, OperatorPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new OperatorPolicy(name, id, state, options);
        }
    }

    public sealed class OperatorPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the operator policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The settings of the operator policy. The structure is
        /// described below.
        /// </summary>
        [Input("policy", required: true)]
        public Input<Inputs.OperatorPolicyPolicyArgs> Policy { get; set; } = null!;

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost", required: true)]
        public Input<string> Vhost { get; set; } = null!;

        public OperatorPolicyArgs()
        {
        }
        public static new OperatorPolicyArgs Empty => new OperatorPolicyArgs();
    }

    public sealed class OperatorPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the operator policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The settings of the operator policy. The structure is
        /// described below.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.OperatorPolicyPolicyGetArgs>? Policy { get; set; }

        /// <summary>
        /// The vhost to create the resource in.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public OperatorPolicyState()
        {
        }
        public static new OperatorPolicyState Empty => new OperatorPolicyState();
    }
}
