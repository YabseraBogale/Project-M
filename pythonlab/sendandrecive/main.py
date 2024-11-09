from pprint import pprint
import imaplib
import email
import os

Email="cheretaaddis@gmail.com"
password=os.getenv("password")
server="imap.gmail.com"

mail=imaplib.IMAP4_SSL(server,993)


result=mail.login(Email,password)


mail.select("inbox")
status,data=mail.search(None,"ALL")
mail_ids = []
for block in data:
    mail_ids+=block.split()


for i in mail_ids:
    try:    
        status, data = mail.fetch(i, '(RFC822)')
        for response_part in data:
            if isinstance(response_part, tuple):
                message = email.message_from_bytes(response_part[1])
                mail_from = message['from']
                mail_subject = message['subject']
                if message.is_multipart():
                    mail_content = ''

                    for part in message.get_payload():
                        if part.get_content_type() == 'text/plain':
                            mail_content += part.get_payload()

                else:
                    mail_content = message.get_payload()
                print(f'From: {mail_from}')
                print(f'Subject: {mail_subject}')
                print(f'Content: {mail_content}')
    except Exception as e:
        print(str(e))