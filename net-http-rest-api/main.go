package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product represents our resource
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name" validate:"required,min=2"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

// In-memory store with a mutex for thread safety
type Store struct {
	mu       sync.RWMutex
	products map[string]Product
}

var validate = validator.New()

// Helper: JSON response wrapper
func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("failed to encode response: %v", err)
		}
	}
}

// Helper: Error response wrapper
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// Handler struct containing dependencies
type ProductHandler struct {
	store *Store
}

// GET /products - List all products
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.store.mu.RLock()
	defer h.store.mu.RUnlock()

	productList := make([]Product, 0, len(h.store.products))
	for _, p := range h.store.products {
		productList = append(productList, p)
	}

	respondJSON(w, http.StatusOK, productList)
}

// GET /products/{id} - Get product by ID
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// r.PathValue extracts wildcards in Go 1.22+
	id := r.PathValue("id")

	h.store.mu.RLock()
	product, exists := h.store.products[id]
	h.store.mu.RUnlock()

	if !exists {
		respondError(w, http.StatusNotFound, "Product not found")
		return
	}

	respondJSON(w, http.StatusOK, product)
}

// POST /products - Create a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate struct fields
	if err := validate.Struct(p); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.store.mu.Lock()
	defer h.store.mu.Unlock()

	if _, exists := h.store.products[p.ID]; exists {
		respondError(w, http.StatusConflict, "Product ID already exists")
		return
	}

	h.store.products[p.ID] = p
	respondJSON(w, http.StatusCreated, p)
}

// PUT /products/{id} - Update product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p.ID = id
	if err := validate.Struct(p); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.store.mu.Lock()
	defer h.store.mu.Unlock()

	if _, exists := h.store.products[id]; !exists {
		respondError(w, http.StatusNotFound, "Product not found")
		return
	}

	h.store.products[id] = p
	respondJSON(w, http.StatusOK, p)
}

// DELETE /products/{id} - Delete product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	h.store.mu.Lock()
	defer h.store.mu.Unlock()

	if _, exists := h.store.products[id]; !exists {
		respondError(w, http.StatusNotFound, "Product not found")
		return
	}

	delete(h.store.products, id)
	respondJSON(w, http.StatusNoContent, nil)
}

// Middleware: Logging middleware
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Initialize store
	store := &Store{
		products: map[string]Product{
			"101": {ID: "101", Name: "Mechanical Keyboard", Price: 89.99},
			"102": {ID: "102", Name: "Wireless Mouse", Price: 29.99},
		},
	}

	handler := &ProductHandler{store: store}

	// Create Go 1.22+ ServeMux with HTTP method routing
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", handler.GetProducts)
	mux.HandleFunc("GET /products/{id}", handler.GetProductByID)
	mux.HandleFunc("POST /products", handler.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", handler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", handler.DeleteProduct)

	// Wrap entire mux with logging middleware
	wrappedMux := loggerMiddleware(mux)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      wrappedMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server listening on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed: %v", err)
	}
}
