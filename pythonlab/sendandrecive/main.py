import os
import imaplib
import email
from email.header import decode_header

# Account credentials
username = "your_email@example.com"
password =  os.getenv("password")

# Connect to the server
# For Gmail, use "imap.gmail.com"
imap_server = "imap.example.com"
mail = imaplib.IMAP4_SSL(imap_server)

# Log in
mail.login(username, password)

# Select the mailbox you want to use ('inbox' in this case)
mail.select("inbox")

# Search for all emails in the inbox
status, messages = mail.search(None, "ALL")

# Get the list of email IDs
email_ids = messages[0].split()

# Fetch the latest email by ID
latest_email_id = email_ids[-1]

# Fetch the email by ID
status, msg_data = mail.fetch(latest_email_id, "(RFC822)")

# Parse the raw email content
for response_part in msg_data:
    if isinstance(response_part, tuple):
        msg = email.message_from_bytes(response_part[1])
        # Decode email sender, subject, and date
        subject, encoding = decode_header(msg["Subject"])[0]
        if isinstance(subject, bytes):
            subject = subject.decode(encoding if encoding else "utf-8")
        from_ = msg.get("From")
        date_ = msg.get("Date")

        print("From:", from_)
        print("Subject:", subject)
        print("Date:", date_)

        # If the email message is multipart
        if msg.is_multipart():
            for part in msg.walk():
                # Check if part is text/plain or text/html
                if part.get_content_type() == "text/plain" and part.get("Content-Disposition") is None:
                    body = part.get_payload(decode=True).decode()
                    print("Body:", body)
                    break
        else:
            # If the email message is not multipart
            body = msg.get_payload(decode=True).decode()
            print("Body:", body)

# Close the connection and logout
mail.close()
mail.logout()
