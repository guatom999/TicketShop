package middlewareRepositories

type MiddlwareRepositoryService interface {
}

type middlewareRepository struct {
}

func NewMiddlewareRepository() MiddlwareRepositoryService {
	return &middlewareRepository{}
}

// func (r *middlewareRepository) AccessTokenSearch(pctx context.Context, accessToken string) error {
// 	_, cancel := context.WithTimeout(pctx, time.Second*15)
// 	defer cancel()

// 	return nil

// }
