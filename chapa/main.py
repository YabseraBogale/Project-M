from flask import Flask, render_template, request, redirect, url_for, jsonify
import requests

app = Flask(__name__)

CHAPA_SECRET_KEY = 'your_chapa_secret_key'
CHAPA_BASE_URL = 'https://api.chapa.co/v1/transaction'

@app.route('/')
def index():
    return render_template('checkout.html')

@app.route('/pay', methods=['POST'])
def initiate_payment():
    data = {
        "amount": request.form['amount'],
        "currency": "ETB",  # Chapa primarily supports Ethiopian Birr
        "email": request.form['email'],
        "first_name": request.form['first_name'],
        "last_name": request.form['last_name'],
        "tx_ref": "tx-" + str(hash(request.form['email'])),  # Unique transaction reference
        "callback_url": url_for('payment_callback', _external=True),
        "return_url": url_for('payment_success', _external=True),
    }

    headers = {
        "Authorization": f"Bearer {CHAPA_SECRET_KEY}"
    }

    response = requests.post(f"{CHAPA_BASE_URL}/initialize", json=data, headers=headers)
    
    if response.status_code == 200:
        return redirect(response.json()['data']['checkout_url'])
    else:
        return render_template('error.html', error=response.json().get('message', 'Payment initiation failed'))

@app.route('/callback', methods=['GET'])
def payment_callback():
    # Handle Chapa webhook callback here
    return jsonify({"status": "Callback received"})

@app.route('/success', methods=['GET'])
def payment_success():
    return render_template('success.html')

if __name__ == "__main__":
    app.run(debug=True)

