# AWS Lambda tags ABAC

Inspired by [this article](https://aws.amazon.com/blogs/compute/scaling-aws-lambda-permissions-with-attribute-based-access-control-abac/).

## Learnings

- It would be neat for the SAM to know how to build `provided.al2` runtime Go lambdas. To my knowledge, the `go1.x` runtime is deprecated – it does not support AWS Lambda extensions!

- AWS SAM will create two stages whenever you deploy an AWS Lambda function. The first stage is called "Stage" and the second one is called "Prod". I have no idea why they chose to use _CapitalCase_ for the names. As for the additional "Stage", [refer to this issue](https://github.com/aws/serverless-application-model/issues/191).

- I always forget that **to add a resource-based policy to the AWS Lambda function, one has to use the `AWS::Lambda::Permission` resource**.

  - Keep in mind that **the identity making the call must also have permission to perform a given action**. The resource policy is optional in cross-account settings.

  - Read more about resource policies and how they relate to AWS Lambda [here](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html).

- To make the condition based on tags work, you must assume a role with the correct tags. I could not find a way for the AWS Lambda to do that automatically for me – I have zero control over the STS call AWS Lambda makes.

- You **have to have permissions to tag the session**. The permissions are set at **the role trust policy level**. Otherwise, you might find yourself creating loops with CFN references.

- It was surprisingly hard to find a good example of passing the credentials returned from the `sts:AssumeRole` call onto the AWS SDK configuration.
