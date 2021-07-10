package controllers

import (
	"strconv"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetResource(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	resource, err := h.Repo.GetResourceById(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "resource with the given ID was not found",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    resource,
	})

}

func (h *Handlers) GetResourceByUrl(ctx *fiber.Ctx) error {
	url := ctx.Params("url")

	// ADD URL VALIDATION

	resource, err := h.Repo.GetResourceByUrl(ctx.Context(), url)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    resource,
	})
}

func (h *Handlers) GetResourcesByLanguage(ctx *fiber.Ctx) error {
	lang := ctx.Params("lang")

	// ADD LANGUAGE VALIDATION

	resources, err := h.Repo.GetResourcesByLanguage(ctx.Context(), lang)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    resources,
	})
}

func (h *Handlers) GetResourcesPostsByUser(ctx *fiber.Ctx) error {
	usrId := ctx.Params("usrId")

	resources, err := h.Repo.GetResourcesPostsByUser(ctx.Context(), usrId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    resources,
	})
}

func (h *Handlers) GetResourcePost(ctx *fiber.Ctx) error {
	postId, err := strconv.ParseInt(ctx.Params("postId"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Something went wrong: " + err.Error(),
			"data":    nil,
		})
	}

	resource, err := h.Repo.GetResourcePost(ctx.Context(), postId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    resource,
	})
}

func (h *Handlers) AddResource(ctx *fiber.Ctx) error {
	resource, err := validation.ValidateResourceDataAndUrlAndAuthorization(ctx, &db.Resource{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error occured: " + err.Error(),
			"data":    nil,
		})
	}

	args := db.AddResourceParams{
		Url:          resource.Url,
		Language:     resource.Language,
		Difficulty:   resource.Difficulty,
		Title:        resource.Title,
		Description:  resource.Description,
		MediaType:    resource.MediaType,
		Category:     resource.Category,
		ThumbnailUrl: resource.ThumbnailUrl,
	}
	addedResource, err := h.Repo.AddResource(ctx.Context(), args)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error occured: " + err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   true,
		"message": nil,
		"data":    addedResource,
	})
}

func (h *Handlers) UpdateResource(ctx *fiber.Ctx) error {
	resource, err := validation.ValidateResourceDataAndUrlAndAuthorization(ctx, &db.Resource{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error occured: " + err.Error(),
			"data":    nil,
		})
	}

	args := db.UpdateResourceParams{
		Url:          resource.Url,
		Language:     resource.Language,
		Difficulty:   resource.Difficulty,
		Title:        resource.Title,
		Description:  resource.Description,
		MediaType:    resource.MediaType,
		Category:     resource.Category,
		ThumbnailUrl: resource.ThumbnailUrl,
	}

	editedResource, err := h.Repo.UpdateResource(ctx.Context(), args)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error occured: " + err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   true,
		"message": nil,
		"data":    editedResource,
	})
}
