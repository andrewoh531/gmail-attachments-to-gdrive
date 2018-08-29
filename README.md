# Save Gmail Attachments to Google Drive

## Prerequisite
### Requirements
* AWS CLI already configured with at least PowerUser permission
* [Docker](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

### Setting Up AWS Parameter Store
The following credentials/keys will need to be stored in AWS Parameter Store as secure strings 
using the default KMS keys:
- Gmail client credentials: This is a JSON file but you can copy the contents as a secure string.
  You can retrieve the JSON credentials from the [Google developer console page](https://console.developers.google.com/apis/credentials)
- GMail user's refresh token: The OAuth2 refresh token is used to authorize the application to search your gmail account. 
  [Here](https://github.com/andrewoh531/gmail-authenticator) is sample app that can do this for you.

### Building and Running
See the `Makefile` to see the various targets that are supported