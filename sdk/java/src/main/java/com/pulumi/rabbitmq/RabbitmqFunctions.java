// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.rabbitmq;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.rabbitmq.Utilities;
import com.pulumi.rabbitmq.inputs.GetExchangeArgs;
import com.pulumi.rabbitmq.inputs.GetExchangePlainArgs;
import com.pulumi.rabbitmq.inputs.GetUserArgs;
import com.pulumi.rabbitmq.inputs.GetUserPlainArgs;
import com.pulumi.rabbitmq.inputs.GetVHostArgs;
import com.pulumi.rabbitmq.inputs.GetVHostPlainArgs;
import com.pulumi.rabbitmq.outputs.GetExchangeResult;
import com.pulumi.rabbitmq.outputs.GetUserResult;
import com.pulumi.rabbitmq.outputs.GetVHostResult;
import java.util.concurrent.CompletableFuture;

public final class RabbitmqFunctions {
    public static Output<GetExchangeResult> getExchange(GetExchangeArgs args) {
        return getExchange(args, InvokeOptions.Empty);
    }
    public static CompletableFuture<GetExchangeResult> getExchangePlain(GetExchangePlainArgs args) {
        return getExchangePlain(args, InvokeOptions.Empty);
    }
    public static Output<GetExchangeResult> getExchange(GetExchangeArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("rabbitmq:index/getExchange:getExchange", TypeShape.of(GetExchangeResult.class), args, Utilities.withVersion(options));
    }
    public static Output<GetExchangeResult> getExchange(GetExchangeArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("rabbitmq:index/getExchange:getExchange", TypeShape.of(GetExchangeResult.class), args, Utilities.withVersion(options));
    }
    public static CompletableFuture<GetExchangeResult> getExchangePlain(GetExchangePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("rabbitmq:index/getExchange:getExchange", TypeShape.of(GetExchangeResult.class), args, Utilities.withVersion(options));
    }
    public static Output<GetUserResult> getUser(GetUserArgs args) {
        return getUser(args, InvokeOptions.Empty);
    }
    public static CompletableFuture<GetUserResult> getUserPlain(GetUserPlainArgs args) {
        return getUserPlain(args, InvokeOptions.Empty);
    }
    public static Output<GetUserResult> getUser(GetUserArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("rabbitmq:index/getUser:getUser", TypeShape.of(GetUserResult.class), args, Utilities.withVersion(options));
    }
    public static Output<GetUserResult> getUser(GetUserArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("rabbitmq:index/getUser:getUser", TypeShape.of(GetUserResult.class), args, Utilities.withVersion(options));
    }
    public static CompletableFuture<GetUserResult> getUserPlain(GetUserPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("rabbitmq:index/getUser:getUser", TypeShape.of(GetUserResult.class), args, Utilities.withVersion(options));
    }
    public static Output<GetVHostResult> getVHost(GetVHostArgs args) {
        return getVHost(args, InvokeOptions.Empty);
    }
    public static CompletableFuture<GetVHostResult> getVHostPlain(GetVHostPlainArgs args) {
        return getVHostPlain(args, InvokeOptions.Empty);
    }
    public static Output<GetVHostResult> getVHost(GetVHostArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("rabbitmq:index/getVHost:getVHost", TypeShape.of(GetVHostResult.class), args, Utilities.withVersion(options));
    }
    public static Output<GetVHostResult> getVHost(GetVHostArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("rabbitmq:index/getVHost:getVHost", TypeShape.of(GetVHostResult.class), args, Utilities.withVersion(options));
    }
    public static CompletableFuture<GetVHostResult> getVHostPlain(GetVHostPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("rabbitmq:index/getVHost:getVHost", TypeShape.of(GetVHostResult.class), args, Utilities.withVersion(options));
    }
}
