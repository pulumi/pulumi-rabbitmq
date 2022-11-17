// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;


public final class PolicyPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final PolicyPolicyArgs Empty = new PolicyPolicyArgs();

    /**
     * Can either be &#34;exchanges&#34;, &#34;queues&#34;, or &#34;all&#34;.
     * 
     */
    @Import(name="applyTo", required=true)
    private Output<String> applyTo;

    /**
     * @return Can either be &#34;exchanges&#34;, &#34;queues&#34;, or &#34;all&#34;.
     * 
     */
    public Output<String> applyTo() {
        return this.applyTo;
    }

    /**
     * Key/value pairs of the policy definition. See the
     * RabbitMQ documentation for definition references and examples.
     * 
     */
    @Import(name="definition", required=true)
    private Output<Map<String,Object>> definition;

    /**
     * @return Key/value pairs of the policy definition. See the
     * RabbitMQ documentation for definition references and examples.
     * 
     */
    public Output<Map<String,Object>> definition() {
        return this.definition;
    }

    /**
     * A pattern to match an exchange or queue name.
     * 
     */
    @Import(name="pattern", required=true)
    private Output<String> pattern;

    /**
     * @return A pattern to match an exchange or queue name.
     * 
     */
    public Output<String> pattern() {
        return this.pattern;
    }

    /**
     * The policy with the greater priority is applied first.
     * 
     */
    @Import(name="priority", required=true)
    private Output<Integer> priority;

    /**
     * @return The policy with the greater priority is applied first.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }

    private PolicyPolicyArgs() {}

    private PolicyPolicyArgs(PolicyPolicyArgs $) {
        this.applyTo = $.applyTo;
        this.definition = $.definition;
        this.pattern = $.pattern;
        this.priority = $.priority;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyPolicyArgs $;

        public Builder() {
            $ = new PolicyPolicyArgs();
        }

        public Builder(PolicyPolicyArgs defaults) {
            $ = new PolicyPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applyTo Can either be &#34;exchanges&#34;, &#34;queues&#34;, or &#34;all&#34;.
         * 
         * @return builder
         * 
         */
        public Builder applyTo(Output<String> applyTo) {
            $.applyTo = applyTo;
            return this;
        }

        /**
         * @param applyTo Can either be &#34;exchanges&#34;, &#34;queues&#34;, or &#34;all&#34;.
         * 
         * @return builder
         * 
         */
        public Builder applyTo(String applyTo) {
            return applyTo(Output.of(applyTo));
        }

        /**
         * @param definition Key/value pairs of the policy definition. See the
         * RabbitMQ documentation for definition references and examples.
         * 
         * @return builder
         * 
         */
        public Builder definition(Output<Map<String,Object>> definition) {
            $.definition = definition;
            return this;
        }

        /**
         * @param definition Key/value pairs of the policy definition. See the
         * RabbitMQ documentation for definition references and examples.
         * 
         * @return builder
         * 
         */
        public Builder definition(Map<String,Object> definition) {
            return definition(Output.of(definition));
        }

        /**
         * @param pattern A pattern to match an exchange or queue name.
         * 
         * @return builder
         * 
         */
        public Builder pattern(Output<String> pattern) {
            $.pattern = pattern;
            return this;
        }

        /**
         * @param pattern A pattern to match an exchange or queue name.
         * 
         * @return builder
         * 
         */
        public Builder pattern(String pattern) {
            return pattern(Output.of(pattern));
        }

        /**
         * @param priority The policy with the greater priority is applied first.
         * 
         * @return builder
         * 
         */
        public Builder priority(Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority The policy with the greater priority is applied first.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        public PolicyPolicyArgs build() {
            $.applyTo = Objects.requireNonNull($.applyTo, "expected parameter 'applyTo' to be non-null");
            $.definition = Objects.requireNonNull($.definition, "expected parameter 'definition' to be non-null");
            $.pattern = Objects.requireNonNull($.pattern, "expected parameter 'pattern' to be non-null");
            $.priority = Objects.requireNonNull($.priority, "expected parameter 'priority' to be non-null");
            return $;
        }
    }

}
