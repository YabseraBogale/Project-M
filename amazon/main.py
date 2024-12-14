import requests
from bs4 import BeautifulSoup
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:114.0) Gecko/20100101 Firefox/114.0",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
    "Accept-Language": "en-US,en;q=0.5",
    "Connection": "keep-alive",
}

url = "https://www.amazon.com/s?k=laptops&crid=3BVZ14BC0I8V7&sprefix=laptop%2Caps%2C970&ref=nb_sb_noss_1"
response = requests.get(url,headers=headers)
soup = BeautifulSoup(response.text, "html.parser")


for i in soup:
    print(i.get_text())