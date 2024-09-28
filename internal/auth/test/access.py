import requests     # pip install requests
from dataclasses import dataclass

# Define the URL of the server
BASE_URL = "http://localhost:8081/"


@dataclass
class Request:
    path: str
    data: dict
    headers: dict


def get_response(req: Request):
    return requests.post(req.path, json=req.data, headers=req.headers)



if __name__ == "__main__":
    req = Request(
        BASE_URL + "auth/register",
        # Define the data to be sent in the request body
        data = {"username": "value", "password": "bar!A5barbarbarbb"},

        # Define the headers for the request
        headers = {"Content-Type": "application/json"}
    )

    response = get_response(req)
    print(response)

    # Check the status code of the response
    if response.status_code == 200:
        print("Request successful!")
    else:
        print("Error:", response.status_code)
