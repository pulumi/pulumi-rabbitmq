// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetExchangePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetExchangePlainArgs Empty = new GetExchangePlainArgs();

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="vhost")
    private @Nullable String vhost;

    public Optional<String> vhost() {
        return Optional.ofNullable(this.vhost);
    }

    private GetExchangePlainArgs() {}

    private GetExchangePlainArgs(GetExchangePlainArgs $) {
        this.name = $.name;
        this.vhost = $.vhost;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetExchangePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetExchangePlainArgs $;

        public Builder() {
            $ = new GetExchangePlainArgs();
        }

        public Builder(GetExchangePlainArgs defaults) {
            $ = new GetExchangePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder vhost(@Nullable String vhost) {
            $.vhost = vhost;
            return this;
        }

        public GetExchangePlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetExchangePlainArgs", "name");
            }
            return $;
        }
    }

}
