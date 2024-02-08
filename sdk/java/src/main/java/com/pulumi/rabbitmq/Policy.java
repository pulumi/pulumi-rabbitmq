// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.PolicyArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.PolicyState;
import com.pulumi.rabbitmq.outputs.PolicyPolicy;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.Policy`` resource creates and manages policies for exchanges
 * and queues.
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
 * import com.pulumi.rabbitmq.Policy;
 * import com.pulumi.rabbitmq.PolicyArgs;
 * import com.pulumi.rabbitmq.inputs.PolicyPolicyArgs;
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
 *         var testPolicy = new Policy(&#34;testPolicy&#34;, PolicyArgs.builder()        
 *             .policy(PolicyPolicyArgs.builder()
 *                 .applyTo(&#34;all&#34;)
 *                 .definition(Map.of(&#34;ha-mode&#34;, &#34;all&#34;))
 *                 .pattern(&#34;.*&#34;)
 *                 .priority(0)
 *                 .build())
 *             .vhost(guest.vhost())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Policies can be imported using the `id` which is composed of `name@vhost`.
 * 
 *  E.g.
 * 
 * ```sh
 * $ pulumi import rabbitmq:index/policy:Policy test name@vhost
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * The name of the policy.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The settings of the policy. The structure is
     * described below.
     * 
     */
    @Export(name="policy", refs={PolicyPolicy.class}, tree="[0]")
    private Output<PolicyPolicy> policy;

    /**
     * @return The settings of the policy. The structure is
     * described below.
     * 
     */
    public Output<PolicyPolicy> policy() {
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
    public Policy(String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(String name, PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(String name, PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/policy:Policy", name, args == null ? PolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Policy(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/policy:Policy", name, state, makeResourceOptions(options, id));
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
    public static Policy get(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}
