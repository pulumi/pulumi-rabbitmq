import * as rabbitmq from "@pulumi/rabbitmq";
import * as random from "@pulumi/random";

const pw = new random.RandomPassword("my-password", {
    length: 30,
    special: true,
}, {
    additionalSecretOutputs: ["result"]
})

const user = new rabbitmq.User("my-typescript-user", {
    password: pw.result
});

export const userName = user.name;
