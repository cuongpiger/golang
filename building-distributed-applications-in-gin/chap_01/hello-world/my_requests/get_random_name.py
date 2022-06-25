import requests
import json

url = "http://127.0.0.1:8080/random-name/"

get_request = requests.get(url)


print(get_request.text)