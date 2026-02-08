package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInitialLocationAreaURL(t *testing.T) {
	client := NewClient("https://example.com/api/v2/")
	expected := "https://example.com/api/v2/location-area/?offset=0&limit=20"
	actual := client.InitialLocationAreaURL()
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestGetLocationAreas(t *testing.T) {
	t.Run("get location areas pass", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp := LocationAreaResponse{
				Next: "next-url",
				Results: []LocationAreaEntry{
					{Name: "pallet-town", URL: "url1"},
					{Name: "viridian-city", URL: "url2"},
				},
			}
			json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := NewClient(server.URL + "/")
		locations, err := client.GetLocationAreas(server.URL + "/location-area/")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if locations.Next != "next-url" {
			t.Errorf("expected Next to be 'next-url', got '%s'", locations.Next)
		}
		if len(locations.Results) != 2 {
			t.Fatalf("expected 2 results, got %d", len(locations.Results))
		}
		if locations.Results[0].Name != "pallet-town" {
			t.Errorf("expected first result to be 'pallet-town', got '%s'", locations.Results[0].Name)
		}
	})

	t.Run("location areas server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		client := NewClient(server.URL + "/")
		_, err := client.GetLocationAreas(server.URL + "/location-area/")
		if err == nil {
			t.Error("expected error for 500 response, got nil")
		}
	})
}

func TestGetLocationAreaPokemon(t *testing.T) {
	t.Run("get pokemon pass", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp := LocationAreaDetailResponse{
				PokemonEncounters: []PokemonEncounter{
					{Pokemon: PokemonEntry{Name: "pikachu"}},
					{Pokemon: PokemonEntry{Name: "bulbasaur"}},
				},
			}
			json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := NewClient(server.URL + "/")
		details, err := client.GetLocationAreaPokemon("pallet-town")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(details.PokemonEncounters) != 2 {
			t.Fatalf("expected 2 encounters, got %d", len(details.PokemonEncounters))
		}
		if details.PokemonEncounters[0].Pokemon.Name != "pikachu" {
			t.Errorf("expected 'pikachu', got '%s'", details.PokemonEncounters[0].Pokemon.Name)
		}
	})

	t.Run("get pokemon not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		client := NewClient(server.URL + "/")
		_, err := client.GetLocationAreaPokemon("fake-town")
		if err == nil {
			t.Fatal("expected error for 404 response, got nil")
		}
		if !strings.Contains(err.Error(), "fake-town") {
			t.Errorf("expected error to contain 'fake-town', got '%s'", err.Error())
		}
	})
}

func TestCatchPokemon(t *testing.T) {
	t.Run("catch pokemon not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		client := NewClient(server.URL + "/")
		_, err := client.CatchPokemon("missingno")
		if err == nil {
			t.Fatal("expected error for 404 response, got nil")
		}
		if !strings.Contains(err.Error(), "missingno") {
			t.Errorf("expected error to contain 'missingno', got '%s'", err.Error())
		}
	})
}
