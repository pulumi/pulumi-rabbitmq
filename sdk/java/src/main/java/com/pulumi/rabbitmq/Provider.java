// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.ProviderArgs;
import com.pulumi.rabbitmq.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the rabbitmq package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:rabbitmq")
public class Provider extends com.pulumi.resources.ProviderResource {
    @Export(name="cacertFile", type=String.class, parameters={})
    private Output</* @Nullable */ String> cacertFile;

    public Output<Optional<String>> cacertFile() {
        return Codegen.optional(this.cacertFile);
    }
    @Export(name="clientcertFile", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientcertFile;

    public Output<Optional<String>> clientcertFile() {
        return Codegen.optional(this.clientcertFile);
    }
    @Export(name="clientkeyFile", type=String.class, parameters={})
    private Output</* @Nullable */ String> clientkeyFile;

    public Output<Optional<String>> clientkeyFile() {
        return Codegen.optional(this.clientkeyFile);
    }
    @Export(name="endpoint", type=String.class, parameters={})
    private Output<String> endpoint;

    public Output<String> endpoint() {
        return this.endpoint;
    }
    @Export(name="password", type=String.class, parameters={})
    private Output<String> password;

    public Output<String> password() {
        return this.password;
    }
    @Export(name="proxy", type=String.class, parameters={})
    private Output</* @Nullable */ String> proxy;

    public Output<Optional<String>> proxy() {
        return Codegen.optional(this.proxy);
    }
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
