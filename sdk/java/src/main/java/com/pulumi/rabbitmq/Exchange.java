// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.ExchangeArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.ExchangeState;
import com.pulumi.rabbitmq.outputs.ExchangeSettings;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.Exchange`` resource creates and manages an exchange.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.rabbitmq.VHost;
 * import com.pulumi.rabbitmq.Permissions;
 * import com.pulumi.rabbitmq.PermissionsArgs;
 * import com.pulumi.rabbitmq.inputs.PermissionsPermissionsArgs;
 * import com.pulumi.rabbitmq.Exchange;
 * import com.pulumi.rabbitmq.ExchangeArgs;
 * import com.pulumi.rabbitmq.inputs.ExchangeSettingsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var testVHost = new VHost(&#34;testVHost&#34;);
 * 
 *         var guest = new Permissions(&#34;guest&#34;, PermissionsArgs.builder()        
 *             .permissions(PermissionsPermissionsArgs.builder()
 *                 .configure(&#34;.*&#34;)
 *                 .read(&#34;.*&#34;)
 *                 .write(&#34;.*&#34;)
 *                 .build())
 *             .user(&#34;guest&#34;)
 *             .vhost(testVHost.name())
 *             .build());
 * 
 *         var testExchange = new Exchange(&#34;testExchange&#34;, ExchangeArgs.builder()        
 *             .settings(ExchangeSettingsArgs.builder()
 *                 .autoDelete(true)
 *                 .durable(false)
 *                 .type(&#34;fanout&#34;)
 *                 .build())
 *             .vhost(guest.vhost())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Exchanges can be imported using the `id` which is composed of  `name@vhost`.
 * 
 * E.g.
 * 
 * ```sh
 * $ pulumi import rabbitmq:index/exchange:Exchange test test@vhost
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/exchange:Exchange")
public class Exchange extends com.pulumi.resources.CustomResource {
    /**
     * The name of the exchange.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the exchange.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The settings of the exchange. The structure is
     * described below.
     * 
     */
    @Export(name="settings", refs={ExchangeSettings.class}, tree="[0]")
    private Output<ExchangeSettings> settings;

    /**
     * @return The settings of the exchange. The structure is
     * described below.
     * 
     */
    public Output<ExchangeSettings> settings() {
        return this.settings;
    }
    /**
     * The vhost to create the resource in.
     * 
     */
    @Export(name="vhost", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vhost;

    /**
     * @return The vhost to create the resource in.
     * 
     */
    public Output<Optional<String>> vhost() {
        return Codegen.optional(this.vhost);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Exchange(String name) {
        this(name, ExchangeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Exchange(String name, ExchangeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Exchange(String name, ExchangeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/exchange:Exchange", name, args == null ? ExchangeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Exchange(String name, Output<String> id, @Nullable ExchangeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/exchange:Exchange", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Exchange get(String name, Output<String> id, @Nullable ExchangeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Exchange(name, id, state, options);
    }
}
