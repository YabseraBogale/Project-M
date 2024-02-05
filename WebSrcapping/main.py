import requests
from bs4 import BeautifulSoup as bts
import pandas as pd
import re
import numpy as np
from pytictoc import TicToc
from datetime import datetime
import time

def getAndParseURL(url):
    result = requests.get(url, headers={"User-Agent":"Chrome/103.0.0.0"}) # Safari/537.36. Mozilla/5.0
    soup = bts(result.text, "html.parser")
    return soup

pages = [] #annual bestsellers&literary fiction
for page in range(1,24):
    pages.append("https://www.kitapyurdu.com/index.php?route=product/best_sellers&page="+str(page)+"&list_id=18&filter_in_stock=1&filter_in_stock=1")




products = []
for page in pages:
    html = getAndParseURL(page)
    for product in html.findAll("div",{"class":"cover"}):
        products.append(product.a.get("href"))
