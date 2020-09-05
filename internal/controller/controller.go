package controller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/anrid/rkspce/internal/domain"
	"github.com/labstack/echo/v4"
)

// We use this in place of a real database.
type basketDB map[string]*domain.Basket

// Controller ...
type Controller struct {
	db  basketDB
	mux *sync.RWMutex
}

// New ...
func New() *Controller {
	return &Controller{
		db:  make(basketDB),
		mux: &sync.RWMutex{},
	}
}

// PostBasket ...
func (co *Controller) PostBasket(c echo.Context) error {
	b := domain.NewBasket()

	co.mux.Lock()
	co.db[b.ID] = b
	co.mux.Unlock()

	if c.QueryParam("format") == "txt" {
		return co.SendBasketAsPlainText(c, b)
	}

	return c.JSON(http.StatusOK, PostResponseV1{b})
}

// GetBasket ...
func (co *Controller) GetBasket(c echo.Context) error {
	id := c.Param("id")

	co.mux.RLock()
	b, found := co.db[id]
	co.mux.RUnlock()

	if !found {
		return c.JSON(http.StatusNotFound, ErrorResponse{"could not find basket id " + id})
	}

	if c.QueryParam("format") == "txt" {
		return co.SendBasketAsPlainText(c, b)
	}

	return c.JSON(http.StatusOK, GetResponseV1{b})
}

// PatchBasket ...
func (co *Controller) PatchBasket(c echo.Context) error {
	id := c.Param("id")
	code := c.Param("code")

	co.mux.RLock()
	basket, foundBasket := co.db[id]
	co.mux.RUnlock()

	if !foundBasket {
		return c.JSON(http.StatusNotFound, ErrorResponse{"could not find basket id " + id})
	}

	product, err := domain.GetProductByCode(code)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"could not find product with code " + code})
	}

	// Add product item to basket.
	basket.Add(product.ToItem())

	// Apply specials.
	domain.ApplySpecials(basket)

	// Update basket.
	co.mux.Lock()
	co.db[basket.ID] = basket
	co.mux.Unlock()

	if c.QueryParam("format") == "txt" {
		return co.SendBasketAsPlainText(c, basket)
	}

	return c.JSON(http.StatusOK, PatchResponseV1{basket})
}

// SetupRoutes ...
func (co *Controller) SetupRoutes(e *echo.Echo) {
	e.POST("/api/basket", co.PostBasket)
	e.GET("/api/basket/:id", co.GetBasket)
	e.PATCH("/api/basket/:id/product/:code", co.PatchBasket)
}

// PostResponseV1 ...
type PostResponseV1 struct {
	Basket *domain.Basket `json:"basket"`
}

// GetResponseV1 ...
type GetResponseV1 struct {
	Basket *domain.Basket `json:"basket"`
}

// PatchResponseV1 ...
type PatchResponseV1 struct {
	Basket *domain.Basket `json:"basket"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"error"`
}

// SendBasketAsPlainText ...
func (co *Controller) SendBasketAsPlainText(c echo.Context, b *domain.Basket) error {
	return c.String(
		http.StatusOK,
		fmt.Sprintf("Basket: %s\n%s", b.ID, b.Dump()),
	)
}
