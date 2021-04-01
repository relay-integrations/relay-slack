#!/usr/bin/env python
# queries the slack api with the email address and gets back a user id

import os
import sys
from relay_sdk import Interface, Dynamic as D
from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError

relay = Interface()

api_token = relay.get(D.connection.apiToken)
email_address = relay.get(D.email)

try:
  client = WebClient(api_token)
  client.auth_test()
except SlackApiError as e:
  sys.stderr.write('Unable to authenticate with Slack')
  exit(1)

try:
    response = client.users_lookupByEmail(email=email_address)
    
    if response and response['user']:
      relay.outputs.set("memberID", response['user']['id'])
except SlackApiError as e:
    sys.stderr.write('The error return from the Slack API is: ' + e.response["error"])
    exit(1)
