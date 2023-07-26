package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"
)

//轮播图
func (s *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	bannerListResponse := proto.BannerListResponse{}

	var banners []model.Banner
	result := global.DB.Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerResponses []*proto.BannerResponse
	for _, banner := range banners {
		bannerResponses = append(bannerResponses, &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerResponses

	return &bannerListResponse, nil
}

func (s *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {

	banner := model.Banner{}

	banner.Image = req.Image
	banner.Index = req.Index
	banner.Url = req.Url

	global.DB.Save(&banner)

	return &proto.BannerResponse{
		Id: banner.ID,
	}, nil

}

func (s *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")

	}
	return &emptypb.Empty{}, nil

}

func (s *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {

	var banner model.Banner
	if result := global.DB.First(&model.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")

	}

	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	global.DB.Save(&banner)
	return &emptypb.Empty{}, nil

}
