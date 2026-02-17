package user

import (
	"context"
	"sikabiz/user-importer/internal/domain"
	"sync"
)

func (u *userService) ImportUsers(ctx context.Context, users []domain.User, workerCount int) []error {
	jobs := make(chan domain.User, 1000)
	var wg sync.WaitGroup
	var errs []error

	worker := func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case user, ok := <-jobs:
				if !ok {
					return
				}

				// inserting user in here
				userDomain, err := u.userRepository.InsertUser(ctx, user)
				if err != nil {
					errs = append(errs, err)
				}

				// insert address
				for _, address := range user.Addresses {
					address.UserId = userDomain.Id
					err := u.addressRepository.InsertAddress(ctx, address)
					if err != nil {
						errs = append(errs, err)
					}
				}
			}
		}
	}

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker()
	}

	// Feed jobs
	for _, user := range users {
		select {
		case <-ctx.Done():
			return errs
		case jobs <- user:
		}
	}

	close(jobs)
	wg.Wait()

	return errs
}

func (u *userService) GetUser(ctx context.Context, userId string) (*domain.User, error) {
	user, err := u.userRepository.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	addresses, err := u.addressRepository.GetAddressByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	user.Addresses = addresses
	return user, nil
}
