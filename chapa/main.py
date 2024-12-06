import os
from chapa import Chapa

chapa_api=os.getenv("api")
chapa=Chapa(chapa_api)
print(chapa_api)


# # Initialize the payment
# response = chapa.initialize(
#     email="yabserapython@gmail.com",
#     amount=1000,
#     first_name="John",
#     last_name="Doe",
#     tx_ref="txn_123456",
   
# )

# if response["status"] == "success":
#     # Redirect user to the payment link
#     print("Payment Link:", response["data"]["checkout_url"])
# else:
#     print("Failed to initialize payment:", response)

transaction_id = ""
verification_response = chapa.verify(transaction_id)
print(verification_response)