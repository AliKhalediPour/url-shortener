package errors

import "github.com/gofiber/fiber/v2"

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	switch t := err.(type) {
	case *BadRequestError:
		return ctx.Status(400).JSON(map[string]any{
			"message": t.Message,
		})

	case *InternalError:
		return ctx.Status(500).JSON(map[string]any{
			"message": t.Message,
		})

	case *NotFoundError:
		return ctx.Status(404).JSON(map[string]any{
			"message": t.Message,
		})

	default:
		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(map[string]any{
			"message": err.Error(),
		})
	}
}

type BadRequestError struct {
	Message string
}

func (b *BadRequestError) Error() string {
	return b.Message
}

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewBadRequest(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
	}
}

func NewInternalError(message string) *InternalError {
	return &InternalError{
		Message: message,
	}
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}
