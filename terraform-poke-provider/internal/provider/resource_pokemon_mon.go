package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePokemon() *schema.Resource {
	return &schema.Resource{
		Create: resourcePokemonCreate,
		Read:   resourcePokemonRead,
		Update: resourcePokemonUpdate,
		Delete: resourcePokemonDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"level": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

func resourcePokemonCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	level := d.Get("level").(int)

	payload := map[string]interface{}{
		"name":  name,
		"level": level,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Sending POST to Flask with payload: %s", string(body))
	resp, err := http.Post("http://localhost:5000/api/pokemon", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error making POST request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API returned unexpected status %d: %s", resp.StatusCode, string(b))
	}

	d.SetId(name)
	return resourcePokemonRead(d, m)
}

func resourcePokemonRead(d *schema.ResourceData, m interface{}) error {
	// No endpoint GET disponible, se asume que el recurso existe si tiene ID.
	return nil
}

func resourcePokemonUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	level := d.Get("level").(int)

	payload := map[string]interface{}{
		"name":  name,
		"level": level,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Sending PUT to Flask with payload: %s", string(body))
	req, err := http.NewRequest("PUT", "http://localhost:5000/api/pokemon", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making PUT request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API returned unexpected status %d: %s", resp.StatusCode, string(b))
	}

	return resourcePokemonRead(d, m)
}

func resourcePokemonDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Id()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:5000/api/pokemon/%s", name), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	log.Printf("[INFO] Sending DELETE to Flask for: %s", name)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotFound {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API returned unexpected status %d: %s", resp.StatusCode, string(b))
	}

	d.SetId("")
	return nil
}
