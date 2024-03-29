// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    @Import(name="cacertFile")
    private @Nullable Output<String> cacertFile;

    public Optional<Output<String>> cacertFile() {
        return Optional.ofNullable(this.cacertFile);
    }

    @Import(name="clientcertFile")
    private @Nullable Output<String> clientcertFile;

    public Optional<Output<String>> clientcertFile() {
        return Optional.ofNullable(this.clientcertFile);
    }

    @Import(name="clientkeyFile")
    private @Nullable Output<String> clientkeyFile;

    public Optional<Output<String>> clientkeyFile() {
        return Optional.ofNullable(this.clientkeyFile);
    }

    @Import(name="endpoint", required=true)
    private Output<String> endpoint;

    public Output<String> endpoint() {
        return this.endpoint;
    }

    @Import(name="insecure", json=true)
    private @Nullable Output<Boolean> insecure;

    public Optional<Output<Boolean>> insecure() {
        return Optional.ofNullable(this.insecure);
    }

    @Import(name="password", required=true)
    private Output<String> password;

    public Output<String> password() {
        return this.password;
    }

    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    @Import(name="username", required=true)
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.cacertFile = $.cacertFile;
        this.clientcertFile = $.clientcertFile;
        this.clientkeyFile = $.clientkeyFile;
        this.endpoint = $.endpoint;
        this.insecure = $.insecure;
        this.password = $.password;
        this.proxy = $.proxy;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        public Builder cacertFile(@Nullable Output<String> cacertFile) {
            $.cacertFile = cacertFile;
            return this;
        }

        public Builder cacertFile(String cacertFile) {
            return cacertFile(Output.of(cacertFile));
        }

        public Builder clientcertFile(@Nullable Output<String> clientcertFile) {
            $.clientcertFile = clientcertFile;
            return this;
        }

        public Builder clientcertFile(String clientcertFile) {
            return clientcertFile(Output.of(clientcertFile));
        }

        public Builder clientkeyFile(@Nullable Output<String> clientkeyFile) {
            $.clientkeyFile = clientkeyFile;
            return this;
        }

        public Builder clientkeyFile(String clientkeyFile) {
            return clientkeyFile(Output.of(clientkeyFile));
        }

        public Builder endpoint(Output<String> endpoint) {
            $.endpoint = endpoint;
            return this;
        }

        public Builder endpoint(String endpoint) {
            return endpoint(Output.of(endpoint));
        }

        public Builder insecure(@Nullable Output<Boolean> insecure) {
            $.insecure = insecure;
            return this;
        }

        public Builder insecure(Boolean insecure) {
            return insecure(Output.of(insecure));
        }

        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderArgs build() {
            $.cacertFile = Codegen.stringProp("cacertFile").output().arg($.cacertFile).env("RABBITMQ_CACERT").getNullable();
            if ($.endpoint == null) {
                throw new MissingRequiredPropertyException("ProviderArgs", "endpoint");
            }
            $.insecure = Codegen.booleanProp("insecure").output().arg($.insecure).env("RABBITMQ_INSECURE").getNullable();
            if ($.password == null) {
                throw new MissingRequiredPropertyException("ProviderArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("ProviderArgs", "username");
            }
            return $;
        }
    }

}
