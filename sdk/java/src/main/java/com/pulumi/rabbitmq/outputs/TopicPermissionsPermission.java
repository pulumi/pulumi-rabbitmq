// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class TopicPermissionsPermission {
    /**
     * @return The exchange to set the permissions for.
     * 
     */
    private String exchange;
    /**
     * @return The &#34;read&#34; ACL.
     * 
     */
    private String read;
    /**
     * @return The &#34;write&#34; ACL.
     * 
     */
    private String write;

    private TopicPermissionsPermission() {}
    /**
     * @return The exchange to set the permissions for.
     * 
     */
    public String exchange() {
        return this.exchange;
    }
    /**
     * @return The &#34;read&#34; ACL.
     * 
     */
    public String read() {
        return this.read;
    }
    /**
     * @return The &#34;write&#34; ACL.
     * 
     */
    public String write() {
        return this.write;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TopicPermissionsPermission defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String exchange;
        private String read;
        private String write;
        public Builder() {}
        public Builder(TopicPermissionsPermission defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.exchange = defaults.exchange;
    	      this.read = defaults.read;
    	      this.write = defaults.write;
        }

        @CustomType.Setter
        public Builder exchange(String exchange) {
            if (exchange == null) {
              throw new MissingRequiredPropertyException("TopicPermissionsPermission", "exchange");
            }
            this.exchange = exchange;
            return this;
        }
        @CustomType.Setter
        public Builder read(String read) {
            if (read == null) {
              throw new MissingRequiredPropertyException("TopicPermissionsPermission", "read");
            }
            this.read = read;
            return this;
        }
        @CustomType.Setter
        public Builder write(String write) {
            if (write == null) {
              throw new MissingRequiredPropertyException("TopicPermissionsPermission", "write");
            }
            this.write = write;
            return this;
        }
        public TopicPermissionsPermission build() {
            final var _resultValue = new TopicPermissionsPermission();
            _resultValue.exchange = exchange;
            _resultValue.read = read;
            _resultValue.write = write;
            return _resultValue;
        }
    }
}
