// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserPlainArgs Empty = new GetUserPlainArgs();

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    private GetUserPlainArgs() {}

    private GetUserPlainArgs(GetUserPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserPlainArgs $;

        public Builder() {
            $ = new GetUserPlainArgs();
        }

        public Builder(GetUserPlainArgs defaults) {
            $ = new GetUserPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetUserPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}