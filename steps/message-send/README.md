# slack-step-message-send

This [Slack](https://slack.com) step sends a message to a desired channel.

This message can either be via [blocks](https://api.slack.com/reference/block-kit/blocks) or as a direct message; at least one must be used. If both are supplied, the behavior is defined by the [Slack API](https://api.slack.com/methods/chat.postMessage#text_usage).

In order to connect to Slack you will need to **Create a new Slack app** at [https://api.slack.com/apps/](https://api.slack.com/apps?new_app=1), then:

* Give it a name like "Notifications from Relay.sh"
* Select the workspace you want to target for notifications
* Navigate to **Add features and functionality**
  * Click **Permissions**
  * On the **OAuth and Permissions** page, go to **Scopes**
  * Under **Bot Token Scopes**, add `chat:write`, `chat:write.public`, `chat:write.customize`
* For some organizations your Slack app may require approval before it can be installed.
* Once approved and installed, you'll see a **Bot User OAuth Access Token** on the site.
* In Relay, go to the **Connections** tab and click **Add connection**. Choose **Slack** from the drop-down.
* Give the connection a name which you'll reference in your workflows and paste the token in.
