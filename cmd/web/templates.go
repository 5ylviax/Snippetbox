package main

import "snippetbox.silviasaavedra.net/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
