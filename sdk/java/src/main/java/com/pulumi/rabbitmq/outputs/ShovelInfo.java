// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ShovelInfo {
    private @Nullable String ackMode;
    /**
     * @deprecated
     * use destination_add_forward_headers instead
     * 
     */
    @Deprecated /* use destination_add_forward_headers instead */
    private @Nullable Boolean addForwardHeaders;
    /**
     * @deprecated
     * use source_delete_after instead
     * 
     */
    @Deprecated /* use source_delete_after instead */
    private @Nullable String deleteAfter;
    private @Nullable Boolean destinationAddForwardHeaders;
    private @Nullable Boolean destinationAddTimestampHeader;
    private @Nullable String destinationAddress;
    private @Nullable String destinationApplicationProperties;
    private @Nullable String destinationExchange;
    private @Nullable String destinationExchangeKey;
    private @Nullable String destinationProperties;
    private @Nullable String destinationProtocol;
    private @Nullable String destinationPublishProperties;
    private @Nullable String destinationQueue;
    private String destinationUri;
    /**
     * @deprecated
     * use source_prefetch_count instead
     * 
     */
    @Deprecated /* use source_prefetch_count instead */
    private @Nullable Integer prefetchCount;
    private @Nullable Integer reconnectDelay;
    private @Nullable String sourceAddress;
    private @Nullable String sourceDeleteAfter;
    private @Nullable String sourceExchange;
    private @Nullable String sourceExchangeKey;
    private @Nullable Integer sourcePrefetchCount;
    private @Nullable String sourceProtocol;
    private @Nullable String sourceQueue;
    private String sourceUri;

    private ShovelInfo() {}
    public Optional<String> ackMode() {
        return Optional.ofNullable(this.ackMode);
    }
    /**
     * @deprecated
     * use destination_add_forward_headers instead
     * 
     */
    @Deprecated /* use destination_add_forward_headers instead */
    public Optional<Boolean> addForwardHeaders() {
        return Optional.ofNullable(this.addForwardHeaders);
    }
    /**
     * @deprecated
     * use source_delete_after instead
     * 
     */
    @Deprecated /* use source_delete_after instead */
    public Optional<String> deleteAfter() {
        return Optional.ofNullable(this.deleteAfter);
    }
    public Optional<Boolean> destinationAddForwardHeaders() {
        return Optional.ofNullable(this.destinationAddForwardHeaders);
    }
    public Optional<Boolean> destinationAddTimestampHeader() {
        return Optional.ofNullable(this.destinationAddTimestampHeader);
    }
    public Optional<String> destinationAddress() {
        return Optional.ofNullable(this.destinationAddress);
    }
    public Optional<String> destinationApplicationProperties() {
        return Optional.ofNullable(this.destinationApplicationProperties);
    }
    public Optional<String> destinationExchange() {
        return Optional.ofNullable(this.destinationExchange);
    }
    public Optional<String> destinationExchangeKey() {
        return Optional.ofNullable(this.destinationExchangeKey);
    }
    public Optional<String> destinationProperties() {
        return Optional.ofNullable(this.destinationProperties);
    }
    public Optional<String> destinationProtocol() {
        return Optional.ofNullable(this.destinationProtocol);
    }
    public Optional<String> destinationPublishProperties() {
        return Optional.ofNullable(this.destinationPublishProperties);
    }
    public Optional<String> destinationQueue() {
        return Optional.ofNullable(this.destinationQueue);
    }
    public String destinationUri() {
        return this.destinationUri;
    }
    /**
     * @deprecated
     * use source_prefetch_count instead
     * 
     */
    @Deprecated /* use source_prefetch_count instead */
    public Optional<Integer> prefetchCount() {
        return Optional.ofNullable(this.prefetchCount);
    }
    public Optional<Integer> reconnectDelay() {
        return Optional.ofNullable(this.reconnectDelay);
    }
    public Optional<String> sourceAddress() {
        return Optional.ofNullable(this.sourceAddress);
    }
    public Optional<String> sourceDeleteAfter() {
        return Optional.ofNullable(this.sourceDeleteAfter);
    }
    public Optional<String> sourceExchange() {
        return Optional.ofNullable(this.sourceExchange);
    }
    public Optional<String> sourceExchangeKey() {
        return Optional.ofNullable(this.sourceExchangeKey);
    }
    public Optional<Integer> sourcePrefetchCount() {
        return Optional.ofNullable(this.sourcePrefetchCount);
    }
    public Optional<String> sourceProtocol() {
        return Optional.ofNullable(this.sourceProtocol);
    }
    public Optional<String> sourceQueue() {
        return Optional.ofNullable(this.sourceQueue);
    }
    public String sourceUri() {
        return this.sourceUri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ShovelInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ackMode;
        private @Nullable Boolean addForwardHeaders;
        private @Nullable String deleteAfter;
        private @Nullable Boolean destinationAddForwardHeaders;
        private @Nullable Boolean destinationAddTimestampHeader;
        private @Nullable String destinationAddress;
        private @Nullable String destinationApplicationProperties;
        private @Nullable String destinationExchange;
        private @Nullable String destinationExchangeKey;
        private @Nullable String destinationProperties;
        private @Nullable String destinationProtocol;
        private @Nullable String destinationPublishProperties;
        private @Nullable String destinationQueue;
        private String destinationUri;
        private @Nullable Integer prefetchCount;
        private @Nullable Integer reconnectDelay;
        private @Nullable String sourceAddress;
        private @Nullable String sourceDeleteAfter;
        private @Nullable String sourceExchange;
        private @Nullable String sourceExchangeKey;
        private @Nullable Integer sourcePrefetchCount;
        private @Nullable String sourceProtocol;
        private @Nullable String sourceQueue;
        private String sourceUri;
        public Builder() {}
        public Builder(ShovelInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ackMode = defaults.ackMode;
    	      this.addForwardHeaders = defaults.addForwardHeaders;
    	      this.deleteAfter = defaults.deleteAfter;
    	      this.destinationAddForwardHeaders = defaults.destinationAddForwardHeaders;
    	      this.destinationAddTimestampHeader = defaults.destinationAddTimestampHeader;
    	      this.destinationAddress = defaults.destinationAddress;
    	      this.destinationApplicationProperties = defaults.destinationApplicationProperties;
    	      this.destinationExchange = defaults.destinationExchange;
    	      this.destinationExchangeKey = defaults.destinationExchangeKey;
    	      this.destinationProperties = defaults.destinationProperties;
    	      this.destinationProtocol = defaults.destinationProtocol;
    	      this.destinationPublishProperties = defaults.destinationPublishProperties;
    	      this.destinationQueue = defaults.destinationQueue;
    	      this.destinationUri = defaults.destinationUri;
    	      this.prefetchCount = defaults.prefetchCount;
    	      this.reconnectDelay = defaults.reconnectDelay;
    	      this.sourceAddress = defaults.sourceAddress;
    	      this.sourceDeleteAfter = defaults.sourceDeleteAfter;
    	      this.sourceExchange = defaults.sourceExchange;
    	      this.sourceExchangeKey = defaults.sourceExchangeKey;
    	      this.sourcePrefetchCount = defaults.sourcePrefetchCount;
    	      this.sourceProtocol = defaults.sourceProtocol;
    	      this.sourceQueue = defaults.sourceQueue;
    	      this.sourceUri = defaults.sourceUri;
        }

        @CustomType.Setter
        public Builder ackMode(@Nullable String ackMode) {

            this.ackMode = ackMode;
            return this;
        }
        @CustomType.Setter
        public Builder addForwardHeaders(@Nullable Boolean addForwardHeaders) {

            this.addForwardHeaders = addForwardHeaders;
            return this;
        }
        @CustomType.Setter
        public Builder deleteAfter(@Nullable String deleteAfter) {

            this.deleteAfter = deleteAfter;
            return this;
        }
        @CustomType.Setter
        public Builder destinationAddForwardHeaders(@Nullable Boolean destinationAddForwardHeaders) {

            this.destinationAddForwardHeaders = destinationAddForwardHeaders;
            return this;
        }
        @CustomType.Setter
        public Builder destinationAddTimestampHeader(@Nullable Boolean destinationAddTimestampHeader) {

            this.destinationAddTimestampHeader = destinationAddTimestampHeader;
            return this;
        }
        @CustomType.Setter
        public Builder destinationAddress(@Nullable String destinationAddress) {

            this.destinationAddress = destinationAddress;
            return this;
        }
        @CustomType.Setter
        public Builder destinationApplicationProperties(@Nullable String destinationApplicationProperties) {

            this.destinationApplicationProperties = destinationApplicationProperties;
            return this;
        }
        @CustomType.Setter
        public Builder destinationExchange(@Nullable String destinationExchange) {

            this.destinationExchange = destinationExchange;
            return this;
        }
        @CustomType.Setter
        public Builder destinationExchangeKey(@Nullable String destinationExchangeKey) {

            this.destinationExchangeKey = destinationExchangeKey;
            return this;
        }
        @CustomType.Setter
        public Builder destinationProperties(@Nullable String destinationProperties) {

            this.destinationProperties = destinationProperties;
            return this;
        }
        @CustomType.Setter
        public Builder destinationProtocol(@Nullable String destinationProtocol) {

            this.destinationProtocol = destinationProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder destinationPublishProperties(@Nullable String destinationPublishProperties) {

            this.destinationPublishProperties = destinationPublishProperties;
            return this;
        }
        @CustomType.Setter
        public Builder destinationQueue(@Nullable String destinationQueue) {

            this.destinationQueue = destinationQueue;
            return this;
        }
        @CustomType.Setter
        public Builder destinationUri(String destinationUri) {
            if (destinationUri == null) {
              throw new MissingRequiredPropertyException("ShovelInfo", "destinationUri");
            }
            this.destinationUri = destinationUri;
            return this;
        }
        @CustomType.Setter
        public Builder prefetchCount(@Nullable Integer prefetchCount) {

            this.prefetchCount = prefetchCount;
            return this;
        }
        @CustomType.Setter
        public Builder reconnectDelay(@Nullable Integer reconnectDelay) {

            this.reconnectDelay = reconnectDelay;
            return this;
        }
        @CustomType.Setter
        public Builder sourceAddress(@Nullable String sourceAddress) {

            this.sourceAddress = sourceAddress;
            return this;
        }
        @CustomType.Setter
        public Builder sourceDeleteAfter(@Nullable String sourceDeleteAfter) {

            this.sourceDeleteAfter = sourceDeleteAfter;
            return this;
        }
        @CustomType.Setter
        public Builder sourceExchange(@Nullable String sourceExchange) {

            this.sourceExchange = sourceExchange;
            return this;
        }
        @CustomType.Setter
        public Builder sourceExchangeKey(@Nullable String sourceExchangeKey) {

            this.sourceExchangeKey = sourceExchangeKey;
            return this;
        }
        @CustomType.Setter
        public Builder sourcePrefetchCount(@Nullable Integer sourcePrefetchCount) {

            this.sourcePrefetchCount = sourcePrefetchCount;
            return this;
        }
        @CustomType.Setter
        public Builder sourceProtocol(@Nullable String sourceProtocol) {

            this.sourceProtocol = sourceProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder sourceQueue(@Nullable String sourceQueue) {

            this.sourceQueue = sourceQueue;
            return this;
        }
        @CustomType.Setter
        public Builder sourceUri(String sourceUri) {
            if (sourceUri == null) {
              throw new MissingRequiredPropertyException("ShovelInfo", "sourceUri");
            }
            this.sourceUri = sourceUri;
            return this;
        }
        public ShovelInfo build() {
            final var _resultValue = new ShovelInfo();
            _resultValue.ackMode = ackMode;
            _resultValue.addForwardHeaders = addForwardHeaders;
            _resultValue.deleteAfter = deleteAfter;
            _resultValue.destinationAddForwardHeaders = destinationAddForwardHeaders;
            _resultValue.destinationAddTimestampHeader = destinationAddTimestampHeader;
            _resultValue.destinationAddress = destinationAddress;
            _resultValue.destinationApplicationProperties = destinationApplicationProperties;
            _resultValue.destinationExchange = destinationExchange;
            _resultValue.destinationExchangeKey = destinationExchangeKey;
            _resultValue.destinationProperties = destinationProperties;
            _resultValue.destinationProtocol = destinationProtocol;
            _resultValue.destinationPublishProperties = destinationPublishProperties;
            _resultValue.destinationQueue = destinationQueue;
            _resultValue.destinationUri = destinationUri;
            _resultValue.prefetchCount = prefetchCount;
            _resultValue.reconnectDelay = reconnectDelay;
            _resultValue.sourceAddress = sourceAddress;
            _resultValue.sourceDeleteAfter = sourceDeleteAfter;
            _resultValue.sourceExchange = sourceExchange;
            _resultValue.sourceExchangeKey = sourceExchangeKey;
            _resultValue.sourcePrefetchCount = sourcePrefetchCount;
            _resultValue.sourceProtocol = sourceProtocol;
            _resultValue.sourceQueue = sourceQueue;
            _resultValue.sourceUri = sourceUri;
            return _resultValue;
        }
    }
}
