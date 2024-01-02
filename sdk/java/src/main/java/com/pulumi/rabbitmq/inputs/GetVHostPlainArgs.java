// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetVHostPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVHostPlainArgs Empty = new GetVHostPlainArgs();

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    private GetVHostPlainArgs() {}

    private GetVHostPlainArgs(GetVHostPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVHostPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVHostPlainArgs $;

        public Builder() {
            $ = new GetVHostPlainArgs();
        }

        public Builder(GetVHostPlainArgs defaults) {
            $ = new GetVHostPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetVHostPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetVHostPlainArgs", "name");
            }
            return $;
        }
    }

}
