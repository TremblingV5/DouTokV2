package videodomain

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/userdata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/videodata"
	service_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/middleware/jwt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type VideoUseCase struct {
	config    *conf.Config
	videoRepo videodata.IVideoRepo
	userRepo  userdata.IUserRepo
	log       *log.Helper
}

func NewVideoUseCase(
	config *conf.Config,
	userRepo userdata.IUserRepo,
	videoRepo videodata.IVideoRepo,
	logger log.Logger,
) *VideoUseCase {
	return &VideoUseCase{
		config:    config,
		videoRepo: videoRepo,
		userRepo:  userRepo,
		log:       log.NewHelper(logger),
	}
}

func (uc *VideoUseCase) FeedShortVideo(ctx context.Context, request *service_dto.FeedShortVideoRequest) (*service_dto.FeedShortVideoResponse, error) {
	latestTime := time.Now().UTC().Unix()
	if request.LatestTime > 0 {
		latestTime = request.LatestTime
	}

	videos, err := uc.videoRepo.GetVideoFeed(ctx, query.Q, request.UserId, latestTime, request.FeedNum)
	if err != nil {
		return nil, err
	}
	user, err := uc.userRepo.FindByID(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	videos := entity.FromVideoModelList(resp.Videos)
	author := entity.ToAuthorEntity(user)
	for _, video := range videos {
		video.Author = author
	}

	return &service_dto.FeedShortVideoResponse{
		Videos: videoList,
	}, nil
}

func (uc *VideoUseCase) PublishVideo(ctx context.Context, in *service_dto.PublishVideoRequest) (int64, error) {
	userId, err := jwt.GetLoginUser(ctx)
	if err != nil {
		return 0, fmt.Errorf("get login user failed: %v", err)
	}
	video := model.Video{
		ID:          utils.GetSnowflakeId(),
		UserID:      userId,
		Title:       in.Title,
		Description: in.Description,
		VideoURL:    in.VideoURL,
		CoverURL:    in.CoverURL,
	}
	err = uc.videoRepo.Save(ctx, query.Q, &video)
	if err != nil {
		return 0, err
	}
	return video.ID, nil
}

func (uc *VideoUseCase) GetVideoById(ctx context.Context, videoId int64) (*entity.Video, error) {
	video, err := uc.videoRepo.FindByID(ctx, query.Q, videoId)
	if err != nil {
		return nil, err
	}
	authorModel, err := uc.userRepo.FindByID(ctx, query.Q, video.UserID)
	if err != nil {
		return nil, err
	}
	videoEntity := entity.FromVideoModel(video)
	videoEntity.Author = entity.ToAuthorEntity(authorModel)
	return videoEntity, nil
}

func (uc *VideoUseCase) ListPublishedVideo(ctx context.Context, request *service_dto.ListPublishedVideoRequest) (*service_dto.ListPublishedVideoResponse, error) {
	latestTime := time.Now().UTC().Unix()
	if request.LatestTime > 0 {
		latestTime = request.LatestTime
	}

	user, err := uc.userRepo.FindByID(ctx, query.Q, request.UserId)
	if err != nil {
		return nil, err
	}

	videos, pageResp, err := uc.videoRepo.GetVideoList(ctx, query.Q, request.UserId, latestTime, request.Pagination)
	if err != nil {
		return nil, err
	}

	videoEntityList := entity.FromVideoModelList(videos)
	author := entity.ToAuthorEntity(user)
	for _, video := range videoEntityList {
		video.Author = author
	}

	return &service_dto.ListPublishedVideoResponse{
		Videos:     videoEntityList,
		Pagination: pageResp,
	}, nil
}
