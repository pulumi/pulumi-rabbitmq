// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ExchangeSettings {
    /**
     * Additional key/value settings for the exchange.
     */
    arguments?: {[key: string]: any};
    /**
     * Whether the exchange will self-delete when all
     * queues have finished using it.
     */
    autoDelete?: boolean;
    /**
     * Whether the exchange survives server restarts.
     * Defaults to `false`.
     */
    durable?: boolean;
    /**
     * The type of exchange.
     */
    type: string;
}

export interface PermissionsPermissions {
    /**
     * The "configure" ACL.
     */
    configure: string;
    /**
     * The "read" ACL.
     */
    read: string;
    /**
     * The "write" ACL.
     */
    write: string;
}

export interface PolicyPolicy {
    /**
     * Can either be "exchanges", "queues", or "all".
     */
    applyTo: string;
    /**
     * Key/value pairs of the policy definition. See the
     * RabbitMQ documentation for definition references and examples.
     */
    definition: {[key: string]: any};
    /**
     * A pattern to match an exchange or queue name.
     */
    pattern: string;
    /**
     * The policy with the greater priority is applied first.
     */
    priority: number;
}

export interface QueueSettings {
    /**
     * Additional key/value settings for the queue.
     * All values will be sent to RabbitMQ as a string. If you require non-string
     * values, use `argumentsJson`.
     */
    arguments?: {[key: string]: any};
    /**
     * A nested JSON string which contains additional
     * settings for the queue. This is useful for when the arguments contain
     * non-string values.
     */
    argumentsJson?: string;
    /**
     * Whether the queue will self-delete when all
     * consumers have unsubscribed.
     */
    autoDelete?: boolean;
    /**
     * Whether the queue survives server restarts.
     * Defaults to `false`.
     */
    durable?: boolean;
}

export interface TopicPermissionsPermission {
    /**
     * The exchange to set the permissions for.
     */
    exchange: string;
    /**
     * The "read" ACL.
     */
    read: string;
    /**
     * The "write" ACL.
     */
    write: string;
}
