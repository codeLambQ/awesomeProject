package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Article struct {
	Id         int64
	Title      string
	Content    string
	Created_at time.Time
	Updated_at time.Time
	Like       int64
}
type ArticleRepo interface {
	CreateArticle(ctx context.Context, article *Article) (*Article, error)
	UpdateArticle(ctx context.Context, id int64, article *Article) (*Article, error)
	DeleteArticle(ctx context.Context, id int64) error
	GetArticleById(ctx context.Context, id int64) (*Article, error)
	ListArticles(ctx context.Context) ([]*Article, error)

	// redis
	GetArticleLink(ctx context.Context, id int64) (int64, error)
	IncreaseLikeCount(ctx context.Context, id int64) error
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (uc *ArticleUsecase) Create(ctx context.Context, article *Article) (art *Article, err error) {
	art, err = uc.repo.CreateArticle(ctx, article)
	if err != nil {
		return
	}
	return
}

func (uc *ArticleUsecase) Update(ctx context.Context, id int64, article *Article) (art *Article, err error) {
	return uc.repo.UpdateArticle(ctx, id, article)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (article *Article, err error) {
	article, err = uc.repo.GetArticleById(ctx, id)
	if err != nil {
		return
	}

	err = uc.repo.IncreaseLikeCount(ctx, id)
	if err != nil {
		return
	}
	article.Like, err = uc.repo.GetArticleLink(ctx, id)
	if err != nil {
		return
	}
	return
}
