import random
import requests

pokedex = []

def add_pokemon(name, level):
    res = requests.get(f"https://pokeapi.co/api/v2/pokemon/{name.lower()}")
    data = res.json()

    is_shiny = random.randint(1, 10) == 1
    poke = {
        "name": data["name"].capitalize(),
        "img": data["sprites"]["front_shiny"] if is_shiny else data["sprites"]["front_default"],
        "weight": data["weight"],
        "height": data["height"],
        "description": get_pokemon_data(data["species"]["url"]),
        "level": level,
        "is_shiny": is_shiny
    }
    pokedex.append(poke)
    return poke

def remove_pokemon(name):
    global pokedex
    before = len(pokedex)
    pokedex = [p for p in pokedex if p["name"].lower() != name.lower()]
    return len(pokedex) != before

def get_pokemon_data(url):
    res = requests.get(url)
    data = res.json()
    entries = data["flavor_text_entries"]
    for entry in entries:
        if entry["language"]["name"] == "en":
            return entry["flavor_text"].replace("\n", " ").replace("\f", " ")
    return "No description available."


def update_pokemon(name, level):
    for poke in pokedex:
        if poke["name"].lower() == name.lower():
            poke["level"] = level
            return poke
    return None