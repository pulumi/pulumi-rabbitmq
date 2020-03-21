using Pulumi;
using RabbitMq = Pulumi.RabbitMQ;
using Random = Pulumi.Random;

class UserStack : Stack
{
    public UserStack()
    {
        var pw = new Random.RandomPassword("my-password", new Random.RandomPasswordArgs
        {
            Length = 30,
            Special = true,
        });

        var user = new RabbitMq.User("my-csharp-user", new RabbitMq.UserArgs
        {
            Password = pw.Result,
        });
    }
}
