package ability369

import (
	"log"
	"errors"
	"github.com/toodo-fun/topsdk"
	"github.com/toodo-fun/topsdk/ability369/request"
	"github.com/toodo-fun/topsdk/ability369/response"
	"github.com/toodo-fun/topsdk/util"
)

type Ability369 struct {
	Client *topsdk.TopClient
}

func NewAbility369(client *topsdk.TopClient) *Ability369 {
	return &Ability369{client}
}

/*
   淘宝客-推广者-物料精选
*/
func (ability *Ability369) TaobaoTbkDgOptimusMaterial(req *request.TaobaoTbkDgOptimusMaterialRequest) (*response.TaobaoTbkDgOptimusMaterialResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability369 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.optimus.material", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgOptimusMaterialResponse{}
	if err != nil {
		log.Println("taobaoTbkDgOptimusMaterial error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-权益物料精选
*/
func (ability *Ability369) TaobaoTbkDgOptimusPromotion(req *request.TaobaoTbkDgOptimusPromotionRequest) (*response.TaobaoTbkDgOptimusPromotionResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability369 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.optimus.promotion", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgOptimusPromotionResponse{}
	if err != nil {
		log.Println("taobaoTbkDgOptimusPromotion error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if (respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
