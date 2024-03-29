// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.rabbitmq.inputs.ExchangeSettingsArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExchangeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExchangeArgs Empty = new ExchangeArgs();

    /**
     * The name of the exchange.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the exchange.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The settings of the exchange. The structure is
     * described below.
     * 
     */
    @Import(name="settings", required=true)
    private Output<ExchangeSettingsArgs> settings;

    /**
     * @return The settings of the exchange. The structure is
     * described below.
     * 
     */
    public Output<ExchangeSettingsArgs> settings() {
        return this.settings;
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

    private ExchangeArgs() {}

    private ExchangeArgs(ExchangeArgs $) {
        this.name = $.name;
        this.settings = $.settings;
        this.vhost = $.vhost;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExchangeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExchangeArgs $;

        public Builder() {
            $ = new ExchangeArgs();
        }

        public Builder(ExchangeArgs defaults) {
            $ = new ExchangeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the exchange.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the exchange.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param settings The settings of the exchange. The structure is
         * described below.
         * 
         * @return builder
         * 
         */
        public Builder settings(Output<ExchangeSettingsArgs> settings) {
            $.settings = settings;
            return this;
        }

        /**
         * @param settings The settings of the exchange. The structure is
         * described below.
         * 
         * @return builder
         * 
         */
        public Builder settings(ExchangeSettingsArgs settings) {
            return settings(Output.of(settings));
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

        public ExchangeArgs build() {
            if ($.settings == null) {
                throw new MissingRequiredPropertyException("ExchangeArgs", "settings");
            }
            return $;
        }
    }

}
