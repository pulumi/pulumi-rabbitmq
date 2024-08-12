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
 *         var test = new VHost("test", VHostArgs.builder()
 *             .name("test")
 *             .build());
 * 
 *         var testUser = new User("testUser", UserArgs.builder()
 *             .name("mctest")
 *             .password("foobar")
 *             .tags("administrator")
 *             .build());
 * 
 *         var testTopicPermissions = new TopicPermissions("testTopicPermissions", TopicPermissionsArgs.builder()
 *             .user(testUser.name())
 *             .vhost(test.name())
 *             .permissions(TopicPermissionsPermissionArgs.builder()
 *                 .exchange("amq.topic")
 *                 .write(".*")
 *                 .read(".*")
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
 * Permissions can be imported using the `id` which is composed of  `user{@literal @}vhost`.
 * 
 * E.g.
 * 
 * ```sh
 * $ pulumi import rabbitmq:index/topicPermissions:TopicPermissions test user{@literal @}vhost
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
    @Export(name="permissions", refs={List.class,TopicPermissionsPermission.class}, tree="[0,1]")
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
    @Export(name="user", refs={String.class}, tree="[0]")
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
    public TopicPermissions(java.lang.String name) {
        this(name, TopicPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TopicPermissions(java.lang.String name, TopicPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TopicPermissions(java.lang.String name, TopicPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/topicPermissions:TopicPermissions", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TopicPermissions(java.lang.String name, Output<java.lang.String> id, @Nullable TopicPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("rabbitmq:index/topicPermissions:TopicPermissions", name, state, makeResourceOptions(options, id), false);
    }

    private static TopicPermissionsArgs makeArgs(TopicPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TopicPermissionsArgs.Empty : args;
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
    public static TopicPermissions get(java.lang.String name, Output<java.lang.String> id, @Nullable TopicPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TopicPermissions(name, id, state, options);
    }
}
