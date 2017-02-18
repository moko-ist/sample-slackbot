# hello-slackbot

My First SlackBot

# setup
### api token

Need to setup your slack api token.
Please generate `.envrc` and set `SLACK_API_TOKEN` environment variables.

```sh
$ cd /path/to/project
$ direnv edit .
# => set SLACK_API_TOKEN

$ cat .envrc
SLACK_API_TOKEN="*** your api token ***"
```

if the error occurs (like `direnv: error .envrc is blocked. Run `direnv allow` to approve its content.` ), please allow direnv like the following:

```sh
$ direnv allow
```
