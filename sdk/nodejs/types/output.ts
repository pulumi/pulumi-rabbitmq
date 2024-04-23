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
     *
     * Applicable to Federated Queues Only
     */
    messageTtl?: number;
    /**
     * Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
     */
    prefetchCount?: number;
    /**
     * The name of the upstream queue.
     *
     * Consult the RabbitMQ [Federation Reference](https://www.rabbitmq.com/federation-reference.html) documentation for detailed information and guidance on setting these values.
     */
    queue?: string;
    /**
     * Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
     */
    reconnectDelay?: number;
    /**
     * Determines how federation should interact with the validated user-id feature. Default is `false`.
     *
     * Applicable to Federated Exchanges Only
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
    ackMode?: string;
    /**
     * @deprecated use destinationAddForwardHeaders instead
     */
    addForwardHeaders?: boolean;
    /**
     * @deprecated use sourceDeleteAfter instead
     */
    deleteAfter?: string;
    destinationAddForwardHeaders?: boolean;
    destinationAddTimestampHeader?: boolean;
    destinationAddress?: string;
    destinationApplicationProperties?: string;
    destinationExchange?: string;
    destinationExchangeKey?: string;
    destinationProperties?: string;
    destinationProtocol?: string;
    destinationPublishProperties?: string;
    destinationQueue?: string;
    destinationUri: string;
    /**
     * @deprecated use sourcePrefetchCount instead
     */
    prefetchCount?: number;
    reconnectDelay?: number;
    sourceAddress?: string;
    sourceDeleteAfter?: string;
    sourceExchange?: string;
    sourceExchangeKey?: string;
    sourcePrefetchCount?: number;
    sourceProtocol?: string;
    sourceQueue?: string;
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

