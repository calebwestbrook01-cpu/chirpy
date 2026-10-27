package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	res, err := cfg.dbQueries.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch chirps", err)
		return
	}
	var chirps []Chirp
	for _, dat := range res {
		chirp := Chirp{
			ID:        dat.ID,
			CreatedAt: dat.CreatedAt,
			UpdatedAt: dat.UpdatedAt,
			Body:      dat.Body,
			UserID:    dat.UserID,
		}
		chirps = append(chirps, chirp)
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpGet(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("chirpID")
	parsedUUID, err := uuid.Parse(path)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid path", err)
		return
	}
	res, err := cfg.dbQueries.GetChirp(r.Context(), parsedUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}
	chirp := Chirp{
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Body:      res.Body,
		UserID:    res.UserID,
	}

	respondWithJSON(w, http.StatusOK, chirp)
}
