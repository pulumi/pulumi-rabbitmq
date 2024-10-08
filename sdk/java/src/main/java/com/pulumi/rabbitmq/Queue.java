// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.QueueArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.QueueState;
import com.pulumi.rabbitmq.outputs.QueueSettings;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.Queue`` resource creates and manages a queue.
 * 
 * ## Example Usage
 * 
 * ### Basic Example
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.rabbitmq.VHost;
 * import com.pulumi.rabbitmq.VHostArgs;
 * import com.pulumi.rabbitmq.Permissions;
 * import com.pulumi.rabbitmq.PermissionsArgs;
 * import com.pulumi.rabbitmq.inputs.PermissionsPermissionsArgs;
 * import com.pulumi.rabbitmq.Queue;
 * import com.pulumi.rabbitmq.QueueArgs;
 * import com.pulumi.rabbitmq.inputs.QueueSettingsArgs;
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
 *         var test = new VHost("test", VHostArgs.builder()
 *             .name("test")
 *             .build());
 * 
 *         var guest = new Permissions("guest", PermissionsArgs.builder()
 *             .user("guest")
 *             .vhost(test.name())
 *             .permissions(PermissionsPermissionsArgs.builder()
 *                 .configure(".*")
 *                 .write(".*")
 *                 .read(".*")
 *                 .build())
 *             .build());
 * 
 *         var testQueue = new Queue("testQueue", QueueArgs.builder()
 *             .name("test")
 *             .vhost(guest.vhost())
 *             .settings(QueueSettingsArgs.builder()
 *                 .durable(false)
 *                 .autoDelete(true)
 *                 .arguments(Map.of("x-queue-type", "quorum"))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Example With JSON Arguments
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.rabbitmq.VHost;
 * import com.pulumi.rabbitmq.VHostArgs;
 * import com.pulumi.rabbitmq.Permissions;
 * import com.pulumi.rabbitmq.PermissionsArgs;
 * import com.pulumi.rabbitmq.inputs.PermissionsPermissionsArgs;
 * import com.pulumi.rabbitmq.Queue;
 * import com.pulumi.rabbitmq.QueueArgs;
 * import com.pulumi.rabbitmq.inputs.QueueSettingsArgs;
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
 *         final var config = ctx.config();
 *         final var arguments = config.get("arguments").orElse("""
 * {
 *   "x-message-ttl": 5000
 * }
 *         """);
 *         var test = new VHost("test", VHostArgs.builder()
 *             .name("test")
 *             .build());
 * 
 *         var guest = new Permissions("guest", PermissionsArgs.builder()
 *             .user("guest")
 *             .vhost(test.name())
 *             .permissions(PermissionsPermissionsArgs.builder()
 *                 .configure(".*")
 *                 .write(".*")
 *                 .read(".*")
 *                 .build())
 *             .build());
 * 
 *         var testQueue = new Queue("testQueue", QueueArgs.builder()
 *             .name("test")
 *             .vhost(guest.vhost())
 *             .settings(QueueSettingsArgs.builder()
 *                 .durable(false)
 *                 .autoDelete(true)
 *                 .argumentsJson(arguments)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Queues can be imported using the `id` which is composed of `name{@literal @}vhost`. E.g.
 * 
 * ```sh
 * $ pulumi import rabbitmq:index/queue:Queue test name{@literal @}vhost
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/queue:Queue")
public class Queue extends com.pulumi.resources.CustomResource {
    /**
     * The name of the queue.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the queue.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The settings of the queue. The structure is
     * described below.
     * 
     */
    @Export(name="settings", refs={QueueSettings.class}, tree="[0]")
    private Output<QueueSettings> settings;

    /**
     * @return The settings of the queue. The structure is
     * described below.
     * 
     */
    public Output<QueueSettings> settings() {
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
    public Queue(java.lang.String name) {
        this(name, QueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Queue(java.lang.String name, QueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Queue(java.lang.String name, QueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/queue:Queue", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Queue(java.lang.String name, Output<java.lang.String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/queue:Queue", name, state, makeResourceOptions(options, id), false);
    }

    private static QueueArgs makeArgs(QueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? QueueArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Queue get(java.lang.String name, Output<java.lang.String> id, @Nullable QueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Queue(name, id, state, options);
    }
}
