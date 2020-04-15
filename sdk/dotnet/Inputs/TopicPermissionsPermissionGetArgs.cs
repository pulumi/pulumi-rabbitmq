// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ.Inputs
{

    public sealed class TopicPermissionsPermissionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The exchange to set the permissions for.
        /// </summary>
        [Input("exchange", required: true)]
        public Input<string> Exchange { get; set; } = null!;

        /// <summary>
        /// The "read" ACL.
        /// </summary>
        [Input("read", required: true)]
        public Input<string> Read { get; set; } = null!;

        /// <summary>
        /// The "write" ACL.
        /// </summary>
        [Input("write", required: true)]
        public Input<string> Write { get; set; } = null!;

        public TopicPermissionsPermissionGetArgs()
        {
        }
    }
}
