import requests
from requests import auth

HOST = "http://127.0.0.1:8000"    
    
def post_videos(dict_data):
    url = f"{HOST}/videos/"
    request = requests.post(url=url, json=dict_data)
    
    return request.json()

def get_videos():
    url = f"{HOST}/videos/"
    request = requests.get(url=url)
    
    return request.json()

def auth_get_videos(username: str, password: str):
    session = requests.Session()
    session.auth = (username, password)    
    url = f"{HOST}/videos/"
    
    session.post(url=url)
    response = session.get(url=url)
    
    return response.json()

def auth_post_videos(username: str, password: str, dict_data):
    url = f"{HOST}/videos/"
    response = requests.post(url=url, 
                             json=dict_data, 
                             auth=auth.HTTPBasicAuth(username, password))
    
    return response.json()