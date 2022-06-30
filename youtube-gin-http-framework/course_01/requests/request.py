import requests
import json

HOST = "http://127.0.0.1:8000"


def get_test():
    url = f"{HOST}/test"    
    rq = requests.get(url=url)
    
    return rq.json()


def print_response(json_data):
    json_object = json.loads(json_data)
    json_formatted_str = json.dumps(json_object, indent=2)

    print(json_formatted_str)