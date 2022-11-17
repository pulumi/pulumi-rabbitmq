// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class OperatorPolicyPolicy {
    /**
     * @return Can be &#34;queues&#34;.
     * 
     */
    private String applyTo;
    /**
     * @return Key/value pairs of the operator policy definition. See the
     * RabbitMQ documentation for definition references and examples.
     * 
     */
    private Map<String,Object> definition;
    /**
     * @return A pattern to match an exchange or queue name.
     * 
     */
    private String pattern;
    /**
     * @return The policy with the greater priority is applied first.
     * 
     */
    private Integer priority;

    private OperatorPolicyPolicy() {}
    /**
     * @return Can be &#34;queues&#34;.
     * 
     */
    public String applyTo() {
        return this.applyTo;
    }
    /**
     * @return Key/value pairs of the operator policy definition. See the
     * RabbitMQ documentation for definition references and examples.
     * 
     */
    public Map<String,Object> definition() {
        return this.definition;
    }
    /**
     * @return A pattern to match an exchange or queue name.
     * 
     */
    public String pattern() {
        return this.pattern;
    }
    /**
     * @return The policy with the greater priority is applied first.
     * 
     */
    public Integer priority() {
        return this.priority;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OperatorPolicyPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String applyTo;
        private Map<String,Object> definition;
        private String pattern;
        private Integer priority;
        public Builder() {}
        public Builder(OperatorPolicyPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.applyTo = defaults.applyTo;
    	      this.definition = defaults.definition;
    	      this.pattern = defaults.pattern;
    	      this.priority = defaults.priority;
        }

        @CustomType.Setter
        public Builder applyTo(String applyTo) {
            this.applyTo = Objects.requireNonNull(applyTo);
            return this;
        }
        @CustomType.Setter
        public Builder definition(Map<String,Object> definition) {
            this.definition = Objects.requireNonNull(definition);
            return this;
        }
        @CustomType.Setter
        public Builder pattern(String pattern) {
            this.pattern = Objects.requireNonNull(pattern);
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        public OperatorPolicyPolicy build() {
            final var o = new OperatorPolicyPolicy();
            o.applyTo = applyTo;
            o.definition = definition;
            o.pattern = pattern;
            o.priority = priority;
            return o;
        }
    }
}
