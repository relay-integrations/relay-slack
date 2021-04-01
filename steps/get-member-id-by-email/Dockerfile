FROM relaysh/core:latest-python
RUN pip --no-cache-dir install slack_sdk
COPY "./get-member-id-by-email.py" "/entrypoint.py"
ENTRYPOINT []
CMD ["python3", "/entrypoint.py"]

LABEL "org.opencontainers.image.title"="Slack Lookup Member ID by Email"
LABEL "org.opencontainers.image.description"="This step gets a slack users member id from an email"
LABEL "sh.relay.sdk.version"="v1"
