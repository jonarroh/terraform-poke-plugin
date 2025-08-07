terraform {
  required_providers {
    pokemon = {
      source  = "jonarroh/pokemon"
      version = "0.1.0"
    }
  }
}
provider "pokemon" {
  # configuración del provider si es que tiene parámetros (normalmente vacío)
}

resource "pokemon_pokemon" "pikachu" {
  name  = "pikachu"
  level = 40
}

resource "pokemon_pokemon" "charmander" {
  name  = "charmander"
  level = 15
}

resource "pokemon_pokemon" "bulbasaur" {
  name  = "bulbasaur"
}

resource "pokemon_pokemon" "umbreon" {
  name  = "umbreon"
  level = 10
}

resource "pokemon_pokemon" "eevee" {
  name  = "eevee"
  level = 5
}