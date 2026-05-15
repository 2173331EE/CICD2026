package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Autorise les requêtes venant du frontend Svelte
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	ctx := context.Background()

	var store ClassStore
	if dynamoEnabled() {
		client, err := newDynamoClient(ctx)
		if err != nil {
			e.Logger.Fatal(err)
		}
		table := strings.TrimSpace(os.Getenv("DYNAMODB_TABLE"))
		if table == "" {
			table = "boat_classes"
		}
		if err := ensureBoatClassesTable(ctx, client, table); err != nil {
			e.Logger.Fatal(err)
		}
		store = NewDynamoClassStore(client, table)
	} else {
		store = NewMemoryClassStore()
	}

	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Backend Go fonctionne !")
	})

	e.GET("/api/classes", func(c echo.Context) error {
		classes, err := store.List(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, classes)
	})

	e.POST("/api/classes", func(c echo.Context) error {
		var newClass BoatClass
		if err := c.Bind(&newClass); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid JSON payload"})
		}

		newClass.Name = strings.TrimSpace(newClass.Name)
		newClass.HandicapType = strings.ToUpper(strings.TrimSpace(newClass.HandicapType))

		if newClass.Name == "" || newClass.HandicapType == "" || newClass.HandicapValue <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "name, handicapType and handicapValue are required"})
		}

		err := store.Add(c.Request().Context(), newClass)
		if err != nil {
			if err == ErrClassAlreadyExists {
				return c.JSON(http.StatusConflict, map[string]string{"error": "class already exists"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, newClass)
	})

	e.DELETE("/api/classes/:name", func(c echo.Context) error {
		name := strings.TrimSpace(c.Param("name"))
		if name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "class name is required"})
		}

		deleted, ok, err := store.DeleteByName(c.Request().Context(), name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !ok {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "class not found"})
		}
		return c.JSON(http.StatusOK, deleted)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
