package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sha1sof/Effective_Mobile_test/internal/db"
	"github.com/sha1sof/Effective_Mobile_test/internal/model"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := db.GetPeople()
	if err != nil {
		log.Printf("Error retrieving people: %v", err)
		http.Error(w, "Error retrieving people", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		log.Printf("Error encoding people data: %v", err)
		http.Error(w, "Error encoding people data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Invalid person ID: %v", err)
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	err = db.DeletePerson(personID)
	if err != nil {
		log.Printf("Error deleting person: %v", err)
		http.Error(w, "Error deleting person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var personRequest struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic,omitempty"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&personRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	age, err := getAge(personRequest.Name)
	if err != nil {
		log.Printf("Error getting age: %v", err)
		http.Error(w, "Error getting age", http.StatusInternalServerError)
		return
	}

	gender, err := getGender(personRequest.Name)
	if err != nil {
		log.Printf("Error getting gender: %v", err)
		http.Error(w, "Error getting gender", http.StatusInternalServerError)
		return
	}

	nationality, err := getNationality(personRequest.Name)
	if err != nil {
		log.Printf("Error getting nationality: %v", err)
		http.Error(w, "Error getting nationality", http.StatusInternalServerError)
		return
	}

	newPerson := model.Person{
		Name:        personRequest.Name,
		Surname:     personRequest.Surname,
		Patronymic:  personRequest.Patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
	}

	err = db.CreatePerson(&newPerson)
	if err != nil {
		log.Printf("Error creating person: %v", err)
		http.Error(w, "Error creating person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	var updatedPerson model.Person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedPerson); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updatedPerson.ID = uint(personID)

	err = db.UpdatePerson(&updatedPerson)
	if err != nil {
		log.Printf("Error updating person: %v", err)
		http.Error(w, "Error updating person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPerson)
}

func getAge(name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		log.Printf("Error getting age for %s: %v", name, err)
		return 0, err
	}
	defer resp.Body.Close()

	var agifyResponse struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}

	err = json.NewDecoder(resp.Body).Decode(&agifyResponse)
	if err != nil {
		log.Printf("Error decoding age response for %s: %v", name, err)
		return 0, err
	}

	return agifyResponse.Age, nil
}

func getGender(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		log.Printf("Error getting gender for %s: %v", name, err)
		return "", err
	}
	defer resp.Body.Close()

	var genderizeResponse struct {
		Count       int     `json:"count"`
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Probability float64 `json:"probability"`
	}
	err = json.NewDecoder(resp.Body).Decode(&genderizeResponse)
	if err != nil {
		log.Printf("Error decoding gender response for %s: %v", name, err)
		return "", err
	}

	return genderizeResponse.Gender, nil
}

func getNationality(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		log.Printf("Error getting nationality for %s: %v", name, err)
		return "", err
	}
	defer resp.Body.Close()

	var nationalizeResponse struct {
		Count   int    `json:"count"`
		Name    string `json:"name"`
		Country []struct {
			CountryID   string  `json:"country_id"`
			Probability float64 `json:"probability"`
		} `json:"country"`
	}
	err = json.NewDecoder(resp.Body).Decode(&nationalizeResponse)
	if err != nil {
		log.Printf("Error decoding nationality response for %s: %v", name, err)
		return "", err
	}

	var highestProbability float64
	var mostLikelyCountry string

	for _, country := range nationalizeResponse.Country {
		if country.Probability > highestProbability {
			highestProbability = country.Probability
			mostLikelyCountry = country.CountryID
		}
	}

	return mostLikelyCountry, nil
}

func Filter(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	patronymic := r.URL.Query().Get("patronymic")
	ageStr := r.URL.Query().Get("age")
	gender := r.URL.Query().Get("gender")
	nationality := r.URL.Query().Get("nationality")
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	pageSizeStr := r.URL.Query().Get("pageSize")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}

	age, _ := strconv.Atoi(ageStr)
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	people, err := db.FilterPeople(name, surname, patronymic, age, gender, nationality, page, pageSize)
	if err != nil {
		log.Printf("Error filtering people: %v", err)
		http.Error(w, "Error filtering people", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(people); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
