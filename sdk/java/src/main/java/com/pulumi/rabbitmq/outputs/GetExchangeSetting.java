// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetExchangeSetting {
    private @Nullable Map<String,Object> arguments;
    private @Nullable Boolean autoDelete;
    private @Nullable Boolean durable;
    private String type;

    private GetExchangeSetting() {}
    public Map<String,Object> arguments() {
        return this.arguments == null ? Map.of() : this.arguments;
    }
    public Optional<Boolean> autoDelete() {
        return Optional.ofNullable(this.autoDelete);
    }
    public Optional<Boolean> durable() {
        return Optional.ofNullable(this.durable);
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetExchangeSetting defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,Object> arguments;
        private @Nullable Boolean autoDelete;
        private @Nullable Boolean durable;
        private String type;
        public Builder() {}
        public Builder(GetExchangeSetting defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arguments = defaults.arguments;
    	      this.autoDelete = defaults.autoDelete;
    	      this.durable = defaults.durable;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder arguments(@Nullable Map<String,Object> arguments) {
            this.arguments = arguments;
            return this;
        }
        @CustomType.Setter
        public Builder autoDelete(@Nullable Boolean autoDelete) {
            this.autoDelete = autoDelete;
            return this;
        }
        @CustomType.Setter
        public Builder durable(@Nullable Boolean durable) {
            this.durable = durable;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetExchangeSetting build() {
            final var o = new GetExchangeSetting();
            o.arguments = arguments;
            o.autoDelete = autoDelete;
            o.durable = durable;
            o.type = type;
            return o;
        }
    }
}
