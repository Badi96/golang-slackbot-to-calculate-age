# Golang slack bot for calculate age
Creating a bot in slack using golang that lets a users calculate their age based on year of birth

## Prerequisites
- Download and install Go (https://golang.org/dl/)
- Slack Installed with a slack account and worspace setup that you are the admin to.

## Setting up the slack bot environment 
In main.go, replace the strings ```"Replace_with_Bot_token_here"``` and ```"Replace_with_socket_token_here"``` below with your specific tokens.
```go 
os.Setenv("SLACK_BOT_TOKEN", "Replace_with_Bot_token_here") // Token from Oauth & Permissions
```
```go
os.Setenv("SLACK_APP_TOKEN", "Replace_with_socket_token_here") //socket token. The channel ID in the slack channel of the workspace
```
1. To find your token, you need to got to api.slack.com/apps. Press ```create new app``` and choose ```From scratch```, name you app with your worspace you are admin to.
2. Choose ```Socket Mode``` in the left menu and enable socket mode. Give a name for the socket/app token. Replace this token with the ```Replace_with_socket_token_here``` above in the code. 
3. Go to ```Event subscriptions```, enable event subscription. Go to Subscribe to bot events, and add all the event the bot should subscribe to. For this example, add: 
    - ```app_mention```
    - ```channels:history``` 
    - ```channels:read``` 
    - ```chat:write``` 
    - ```chat:write.public``` 
    - ```im:history``` 
    - ```mpim:history``` 
4. Finally, press on ```Install to Workspace``` and accept. Take the bot User Oauth Token and replace with ```"Replace_with_socket_token_here"``` above in the code. 
## Commands 
In this project folder, run in terminal:
```go mod init github.com/badi96/slack-age-bot```
```go get "github.com/shomali11/slacker"```
```go mod tidy```
```go get github.com/shomali11/slacker```
```go build```
```go run main.go```

## Runing the bot in Slack
To run the bot in slack, open the slack application and go to you general channel in your woekspace. Add the bot to your channel (if not alreaady added). To get you age, write for example ```@age-bot my yob is 1996``` and you should get your age as a response. 
