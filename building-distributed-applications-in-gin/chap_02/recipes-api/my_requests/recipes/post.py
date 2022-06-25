import requests
import json

data = {
    "id": "c0283p3d0cvuglq85log",
    "name": "Oregano Marinated Chicken",
    "tags": [
            "main",
            "chicken"
    ],
    "ingredients": [
        "4 (6 to 7-ounce) boneless skinless chicken breasts\r",
        "10 grinds black pepper\r",
        "1/2 tsp salt\r",
        "2 tablespoon extra-virgin olive oil\r",
        "1 teaspoon dried oregano\r",
        "1 lemon, juiced"
    ],
    "instructions": [
        "To marinate the chicken: In a non-reactive dish, combine the lemon juice, olive oil, oregano, salt, and pepper and mix together",
        " Add the chicken breasts to the dish and rub both sides in the mixture",
        " Cover the dish with plastic wrap and let marinate in the refrigerator for at least 30 minutes and up to 4 hours",
        "\r\n\r\nTo cook the chicken: Heat a nonstick skillet or grill pan over high heat",
        " Add the chicken breasts and cook, turning once, until well browned, about 4 to 5 minutes on each side or until cooked through",
        " Let the chicken rest on a cutting board for a few minutes before slicing it into thin strips"
    ],
    "publishedAt": "2021-01-17T19:28:52.803062+01:00"
}


endpoint = f"http://127.0.0.1:8080/recipes/"

post_request = requests.post(endpoint, json=data)

response = post_request.json()
json_format = json.dumps(response, indent=4, sort_keys=True)

print(json_format)
