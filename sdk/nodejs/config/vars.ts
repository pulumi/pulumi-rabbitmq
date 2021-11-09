// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("rabbitmq");

export declare const cacertFile: string | undefined;
Object.defineProperty(exports, "cacertFile", {
    get() {
        return __config.get("cacertFile") ?? utilities.getEnv("RABBITMQ_CACERT");
    },
    enumerable: true,
});

export declare const clientcertFile: string | undefined;
Object.defineProperty(exports, "clientcertFile", {
    get() {
        return __config.get("clientcertFile");
    },
    enumerable: true,
});

export declare const clientkeyFile: string | undefined;
Object.defineProperty(exports, "clientkeyFile", {
    get() {
        return __config.get("clientkeyFile");
    },
    enumerable: true,
});

export declare const endpoint: string | undefined;
Object.defineProperty(exports, "endpoint", {
    get() {
        return __config.get("endpoint");
    },
    enumerable: true,
});

export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure") ?? <any>utilities.getEnvBoolean("RABBITMQ_INSECURE");
    },
    enumerable: true,
});

export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password");
    },
    enumerable: true,
});

export declare const proxy: string | undefined;
Object.defineProperty(exports, "proxy", {
    get() {
        return __config.get("proxy");
    },
    enumerable: true,
});

export declare const username: string | undefined;
Object.defineProperty(exports, "username", {
    get() {
        return __config.get("username");
    },
    enumerable: true,
});

