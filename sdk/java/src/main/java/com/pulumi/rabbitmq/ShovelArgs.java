// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.rabbitmq.inputs.ShovelInfoArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ShovelArgs extends com.pulumi.resources.ResourceArgs {

    public static final ShovelArgs Empty = new ShovelArgs();

    /**
     * The settings of the dynamic shovel. The structure is
     * described below.
     * 
     */
    @Import(name="info", required=true)
    private Output<ShovelInfoArgs> info;

    /**
     * @return The settings of the dynamic shovel. The structure is
     * described below.
     * 
     */
    public Output<ShovelInfoArgs> info() {
        return this.info;
    }

    /**
     * The shovel name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The shovel name.
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

    private ShovelArgs() {}

    private ShovelArgs(ShovelArgs $) {
        this.info = $.info;
        this.name = $.name;
        this.vhost = $.vhost;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ShovelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ShovelArgs $;

        public Builder() {
            $ = new ShovelArgs();
        }

        public Builder(ShovelArgs defaults) {
            $ = new ShovelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param info The settings of the dynamic shovel. The structure is
         * described below.
         * 
         * @return builder
         * 
         */
        public Builder info(Output<ShovelInfoArgs> info) {
            $.info = info;
            return this;
        }

        /**
         * @param info The settings of the dynamic shovel. The structure is
         * described below.
         * 
         * @return builder
         * 
         */
        public Builder info(ShovelInfoArgs info) {
            return info(Output.of(info));
        }

        /**
         * @param name The shovel name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The shovel name.
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

        public ShovelArgs build() {
            if ($.info == null) {
                throw new MissingRequiredPropertyException("ShovelArgs", "info");
            }
            if ($.vhost == null) {
                throw new MissingRequiredPropertyException("ShovelArgs", "vhost");
            }
            return $;
        }
    }

}
