package handler

import (
	mysql "antinolabs/mysqldb"
	"antinolabs/structure"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var start time.Time

func GetblogpostsInfoHandler(ctx *gin.Context) {
	start = time.Now()
	Id, _ := ctx.GetQuery("Id")
	blogs := []structure.BlogPostInfoResponse{}

	var scanError error

	var response structure.BlogPostInfoResponse

	Id_int, err := strconv.ParseInt(Id, 0, 64)
	if err != nil {
		log.Println("Id should be an integer. Error-", err)
		constructResponse(400, true,
			ctx,
			nil, Id_int, "Id should be an integer value.")
		return
	}

	if Id_int <= 0 {
		log.Println("Id should be greater than zero")
		constructResponse(400, true,
			ctx,
			nil, Id_int, "Id should be greater than zero.")
		return
	}

	query := `SELECT * FROM antinolabs.posts 
	where id = ? `

	rows, err := mysql.DB.Query(query, Id_int)
	if err != nil {
		log.Println("Error while fetching posts", err)

	}
	defer rows.Close()
	for rows.Next() {
		if scanError = rows.Scan(&response.Id,
			&response.BlogPost,
		); scanError != nil {
			log.Println("Error while scanning for blogpost columns", err)
			constructResponse(http.StatusInternalServerError, true, ctx, nil, Id_int, "Error while scanning for blogpost columns.")
		}
		blogs = append(blogs, response)
	}
	if err1 := rows.Err(); err1 != nil {
		log.Println("Error while fetching blogs", err1)
		constructResponse(http.StatusInternalServerError, true, ctx, nil, Id_int, "Error while fetching blogs.")
	}

	if err == sql.ErrNoRows {
		log.Println("Error -No rows exists", err)
		constructResponse(http.StatusInternalServerError, true, ctx, nil, Id_int, "No rows exists for input id.")
		return
	}
	//Return response based on the data retrieved from DB else throw error
	if scanError == nil {
		constructResponse(200, false,
			ctx,
			&response, Id_int, "")
		return
	}

}

/*
common logic to construct controlvaluesets response
*/
func constructResponse(statusCode int, haserror bool,
	c *gin.Context, blogposttresp *structure.BlogPostInfoResponse,
	Id_int int64, errResp string) {

	response := structure.BlogPostGetResponse{
		StatusCode:       statusCode,
		TimeElapsed:      time.Since(start).Milliseconds(),
		Id:               Id_int,
		BlogPost:         blogposttresp,
		HasErrorResponse: haserror,
		ErrorResponse:    errResp,
	}
	c.JSON(statusCode, response)
}
