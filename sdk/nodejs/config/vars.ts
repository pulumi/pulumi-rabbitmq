// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("rabbitmq");

export let cacertFile: string | undefined = __config.get("cacertFile") || utilities.getEnv("RABBITMQ_CACERT");
export let endpoint: string | undefined = __config.get("endpoint") || utilities.getEnv("RABBITMQ_ENDPOINT");
export let insecure: boolean | undefined = __config.getObject<boolean>("insecure") || utilities.getEnvBoolean("RABBITMQ_INSECURE");
export let password: string | undefined = __config.get("password") || utilities.getEnv("RABBITMQ_PASSWORD");
export let username: string | undefined = __config.get("username") || utilities.getEnv("RABBITMQ_USERNAME");
