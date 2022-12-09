package response

import (
	"github.com/toodo-fun/topsdk/ability2138/domain"
)

type TaobaoTbkDgNewuserOrderGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   data
	*/
	Results domain.TaobaoTbkDgNewuserOrderGetResults `json:"results,omitempty" `
}
