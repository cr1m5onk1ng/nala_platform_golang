package controllers

import (
	"strconv"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
)

// UTILS
func parseResourceData(ctx *fiber.Ctx) (*domain.MappedResource, error) {
	resourceData, err := validation.ValidateBodyData(ctx, &domain.MappedResource{})
	if err != nil {
		return nil, err
	}
	post, ok := resourceData.(*domain.MappedResource)

	if !ok {
		return nil, SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"Could not cast input data to data model",
		)
	}
	return post, nil
}

// CONTROLLERS
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
	var err error
	var resource db.Resource

	err = validation.ValidateResourceUrl(url)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusBadRequest, "invalid url")
	}

	resource, err = h.Repo.GetResourceByUrl(ctx.Context(), url)
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
	usrId := ctx.Params("usr-id")

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
	postId, err := strconv.ParseInt(ctx.Params("post-id"), 10, 64)
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
	resource, err := parseResourceData(ctx)
	if err != nil {
		return err
	}

	args := db.AddResourceParams{
		Url:       resource.Url,
		Language:  resource.Language,
		MediaType: resource.MediaType,
		Category:  resource.Category,
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

func (h *Handlers) AddResourceNotSecure(ctx *fiber.Ctx) error {
	resource, err := parseResourceData(ctx)
	if err != nil {
		return err
	}

	args := db.AddResourceParams{
		Url:       resource.Url,
		Language:  resource.Language,
		MediaType: resource.MediaType,
		Category:  resource.Category,
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
	resource, err := validation.ValidateResourceDataAndUrlAndAuthorization(ctx, &domain.MappedResource{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error occured: " + err.Error(),
			"data":    nil,
		})
	}

	args := db.UpdateResourceParams{
		Url:       resource.Url,
		Language:  resource.Language,
		MediaType: resource.MediaType,
		Category:  resource.Category,
	}

	editedResource, err := h.Repo.UpdateResourceTrans(ctx.Context(), args)

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
