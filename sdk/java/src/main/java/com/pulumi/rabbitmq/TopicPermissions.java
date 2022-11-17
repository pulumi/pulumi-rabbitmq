// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.rabbitmq.TopicPermissionsArgs;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.TopicPermissionsState;
import com.pulumi.rabbitmq.outputs.TopicPermissionsPermission;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The ``rabbitmq.TopicPermissions`` resource creates and manages a user&#39;s set of
 * topic permissions.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.rabbitmq.VHost;
 * import com.pulumi.rabbitmq.User;
 * import com.pulumi.rabbitmq.UserArgs;
 * import com.pulumi.rabbitmq.TopicPermissions;
 * import com.pulumi.rabbitmq.TopicPermissionsArgs;
 * import com.pulumi.rabbitmq.inputs.TopicPermissionsPermissionArgs;
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
 *         var testUser = new User(&#34;testUser&#34;, UserArgs.builder()        
 *             .password(&#34;foobar&#34;)
 *             .tags(&#34;administrator&#34;)
 *             .build());
 * 
 *         var testTopicPermissions = new TopicPermissions(&#34;testTopicPermissions&#34;, TopicPermissionsArgs.builder()        
 *             .permissions(TopicPermissionsPermissionArgs.builder()
 *                 .exchange(&#34;amq.topic&#34;)
 *                 .read(&#34;.*&#34;)
 *                 .write(&#34;.*&#34;)
 *                 .build())
 *             .user(testUser.name())
 *             .vhost(testVHost.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Permissions can be imported using the `id` which is composed of
 * 
 * `user@vhost`. E.g.
 * 
 * ```sh
 *  $ pulumi import rabbitmq:index/topicPermissions:TopicPermissions test user@vhost
 * ```
 * 
 */
@ResourceType(type="rabbitmq:index/topicPermissions:TopicPermissions")
public class TopicPermissions extends com.pulumi.resources.CustomResource {
    /**
     * The settings of the permissions. The structure is
     * described below.
     * 
     */
    @Export(name="permissions", type=List.class, parameters={TopicPermissionsPermission.class})
    private Output<List<TopicPermissionsPermission>> permissions;

    /**
     * @return The settings of the permissions. The structure is
     * described below.
     * 
     */
    public Output<List<TopicPermissionsPermission>> permissions() {
        return this.permissions;
    }
    /**
     * The user to apply the permissions to.
     * 
     */
    @Export(name="user", type=String.class, parameters={})
    private Output<String> user;

    /**
     * @return The user to apply the permissions to.
     * 
     */
    public Output<String> user() {
        return this.user;
    }
    /**
     * The vhost to create the resource in.
     * 
     */
    @Export(name="vhost", type=String.class, parameters={})
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
    public TopicPermissions(String name) {
        this(name, TopicPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TopicPermissions(String name, TopicPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TopicPermissions(String name, TopicPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/topicPermissions:TopicPermissions", name, args == null ? TopicPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TopicPermissions(String name, Output<String> id, @Nullable TopicPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/topicPermissions:TopicPermissions", name, state, makeResourceOptions(options, id));
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
    public static TopicPermissions get(String name, Output<String> id, @Nullable TopicPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TopicPermissions(name, id, state, options);
    }
}
