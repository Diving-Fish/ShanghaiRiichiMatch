import requests

resp = requests.post("http://47.100.50.175:8088/api/public/push_score", json={
    "name": "2-3",
    "round": 1,
    "point": 14.5
})
print(resp)