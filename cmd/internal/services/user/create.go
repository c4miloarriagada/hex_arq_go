package services

func CreateUserService(user domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := db.GetDB()
	if err != nil {
		return nil, err 
	}

	done := make(chan error, 1)
	var createdUser domain.User 

	go func() {
		var existingUser domain.User
		if err := database.WithContext(ctx).Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			done <- errors.New("user with this email already exists")
			return
		}

		if err := database.WithContext(ctx).Create(&user).Error; err != nil {
			done <- err
			return
		}

		createdUser = user 
		done <- nil
	}()

	select {
	case <-ctx.Done():
		log.Println("Request timed out")
		return nil, ctx.Err() 
	case err := <-done:
		if err != nil {
			log.Printf("Error creating user: %v", err)
			return nil, err 
		}
	}

	log.Println("User created successfully")
	return &createdUser, nil 
