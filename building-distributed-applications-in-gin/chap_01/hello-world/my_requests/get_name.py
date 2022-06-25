from faker import Faker

import requests
import json

fake = Faker()

name = fake.name()
url = f"http://127.0.0.1:8080/{name}/"

get_request = requests.get(url)

response = get_request.json()
json_format = json.dumps(response, indent=4, sort_keys=True)

print(json_format)