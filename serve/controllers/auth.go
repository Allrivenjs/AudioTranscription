package controllers

import (
	"AudioTranscription/serve/models"
	"AudioTranscription/serve/repository"
	"AudioTranscription/serve/security"
	"AudioTranscription/serve/util"
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/asaskevich/govalidator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AuthController interface {
	// swagger:route POST /auth/signup Auth SignUp
	// Sign up to the application
	// responses:
	// 	201: userResponse
	SignUp(ctx *fiber.Ctx) error
	// swagger:route POST /auth/signin Auth SignIn
	// Sign in to the application
	// responses:
	// 	200: userResponse
	SignIn(ctx *fiber.Ctx) error
	// swagger:route GET /auth/{id} Auth GetUser
	// Get a user by id
	// responses:
	// 	200: userResponse
	GetUser(ctx *fiber.Ctx) error
	// swagger:route GET /auth Auth GetUsers
	// Get all users
	// responses:
	// 	200: userResponse
	GetUsers(ctx *fiber.Ctx) error
	// swagger:route PUT /auth/{id} Auth PutUser
	// Update a user by id
	// responses:
	// 	200: userResponse
	PutUser(ctx *fiber.Ctx) error
	// swagger:route DELETE /auth/{id} Auth DeleteUser
	// Delete a user by id
	// responses:
	// 	204: noContentResponse
	DeleteUser(ctx *fiber.Ctx) error
}

// swagger:body listOneService
type SignIn struct {
	// In: body
	Email string `valid:"required,stringlength(3|100)" json:"email"`
	// In: body
	Password string `valid:"required" json:"password"`
}

type authController struct {
	usersRepo repository.UsersRepository
}

func NewAuthController(usersRepo repository.UsersRepository) AuthController {
	return &authController{usersRepo}
}

type SignUp struct {
	Email    string `valid:"required,stringlength(3|100)" json:"email"`
	Password string `valid:"required,stringlength(3|100)" json:"password"`
}

func (c *authController) SignUp(ctx *fiber.Ctx) error {
	var data SignUp
	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("Error parsing body")
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	var newUser models.User
	newUser.Email = util.NormalizeEmail(data.Email)

	if !govalidator.IsEmail(newUser.Email) {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrInvalidEmail))
	}

	exists, err := c.usersRepo.GetByEmail(newUser.Email)
	fmt.Println(fmt.Sprintf("exists: %v and error is %v", exists, err))
	if err.Error() == "record not found" {
		fmt.Println("password: ", data.Password)
		if strings.TrimSpace(data.Password) == "" {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyPassword))
		}
		newUser.Password, err = security.EncryptPassword(data.Password)

		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = newUser.CreatedAt
		err = c.usersRepo.SaveOrUpdate(&newUser)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		return ctx.
			Status(http.StatusCreated).
			JSON(newUser)
	}

	fmt.Println("user")
	if exists != nil {
		err = util.ErrEmailAlreadyExists
	}

	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

// SignIn godoc
// @Summary Sign in to the application
// @Description sign in to the application
// @Tags auth
// @Accept  json
// @Produce  json
// @body SignIn true "User object that needs to be added"
// @Success 200 {array} models.User
// @Router /auth/signup [post]
func (c *authController) SignIn(ctx *fiber.Ctx) error {
	var input SignIn

	err := ctx.BodyParser(&input)

	if err != nil {
		fmt.Println(err, "input")
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.ErrorResponse(util.NewJError(err)))
	}
	validateErrors := util.ValidateInput(ctx, input)
	if validateErrors != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(util.ErrorResponse(validateErrors))
	}
	input.Email = util.NormalizeEmail(input.Email)
	user, err := c.usersRepo.GetByEmail(input.Email)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.ErrorResponse(util.NewJError(err)))
	}
	err = security.VerifyPassword(user.Password, input.Password)

	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.ErrorResponse(util.NewJError(util.ErrInvalidCredentials)))
	}

	token, err := security.NewToken(strconv.Itoa(int(user.ID)))

	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.ErrorResponse(util.NewJError(err)))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(util.SuccessResponse(&fiber.Map{
			"user":  user,
			"token": fmt.Sprintf("Bearer %s", token),
		}))
}

func (c *authController) GetUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	user, err := c.usersRepo.GetById(payload.Id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(user)
}

func (c *authController) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.usersRepo.GetAll()
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(users)
}

func (c *authController) PutUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	var update models.User
	err = ctx.BodyParser(&update)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}
	update.Email = util.NormalizeEmail(update.Email)
	if !govalidator.IsEmail(update.Email) {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrInvalidEmail))
	}
	exists, err := c.usersRepo.GetByEmail(update.Email)
	if errors.Is(err, mgo.ErrNotFound) || strconv.Itoa(int(exists.ID)) == payload.Id {
		user, err := c.usersRepo.GetById(payload.Id)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		user.Email = update.Email
		user.UpdatedAt = time.Now()
		err = c.usersRepo.SaveOrUpdate(user)
		if err != nil {
			return ctx.
				Status(http.StatusUnprocessableEntity).
				JSON(util.NewJError(err))
		}
		return ctx.
			Status(http.StatusOK).
			JSON(user)
	}

	if exists != nil {
		err = util.ErrEmailAlreadyExists
	}

	return ctx.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))
}

func (c *authController) DeleteUser(ctx *fiber.Ctx) error {
	payload, err := AuthRequestWithId(ctx)
	if err != nil {
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	err = c.usersRepo.Delete(payload.Id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	ctx.Set("Entity", payload.Id)
	return ctx.SendStatus(http.StatusNoContent)
}

func AuthRequestWithId(ctx *fiber.Ctx) (*jwt.StandardClaims, error) {
	id := ctx.Params("id")
	if !bson.IsObjectIdHex(id) {
		return nil, util.ErrUnauthorized
	}
	token := ctx.Locals("user").(*jwt.Token)
	payload, err := security.ParseToken(token.Raw)
	if err != nil {
		return nil, err
	}
	if payload.Id != id || payload.Issuer != id {
		return nil, util.ErrUnauthorized
	}
	return payload, nil
}
