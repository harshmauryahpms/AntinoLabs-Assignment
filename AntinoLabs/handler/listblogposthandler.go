package handler

import (
	mysql "antinolabs/mysqldb"
	"antinolabs/structure"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func GetListblogposts(ctx *gin.Context) {
	start = time.Now()
	var blogs []structure.BlogPostInfoResponse
	var scanListError error
	var response structure.BlogPostInfoResponse

	totalRecords, scanError := TotalRecordCount(ctx)
	if scanError != nil {
		log.Println("Error while counting blogpost rows", scanError)
		constructListResponse(http.StatusInternalServerError, ctx, startTime, totalRecords, true, blogs, "Error while counting blogpost rows.")
	}

	query := `SELECT * FROM antinolabs.posts`

	rows, err := mysql.DB.Query(query)
	if err != nil {
		log.Println("Error while fetching lists of posts", err)
	}
	defer rows.Close()

	for rows.Next() {

		if scanListError = rows.Scan(&response.Id,
			&response.BlogPost,
		); scanListError != nil {
			log.Println("Error while scanning for blogpost rows", err)
			constructListResponse(http.StatusInternalServerError, ctx, startTime, totalRecords, true, blogs, "Error while scanning for blogpost rows.")

			return
		}
		blogs = append(blogs, response)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error while fetching blogs", err)
		constructListResponse(http.StatusInternalServerError, ctx, startTime, totalRecords, true, blogs, "Error while fetching blogs.")
	}
	if err == sql.ErrNoRows {
		log.Println("Error -No rows exists", err)
		constructListResponse(http.StatusInternalServerError, ctx, startTime, totalRecords, true, blogs, "No rows exists in the table.")
	}

	//Return response based on the data retrieved from DB else throw error
	if scanListError == nil {
		constructListResponse(200, ctx, startTime, totalRecords, false, blogs, "")
		return
	}

}

func TotalRecordCount(ctx *gin.Context) (int, error) {

	numberofrecord := 0
	var scanError error

	countquery := `SELECT count(*)
	from antinolabs.posts `

	rows, err := mysql.DB.Query(countquery)

	if err != nil {
		log.Println("Error while fetching count of posts", err)

	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.

	for rows.Next() {
		scanError = rows.Scan(&numberofrecord)
	}
	log.Println("number of records are-->", numberofrecord)
	return numberofrecord, scanError
}

/*
Construct response for list service
*/
func constructListResponse(statusCode int, ctx *gin.Context, startTime time.Time, totalreccount int,
	haserror bool, standComminforesponse []structure.BlogPostInfoResponse, errResp string) {
	response := structure.BlogPostListResponse{
		StatusCode:       statusCode,
		TimeElapsed:      time.Since(startTime).Milliseconds(),
		TotalRecords:     totalreccount,
		HasErrorResponse: haserror,
		ErrorResponse:    errResp,
		BlogPostInfoSet:  standComminforesponse,
	}
	ctx.JSON(statusCode, response)
}
