// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FederationUpstreamDefinition {
    /**
     * @return Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
     * 
     */
    private @Nullable String ackMode;
    /**
     * @return The name of the upstream exchange.
     * 
     */
    private @Nullable String exchange;
    /**
     * @return The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
     * 
     */
    private @Nullable Integer expires;
    /**
     * @return Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
     * 
     */
    private @Nullable Integer maxHops;
    /**
     * @return The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
     * 
     * Applicable to Federated Queues Only
     * 
     */
    private @Nullable Integer messageTtl;
    /**
     * @return Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
     * 
     */
    private @Nullable Integer prefetchCount;
    /**
     * @return The name of the upstream queue.
     * 
     * Consult the RabbitMQ [Federation Reference](https://www.rabbitmq.com/federation-reference.html) documentation for detailed information and guidance on setting these values.
     * 
     */
    private @Nullable String queue;
    /**
     * @return Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
     * 
     */
    private @Nullable Integer reconnectDelay;
    /**
     * @return Determines how federation should interact with the validated user-id feature. Default is `false`.
     * 
     * Applicable to Federated Exchanges Only
     * 
     */
    private @Nullable Boolean trustUserId;
    /**
     * @return The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
     * 
     */
    private String uri;

    private FederationUpstreamDefinition() {}
    /**
     * @return Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
     * 
     */
    public Optional<String> ackMode() {
        return Optional.ofNullable(this.ackMode);
    }
    /**
     * @return The name of the upstream exchange.
     * 
     */
    public Optional<String> exchange() {
        return Optional.ofNullable(this.exchange);
    }
    /**
     * @return The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
     * 
     */
    public Optional<Integer> expires() {
        return Optional.ofNullable(this.expires);
    }
    /**
     * @return Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
     * 
     */
    public Optional<Integer> maxHops() {
        return Optional.ofNullable(this.maxHops);
    }
    /**
     * @return The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
     * 
     * Applicable to Federated Queues Only
     * 
     */
    public Optional<Integer> messageTtl() {
        return Optional.ofNullable(this.messageTtl);
    }
    /**
     * @return Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
     * 
     */
    public Optional<Integer> prefetchCount() {
        return Optional.ofNullable(this.prefetchCount);
    }
    /**
     * @return The name of the upstream queue.
     * 
     * Consult the RabbitMQ [Federation Reference](https://www.rabbitmq.com/federation-reference.html) documentation for detailed information and guidance on setting these values.
     * 
     */
    public Optional<String> queue() {
        return Optional.ofNullable(this.queue);
    }
    /**
     * @return Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
     * 
     */
    public Optional<Integer> reconnectDelay() {
        return Optional.ofNullable(this.reconnectDelay);
    }
    /**
     * @return Determines how federation should interact with the validated user-id feature. Default is `false`.
     * 
     * Applicable to Federated Exchanges Only
     * 
     */
    public Optional<Boolean> trustUserId() {
        return Optional.ofNullable(this.trustUserId);
    }
    /**
     * @return The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
     * 
     */
    public String uri() {
        return this.uri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FederationUpstreamDefinition defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ackMode;
        private @Nullable String exchange;
        private @Nullable Integer expires;
        private @Nullable Integer maxHops;
        private @Nullable Integer messageTtl;
        private @Nullable Integer prefetchCount;
        private @Nullable String queue;
        private @Nullable Integer reconnectDelay;
        private @Nullable Boolean trustUserId;
        private String uri;
        public Builder() {}
        public Builder(FederationUpstreamDefinition defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ackMode = defaults.ackMode;
    	      this.exchange = defaults.exchange;
    	      this.expires = defaults.expires;
    	      this.maxHops = defaults.maxHops;
    	      this.messageTtl = defaults.messageTtl;
    	      this.prefetchCount = defaults.prefetchCount;
    	      this.queue = defaults.queue;
    	      this.reconnectDelay = defaults.reconnectDelay;
    	      this.trustUserId = defaults.trustUserId;
    	      this.uri = defaults.uri;
        }

        @CustomType.Setter
        public Builder ackMode(@Nullable String ackMode) {
            this.ackMode = ackMode;
            return this;
        }
        @CustomType.Setter
        public Builder exchange(@Nullable String exchange) {
            this.exchange = exchange;
            return this;
        }
        @CustomType.Setter
        public Builder expires(@Nullable Integer expires) {
            this.expires = expires;
            return this;
        }
        @CustomType.Setter
        public Builder maxHops(@Nullable Integer maxHops) {
            this.maxHops = maxHops;
            return this;
        }
        @CustomType.Setter
        public Builder messageTtl(@Nullable Integer messageTtl) {
            this.messageTtl = messageTtl;
            return this;
        }
        @CustomType.Setter
        public Builder prefetchCount(@Nullable Integer prefetchCount) {
            this.prefetchCount = prefetchCount;
            return this;
        }
        @CustomType.Setter
        public Builder queue(@Nullable String queue) {
            this.queue = queue;
            return this;
        }
        @CustomType.Setter
        public Builder reconnectDelay(@Nullable Integer reconnectDelay) {
            this.reconnectDelay = reconnectDelay;
            return this;
        }
        @CustomType.Setter
        public Builder trustUserId(@Nullable Boolean trustUserId) {
            this.trustUserId = trustUserId;
            return this;
        }
        @CustomType.Setter
        public Builder uri(String uri) {
            this.uri = Objects.requireNonNull(uri);
            return this;
        }
        public FederationUpstreamDefinition build() {
            final var _resultValue = new FederationUpstreamDefinition();
            _resultValue.ackMode = ackMode;
            _resultValue.exchange = exchange;
            _resultValue.expires = expires;
            _resultValue.maxHops = maxHops;
            _resultValue.messageTtl = messageTtl;
            _resultValue.prefetchCount = prefetchCount;
            _resultValue.queue = queue;
            _resultValue.reconnectDelay = reconnectDelay;
            _resultValue.trustUserId = trustUserId;
            _resultValue.uri = uri;
            return _resultValue;
        }
    }
}
