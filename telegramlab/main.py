from pyrogram import Client, filters

import config

config=config.Config()

app_id=config.app_id
app_hash=config.app_hash

app=Client("my_account",api_id=app_id,api_hash=app_hash)

@app.on_message(filters.chat("freelance_ethio"))  # Filter messages from a specific channel
def message_listener(client, message):
    # Check if the message contains the target word
    try:
        if "Job Title".lower() in message.text.lower():
        # Print and respond to the message
            print(f"message: {message.text}")
        
        # Reply to the message (this sends a response in the channel)
    except Exception as e:
        print(str(e))
# Start the bot
app.run()
