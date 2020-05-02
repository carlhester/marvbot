# Marvbot

Marvbot is a Discord bot that can read and send messages.  

It uses websockets to receive data from Discord and Webhooks to send chat messages.   


## Config
You must first register as a bot in the discord developer site to get the token information.

The environment variables are required:

# DISCORD_API is the primary http api endpoint
- export DISCORD_API='https://discord.com/api'

# DISCORDTOKEN is used to send data to the Discord websocket api
- export DISCORDTOKEN='xxxxx'

# DISCORD_BOT_TOKEN is used to talk to the Discord http api
- export DISCORD_BOT_TOKEN='Bot yyyyyy'
