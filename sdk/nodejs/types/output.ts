// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

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

export interface FederationUpstreamDefinition {
    /**
     * Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
     */
    ackMode?: string;
    /**
     * The name of the upstream exchange.
     */
    exchange?: string;
    /**
     * The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
     */
    expires?: number;
    /**
     * Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
     */
    maxHops?: number;
    /**
     * The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
     */
    messageTtl?: number;
    /**
     * Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
     */
    prefetchCount?: number;
    /**
     * The name of the upstream queue.
     */
    queue?: string;
    /**
     * Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
     */
    reconnectDelay?: number;
    /**
     * Determines how federation should interact with the validated user-id feature. Default is `false`.
     */
    trustUserId?: boolean;
    /**
     * The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
     */
    uri: string;
}

export interface GetExchangeSetting {
    arguments?: {[key: string]: any};
    autoDelete?: boolean;
    durable?: boolean;
    type: string;
}

export interface OperatorPolicyPolicy {
    /**
     * Can be "queues".
     */
    applyTo: string;
    /**
     * Key/value pairs of the operator policy definition. See the
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

export interface ShovelInfo {
    /**
     * Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
     * Defaults to `on-confirm`.
     */
    ackMode?: string;
    /**
     * Whether to add `x-shovelled` headers to shovelled messages.
     *
     * @deprecated use destination_add_forward_headers instead
     */
    addForwardHeaders?: boolean;
    /**
     * Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
     *
     * @deprecated use source_delete_after instead
     */
    deleteAfter?: string;
    /**
     * Whether to add `x-shovelled` headers to shovelled messages.
     */
    destinationAddForwardHeaders?: boolean;
    destinationAddTimestampHeader?: boolean;
    /**
     * The AMQP 1.0 destination link address.
     */
    destinationAddress?: string;
    /**
     * Application properties to set when shovelling messages.
     */
    destinationApplicationProperties?: string;
    /**
     * The exchange to which messages should be published.
     * Either this or `destinationQueue` must be specified but not both.
     */
    destinationExchange?: string;
    /**
     * The routing key when using `destinationExchange`.
     */
    destinationExchangeKey?: string;
    /**
     * Properties to overwrite when shovelling messages.
     */
    destinationProperties?: string;
    /**
     * The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
     * Defaults to `amqp091`.
     */
    destinationProtocol?: string;
    /**
     * A map of properties to overwrite when shovelling messages.
     */
    destinationPublishProperties?: string;
    /**
     * The queue to which messages should be published.
     * Either this or `destinationExchange` must be specified but not both.
     */
    destinationQueue?: string;
    /**
     * The amqp uri for the destination .
     */
    destinationUri: string;
    /**
     * The maximum number of unacknowledged messages copied over a shovel at any one time.
     *
     * @deprecated use source_prefetch_count instead
     */
    prefetchCount?: number;
    /**
     * The duration in seconds to reconnect to a broker after disconnected.
     * Defaults to `1`.
     */
    reconnectDelay?: number;
    /**
     * The AMQP 1.0 source link address.
     */
    sourceAddress?: string;
    /**
     * Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
     */
    sourceDeleteAfter?: string;
    /**
     * The exchange from which to consume.
     * Either this or `sourceQueue` must be specified but not both.
     */
    sourceExchange?: string;
    /**
     * The routing key when using `sourceExchange`.
     */
    sourceExchangeKey?: string;
    /**
     * The maximum number of unacknowledged messages copied over a shovel at any one time.
     */
    sourcePrefetchCount?: number;
    /**
     * The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
     * Defaults to `amqp091`.
     */
    sourceProtocol?: string;
    /**
     * The queue from which to consume.
     * Either this or `sourceExchange` must be specified but not both.
     */
    sourceQueue?: string;
    /**
     * The amqp uri for the source.
     */
    sourceUri: string;
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

