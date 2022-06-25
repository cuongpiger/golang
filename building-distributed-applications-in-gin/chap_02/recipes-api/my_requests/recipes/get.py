import requests
import json

endpoint = f"http://127.0.0.1:8080/recipes/"

request = requests.get(endpoint)

response = request.json()
json_format = json.dumps(response, indent=4, sort_keys=True)

print(json_format)
