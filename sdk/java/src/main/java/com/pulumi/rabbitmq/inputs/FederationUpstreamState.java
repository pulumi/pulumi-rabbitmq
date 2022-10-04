// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.rabbitmq.inputs.FederationUpstreamDefinitionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FederationUpstreamState extends com.pulumi.resources.ResourceArgs {

    public static final FederationUpstreamState Empty = new FederationUpstreamState();

    /**
     * Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
     * 
     */
    @Import(name="component")
    private @Nullable Output<String> component;

    /**
     * @return Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
     * 
     */
    public Optional<Output<String>> component() {
        return Optional.ofNullable(this.component);
    }

    /**
     * The configuration of the federation upstream. The structure is described below.
     * 
     */
    @Import(name="definition")
    private @Nullable Output<FederationUpstreamDefinitionArgs> definition;

    /**
     * @return The configuration of the federation upstream. The structure is described below.
     * 
     */
    public Optional<Output<FederationUpstreamDefinitionArgs>> definition() {
        return Optional.ofNullable(this.definition);
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
    @Import(name="vhost")
    private @Nullable Output<String> vhost;

    /**
     * @return The vhost to create the resource in.
     * 
     */
    public Optional<Output<String>> vhost() {
        return Optional.ofNullable(this.vhost);
    }

    private FederationUpstreamState() {}

    private FederationUpstreamState(FederationUpstreamState $) {
        this.component = $.component;
        this.definition = $.definition;
        this.name = $.name;
        this.vhost = $.vhost;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FederationUpstreamState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FederationUpstreamState $;

        public Builder() {
            $ = new FederationUpstreamState();
        }

        public Builder(FederationUpstreamState defaults) {
            $ = new FederationUpstreamState(Objects.requireNonNull(defaults));
        }

        /**
         * @param component Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
         * 
         * @return builder
         * 
         */
        public Builder component(@Nullable Output<String> component) {
            $.component = component;
            return this;
        }

        /**
         * @param component Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
         * 
         * @return builder
         * 
         */
        public Builder component(String component) {
            return component(Output.of(component));
        }

        /**
         * @param definition The configuration of the federation upstream. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder definition(@Nullable Output<FederationUpstreamDefinitionArgs> definition) {
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
        public Builder vhost(@Nullable Output<String> vhost) {
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

        public FederationUpstreamState build() {
            return $;
        }
    }

}