import requests
from bs4 import BeautifulSoup

url = "https://www.airbnb.com/"
search_params = {
    "location": "Los Angeles, California",
    "checkin": "2023-10-10",
    "checkout": "2023-10-25"
}

response = requests.get(url, params=search_params)
response.raise_for_status()

soup = BeautifulSoup(response.content, "html.parser")
results = soup.find_all("div", class_="c4mnd7m dir dir-ltr")[:10]

for result in results:
    print(result.get_text())
    print("---------------")
