import firebase_admin
from firebase_admin import credentials
from firebase_admin import firestore

# Use a service account.
cred = credentials.Certificate('./serviceAccountKey.json')

app = firebase_admin.initialize_app(cred)

db = firestore.client()



data = {"name": "Los Angeles", "state": "CA", "country": "USA"}

# Add a new doc in collection 'cities' with ID 'LA'
db.collection("cities").document("LA").set(data)
