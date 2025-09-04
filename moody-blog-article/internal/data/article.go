package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"moody-blog-article/internal/biz"
	"time"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func newArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	_, err := ar.data.mdb.Article.Create().SetTitle(article.Title).SetContent(article.Content).Save(ctx)
	if err != nil {

		return err
	}
	return err
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) (*biz.Article, error) {
	art, err := ar.data.mdb.Article.UpdateOneID(id).SetTitle(article.Title).SetContent(article.Content).SetUpdateAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Article{
		Id:         art.ID,
		Title:      art.Title,
		Content:    art.Content,
		Created_at: art.CreatedAt,
		Updated_at: art.UpdateAt,
	}, err
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return ar.data.mdb.Article.DeleteOneID(id).Exec(ctx)
}

func (ar *articleRepo) GetArticleById(ctx context.Context, id int64) (*biz.Article, error) {
	art, err := ar.data.mdb.Article.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Article{
		Id:         art.ID,
		Title:      art.Title,
		Content:    art.Content,
		Created_at: art.CreatedAt,
		Updated_at: art.UpdateAt,
	}, err
}

func (ar *articleRepo) ListArticles(ctx context.Context) ([]*biz.Article, error) {
	articles, err := ar.data.mdb.Article.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	listArticles := make([]*biz.Article, 0, len(articles))
	for _, art := range articles {
		listArticles = append(listArticles, &biz.Article{
			Id:         art.ID,
			Title:      art.Title,
			Content:    art.Content,
			Created_at: art.CreatedAt,
			Updated_at: art.UpdateAt,
		})
	}
	return listArticles, err
}

func stringLike(id int64) string {
	return fmt.Sprintf("like:%d", id)
}

func (ar *articleRepo) GetArticleLink(ctx context.Context, id int64) (int64, error) {
	count := ar.data.rdb.Get(ctx, stringLike(id))
	return count.Int64()
}

func (ar *articleRepo) IncreaseLikeCount(ctx context.Context, id int64) error {
	_, err := ar.data.rdb.Incr(ctx, stringLike(id)).Result()
	return err
}
