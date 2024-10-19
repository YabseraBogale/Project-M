from pyrogram import Client
from config import Config
# API ID and Hash from my.telegram.org
api_id = Config().app_id
api_hash = Config().app_hash
phone_number = "+251920201161"

# Create the Pyrogram Client
app = Client("my_account", api_id=api_id, api_hash=api_hash)



# List of usernames to send the message to
users = ["username1", "username2", "username3"]  # Add Telegram usernames here
message_text = "Hello! This is a broadcast message."


# Function to send message to a list of users
@app.on_message()
def send_messages_to_users(client, message):
    for user in users:
        try:
            app.send_message(user, message_text)
            print(f"Message sent to {user}")
        except Exception as e:
            print(f"Failed to send message to {user}: {e}")

# Start the broadcast



app.run()

