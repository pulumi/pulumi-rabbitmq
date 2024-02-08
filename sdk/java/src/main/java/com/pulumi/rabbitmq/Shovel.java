// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.ShovelArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.ShovelState;
import com.pulumi.rabbitmq.outputs.ShovelInfo;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.Shovel`` resource creates and manages a dynamic shovel.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.rabbitmq.VHost;
 * import com.pulumi.rabbitmq.Exchange;
 * import com.pulumi.rabbitmq.ExchangeArgs;
 * import com.pulumi.rabbitmq.inputs.ExchangeSettingsArgs;
 * import com.pulumi.rabbitmq.Queue;
 * import com.pulumi.rabbitmq.QueueArgs;
 * import com.pulumi.rabbitmq.inputs.QueueSettingsArgs;
 * import com.pulumi.rabbitmq.Shovel;
 * import com.pulumi.rabbitmq.ShovelArgs;
 * import com.pulumi.rabbitmq.inputs.ShovelInfoArgs;
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
 *         var testExchange = new Exchange(&#34;testExchange&#34;, ExchangeArgs.builder()        
 *             .settings(ExchangeSettingsArgs.builder()
 *                 .autoDelete(true)
 *                 .durable(false)
 *                 .type(&#34;fanout&#34;)
 *                 .build())
 *             .vhost(testVHost.name())
 *             .build());
 * 
 *         var testQueue = new Queue(&#34;testQueue&#34;, QueueArgs.builder()        
 *             .settings(QueueSettingsArgs.builder()
 *                 .autoDelete(true)
 *                 .durable(false)
 *                 .build())
 *             .vhost(testVHost.name())
 *             .build());
 * 
 *         var shovelTest = new Shovel(&#34;shovelTest&#34;, ShovelArgs.builder()        
 *             .info(ShovelInfoArgs.builder()
 *                 .destinationQueue(testQueue.name())
 *                 .destinationUri(&#34;amqp:///test&#34;)
 *                 .sourceExchange(testExchange.name())
 *                 .sourceExchangeKey(&#34;test&#34;)
 *                 .sourceUri(&#34;amqp:///test&#34;)
 *                 .build())
 *             .vhost(testVHost.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Shovels can be imported using the `name` and `vhost`
 * 
 *  E.g.
 * 
 * ```sh
 * $ pulumi import rabbitmq:index/shovel:Shovel test shovelTest@test
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/shovel:Shovel")
public class Shovel extends com.pulumi.resources.CustomResource {
    /**
     * The settings of the dynamic shovel. The structure is
     * described below.
     * 
     */
    @Export(name="info", refs={ShovelInfo.class}, tree="[0]")
    private Output<ShovelInfo> info;

    /**
     * @return The settings of the dynamic shovel. The structure is
     * described below.
     * 
     */
    public Output<ShovelInfo> info() {
        return this.info;
    }
    /**
     * The shovel name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The shovel name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The vhost to create the resource in.
     * 
     */
    @Export(name="vhost", refs={String.class}, tree="[0]")
    private Output<String> vhost;

    /**
     * @return The vhost to create the resource in.
     * 
     */
    public Output<String> vhost() {
        return this.vhost;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Shovel(String name) {
        this(name, ShovelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Shovel(String name, ShovelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Shovel(String name, ShovelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/shovel:Shovel", name, args == null ? ShovelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Shovel(String name, Output<String> id, @Nullable ShovelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/shovel:Shovel", name, state, makeResourceOptions(options, id));
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
    public static Shovel get(String name, Output<String> id, @Nullable ShovelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Shovel(name, id, state, options);
    }
}
