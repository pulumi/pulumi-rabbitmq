// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
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
    private Map<String,String> definition;
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
    public Map<String,String> definition() {
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
        private Map<String,String> definition;
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
            if (applyTo == null) {
              throw new MissingRequiredPropertyException("OperatorPolicyPolicy", "applyTo");
            }
            this.applyTo = applyTo;
            return this;
        }
        @CustomType.Setter
        public Builder definition(Map<String,String> definition) {
            if (definition == null) {
              throw new MissingRequiredPropertyException("OperatorPolicyPolicy", "definition");
            }
            this.definition = definition;
            return this;
        }
        @CustomType.Setter
        public Builder pattern(String pattern) {
            if (pattern == null) {
              throw new MissingRequiredPropertyException("OperatorPolicyPolicy", "pattern");
            }
            this.pattern = pattern;
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            if (priority == null) {
              throw new MissingRequiredPropertyException("OperatorPolicyPolicy", "priority");
            }
            this.priority = priority;
            return this;
        }
        public OperatorPolicyPolicy build() {
            final var _resultValue = new OperatorPolicyPolicy();
            _resultValue.applyTo = applyTo;
            _resultValue.definition = definition;
            _resultValue.pattern = pattern;
            _resultValue.priority = priority;
            return _resultValue;
        }
    }
}
