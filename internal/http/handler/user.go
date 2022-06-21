package handler

//type UserHandler struct {
//	UserService service.UserService
//}

//func (u *UserHandler) Me(c *fiber.Ctx) error {
//	uuid := c.Locals("user_uuid")
//	user, err := u.UserService.Me(c.Context(), uuid.(string))
//	if err != nil {
//		return c.Status(http.StatusUnauthorized).JSON(utils.DefaultResponse(nil, "", 0))
//	}
//
//	data := &model.UserResource{
//		Phone: *user.Phone,
//		UUID:  user.UUID,
//	}
//
//	return c.JSON(utils.DefaultResponse(data, "", 1))
//}
//
//func (u *UserHandler) Update(c *fiber.Ctx) error {
//	req := new(request.UserUpdateRequest)
//	uuid := c.Locals("user_uuid")
//	if err := c.BodyParser(req); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(utils.DefaultResponse(nil, err.Error(), 0))
//	}
//
//	errors := validator.Check(*req)
//	if errors != nil {
//		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
//	}
//
//	err := u.UserService.Update(c.Context(), uuid.(string), req.Password, req.ConfirmPassword, createActivity(c))
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(utils.DefaultResponse(nil, err.Error(), 0))
//	}
//
//	return c.JSON(utils.DefaultResponse("", "", 1))
//}
