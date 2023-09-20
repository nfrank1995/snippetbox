package main

import "github.com/nfrank1995/snippetbox/internal/models"

type templateData struct {
  Snippet models.Snippet
  Snippets []models.Snippet
}
