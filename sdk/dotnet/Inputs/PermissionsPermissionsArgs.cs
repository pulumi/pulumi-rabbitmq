// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.RabbitMQ.Inputs
{

    public sealed class PermissionsPermissionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The "configure" ACL.
        /// </summary>
        [Input("configure", required: true)]
        public Input<string> Configure { get; set; } = null!;

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

        public PermissionsPermissionsArgs()
        {
        }
    }
}
