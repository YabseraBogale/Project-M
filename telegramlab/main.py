from pyrogram import Client

import config

config=config.Config()

app_id=config.app_id
app_hash=config.app_hash

app=Client("my_account",app_id=app_id,app_hash=app_hash)

app.run()
