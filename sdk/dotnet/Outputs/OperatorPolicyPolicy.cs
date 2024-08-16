// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ.Outputs
{

    [OutputType]
    public sealed class OperatorPolicyPolicy
    {
        /// <summary>
        /// Can be "queues".
        /// </summary>
        public readonly string ApplyTo;
        /// <summary>
        /// Key/value pairs of the operator policy definition. See the
        /// RabbitMQ documentation for definition references and examples.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Definition;
        /// <summary>
        /// A pattern to match an exchange or queue name.
        /// </summary>
        public readonly string Pattern;
        /// <summary>
        /// The policy with the greater priority is applied first.
        /// </summary>
        public readonly int Priority;

        [OutputConstructor]
        private OperatorPolicyPolicy(
            string applyTo,

            ImmutableDictionary<string, string> definition,

            string pattern,

            int priority)
        {
            ApplyTo = applyTo;
            Definition = definition;
            Pattern = pattern;
            Priority = priority;
        }
    }
}
