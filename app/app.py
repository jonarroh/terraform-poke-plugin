
from flask import Flask, request, jsonify, render_template
from poke import pokedex, add_pokemon, remove_pokemon, get_pokemon_data, update_pokemon

app = Flask(__name__)

@app.route("/")
def index():
    return render_template("index.html", pokemons=pokedex)


@app.route("/api/pokemons", methods=["GET"])
def api_list():
    from poke import pokedex
    return jsonify(pokedex)

@app.route("/api/pokemon", methods=["POST"])
def api_add():
    data = request.json
    name = data.get("name")
    level = data.get("level", 1)
    poke = add_pokemon(name, level)
    return jsonify(poke), 201

@app.route("/api/pokemon/<name>", methods=["DELETE"])
def api_delete(name):
    success = remove_pokemon(name)
    return ("", 204) if success else ("Not found", 404)

@app.route("/api/pokemon", methods=["PUT"])
def api_update():
    data = request.json
    name = data.get("name")
    level = data.get("level", 1)
    poke = update_pokemon(name, level)
    return jsonify(poke), 200

if __name__ == "__main__":
    app.run(debug=True)
