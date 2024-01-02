// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.rabbitmq.inputs.PolicyPolicyArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final PolicyArgs Empty = new PolicyArgs();

    /**
     * The name of the policy.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the policy.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The settings of the policy. The structure is
     * described below.
     * 
     */
    @Import(name="policy", required=true)
    private Output<PolicyPolicyArgs> policy;

    /**
     * @return The settings of the policy. The structure is
     * described below.
     * 
     */
    public Output<PolicyPolicyArgs> policy() {
        return this.policy;
    }

    /**
     * The vhost to create the resource in.
     * 
     */
    @Import(name="vhost", required=true)
    private Output<String> vhost;

    /**
     * @return The vhost to create the resource in.
     * 
     */
    public Output<String> vhost() {
        return this.vhost;
    }

    private PolicyArgs() {}

    private PolicyArgs(PolicyArgs $) {
        this.name = $.name;
        this.policy = $.policy;
        this.vhost = $.vhost;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyArgs $;

        public Builder() {
            $ = new PolicyArgs();
        }

        public Builder(PolicyArgs defaults) {
            $ = new PolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the policy.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policy The settings of the policy. The structure is
         * described below.
         * 
         * @return builder
         * 
         */
        public Builder policy(Output<PolicyPolicyArgs> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The settings of the policy. The structure is
         * described below.
         * 
         * @return builder
         * 
         */
        public Builder policy(PolicyPolicyArgs policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param vhost The vhost to create the resource in.
         * 
         * @return builder
         * 
         */
        public Builder vhost(Output<String> vhost) {
            $.vhost = vhost;
            return this;
        }

        /**
         * @param vhost The vhost to create the resource in.
         * 
         * @return builder
         * 
         */
        public Builder vhost(String vhost) {
            return vhost(Output.of(vhost));
        }

        public PolicyArgs build() {
            if ($.policy == null) {
                throw new MissingRequiredPropertyException("PolicyArgs", "policy");
            }
            if ($.vhost == null) {
                throw new MissingRequiredPropertyException("PolicyArgs", "vhost");
            }
            return $;
        }
    }

}
