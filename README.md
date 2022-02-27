# Slack Docker Action

This action sends a custom \ pre-defined slack message to A slack channel.

## Inputs

## `status`

**Required** The wanted slack message - started , success, failed, custom

## Outputs:
Insert slack message image



**Required** Env variables: The following env variables MUST be defined in your workflow
```
env:
  # Statics:
  PROJECT_NAME:           ${{ github.event.repository.name }}
  PROJECT_URL:            ${{ github.event.repository.url }}
  COMMITER:               ${{ github.event.sender.login }}
  SLACK_URL:              <SLACK_URL>
  CHANNEL_ID:             <CHANNEL_ID>
  USERS_FILE:             <NAME_OF_USERS_FILE>
  USERS_S3_FILE_PATH:     <PATH_TO_S3_BUCKET>
  COMMIT_SHA:             ${{ github.event.pull_request.base.sha }}
  PR_BUILD_URL:           ${{ github.event.pull_request.diff_url }}
  PUSH_BUILD_URL:         ${{ github.event.repository.owner.html_url }}
  AWS_ACCESS_KEY_ID:      ${{ secrets.aws_access_key }}
  AWS_SECRET_ACCESS_KEY:  ${{ secrets.aws_secret_key }}
  AWS_REGION:             <AWS_REGION>
  # Dynamics - depends on pr \ push:
  RUN_ID:                 ${{ github.run_id }}
  ENVIRONMENT:            ${{ github.event.repository.default_branch }}
  VERSION:                ${{ github.event.push.base.ref }}
  TEAM:                   Core
```
**Required** Pull Request Env variables: The following env variables MUST be defined in your workflow
```
env:
  COMMIT_MESSAGE:         ${{ github.event.pull_request.body }}
  ENVIRONMENT:            ${{ github.event.pull_request.base.ref }}
  VERSION:                ${{ github.event.pull_request.base.ref }}

```
**Required** Push Env variables: The following env variables MUST be defined in your workflow
```
env:
  COMMIT_MESSAGE:         ${{ github.event.head_commit.message }}
  RUN_ID:                 ${{ github.run_id }}
  ENVIRONMENT:            ${{ github.event.repository.default_branch }}
  VERSION:                ${{ github.event.push.base.ref }}
```
**Optional** Env variables: The following env variables are optional
```
env:
  CUSTOM_PAYLOAD_PATH:    ./examples/custom_payload.json
```


## Example usage 1 - send a pre defined template of started message
```
uses: actions/slack-action@v1
with:
  status: started
```

## Example usage 2 - send a pre defined template of failed message
```
uses: actions/slack-action@v1
with:
  status: failed
```

## Example usage 3 - send a pre defined template of success message
```
uses: actions/slack-action@v1
with:
  status: success
```

## Example usage 4 - send a custom template
### Add the following env variable specifying the path to your custom payload
### Please refer to ./examples/custom_payload.json for reference
```
env:
  CUSTOM_PAYLOAD_PATH:    ./my_custom_payload.json

uses: actions/slack-action@v1
with:
  status: custom
```
