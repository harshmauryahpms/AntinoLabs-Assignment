package util

import (
	"antinolabs/structure"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
validates the necessary request attributes for create request
*/

func ValidateCreateRequest(ctx *gin.Context, req structure.BlogPostCreateRequest) ([]string, bool) {
	errorlist := []string{}
	isvalid := true

	if req.Id == 0 {
		isvalid = false
		errorlist = append(errorlist, "Id can not be Blank/Zero.")
	}

	if req.Id < 0 {
		isvalid = false
		errorlist = append(errorlist, "Id can not be negative value.")
	}

	req.BlogPost = strings.TrimSpace(req.BlogPost)
	if req.BlogPost == "" {
		isvalid = false
		errorlist = append(errorlist, "Blog Field can not be blank.")
	}

	return errorlist, isvalid
}

// /*
// validates the necessary request attributes for update request
// */
func ValidateUpdateRequest(ctx *gin.Context, req structure.BlogPostUpdateRequest) ([]string, bool) {

	errorlist := []string{}
	isvalid := true

	if req.Id == 0 {
		isvalid = false
		errorlist = append(errorlist, "Id can not be Blank/Zero.")
	}

	if req.Id < 0 {
		isvalid = false
		errorlist = append(errorlist, "Id can not be negative value.")
	}

	req.BlogPost = strings.TrimSpace(req.BlogPost)
	if req.BlogPost == "" {
		isvalid = false
		errorlist = append(errorlist, "Blog Field can not be blank.")
	}
	return errorlist, isvalid
}

// //validate function for delete request

func ValidateDeleteRequest(ctx *gin.Context, req structure.BlogPostDeleteRequest) ([]string, bool) {

	errorlist := []string{}
	isvalid := true
	if req.Id == 0 {
		isvalid = false
		errorlist = append(errorlist, "Id can not be Blank/Zero.")
	}

	if req.Id < 0 {
		isvalid = false
		errorlist = append(errorlist, "Id can not be negative value.")
	}
	return errorlist, isvalid
}
