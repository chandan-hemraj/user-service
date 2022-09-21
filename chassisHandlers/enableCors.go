/**
 * Sample Chassis Handler to print log
 *
**/

package chassisHandlers

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/openlog"
)

const HName = "Enable-Cors"

type EnableCorsHanldlerHandler struct{}

func init() { handler.RegisterHandler(HName, New1) }

func New1() handler.Handler { return &EnableCorsHanldlerHandler{} }

func (h *EnableCorsHanldlerHandler) Name() string { return HName }

func (h *EnableCorsHanldlerHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	// request object
	// /*var req *http.Request

	// if r, ok := inv.Args.(*http.Request); ok {
	// 	req = r
	// } else if r, ok := inv.Args.(*restful.Request); ok {
	// 	req = r.Request
	// } else {
	// 	openlog.Error(fmt.Sprintf("this handler only works for http protocol, wrong type: %t", inv.Args))
	// 	return
	// }
	// payload_bytes, _ := ioutil.ReadAll(req.Body)
	// fmt.Println("#######################################################", string(payload_bytes))*/
	var resp *http.ResponseWriter
	if r, ok := inv.Reply.(*http.ResponseWriter); ok {
		resp = r
	} else if r, ok := inv.Reply.(*restful.Response); ok {
		resp = &r.ResponseWriter
	} else {
		openlog.Error(fmt.Sprintf("this handler only works for http protocol, wrong type: %t", inv.Args))
		return
	}

	(*resp).Header().Set("Access-Control-Allow-Origin", "*")
	(*resp).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*resp).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, application/json , Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, multipart/form-data")
	(*resp).Header().Set("Access-Control-Allow-Credentials", "true")
	//(*resp).Header().Set("Access-Control-Allow-Headers", "Content-Type")

	chain.Next(inv, cb)
}
