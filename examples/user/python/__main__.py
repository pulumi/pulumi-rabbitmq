from pulumi import export
import pulumi_random as random
import pulumi_rabbitmq as rabbitmq

password = random.RandomPassword("my-password",
                                    length=30,
                                    special="true")

user = rabbitmq.User("my-python-user",
                        password=password.result)

export("username", user.name)
