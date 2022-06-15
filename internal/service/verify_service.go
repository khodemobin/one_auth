package service

type verifyService struct {
}

//
//func NewVerifyService(repo *repository.Repository) VerifyService {
//	return &verifyService{}
//}
//
//func (v verifyService) Request(ctx context.Context, req request.VerifyRequest) error {
//	// TODO send verify code
//	// TODO check send limit
//	exists, err := r.repo.UserRepo.ExistsByPhone(ctx, phone)
//	if err != nil {
//		panic(fmt.Sprintf("internal error, can not check user exists in  db. err : %s", err.Error()))
//	} else if exists {
//		return errors.New("phone taken before")
//	}
//
//	err = r.repo.ConfirmCodeRepo.Create(phone)
//	if err := r.repo.ActivityRepo.Create(ac); err != nil {
//		panic(err)
//	}
//	return err
//}
//
//func (v verifyService) Verify(ctx context.Context, req request.VerifyConfirmRequest) (*Auth, error) {
//	// TODO check limits
//	exists, err := r.repo.UserRepo.ExistsByPhone(ctx, phone)
//	if err != nil {
//		panic(fmt.Sprintf("internal error, can not check user exists in  db. err : %s", err.Error()))
//	}
//
//	if exists {
//		return nil, errors.New("user verified!")
//	}
//
//	if err := r.checkConfirmCode(phone, code); err != nil {
//		return nil, err
//	}
//
//	user := r.createUser(ctx, phone)
//	refreshToken, token := r.generateToken(ctx, user)
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go func() {
//		if err := r.repo.ConfirmCodeRepo.Delete(phone); err != nil {
//			panic(fmt.Sprintf("internal error, can not delete confirm code. err : %s", err.Error()))
//		}
//		wg.Done()
//	}()
//
//	go func() {
//		if err := r.repo.ActivityRepo.Create(ac); err != nil {
//			panic(fmt.Sprintf("internal error, can not create activity log. err : %s", err.Error()))
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//
//	return &Auth{
//		Token: token,
//		RefreshToken: model.RefreshToken{
//			Token: refreshToken.Token,
//		},
//		ExpiresIn: 3600, // 1 hour
//		ID:        user.UUID,
//	}, nil
//}

//func GenerateConfirmCode(phone string) (*model.ConfirmCode, int, error) {
//	rand.Seed(time.Now().UnixNano())
//	min := 100000
//	max := 900000
//	code := rand.Intn(max-min+1) + min
//	hash, err := Hash(fmt.Sprint(code))
//	if err != nil {
//		return nil, 0, err
//	}
//	time := 3 * time.Minute
//
//	return &model.ConfirmCode{
//		Phone:     phone,
//		Hash:      hash,
//		ExpiresIn: time,
//	}, code, err
//}
