import os
from flask import Flask,render_template
from chapa import Chapa
import requests
import json

chapa_api=os.getenv("api")
app=Flask(__name__)

@app.route("/")
def index():
    return render_template("index.html")

if __name__=="__main__":
    app.run(debug=True)