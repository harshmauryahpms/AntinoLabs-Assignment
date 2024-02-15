package handler

import (
	"antinolabs/structure"
	"antinolabs/util"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var start2 time.Time

/*
handler logic to insert a record with given data
*/
func CreateblogpostsRecord(ctx *gin.Context) {
	start2 = time.Now()
	var BlogPostCreateStructRequest structure.BlogPostCreateRequest

	err := ctx.BindJSON(&BlogPostCreateStructRequest)

	if err != nil {
		log.Println("Error - invalid data types ")
		constructResponseForMaintain(400, true, ctx,
			0, "don't insert", "Error - invalid data types in request. Id should be int and BlogPost should be string.")
		return
	} else {

		errorlist, isvalid := util.ValidateCreateRequest(ctx, BlogPostCreateStructRequest)

		if isvalid {

			id, err := util.ExecuteInsertQuery(ctx, BlogPostCreateStructRequest.Id, BlogPostCreateStructRequest.BlogPost)
			switch err {
			case nil:
				constructResponseForMaintain(200, false,
					ctx, id, "insert", "")
				return

			default:
				log.Println(" -> Error while inserting the data ")
				constructResponseForMaintain(500, true, ctx,
					0, "don't insert", "Error while inserting the data.")
				return
			}

		} else {
			log.Println("Error - missing data: ", errorlist)

			for _, errorval := range errorlist {
				log.Println("Error - missing data: ", errorval)
			}
			constructResponseForMaintain(400, true, ctx,
				0, "don't insert", "Error - Id is missing/zero/negative value and/or BlogPost is missing/blank.")
			return
		}
	}
}

/*
handler logic to update a record with the given data
*/
func UpdateblogpostsRecord(ctx *gin.Context) {
	start2 = time.Now()
	var BlogPostUpdateStructRequest structure.BlogPostUpdateRequest

	if err := ctx.BindJSON(&BlogPostUpdateStructRequest); err != nil {
		log.Println("Error - invalid data types: ")
		constructResponseForMaintain(400, true, ctx,
			0, "don't update", "Error - invalid data types in request. Id should be int and BlogPost should be string.")
		return
	} else {

		errorlist, isvalid := util.ValidateUpdateRequest(ctx, BlogPostUpdateStructRequest)

		if isvalid {

			rowsAffected, err := util.ExecuteUpdateQuery(ctx, BlogPostUpdateStructRequest.Id, BlogPostUpdateStructRequest.BlogPost)
			switch err {
			case nil:
				if affect, _ := rowsAffected.RowsAffected(); affect > 0 {
					constructResponseForMaintain(200, false,
						ctx, BlogPostUpdateStructRequest.Id, "update", "")
					return
				} else {
					log.Println("Error - No rows affected in update. ")
					constructResponseForMaintain(500, true, ctx,
						0, "don't update", "No row exists for the input Id, So UPDATE can not be performed.")
					return
				}
			default:
				log.Println(" -> Error while updating the data ")
				constructResponseForMaintain(500, true, ctx,
					0, "don't update", "Error while updating the data")
				return
			}

		} else {
			log.Println("Error - missing data: ", errorlist)

			for _, errorval := range errorlist {
				log.Println("Error - missing data: ", errorval)
			}
			constructResponseForMaintain(400, true, ctx,
				0, "don't update", "Error - Id is missing/zero/negative value and/or BlogPost is missing/blank.")
			return
		}
	}
}

// /*
//
//	handler logic to delete a record for the given data
//
// */
func DeleteblogpostsRecord(ctx *gin.Context) {
	start2 = time.Now()
	var BlogPostDeleteStructRequest structure.BlogPostDeleteRequest

	if err := ctx.BindJSON(&BlogPostDeleteStructRequest); err != nil {
		log.Println("Error - invalid data types: ")
		constructResponseForMaintain(400, true, ctx,
			0, "don't delete", "Error - invalid data types in request. Id should be int.")
		return
	} else {

		errorlist, isvalid := util.ValidateDeleteRequest(ctx, BlogPostDeleteStructRequest)

		if isvalid {

			rowsAffected, err := util.ExecuteDeleteQuery(ctx, BlogPostDeleteStructRequest.Id)
			switch err {
			case nil:
				if affect, _ := rowsAffected.RowsAffected(); affect > 0 {
					constructResponseForMaintain(200, false,
						ctx, BlogPostDeleteStructRequest.Id, "delete", "")
					return
				} else {
					log.Println("Error - No rows affected in delete. ")
					constructResponseForMaintain(500, true, ctx,
						0, "don't update", "No row exists for the input Id, So DELETE can not be performed.")
					return
				}
			default:
				log.Println(" -> Error while deleting the data ")
				constructResponseForMaintain(500, true, ctx,
					0, "don't delete", "Error while deleting the data.")
				return
			}

		} else {
			log.Println("Error - missing data: ", errorlist)

			for _, errorval := range errorlist {
				log.Println("Error - missing data: ", errorval)
			}
			constructResponseForMaintain(400, true, ctx,
				0, "don't delete", "Error - Id is missing/zero/negative value.")
			return
		}
	}
}

/*
common logic to construct response
*/
func constructResponseForMaintain(statusCode int, haserror bool,
	c *gin.Context, idr int64, operation string, errResp string) {
	var msgStr string
	if !haserror && operation == "insert" {
		msgStr = "Record created in the database."
	}
	if !haserror && operation == "update" {
		msgStr = "Record updated with given data in the database."
	}
	if !haserror && operation == "delete" {
		msgStr = "Record deleted from the database."
	}

	response := structure.BlogPostSetMaintainResponse{
		StatusCode:       statusCode,
		TimeElapsed:      time.Since(start2).Milliseconds(),
		Id:               idr,
		HasErrorResponse: haserror,
		ErrorResponse:    errResp,
		MessageString:    msgStr,
	}
	c.JSON(statusCode, response)
}
