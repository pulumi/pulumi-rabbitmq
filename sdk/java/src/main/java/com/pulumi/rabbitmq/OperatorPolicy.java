// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.OperatorPolicyArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.OperatorPolicyState;
import com.pulumi.rabbitmq.outputs.OperatorPolicyPolicy;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.OperatorPolicy`` resource creates and manages operator policies for queues.
 * 
 * ## Example Usage
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
 * import com.pulumi.rabbitmq.OperatorPolicy;
 * import com.pulumi.rabbitmq.OperatorPolicyArgs;
 * import com.pulumi.rabbitmq.inputs.OperatorPolicyPolicyArgs;
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
 *         var testOperatorPolicy = new OperatorPolicy("testOperatorPolicy", OperatorPolicyArgs.builder()
 *             .name("test")
 *             .vhost(guest.vhost())
 *             .policy(OperatorPolicyPolicyArgs.builder()
 *                 .pattern(".*")
 *                 .priority(0)
 *                 .applyTo("queues")
 *                 .definition(Map.ofEntries(
 *                     Map.entry("message-ttl", "3600000"),
 *                     Map.entry("expires", "1800000")
 *                 ))
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
 * Operator policies can be imported using the `id` which is composed of `name{@literal @}vhost`.
 * 
 * E.g.
 * 
 * ```sh
 * $ pulumi import rabbitmq:index/operatorPolicy:OperatorPolicy test name{@literal @}vhost
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/operatorPolicy:OperatorPolicy")
public class OperatorPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The name of the operator policy.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the operator policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The settings of the operator policy. The structure is
     * described below.
     * 
     */
    @Export(name="policy", refs={OperatorPolicyPolicy.class}, tree="[0]")
    private Output<OperatorPolicyPolicy> policy;

    /**
     * @return The settings of the operator policy. The structure is
     * described below.
     * 
     */
    public Output<OperatorPolicyPolicy> policy() {
        return this.policy;
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
    public OperatorPolicy(java.lang.String name) {
        this(name, OperatorPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OperatorPolicy(java.lang.String name, OperatorPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OperatorPolicy(java.lang.String name, OperatorPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/operatorPolicy:OperatorPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private OperatorPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable OperatorPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/operatorPolicy:OperatorPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static OperatorPolicyArgs makeArgs(OperatorPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? OperatorPolicyArgs.Empty : args;
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
    public static OperatorPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable OperatorPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OperatorPolicy(name, id, state, options);
    }
}
