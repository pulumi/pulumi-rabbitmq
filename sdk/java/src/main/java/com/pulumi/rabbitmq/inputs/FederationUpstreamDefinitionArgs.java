// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FederationUpstreamDefinitionArgs extends com.pulumi.resources.ResourceArgs {

    public static final FederationUpstreamDefinitionArgs Empty = new FederationUpstreamDefinitionArgs();

    /**
     * Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
     * 
     */
    @Import(name="ackMode")
    private @Nullable Output<String> ackMode;

    /**
     * @return Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
     * 
     */
    public Optional<Output<String>> ackMode() {
        return Optional.ofNullable(this.ackMode);
    }

    /**
     * The name of the upstream exchange.
     * 
     */
    @Import(name="exchange")
    private @Nullable Output<String> exchange;

    /**
     * @return The name of the upstream exchange.
     * 
     */
    public Optional<Output<String>> exchange() {
        return Optional.ofNullable(this.exchange);
    }

    /**
     * The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
     * 
     */
    @Import(name="expires")
    private @Nullable Output<Integer> expires;

    /**
     * @return The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
     * 
     */
    public Optional<Output<Integer>> expires() {
        return Optional.ofNullable(this.expires);
    }

    /**
     * Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
     * 
     */
    @Import(name="maxHops")
    private @Nullable Output<Integer> maxHops;

    /**
     * @return Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
     * 
     */
    public Optional<Output<Integer>> maxHops() {
        return Optional.ofNullable(this.maxHops);
    }

    /**
     * The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
     * 
     * Applicable to Federated Queues Only
     * 
     */
    @Import(name="messageTtl")
    private @Nullable Output<Integer> messageTtl;

    /**
     * @return The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
     * 
     * Applicable to Federated Queues Only
     * 
     */
    public Optional<Output<Integer>> messageTtl() {
        return Optional.ofNullable(this.messageTtl);
    }

    /**
     * Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
     * 
     */
    @Import(name="prefetchCount")
    private @Nullable Output<Integer> prefetchCount;

    /**
     * @return Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
     * 
     */
    public Optional<Output<Integer>> prefetchCount() {
        return Optional.ofNullable(this.prefetchCount);
    }

    /**
     * The name of the upstream queue.
     * 
     * Consult the RabbitMQ [Federation Reference](https://www.rabbitmq.com/federation-reference.html) documentation for detailed information and guidance on setting these values.
     * 
     */
    @Import(name="queue")
    private @Nullable Output<String> queue;

    /**
     * @return The name of the upstream queue.
     * 
     * Consult the RabbitMQ [Federation Reference](https://www.rabbitmq.com/federation-reference.html) documentation for detailed information and guidance on setting these values.
     * 
     */
    public Optional<Output<String>> queue() {
        return Optional.ofNullable(this.queue);
    }

    /**
     * Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
     * 
     */
    @Import(name="reconnectDelay")
    private @Nullable Output<Integer> reconnectDelay;

    /**
     * @return Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
     * 
     */
    public Optional<Output<Integer>> reconnectDelay() {
        return Optional.ofNullable(this.reconnectDelay);
    }

    /**
     * Determines how federation should interact with the validated user-id feature. Default is `false`.
     * 
     * Applicable to Federated Exchanges Only
     * 
     */
    @Import(name="trustUserId")
    private @Nullable Output<Boolean> trustUserId;

    /**
     * @return Determines how federation should interact with the validated user-id feature. Default is `false`.
     * 
     * Applicable to Federated Exchanges Only
     * 
     */
    public Optional<Output<Boolean>> trustUserId() {
        return Optional.ofNullable(this.trustUserId);
    }

    /**
     * The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
     * 
     */
    @Import(name="uri", required=true)
    private Output<String> uri;

    /**
     * @return The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
     * 
     */
    public Output<String> uri() {
        return this.uri;
    }

    private FederationUpstreamDefinitionArgs() {}

    private FederationUpstreamDefinitionArgs(FederationUpstreamDefinitionArgs $) {
        this.ackMode = $.ackMode;
        this.exchange = $.exchange;
        this.expires = $.expires;
        this.maxHops = $.maxHops;
        this.messageTtl = $.messageTtl;
        this.prefetchCount = $.prefetchCount;
        this.queue = $.queue;
        this.reconnectDelay = $.reconnectDelay;
        this.trustUserId = $.trustUserId;
        this.uri = $.uri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FederationUpstreamDefinitionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FederationUpstreamDefinitionArgs $;

        public Builder() {
            $ = new FederationUpstreamDefinitionArgs();
        }

        public Builder(FederationUpstreamDefinitionArgs defaults) {
            $ = new FederationUpstreamDefinitionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ackMode Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
         * 
         * @return builder
         * 
         */
        public Builder ackMode(@Nullable Output<String> ackMode) {
            $.ackMode = ackMode;
            return this;
        }

        /**
         * @param ackMode Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
         * 
         * @return builder
         * 
         */
        public Builder ackMode(String ackMode) {
            return ackMode(Output.of(ackMode));
        }

        /**
         * @param exchange The name of the upstream exchange.
         * 
         * @return builder
         * 
         */
        public Builder exchange(@Nullable Output<String> exchange) {
            $.exchange = exchange;
            return this;
        }

        /**
         * @param exchange The name of the upstream exchange.
         * 
         * @return builder
         * 
         */
        public Builder exchange(String exchange) {
            return exchange(Output.of(exchange));
        }

        /**
         * @param expires The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
         * 
         * @return builder
         * 
         */
        public Builder expires(@Nullable Output<Integer> expires) {
            $.expires = expires;
            return this;
        }

        /**
         * @param expires The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
         * 
         * @return builder
         * 
         */
        public Builder expires(Integer expires) {
            return expires(Output.of(expires));
        }

        /**
         * @param maxHops Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
         * 
         * @return builder
         * 
         */
        public Builder maxHops(@Nullable Output<Integer> maxHops) {
            $.maxHops = maxHops;
            return this;
        }

        /**
         * @param maxHops Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
         * 
         * @return builder
         * 
         */
        public Builder maxHops(Integer maxHops) {
            return maxHops(Output.of(maxHops));
        }

        /**
         * @param messageTtl The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
         * 
         * Applicable to Federated Queues Only
         * 
         * @return builder
         * 
         */
        public Builder messageTtl(@Nullable Output<Integer> messageTtl) {
            $.messageTtl = messageTtl;
            return this;
        }

        /**
         * @param messageTtl The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
         * 
         * Applicable to Federated Queues Only
         * 
         * @return builder
         * 
         */
        public Builder messageTtl(Integer messageTtl) {
            return messageTtl(Output.of(messageTtl));
        }

        /**
         * @param prefetchCount Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
         * 
         * @return builder
         * 
         */
        public Builder prefetchCount(@Nullable Output<Integer> prefetchCount) {
            $.prefetchCount = prefetchCount;
            return this;
        }

        /**
         * @param prefetchCount Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
         * 
         * @return builder
         * 
         */
        public Builder prefetchCount(Integer prefetchCount) {
            return prefetchCount(Output.of(prefetchCount));
        }

        /**
         * @param queue The name of the upstream queue.
         * 
         * Consult the RabbitMQ [Federation Reference](https://www.rabbitmq.com/federation-reference.html) documentation for detailed information and guidance on setting these values.
         * 
         * @return builder
         * 
         */
        public Builder queue(@Nullable Output<String> queue) {
            $.queue = queue;
            return this;
        }

        /**
         * @param queue The name of the upstream queue.
         * 
         * Consult the RabbitMQ [Federation Reference](https://www.rabbitmq.com/federation-reference.html) documentation for detailed information and guidance on setting these values.
         * 
         * @return builder
         * 
         */
        public Builder queue(String queue) {
            return queue(Output.of(queue));
        }

        /**
         * @param reconnectDelay Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
         * 
         * @return builder
         * 
         */
        public Builder reconnectDelay(@Nullable Output<Integer> reconnectDelay) {
            $.reconnectDelay = reconnectDelay;
            return this;
        }

        /**
         * @param reconnectDelay Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
         * 
         * @return builder
         * 
         */
        public Builder reconnectDelay(Integer reconnectDelay) {
            return reconnectDelay(Output.of(reconnectDelay));
        }

        /**
         * @param trustUserId Determines how federation should interact with the validated user-id feature. Default is `false`.
         * 
         * Applicable to Federated Exchanges Only
         * 
         * @return builder
         * 
         */
        public Builder trustUserId(@Nullable Output<Boolean> trustUserId) {
            $.trustUserId = trustUserId;
            return this;
        }

        /**
         * @param trustUserId Determines how federation should interact with the validated user-id feature. Default is `false`.
         * 
         * Applicable to Federated Exchanges Only
         * 
         * @return builder
         * 
         */
        public Builder trustUserId(Boolean trustUserId) {
            return trustUserId(Output.of(trustUserId));
        }

        /**
         * @param uri The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
         * 
         * @return builder
         * 
         */
        public Builder uri(Output<String> uri) {
            $.uri = uri;
            return this;
        }

        /**
         * @param uri The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
         * 
         * @return builder
         * 
         */
        public Builder uri(String uri) {
            return uri(Output.of(uri));
        }

        public FederationUpstreamDefinitionArgs build() {
            if ($.uri == null) {
                throw new MissingRequiredPropertyException("FederationUpstreamDefinitionArgs", "uri");
            }
            return $;
        }
    }

}
