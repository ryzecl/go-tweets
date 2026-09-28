package post

import (
	"context"
	"go-tweets/internal/dto"
	"math"
	"net/http"
)

func (s *postService) GetAllPost(ctx context.Context, param *dto.GetAllPostRequest) (*dto.GetAllPostResponse, int, error) {
	// get total post
	totalPost, err := s.postRepo.TotalPost(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// get all post
	offset := param.Limit * (param.Page - 1)
	posts, err := s.postRepo.GetAllPost(ctx, param, int(offset))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// get all comments related to posts
	postIDs := make([]int64, len(posts))
	for _, post := range posts {
		postIDs = append(postIDs, post.ID)
	}
	comments, err := s.commentRepo.GetCommentsByPostIDs(ctx, postIDs)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// mapping all comments with post basen on key value of post_id
	commentsMap := make(map[int64][]dto.Comment)
	for _, comment := range comments {
		commentsMap[comment.PostID] = append(commentsMap[comment.PostID], dto.Comment{
			ID:        comment.ID,
			Username:  comment.Username,
			Content:   comment.Content,
			LikeCount: comment.LikeCount,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
		})
	}

	// mapping response
	var data []dto.DetailPostResponse
	for _, post := range posts {
		comments := commentsMap[post.ID]
		if comments == nil {
			comments = []dto.Comment{}
		}
		data = append(data, dto.DetailPostResponse{
			ID:        post.ID,
			Username:  post.Username,
			Title:     post.Title,
			Content:   post.Content,
			LikeCount: post.LikeCount,
			Comments:  comments,
			CreatedAt: post.CreatedAt.String(),
			UpdatedAt: post.UpdatedAt.String(),
		})
	}

	totalPage := int64(math.Ceil(float64(totalPost) / float64(param.Limit)))
	//
	// return
	result := dto.GetAllPostResponse{
		Limit:       param.Limit,
		CurrentPage: param.Page,
		TotalPage:   totalPage,
		Data:        data,
	}

	return &result, http.StatusOK, nil
}
