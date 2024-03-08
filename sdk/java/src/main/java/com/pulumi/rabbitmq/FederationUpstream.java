// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.FederationUpstreamArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.FederationUpstreamState;
import com.pulumi.rabbitmq.outputs.FederationUpstreamDefinition;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.FederationUpstream`` resource creates and manages a federation upstream parameter.
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
 * import com.pulumi.rabbitmq.FederationUpstream;
 * import com.pulumi.rabbitmq.FederationUpstreamArgs;
 * import com.pulumi.rabbitmq.inputs.FederationUpstreamDefinitionArgs;
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
 *         var test = new VHost(&#34;test&#34;);
 * 
 *         var guest = new Permissions(&#34;guest&#34;, PermissionsArgs.builder()        
 *             .user(&#34;guest&#34;)
 *             .vhost(test.name())
 *             .permissions(PermissionsPermissionsArgs.builder()
 *                 .configure(&#34;.*&#34;)
 *                 .write(&#34;.*&#34;)
 *                 .read(&#34;.*&#34;)
 *                 .build())
 *             .build());
 * 
 *         var fooExchange = new Exchange(&#34;fooExchange&#34;, ExchangeArgs.builder()        
 *             .vhost(guest.vhost())
 *             .settings(ExchangeSettingsArgs.builder()
 *                 .type(&#34;topic&#34;)
 *                 .durable(&#34;true&#34;)
 *                 .build())
 *             .build());
 * 
 *         var fooFederationUpstream = new FederationUpstream(&#34;fooFederationUpstream&#34;, FederationUpstreamArgs.builder()        
 *             .vhost(guest.vhost())
 *             .definition(FederationUpstreamDefinitionArgs.builder()
 *                 .uri(&#34;amqp://guest:guest@upstream-server-name:5672/%2f&#34;)
 *                 .prefetchCount(1000)
 *                 .reconnectDelay(5)
 *                 .ackMode(&#34;on-confirm&#34;)
 *                 .trustUserId(false)
 *                 .maxHops(1)
 *                 .build())
 *             .build());
 * 
 *         var fooPolicy = new Policy(&#34;fooPolicy&#34;, PolicyArgs.builder()        
 *             .vhost(guest.vhost())
 *             .policy(PolicyPolicyArgs.builder()
 *                 .pattern(fooExchange.name().applyValue(name -&gt; String.format(&#34;(^%s$)&#34;, name)))
 *                 .priority(1)
 *                 .applyTo(&#34;exchanges&#34;)
 *                 .definition(Map.of(&#34;federation-upstream&#34;, fooFederationUpstream.name()))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * A Federation upstream can be imported using the resource `id` which is composed of `name@vhost`, e.g.
 * 
 * ```sh
 * $ pulumi import rabbitmq:index/federationUpstream:FederationUpstream foo foo@test
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/federationUpstream:FederationUpstream")
public class FederationUpstream extends com.pulumi.resources.CustomResource {
    /**
     * Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
     * 
     */
    @Export(name="component", refs={String.class}, tree="[0]")
    private Output<String> component;

    /**
     * @return Set to `federation-upstream` by the underlying RabbitMQ provider. You do not set this attribute but will see it in state and plan output.
     * 
     */
    public Output<String> component() {
        return this.component;
    }
    /**
     * The configuration of the federation upstream. The structure is described below.
     * 
     */
    @Export(name="definition", refs={FederationUpstreamDefinition.class}, tree="[0]")
    private Output<FederationUpstreamDefinition> definition;

    /**
     * @return The configuration of the federation upstream. The structure is described below.
     * 
     */
    public Output<FederationUpstreamDefinition> definition() {
        return this.definition;
    }
    /**
     * The name of the federation upstream.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the federation upstream.
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
    public FederationUpstream(String name) {
        this(name, FederationUpstreamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FederationUpstream(String name, FederationUpstreamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FederationUpstream(String name, FederationUpstreamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/federationUpstream:FederationUpstream", name, args == null ? FederationUpstreamArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FederationUpstream(String name, Output<String> id, @Nullable FederationUpstreamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/federationUpstream:FederationUpstream", name, state, makeResourceOptions(options, id));
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
    public static FederationUpstream get(String name, Output<String> id, @Nullable FederationUpstreamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FederationUpstream(name, id, state, options);
    }
}
