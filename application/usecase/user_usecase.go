package usecase

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
	"golang.org/x/sync/errgroup"
)

// UserUseCase :
type UserUseCase interface {
	GetMe(c *gin.Context) (*output.UserResBody, error)
	GetByID(c *gin.Context) (*output.UserResBody, error)
	Create(c *gin.Context) (*models.UserDetail, error)
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type userUsecase struct {
	userRepository         repository.UserRepository
	userDetailRepository   repository.UserDetailRepository
	userService            service.IUserService
	favoriteGameRepository repository.FavoriteGameRepository
}

// NewUserUsecase :
func NewUserUsecase(
	uR repository.UserRepository,
	uDR repository.UserDetailRepository,
	uS service.IUserService,
	fGR repository.FavoriteGameRepository) UserUseCase {
	return &userUsecase{
		userRepository:         uR,
		userDetailRepository:   uDR,
		userService:            uS,
		favoriteGameRepository: fGR,
	}
}

func (uU userUsecase) GetMe(c *gin.Context) (*output.UserResBody, error) {
	claims, err := auth.GetTokenClaims(c)
	if err != nil {
		return nil, errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	claimsID := claims["sub"].(string)
	u, err := uU.userRepository.FindByID(claimsID)
	if err != nil {
		return nil, err
	}
	ok, err := uU.userService.IsDelete(claimsID)
	if !ok {
		return nil, errors.ErrNotFound{}
	}
	uD, err := uU.userDetailRepository.FindByID(claimsID)
	if err != nil {
		return nil, err
	}
	fG, err := uU.favoriteGameRepository.FindByID(claimsID)
	if err != nil {
		return nil, err
	}
	b := &output.UserResBody{
		UserID:    u.UserID,
		UserName:  uD.UserName,
		Icon:      uD.Icon,
		CreatedAt: u.CreatedAt,
	}
	for _, g := range fG {
		d := output.FavoriteGamesRes{
			GameTitle: g.GameTitle,
			CreatedAt: g.CreatedAt,
		}
		b.FavoriteGames = append(b.FavoriteGames, d)
	}
	return b, nil
}

func (uU userUsecase) GetByID(c *gin.Context) (*output.UserResBody, error) {
	uid := c.Params.ByName("id")
	ok, err := uU.userService.IsDelete(uid)
	if !ok {
		return nil, errors.ErrNotFound{}
	}
	u, err := uU.userRepository.FindByID(uid)
	if err != nil {
		return nil, err
	}
	uD, err := uU.userDetailRepository.FindByID(uid)
	if err != nil {
		return nil, err
	}
	fG, err := uU.favoriteGameRepository.FindByID(uid)
	if err != nil {
		return nil, err
	}
	b := &output.UserResBody{
		UserID:   u.UserID,
		UserName: uD.UserName,
		Icon:     uD.Icon,
	}
	for _, g := range fG {
		r := output.FavoriteGamesRes{
			GameTitle: g.GameTitle,
			CreatedAt: g.CreatedAt,
		}
		b.FavoriteGames = append(b.FavoriteGames, r)
	}
	return b, nil
}

func (uU userUsecase) Create(c *gin.Context) (*models.UserDetail, error) {
	b := input.CreateReqBody{}
	if err := c.Bind(&b); err != nil {
		return nil, errors.ErrCreateReqBinding{UserName: b.UserName, Email: b.Email}
	}
	pid, err := c.Cookie("pid")
	if err != nil {
		return nil, errors.ErrCoockie{Key: "pid", Value: pid}
	}
	u, err := models.NewUser(b.Email)
	if err != nil {
		return nil, err
	}
	provider := c.Query("provider")
	switch provider {
	case "google":
		ok, err := uU.userService.IsExist(provider, pid)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.ErrIDAlreadyExists{}
		}
		u.GoogleID = sql.NullString{String: pid, Valid: true}
	default:
		return nil, errors.ErrUnexpectedQueryProvider{Provider: provider}
	}
	ud, err := models.NewUserDetail(u.UserID, b.UserName)
	if err != nil {
		return nil, err
	}

	if err := uU.userRepository.Insert(u); err != nil {
		return nil, err
	}
	if err := uU.userDetailRepository.Insert(ud); err != nil {
		return nil, err
	}
	return ud, nil
}

//todo 存在しないユーザーでも正常処理される
func (uU userUsecase) Update(c *gin.Context) error {
	b := input.UpdateReqBody{}
	if err := c.Bind(&b); err != nil {
		return errors.ErrUpdateReqBinding{UserName: b.UserName, Icon: b.Icon, FavoriteGames: b.FavoriteGames}
	}
	userID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaims(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	claimsID := claims["sub"].(string)
	if userID != claimsID {
		return errors.ErrParams{Need: claimsID, Got: userID}
	}
	if err := uU.userDetailRepository.Update(claimsID, b.UserName, b.Icon); err != nil {
		return err
	}
	bfg, err := uU.favoriteGameRepository.FindByID(claimsID)
	if err != nil {
		return err
	}
	beforeFavoriteGames := []input.FavoriteGameRecord{}
	for _, v := range bfg {
		beforeFavoriteGames = append(beforeFavoriteGames, input.FavoriteGameRecord{GameTitle: v.GameTitle})
	}
	afterFavoriteGames := b.FavoriteGames
	var insertGames []input.FavoriteGameRecord
	var deleteGames []input.FavoriteGameRecord
	for _, a := range afterFavoriteGames {
		if !contains(beforeFavoriteGames, a.GameTitle) {
			insertGames = append(insertGames, a)
		}
	}
	for _, a := range beforeFavoriteGames {
		if !contains(afterFavoriteGames, a.GameTitle) {
			deleteGames = append(deleteGames, a)
		}
	}
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, i := range insertGames {
		f, err := models.NewFavoriteGame(claimsID, i.GameTitle)
		if err != nil {
			return err
		}
		eg.Go(func() error {
			if err := uU.favoriteGameRepository.Insert(f); err != nil {
				return err
			}
			return nil
		})
	}
	for _, g := range deleteGames {
		g := g
		eg.Go(func() error {
			if err := uU.favoriteGameRepository.Delete(claimsID, g.GameTitle); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		cancel()
		return err
	}
	return nil
}

func (uU userUsecase) Delete(c *gin.Context) error {
	userID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaims(c)
	if err != nil {
		return err
	}
	claimsID := claims["sub"].(string)
	if userID != claimsID {
		return errors.ErrParams{Need: claimsID, Got: userID}
	}
	if err := uU.userRepository.Delete(claimsID); err != nil {
		return err
	}
	return nil
}

func contains(gr []input.FavoriteGameRecord, id string) bool {
	for _, r := range gr {
		if id == r.GameTitle {
			return true
		}
	}
	return false
}
