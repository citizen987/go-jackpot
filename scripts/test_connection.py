import json
import sys
import requests # pip install requests


url = "http://localhost:8080/jackpot-draw"
headers = {"Content-Type": "application/json"}
bet = 1

if len(sys.argv) > 1:
    try:
        bet = int(sys.argv[1])
    except ValueError:
        print("Error: Argument bet value must be a number.")
        sys.exit(1)
data = {"bet": bet}

print("Request data:", data)
try:
    response = requests.post(url, headers=headers, data=json.dumps(data), timeout=5)
    print("Request completed.")
    print("Status code:", response.status_code)
    print("Server response:", response.text)
except requests.exceptions.ConnectionError:
    print("Error: could not connect to the server at localhost:8080.")
except requests.exceptions.Timeout:
    print("Error: the request timed out.")
except requests.exceptions.RequestException as e:
    print(f"Unexpected error during request: {e}")