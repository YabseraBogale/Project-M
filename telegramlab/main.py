from telegram.ext import Updater, MessageHandler, Filters
from telegram import ParseMode

# Your bot token from BotFather
BOT_TOKEN = 'YOUR_BOT_TOKEN'
TARGET_WORD = 'your_word'  # The word to listen for
RESPONSE = 'This is the response message'  # The response message

# Specify the channel ID or username (e.g., -1001234567890 or @channelusername)
CHANNEL_ID = '@your_channel_username'  # Or use the numeric ID like -1001234567890

# Function to handle new messages
def message_handler(update, context):
    # Check if the message is from the specified channel
    if update.message.chat.username == CHANNEL_ID.replace('@', '') or update.message.chat.id == CHANNEL_ID:
        message_text = update.message.text
        
        # Check if the target word is in the message
        if TARGET_WORD.lower() in message_text.lower():
            # Respond to the message
            update.message.reply_text(RESPONSE, parse_mode=ParseMode.HTML)

def main():
    # Create Updater and pass the bot's token
    updater = Updater(token=BOT_TOKEN, use_context=True)
    
    # Get the dispatcher to register handlers
    dispatcher = updater.dispatcher
    
    # Add a handler for messages
    dispatcher.add_handler(MessageHandler(Filters.text & ~Filters.command, message_handler))
    
    # Start polling to listen for new messages
    updater.start_polling()
    
    # Block until you press Ctrl+C
    updater.idle()

if __name__ == '__main__':
    main()

