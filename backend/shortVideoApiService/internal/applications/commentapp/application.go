package commentapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/errorx"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/commentoptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/useroptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	core *svcoreadapter.Adapter
}

func New(core *svcoreadapter.Adapter) *Application {
	return &Application{
		core: core,
	}
}

func (a *Application) generateCommentUserInfo(userInfo *v1.User) (commentUser *svapi.CommentUser) {
	if userInfo == nil {
		return commentUser
	}

	commentUser.Name = userInfo.Name
	commentUser.Avatar = userInfo.Avatar
	commentUser.Id = userInfo.Id
	// TODO: 增加是否已关注
	commentUser.IsFollowing = true
	return commentUser
}

func (a *Application) CreateComment(ctx context.Context, request *svapi.CreateCommentRequest) (*svapi.CreateCommentResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	created, err := a.core.CreateComment(
		ctx,
		commentoptions.CreateCommentWithUserId(userId),
		commentoptions.CreateCommentWithContent(request.Content),
		commentoptions.CreateCommentWithVideoId(request.VideoId),
		commentoptions.CreateCommentWithParentId(request.ParentId),
		commentoptions.CreateCommentWithReplyTo(request.ReplyUserId),
	)
	if err != nil {
		log.Context(ctx).Errorf("failed to create comment: %v", err)
		return nil, errorx.New(1, "创建评论失败")
	}

	userInfo, err := a.core.GetUserInfo(ctx, useroptions.GetUserInfoWithUserId(userId))
	if err != nil {
		log.Context(ctx).Warnf("failed to get user info: %v", err)
	}
	userResp := a.generateCommentUserInfo(userInfo)

	var replyUserResp *svapi.CommentUser
	if created.ReplyUserId != 0 && &created.ReplyUserId != nil {
		userInfo, err = a.core.GetUserInfo(ctx, useroptions.GetUserInfoWithUserId(created.ReplyUserId))
		if err != nil {
			// 弱依赖
			log.Context(ctx).Warnf("failed to get user info: %v", err)
		} else {
			replyUserResp = a.generateCommentUserInfo(userInfo)
		}
	}

	return &svapi.CreateCommentResponse{
		Comment: &svapi.Comment{
			Id:         created.Id,
			VideoId:    created.VideoId,
			ParentId:   created.ParentId,
			User:       userResp,
			ReplyUser:  replyUserResp,
			Content:    created.Content,
			Date:       created.Date,
			LikeCount:  created.LikeCount,
			ReplyCount: created.ReplyCount,
		},
	}, nil
}

func (a *Application) assembleCommentListResult(ctx context.Context, data []*v1.Comment) []*svapi.Comment {
	var userIdList []int64
	for _, item := range data {
		userIdList = append(userIdList, item.UserId)
		if item.ReplyUserId != 0 && &item.ReplyUserId != nil {
			userIdList = append(userIdList, item.ReplyUserId)
		}
	}

	userInfoList, err := a.core.GetUserInfoByIdList(ctx, userIdList)
	if err != nil {
		// 弱依赖
		log.Context(ctx).Warnf("failed to get user info list: %v", err)
	}

	userInfoMap := make(map[int64]*v1.User)
	for _, item := range userInfoList {
		userInfoMap[item.Id] = item
	}

	var result []*svapi.Comment
	for _, item := range data {
		var userResp *svapi.CommentUser
		userInfo, ok := userInfoMap[item.UserId]
		if ok {
			userResp = a.generateCommentUserInfo(userInfo)
		}

		var replyUserResp *svapi.CommentUser
		if item.ReplyUserId != 0 && &item.ReplyUserId != nil {
			userInfo, ok = userInfoMap[item.ReplyUserId]
			if ok {
				replyUserResp = a.generateCommentUserInfo(userInfo)
			}
		}

		result = append(result, &svapi.Comment{
			Id:         item.Id,
			VideoId:    item.VideoId,
			ParentId:   item.ParentId,
			User:       userResp,
			ReplyUser:  replyUserResp,
			Content:    item.Content,
			Date:       item.Date,
			LikeCount:  item.LikeCount,
			ReplyCount: item.ReplyCount,
		})
	}

	return result
}

func (a *Application) ListComment4Video(ctx context.Context, request *svapi.ListComment4VideoRequest) (*svapi.ListComment4VideoResponse, error) {
	data, err := a.core.ListComment4Video(ctx, request.VideoId)
	if err != nil {
		log.Context(ctx).Errorf("failed to list comment for video: %v", err)
		return nil, errorx.New(1, "获取评论失败")
	}

	result := a.assembleCommentListResult(ctx, data)

	return &svapi.ListComment4VideoResponse{
		Comments: result,
	}, nil
}

func (a *Application) RemoveComment(ctx context.Context, request *svapi.RemoveCommentRequest) (*svapi.RemoveCommentResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	commentInfo, err := a.core.GetCommentById(ctx, request.Id)
	if err != nil {
		log.Context(ctx).Errorf("failed to get comment info: %v", err)
		return nil, errorx.New(1, "评论不存在")
	}

	if commentInfo.UserId != userId {
		return nil, errorx.New(1, "无权删除评论")
	}

	err = a.core.RemoveComment(ctx, request.Id)
	if err != nil {
		log.Context(ctx).Errorf("failed to remove comment: %v", err)
		return nil, errorx.New(1, "删除评论失败")
	}

	return &svapi.RemoveCommentResponse{}, nil
}

func (a *Application) ListChildComment(ctx context.Context, request *svapi.ListChildCommentRequest) (*svapi.ListChildCommentResponse, error) {
	data, err := a.core.ListChildComments(ctx, request.CommentId)
	if err != nil {
		log.Context(ctx).Errorf("failed to list child comment: %v", err)
		return nil, errorx.New(1, "获取回复失败")
	}

	result := a.assembleCommentListResult(ctx, data)

	return &svapi.ListChildCommentResponse{
		Comments: result,
	}, nil
}

var _ svapi.CommentServiceHTTPServer = (*Application)(nil)
