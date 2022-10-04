// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.rabbitmq.inputs.FederationUpstreamDefinitionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FederationUpstreamArgs extends com.pulumi.resources.ResourceArgs {

    public static final FederationUpstreamArgs Empty = new FederationUpstreamArgs();

    /**
     * The configuration of the federation upstream. The structure is described below.
     * 
     */
    @Import(name="definition", required=true)
    private Output<FederationUpstreamDefinitionArgs> definition;

    /**
     * @return The configuration of the federation upstream. The structure is described below.
     * 
     */
    public Output<FederationUpstreamDefinitionArgs> definition() {
        return this.definition;
    }

    /**
     * The name of the federation upstream.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the federation upstream.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The vhost to create the resource in.
     * 
     */
    @Import(name="vhost", required=true)
    private Output<String> vhost;

    /**
     * @return The vhost to create the resource in.
     * 
     */
    public Output<String> vhost() {
        return this.vhost;
    }

    private FederationUpstreamArgs() {}

    private FederationUpstreamArgs(FederationUpstreamArgs $) {
        this.definition = $.definition;
        this.name = $.name;
        this.vhost = $.vhost;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FederationUpstreamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FederationUpstreamArgs $;

        public Builder() {
            $ = new FederationUpstreamArgs();
        }

        public Builder(FederationUpstreamArgs defaults) {
            $ = new FederationUpstreamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param definition The configuration of the federation upstream. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder definition(Output<FederationUpstreamDefinitionArgs> definition) {
            $.definition = definition;
            return this;
        }

        /**
         * @param definition The configuration of the federation upstream. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder definition(FederationUpstreamDefinitionArgs definition) {
            return definition(Output.of(definition));
        }

        /**
         * @param name The name of the federation upstream.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the federation upstream.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param vhost The vhost to create the resource in.
         * 
         * @return builder
         * 
         */
        public Builder vhost(Output<String> vhost) {
            $.vhost = vhost;
            return this;
        }

        /**
         * @param vhost The vhost to create the resource in.
         * 
         * @return builder
         * 
         */
        public Builder vhost(String vhost) {
            return vhost(Output.of(vhost));
        }

        public FederationUpstreamArgs build() {
            $.definition = Objects.requireNonNull($.definition, "expected parameter 'definition' to be non-null");
            $.vhost = Objects.requireNonNull($.vhost, "expected parameter 'vhost' to be non-null");
            return $;
        }
    }

}