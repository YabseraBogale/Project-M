import requests
from bs4 import BeautifulSoup

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
}

url = "https://www.amazon.com/s?k=laptop"
response = requests.get(url, headers=headers)
soup = BeautifulSoup(response.text, "html.parser")

# Extract product titles
titles = [item.text for item in soup.find_all("span", class_="a-size-medium")]
print(titles)