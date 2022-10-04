// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ShovelInfoArgs extends com.pulumi.resources.ResourceArgs {

    public static final ShovelInfoArgs Empty = new ShovelInfoArgs();

    /**
     * Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
     * Defaults to `on-confirm`.
     * 
     */
    @Import(name="ackMode")
    private @Nullable Output<String> ackMode;

    /**
     * @return Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
     * Defaults to `on-confirm`.
     * 
     */
    public Optional<Output<String>> ackMode() {
        return Optional.ofNullable(this.ackMode);
    }

    /**
     * Whether to add `x-shovelled` headers to shovelled messages.
     * 
     * @deprecated
     * use destination_add_forward_headers instead
     * 
     */
    @Deprecated /* use destination_add_forward_headers instead */
    @Import(name="addForwardHeaders")
    private @Nullable Output<Boolean> addForwardHeaders;

    /**
     * @return Whether to add `x-shovelled` headers to shovelled messages.
     * 
     * @deprecated
     * use destination_add_forward_headers instead
     * 
     */
    @Deprecated /* use destination_add_forward_headers instead */
    public Optional<Output<Boolean>> addForwardHeaders() {
        return Optional.ofNullable(this.addForwardHeaders);
    }

    /**
     * Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
     * 
     * @deprecated
     * use source_delete_after instead
     * 
     */
    @Deprecated /* use source_delete_after instead */
    @Import(name="deleteAfter")
    private @Nullable Output<String> deleteAfter;

    /**
     * @return Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
     * 
     * @deprecated
     * use source_delete_after instead
     * 
     */
    @Deprecated /* use source_delete_after instead */
    public Optional<Output<String>> deleteAfter() {
        return Optional.ofNullable(this.deleteAfter);
    }

    /**
     * Whether to add `x-shovelled` headers to shovelled messages.
     * 
     */
    @Import(name="destinationAddForwardHeaders")
    private @Nullable Output<Boolean> destinationAddForwardHeaders;

    /**
     * @return Whether to add `x-shovelled` headers to shovelled messages.
     * 
     */
    public Optional<Output<Boolean>> destinationAddForwardHeaders() {
        return Optional.ofNullable(this.destinationAddForwardHeaders);
    }

    @Import(name="destinationAddTimestampHeader")
    private @Nullable Output<Boolean> destinationAddTimestampHeader;

    public Optional<Output<Boolean>> destinationAddTimestampHeader() {
        return Optional.ofNullable(this.destinationAddTimestampHeader);
    }

    /**
     * The AMQP 1.0 destination link address.
     * 
     */
    @Import(name="destinationAddress")
    private @Nullable Output<String> destinationAddress;

    /**
     * @return The AMQP 1.0 destination link address.
     * 
     */
    public Optional<Output<String>> destinationAddress() {
        return Optional.ofNullable(this.destinationAddress);
    }

    /**
     * Application properties to set when shovelling messages.
     * 
     */
    @Import(name="destinationApplicationProperties")
    private @Nullable Output<String> destinationApplicationProperties;

    /**
     * @return Application properties to set when shovelling messages.
     * 
     */
    public Optional<Output<String>> destinationApplicationProperties() {
        return Optional.ofNullable(this.destinationApplicationProperties);
    }

    /**
     * The exchange to which messages should be published.
     * Either this or `destination_queue` must be specified but not both.
     * 
     */
    @Import(name="destinationExchange")
    private @Nullable Output<String> destinationExchange;

    /**
     * @return The exchange to which messages should be published.
     * Either this or `destination_queue` must be specified but not both.
     * 
     */
    public Optional<Output<String>> destinationExchange() {
        return Optional.ofNullable(this.destinationExchange);
    }

    /**
     * The routing key when using `destination_exchange`.
     * 
     */
    @Import(name="destinationExchangeKey")
    private @Nullable Output<String> destinationExchangeKey;

    /**
     * @return The routing key when using `destination_exchange`.
     * 
     */
    public Optional<Output<String>> destinationExchangeKey() {
        return Optional.ofNullable(this.destinationExchangeKey);
    }

    /**
     * Properties to overwrite when shovelling messages.
     * 
     */
    @Import(name="destinationProperties")
    private @Nullable Output<String> destinationProperties;

    /**
     * @return Properties to overwrite when shovelling messages.
     * 
     */
    public Optional<Output<String>> destinationProperties() {
        return Optional.ofNullable(this.destinationProperties);
    }

    /**
     * The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
     * Defaults to `amqp091`.
     * 
     */
    @Import(name="destinationProtocol")
    private @Nullable Output<String> destinationProtocol;

    /**
     * @return The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
     * Defaults to `amqp091`.
     * 
     */
    public Optional<Output<String>> destinationProtocol() {
        return Optional.ofNullable(this.destinationProtocol);
    }

    /**
     * A map of properties to overwrite when shovelling messages.
     * 
     */
    @Import(name="destinationPublishProperties")
    private @Nullable Output<String> destinationPublishProperties;

    /**
     * @return A map of properties to overwrite when shovelling messages.
     * 
     */
    public Optional<Output<String>> destinationPublishProperties() {
        return Optional.ofNullable(this.destinationPublishProperties);
    }

    /**
     * The queue to which messages should be published.
     * Either this or `destination_exchange` must be specified but not both.
     * 
     */
    @Import(name="destinationQueue")
    private @Nullable Output<String> destinationQueue;

    /**
     * @return The queue to which messages should be published.
     * Either this or `destination_exchange` must be specified but not both.
     * 
     */
    public Optional<Output<String>> destinationQueue() {
        return Optional.ofNullable(this.destinationQueue);
    }

    /**
     * The amqp uri for the destination .
     * 
     */
    @Import(name="destinationUri", required=true)
    private Output<String> destinationUri;

    /**
     * @return The amqp uri for the destination .
     * 
     */
    public Output<String> destinationUri() {
        return this.destinationUri;
    }

    /**
     * The maximum number of unacknowledged messages copied over a shovel at any one time.
     * 
     * @deprecated
     * use source_prefetch_count instead
     * 
     */
    @Deprecated /* use source_prefetch_count instead */
    @Import(name="prefetchCount")
    private @Nullable Output<Integer> prefetchCount;

    /**
     * @return The maximum number of unacknowledged messages copied over a shovel at any one time.
     * 
     * @deprecated
     * use source_prefetch_count instead
     * 
     */
    @Deprecated /* use source_prefetch_count instead */
    public Optional<Output<Integer>> prefetchCount() {
        return Optional.ofNullable(this.prefetchCount);
    }

    /**
     * The duration in seconds to reconnect to a broker after disconnected.
     * Defaults to `1`.
     * 
     */
    @Import(name="reconnectDelay")
    private @Nullable Output<Integer> reconnectDelay;

    /**
     * @return The duration in seconds to reconnect to a broker after disconnected.
     * Defaults to `1`.
     * 
     */
    public Optional<Output<Integer>> reconnectDelay() {
        return Optional.ofNullable(this.reconnectDelay);
    }

    /**
     * The AMQP 1.0 source link address.
     * 
     */
    @Import(name="sourceAddress")
    private @Nullable Output<String> sourceAddress;

    /**
     * @return The AMQP 1.0 source link address.
     * 
     */
    public Optional<Output<String>> sourceAddress() {
        return Optional.ofNullable(this.sourceAddress);
    }

    /**
     * Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
     * 
     */
    @Import(name="sourceDeleteAfter")
    private @Nullable Output<String> sourceDeleteAfter;

    /**
     * @return Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
     * 
     */
    public Optional<Output<String>> sourceDeleteAfter() {
        return Optional.ofNullable(this.sourceDeleteAfter);
    }

    /**
     * The exchange from which to consume.
     * Either this or `source_queue` must be specified but not both.
     * 
     */
    @Import(name="sourceExchange")
    private @Nullable Output<String> sourceExchange;

    /**
     * @return The exchange from which to consume.
     * Either this or `source_queue` must be specified but not both.
     * 
     */
    public Optional<Output<String>> sourceExchange() {
        return Optional.ofNullable(this.sourceExchange);
    }

    /**
     * The routing key when using `source_exchange`.
     * 
     */
    @Import(name="sourceExchangeKey")
    private @Nullable Output<String> sourceExchangeKey;

    /**
     * @return The routing key when using `source_exchange`.
     * 
     */
    public Optional<Output<String>> sourceExchangeKey() {
        return Optional.ofNullable(this.sourceExchangeKey);
    }

    /**
     * The maximum number of unacknowledged messages copied over a shovel at any one time.
     * 
     */
    @Import(name="sourcePrefetchCount")
    private @Nullable Output<Integer> sourcePrefetchCount;

    /**
     * @return The maximum number of unacknowledged messages copied over a shovel at any one time.
     * 
     */
    public Optional<Output<Integer>> sourcePrefetchCount() {
        return Optional.ofNullable(this.sourcePrefetchCount);
    }

    /**
     * The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
     * Defaults to `amqp091`.
     * 
     */
    @Import(name="sourceProtocol")
    private @Nullable Output<String> sourceProtocol;

    /**
     * @return The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
     * Defaults to `amqp091`.
     * 
     */
    public Optional<Output<String>> sourceProtocol() {
        return Optional.ofNullable(this.sourceProtocol);
    }

    /**
     * The queue from which to consume.
     * Either this or `source_exchange` must be specified but not both.
     * 
     */
    @Import(name="sourceQueue")
    private @Nullable Output<String> sourceQueue;

    /**
     * @return The queue from which to consume.
     * Either this or `source_exchange` must be specified but not both.
     * 
     */
    public Optional<Output<String>> sourceQueue() {
        return Optional.ofNullable(this.sourceQueue);
    }

    /**
     * The amqp uri for the source.
     * 
     */
    @Import(name="sourceUri", required=true)
    private Output<String> sourceUri;

    /**
     * @return The amqp uri for the source.
     * 
     */
    public Output<String> sourceUri() {
        return this.sourceUri;
    }

    private ShovelInfoArgs() {}

    private ShovelInfoArgs(ShovelInfoArgs $) {
        this.ackMode = $.ackMode;
        this.addForwardHeaders = $.addForwardHeaders;
        this.deleteAfter = $.deleteAfter;
        this.destinationAddForwardHeaders = $.destinationAddForwardHeaders;
        this.destinationAddTimestampHeader = $.destinationAddTimestampHeader;
        this.destinationAddress = $.destinationAddress;
        this.destinationApplicationProperties = $.destinationApplicationProperties;
        this.destinationExchange = $.destinationExchange;
        this.destinationExchangeKey = $.destinationExchangeKey;
        this.destinationProperties = $.destinationProperties;
        this.destinationProtocol = $.destinationProtocol;
        this.destinationPublishProperties = $.destinationPublishProperties;
        this.destinationQueue = $.destinationQueue;
        this.destinationUri = $.destinationUri;
        this.prefetchCount = $.prefetchCount;
        this.reconnectDelay = $.reconnectDelay;
        this.sourceAddress = $.sourceAddress;
        this.sourceDeleteAfter = $.sourceDeleteAfter;
        this.sourceExchange = $.sourceExchange;
        this.sourceExchangeKey = $.sourceExchangeKey;
        this.sourcePrefetchCount = $.sourcePrefetchCount;
        this.sourceProtocol = $.sourceProtocol;
        this.sourceQueue = $.sourceQueue;
        this.sourceUri = $.sourceUri;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ShovelInfoArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ShovelInfoArgs $;

        public Builder() {
            $ = new ShovelInfoArgs();
        }

        public Builder(ShovelInfoArgs defaults) {
            $ = new ShovelInfoArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ackMode Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
         * Defaults to `on-confirm`.
         * 
         * @return builder
         * 
         */
        public Builder ackMode(@Nullable Output<String> ackMode) {
            $.ackMode = ackMode;
            return this;
        }

        /**
         * @param ackMode Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
         * Defaults to `on-confirm`.
         * 
         * @return builder
         * 
         */
        public Builder ackMode(String ackMode) {
            return ackMode(Output.of(ackMode));
        }

        /**
         * @param addForwardHeaders Whether to add `x-shovelled` headers to shovelled messages.
         * 
         * @return builder
         * 
         * @deprecated
         * use destination_add_forward_headers instead
         * 
         */
        @Deprecated /* use destination_add_forward_headers instead */
        public Builder addForwardHeaders(@Nullable Output<Boolean> addForwardHeaders) {
            $.addForwardHeaders = addForwardHeaders;
            return this;
        }

        /**
         * @param addForwardHeaders Whether to add `x-shovelled` headers to shovelled messages.
         * 
         * @return builder
         * 
         * @deprecated
         * use destination_add_forward_headers instead
         * 
         */
        @Deprecated /* use destination_add_forward_headers instead */
        public Builder addForwardHeaders(Boolean addForwardHeaders) {
            return addForwardHeaders(Output.of(addForwardHeaders));
        }

        /**
         * @param deleteAfter Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
         * 
         * @return builder
         * 
         * @deprecated
         * use source_delete_after instead
         * 
         */
        @Deprecated /* use source_delete_after instead */
        public Builder deleteAfter(@Nullable Output<String> deleteAfter) {
            $.deleteAfter = deleteAfter;
            return this;
        }

        /**
         * @param deleteAfter Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
         * 
         * @return builder
         * 
         * @deprecated
         * use source_delete_after instead
         * 
         */
        @Deprecated /* use source_delete_after instead */
        public Builder deleteAfter(String deleteAfter) {
            return deleteAfter(Output.of(deleteAfter));
        }

        /**
         * @param destinationAddForwardHeaders Whether to add `x-shovelled` headers to shovelled messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationAddForwardHeaders(@Nullable Output<Boolean> destinationAddForwardHeaders) {
            $.destinationAddForwardHeaders = destinationAddForwardHeaders;
            return this;
        }

        /**
         * @param destinationAddForwardHeaders Whether to add `x-shovelled` headers to shovelled messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationAddForwardHeaders(Boolean destinationAddForwardHeaders) {
            return destinationAddForwardHeaders(Output.of(destinationAddForwardHeaders));
        }

        public Builder destinationAddTimestampHeader(@Nullable Output<Boolean> destinationAddTimestampHeader) {
            $.destinationAddTimestampHeader = destinationAddTimestampHeader;
            return this;
        }

        public Builder destinationAddTimestampHeader(Boolean destinationAddTimestampHeader) {
            return destinationAddTimestampHeader(Output.of(destinationAddTimestampHeader));
        }

        /**
         * @param destinationAddress The AMQP 1.0 destination link address.
         * 
         * @return builder
         * 
         */
        public Builder destinationAddress(@Nullable Output<String> destinationAddress) {
            $.destinationAddress = destinationAddress;
            return this;
        }

        /**
         * @param destinationAddress The AMQP 1.0 destination link address.
         * 
         * @return builder
         * 
         */
        public Builder destinationAddress(String destinationAddress) {
            return destinationAddress(Output.of(destinationAddress));
        }

        /**
         * @param destinationApplicationProperties Application properties to set when shovelling messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationApplicationProperties(@Nullable Output<String> destinationApplicationProperties) {
            $.destinationApplicationProperties = destinationApplicationProperties;
            return this;
        }

        /**
         * @param destinationApplicationProperties Application properties to set when shovelling messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationApplicationProperties(String destinationApplicationProperties) {
            return destinationApplicationProperties(Output.of(destinationApplicationProperties));
        }

        /**
         * @param destinationExchange The exchange to which messages should be published.
         * Either this or `destination_queue` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder destinationExchange(@Nullable Output<String> destinationExchange) {
            $.destinationExchange = destinationExchange;
            return this;
        }

        /**
         * @param destinationExchange The exchange to which messages should be published.
         * Either this or `destination_queue` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder destinationExchange(String destinationExchange) {
            return destinationExchange(Output.of(destinationExchange));
        }

        /**
         * @param destinationExchangeKey The routing key when using `destination_exchange`.
         * 
         * @return builder
         * 
         */
        public Builder destinationExchangeKey(@Nullable Output<String> destinationExchangeKey) {
            $.destinationExchangeKey = destinationExchangeKey;
            return this;
        }

        /**
         * @param destinationExchangeKey The routing key when using `destination_exchange`.
         * 
         * @return builder
         * 
         */
        public Builder destinationExchangeKey(String destinationExchangeKey) {
            return destinationExchangeKey(Output.of(destinationExchangeKey));
        }

        /**
         * @param destinationProperties Properties to overwrite when shovelling messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationProperties(@Nullable Output<String> destinationProperties) {
            $.destinationProperties = destinationProperties;
            return this;
        }

        /**
         * @param destinationProperties Properties to overwrite when shovelling messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationProperties(String destinationProperties) {
            return destinationProperties(Output.of(destinationProperties));
        }

        /**
         * @param destinationProtocol The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
         * Defaults to `amqp091`.
         * 
         * @return builder
         * 
         */
        public Builder destinationProtocol(@Nullable Output<String> destinationProtocol) {
            $.destinationProtocol = destinationProtocol;
            return this;
        }

        /**
         * @param destinationProtocol The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
         * Defaults to `amqp091`.
         * 
         * @return builder
         * 
         */
        public Builder destinationProtocol(String destinationProtocol) {
            return destinationProtocol(Output.of(destinationProtocol));
        }

        /**
         * @param destinationPublishProperties A map of properties to overwrite when shovelling messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationPublishProperties(@Nullable Output<String> destinationPublishProperties) {
            $.destinationPublishProperties = destinationPublishProperties;
            return this;
        }

        /**
         * @param destinationPublishProperties A map of properties to overwrite when shovelling messages.
         * 
         * @return builder
         * 
         */
        public Builder destinationPublishProperties(String destinationPublishProperties) {
            return destinationPublishProperties(Output.of(destinationPublishProperties));
        }

        /**
         * @param destinationQueue The queue to which messages should be published.
         * Either this or `destination_exchange` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder destinationQueue(@Nullable Output<String> destinationQueue) {
            $.destinationQueue = destinationQueue;
            return this;
        }

        /**
         * @param destinationQueue The queue to which messages should be published.
         * Either this or `destination_exchange` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder destinationQueue(String destinationQueue) {
            return destinationQueue(Output.of(destinationQueue));
        }

        /**
         * @param destinationUri The amqp uri for the destination .
         * 
         * @return builder
         * 
         */
        public Builder destinationUri(Output<String> destinationUri) {
            $.destinationUri = destinationUri;
            return this;
        }

        /**
         * @param destinationUri The amqp uri for the destination .
         * 
         * @return builder
         * 
         */
        public Builder destinationUri(String destinationUri) {
            return destinationUri(Output.of(destinationUri));
        }

        /**
         * @param prefetchCount The maximum number of unacknowledged messages copied over a shovel at any one time.
         * 
         * @return builder
         * 
         * @deprecated
         * use source_prefetch_count instead
         * 
         */
        @Deprecated /* use source_prefetch_count instead */
        public Builder prefetchCount(@Nullable Output<Integer> prefetchCount) {
            $.prefetchCount = prefetchCount;
            return this;
        }

        /**
         * @param prefetchCount The maximum number of unacknowledged messages copied over a shovel at any one time.
         * 
         * @return builder
         * 
         * @deprecated
         * use source_prefetch_count instead
         * 
         */
        @Deprecated /* use source_prefetch_count instead */
        public Builder prefetchCount(Integer prefetchCount) {
            return prefetchCount(Output.of(prefetchCount));
        }

        /**
         * @param reconnectDelay The duration in seconds to reconnect to a broker after disconnected.
         * Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder reconnectDelay(@Nullable Output<Integer> reconnectDelay) {
            $.reconnectDelay = reconnectDelay;
            return this;
        }

        /**
         * @param reconnectDelay The duration in seconds to reconnect to a broker after disconnected.
         * Defaults to `1`.
         * 
         * @return builder
         * 
         */
        public Builder reconnectDelay(Integer reconnectDelay) {
            return reconnectDelay(Output.of(reconnectDelay));
        }

        /**
         * @param sourceAddress The AMQP 1.0 source link address.
         * 
         * @return builder
         * 
         */
        public Builder sourceAddress(@Nullable Output<String> sourceAddress) {
            $.sourceAddress = sourceAddress;
            return this;
        }

        /**
         * @param sourceAddress The AMQP 1.0 source link address.
         * 
         * @return builder
         * 
         */
        public Builder sourceAddress(String sourceAddress) {
            return sourceAddress(Output.of(sourceAddress));
        }

        /**
         * @param sourceDeleteAfter Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
         * 
         * @return builder
         * 
         */
        public Builder sourceDeleteAfter(@Nullable Output<String> sourceDeleteAfter) {
            $.sourceDeleteAfter = sourceDeleteAfter;
            return this;
        }

        /**
         * @param sourceDeleteAfter Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
         * 
         * @return builder
         * 
         */
        public Builder sourceDeleteAfter(String sourceDeleteAfter) {
            return sourceDeleteAfter(Output.of(sourceDeleteAfter));
        }

        /**
         * @param sourceExchange The exchange from which to consume.
         * Either this or `source_queue` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder sourceExchange(@Nullable Output<String> sourceExchange) {
            $.sourceExchange = sourceExchange;
            return this;
        }

        /**
         * @param sourceExchange The exchange from which to consume.
         * Either this or `source_queue` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder sourceExchange(String sourceExchange) {
            return sourceExchange(Output.of(sourceExchange));
        }

        /**
         * @param sourceExchangeKey The routing key when using `source_exchange`.
         * 
         * @return builder
         * 
         */
        public Builder sourceExchangeKey(@Nullable Output<String> sourceExchangeKey) {
            $.sourceExchangeKey = sourceExchangeKey;
            return this;
        }

        /**
         * @param sourceExchangeKey The routing key when using `source_exchange`.
         * 
         * @return builder
         * 
         */
        public Builder sourceExchangeKey(String sourceExchangeKey) {
            return sourceExchangeKey(Output.of(sourceExchangeKey));
        }

        /**
         * @param sourcePrefetchCount The maximum number of unacknowledged messages copied over a shovel at any one time.
         * 
         * @return builder
         * 
         */
        public Builder sourcePrefetchCount(@Nullable Output<Integer> sourcePrefetchCount) {
            $.sourcePrefetchCount = sourcePrefetchCount;
            return this;
        }

        /**
         * @param sourcePrefetchCount The maximum number of unacknowledged messages copied over a shovel at any one time.
         * 
         * @return builder
         * 
         */
        public Builder sourcePrefetchCount(Integer sourcePrefetchCount) {
            return sourcePrefetchCount(Output.of(sourcePrefetchCount));
        }

        /**
         * @param sourceProtocol The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
         * Defaults to `amqp091`.
         * 
         * @return builder
         * 
         */
        public Builder sourceProtocol(@Nullable Output<String> sourceProtocol) {
            $.sourceProtocol = sourceProtocol;
            return this;
        }

        /**
         * @param sourceProtocol The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
         * Defaults to `amqp091`.
         * 
         * @return builder
         * 
         */
        public Builder sourceProtocol(String sourceProtocol) {
            return sourceProtocol(Output.of(sourceProtocol));
        }

        /**
         * @param sourceQueue The queue from which to consume.
         * Either this or `source_exchange` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder sourceQueue(@Nullable Output<String> sourceQueue) {
            $.sourceQueue = sourceQueue;
            return this;
        }

        /**
         * @param sourceQueue The queue from which to consume.
         * Either this or `source_exchange` must be specified but not both.
         * 
         * @return builder
         * 
         */
        public Builder sourceQueue(String sourceQueue) {
            return sourceQueue(Output.of(sourceQueue));
        }

        /**
         * @param sourceUri The amqp uri for the source.
         * 
         * @return builder
         * 
         */
        public Builder sourceUri(Output<String> sourceUri) {
            $.sourceUri = sourceUri;
            return this;
        }

        /**
         * @param sourceUri The amqp uri for the source.
         * 
         * @return builder
         * 
         */
        public Builder sourceUri(String sourceUri) {
            return sourceUri(Output.of(sourceUri));
        }

        public ShovelInfoArgs build() {
            $.destinationUri = Objects.requireNonNull($.destinationUri, "expected parameter 'destinationUri' to be non-null");
            $.sourceUri = Objects.requireNonNull($.sourceUri, "expected parameter 'sourceUri' to be non-null");
            return $;
        }
    }

}