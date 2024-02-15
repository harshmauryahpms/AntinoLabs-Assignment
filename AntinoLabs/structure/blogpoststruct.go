package structure

type BlogPostInfoResponse struct {
	Id       int64  `json:"id,omitempty"`
	BlogPost string `json:"blogpost,omitempty"`
}

type BlogPostGetResponse struct {
	StatusCode       int                   `json:"statusCode,omitempty"`
	TimeElapsed      int64                 `json:"timeElapsed,omitempty"`
	Id               int64                 `json:"id,omitempty"`
	BlogPost         *BlogPostInfoResponse `json:"BlogPostInfo,omitempty"`
	HasErrorResponse bool                  `json:"hasErrorResponse"`
	ErrorResponse    string                `json:"errorResponse,omitempty"`
}

//-------ListResponse---------//

type BlogPostListResponse struct {
	StatusCode       int                    `json:"statusCode,omitempty"`
	TimeElapsed      int64                  `json:"timeElapsed,omitempty"`
	TotalRecords     int                    `json:"totalRecords,omitempty"`
	HasErrorResponse bool                   `json:"hasErrorResponse"`
	ErrorResponse    string                 `json:"errorResponse,omitempty"`
	BlogPostInfoSet  []BlogPostInfoResponse `json:"BlogPost_ListResponse,omitempty"`
}

// //------Maintain Structs-------------------

type BlogPostCreateRequest struct {
	Id       int64  `json:"id,omitempty"`
	BlogPost string `json:"blogpost,omitempty"`
}

type BlogPostUpdateRequest struct {
	Id       int64  `json:"id,omitempty"`
	BlogPost string `json:"blogpost,omitempty"`
}

type BlogPostDeleteRequest struct {
	Id int64 `json:"id,omitempty"`
}

type BlogPostSetMaintainResponse struct {
	StatusCode       int    `json:"statusCode,omitempty"`
	TimeElapsed      int64  `json:"timeElapsed,omitempty"`
	Id               int64  `json:"id,omitempty"`
	HasErrorResponse bool   `json:"hasErrorResponse"`
	ErrorResponse    string `json:"errorResponse,omitempty"`
	MessageString    string `json:"SuccessMessage,omitempty"`
}
