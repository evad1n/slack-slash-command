# slack-slash-command
A Slack slash command endpoint written in Go that can be deployed to AWS Lambda




## Limitations

Slack can't add headers to slash commands so you can't use an AWS API key to guard the functions.

https://api.slack.com/authentication/verifying-requests-from-slack