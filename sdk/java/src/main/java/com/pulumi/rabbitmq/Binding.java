// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.BindingArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.BindingState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.Binding`` resource creates and manages a binding relationship
 * between a queue an exchange.
 * 
 * ## Example Usage
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
 * import com.pulumi.rabbitmq.Queue;
 * import com.pulumi.rabbitmq.QueueArgs;
 * import com.pulumi.rabbitmq.inputs.QueueSettingsArgs;
 * import com.pulumi.rabbitmq.Binding;
 * import com.pulumi.rabbitmq.BindingArgs;
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
 *         var testQueue = new Queue(&#34;testQueue&#34;, QueueArgs.builder()        
 *             .settings(QueueSettingsArgs.builder()
 *                 .autoDelete(false)
 *                 .durable(true)
 *                 .build())
 *             .vhost(guest.vhost())
 *             .build());
 * 
 *         var testBinding = new Binding(&#34;testBinding&#34;, BindingArgs.builder()        
 *             .destination(testQueue.name())
 *             .destinationType(&#34;queue&#34;)
 *             .routingKey(&#34;#&#34;)
 *             .source(testExchange.name())
 *             .vhost(testVHost.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Bindings can be imported using the `id` which is composed of
 * 
 *  `vhost/source/destination/destination_type/properties_key`. E.g.
 * 
 * ```sh
 *  $ pulumi import rabbitmq:index/binding:Binding test test/test/test/queue/%23
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/binding:Binding")
public class Binding extends com.pulumi.resources.CustomResource {
    /**
     * Additional key/value arguments for the binding.
     * 
     */
    @Export(name="arguments", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> arguments;

    /**
     * @return Additional key/value arguments for the binding.
     * 
     */
    public Output<Optional<Map<String,Object>>> arguments() {
        return Codegen.optional(this.arguments);
    }
    @Export(name="argumentsJson", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> argumentsJson;

    public Output<Optional<String>> argumentsJson() {
        return Codegen.optional(this.argumentsJson);
    }
    /**
     * The destination queue or exchange.
     * 
     */
    @Export(name="destination", refs={String.class}, tree="[0]")
    private Output<String> destination;

    /**
     * @return The destination queue or exchange.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }
    /**
     * The type of destination (queue or exchange).
     * 
     */
    @Export(name="destinationType", refs={String.class}, tree="[0]")
    private Output<String> destinationType;

    /**
     * @return The type of destination (queue or exchange).
     * 
     */
    public Output<String> destinationType() {
        return this.destinationType;
    }
    /**
     * A unique key to refer to the binding.
     * 
     */
    @Export(name="propertiesKey", refs={String.class}, tree="[0]")
    private Output<String> propertiesKey;

    /**
     * @return A unique key to refer to the binding.
     * 
     */
    public Output<String> propertiesKey() {
        return this.propertiesKey;
    }
    /**
     * A routing key for the binding.
     * 
     */
    @Export(name="routingKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> routingKey;

    /**
     * @return A routing key for the binding.
     * 
     */
    public Output<Optional<String>> routingKey() {
        return Codegen.optional(this.routingKey);
    }
    /**
     * The source exchange.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output<String> source;

    /**
     * @return The source exchange.
     * 
     */
    public Output<String> source() {
        return this.source;
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
    public Binding(String name) {
        this(name, BindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Binding(String name, BindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Binding(String name, BindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/binding:Binding", name, args == null ? BindingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Binding(String name, Output<String> id, @Nullable BindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/binding:Binding", name, state, makeResourceOptions(options, id));
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
    public static Binding get(String name, Output<String> id, @Nullable BindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Binding(name, id, state, options);
    }
}
