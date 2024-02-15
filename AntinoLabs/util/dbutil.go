package util

import (
	mysql "antinolabs/mysqldb"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func ExecuteInsertQuery(ctx *gin.Context, id int64, post string) (lastId int64, err error) {
	query := "INSERT INTO antinolabs.posts(id, posts) VALUES (?, ?)"
	stmt, err := mysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, id, post)
	if err != nil {
		log.Printf("Error %s when inserting row into posts table", err)
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0, err
	}
	log.Printf("%d posts created ", rows)
	lastId, err = res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when getting last inserted id", err)
		return 0, err
	}
	log.Printf("Posts with ID %d created", lastId)
	return lastId, nil
}

// /*
// function to execute update query within a transaction
// */
func ExecuteUpdateQuery(ctx *gin.Context, id int64, post string) (sql.Result, error) {
	var res sql.Result
	query := "UPDATE antinolabs.posts SET posts = ? WHERE id = ?"
	stmt, err := mysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return res, err
	}
	defer stmt.Close()
	res, err = stmt.ExecContext(ctx, post, id)
	if err != nil {
		log.Printf("Error %s when updating row into posts table", err)
		return res, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return res, err
	}
	log.Printf("%d posts updated ", rows)

	return res, nil

}

// /*
// function to execute delete query withing a transaction
// */
func ExecuteDeleteQuery(ctx *gin.Context, id int64) (sql.Result, error) {
	var res sql.Result
	query := "DELETE FROM antinolabs.posts WHERE id = ?"
	stmt, err := mysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return res, err
	}
	defer stmt.Close()
	res, err = stmt.ExecContext(ctx, id)
	if err != nil {
		log.Printf("Error %s when deleting row into posts table", err)
		return res, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return res, err
	}
	log.Printf("%d posts deleted ", rows)

	return res, nil

}
