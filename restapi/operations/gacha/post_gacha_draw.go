// Code generated by go-swagger; DO NOT EDIT.

package gacha

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostGachaDrawHandlerFunc turns a function with the right signature into a post gacha draw handler
type PostGachaDrawHandlerFunc func(PostGachaDrawParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGachaDrawHandlerFunc) Handle(params PostGachaDrawParams) middleware.Responder {
	return fn(params)
}

// PostGachaDrawHandler interface for that can handle valid post gacha draw params
type PostGachaDrawHandler interface {
	Handle(PostGachaDrawParams) middleware.Responder
}

// NewPostGachaDraw creates a new http.Handler for the post gacha draw operation
func NewPostGachaDraw(ctx *middleware.Context, handler PostGachaDrawHandler) *PostGachaDraw {
	return &PostGachaDraw{Context: ctx, Handler: handler}
}

/* PostGachaDraw swagger:route POST /gacha/draw gacha postGachaDraw

ガチャ実行API

ガチャを引いてキャラクターを取得する処理を実装します。
 獲得したキャラクターはユーザ所持キャラクターテーブルへ保存します。
 同じ種類のキャラクターでもユーザは複数所持することができます。

 キャラクターの確率は等倍ではなく、任意に変更できるようテーブルを設計しましょう。

*/
type PostGachaDraw struct {
	Context *middleware.Context
	Handler PostGachaDrawHandler
}

func (o *PostGachaDraw) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostGachaDrawParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
