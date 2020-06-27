#!/usr/local/bin/python3
import requests
import json
import urllib3
urllib3.disable_warnings()
#curlUrl = 'https://prod1-aus2.tnt34-zone1.aus2/api/v1/pods'
curlUrl = 'https://stg1-phl1.tnt34-zone2.phl1/api/v1/pods'
myToken = ''
headers = {
    'Authorization': 'Bearer {}'.format(myToken),
    'Content-Type': 'application/json'
    }
response = requests.get(url=curlUrl, headers=headers, verify=False)
out = response.json()
rc = 0
print(out)
for pod in out['items']:
    if 'waiting' not in pod['status']['containerStatuses'][0]['state']:
        rc == 0
    elif 'waiting' in pod['status']['containerStatuses'][0]['state']:
        print(pod['metadata']['namespace'] + "   " +pod['metadata']['name'],   pod['status']['containerStatuses'][0]['state']['waiting']['reason'],)
        rc += 1
if rc == 0:
    print("OK")