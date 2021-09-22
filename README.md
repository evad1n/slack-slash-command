# Slack slash commands with AWS Lambda

*As of 9/21/2021*

A guide on how to create [slack slash commands](https://api.slack.com/interactivity/slash-commands) with [AWS Lambda](https://aws.amazon.com/lambda/) so you don't need a server.

I found the [official docs](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html) to be somewhat lacking for my purposes.

The template file `main.go` contains code for a simple slack slash command that echoes what you write to it. 

For example, if the command was 
```
/echo
``` 
then typing 
```
/echo peanut butter
```
would print
```
You said 'peanut butter'
```

## The code

The code uses the aws lambda go package to interface with AWS. Specifically, it uses the lambda package for the hooking up the function to be executed, and the events package to connect with an exposable API. You may notice that the [official docs](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html) have somewhat different handler signatures then the ones here, and that's due to using the API gateway to expose the function with a public REST API.

Slack sends its parameters as url encoded data in the request body. These need to be parsed to obtain the key/values. A list of the data slack sends can be seen [here](https://api.slack.com/interactivity/slash-commands). You can also view the accepted responses and possible formatting there.

## The Lambda function

To create an AWS Lambda function you will obviously need an AWS account. AWS has some generous free tiers, so you won't have to pay any money yet.

Once you have an account you can go the [AWS Console](https://aws.amazon.com/console/) and navigate to the Lambda service. There you will click `Create function` and select `Author from scratch`. There are some cool blueprints they have, including a python slack slash command.

Make sure to name your function well, as it will become the endpoint for your REST API. For example, naming it "kanye" would have the API gateway endpoint be `/kanye`.

Choose the Go 1.x runtime, which supports every version of Go past 1.0. Continue to the next page which is the function overview.

Here you'll be able to upload your code in [zip form](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html). The Makefile in this repo can be consulted to see how that is done, but it's quite simple. A **VERY** important note is that whatever your executable is called (when you call go build) needs to be the same as the Handler name in the Runtime settings. By default the starter function has it as `hello`, but you'll probably want `main`, which is the default executable name if you don't specify it with the `-o` flag with `go build`. This can be changed although it really doesn't matter what it is as long as it matches. Otherwise you may get errors about "No such file...." etc.

Next you'll want to hook the function up to a REST API, by clicking `Add trigger` on the function overview page. You'll want to select API Gateway, and then probably have to create a new API. If you have an existing one, then generally it just adds another endpoint corresponding to the function name. For example, it would be ```<API_URL>/<DEPLOYMENT_STAGE/<FUNCTION_NAME>```. For deployment stage just select default, but it essentially is a [snapshot of an API](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-stages.html). For security you will need to make it Public, as slack slash commands are unable to add any additional data/headers such as an API key.

At this point, you can test your API using Postman or cURL. However it would be easier to just hook up the slash command and try it out directly with real data/requests. 

*This can also be done using the [AWS CLI](https://aws.amazon.com/cli/), but I find it easier to just use the console as it handles roles/policies and all that nasty stuff for you. Make sure to stay in the same region, as you can only see the functions while in the region you created them in.*

## The slash command

Create a slack app with a slash command. For the request URL, simply enter the API endpoint for the API gateway attached to your lambda function. That's really it. Slack has [pretty good docs](https://api.slack.com/apps) for creating an app.


## Limitations

Slack can't add headers to slash commands so you can't use an AWS API key to guard the functions.

You can try [this](https://api.slack.com/authentication/verifying-requests-from-slack), but it happens after the request so you can't really use it completely block access to a public API, just stop some sensitive logic.
