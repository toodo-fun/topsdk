package response

import (
	"github.com/toodo-fun/topsdk/ability372/domain"
)

type TaobaoTbkShopRecommendGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   淘宝客店铺
	*/
	Results []domain.TaobaoTbkShopRecommendGetNTbkShop `json:"results,omitempty" `
}
