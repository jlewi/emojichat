import json
import requests

if __name__== "__main__":
    data = {
        "sender": "test_user",
        "message": "Hello",
    }
    resp = requests.post("http://localhost:5005/webhooks/rest/webhook", json=data)
    
    print(resp)
    if resp.ok:
        print(resp.content)
    