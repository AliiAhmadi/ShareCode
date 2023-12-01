package main

import "github.com/AliiAhmadi/ShareCode/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
