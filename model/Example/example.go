package example

import (
	"errors"
	"ginexample/model/Pagination"
	"ginexample/model/Query"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Example struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ExampleFilter struct {
	ID    int
	Limit int
	Page  int
	Order string
	By    string
}

func HandleGetExampleParams(ctx *gin.Context) (ExampleFilter, error) {
	var exampleFilter ExampleFilter

	id, _ := ctx.GetQuery("id")
	limit, _ := ctx.GetQuery("limit")
	page, _ := ctx.GetQuery("page")

	if len(id) > 0 {
		id, err := strconv.Atoi(id)
		if err != nil {
			return ExampleFilter{}, err
		}
		exampleFilter.ID = id
	}

	if len(limit) > 0 {
		limit, err := strconv.Atoi(limit)
		if err != nil {
			return ExampleFilter{}, err
		}
		exampleFilter.Limit = limit
	} else {
		exampleFilter.Limit = 10
	}

	if len(page) > 0 {
		page, err := strconv.Atoi(page)
		if err != nil {
			return ExampleFilter{}, err
		}
		exampleFilter.Page = page
	} else {
		exampleFilter.Page = 1
	}

	return exampleFilter, nil
}

func GetExamples(filter ExampleFilter) (Pagination.PaginationResponse, error) {
	query, queryCount := constructQuery(filter)
	query = "SELECT * FROM example" + query
	results, err := Query.RunQuery(query)
	if err != nil {
		return Pagination.PaginationResponse{}, err
	}

	countQuery := "SELECT COUNT(*) FROM example" + queryCount
	totalCount, err := Query.RunQueryCount(countQuery)
	if err != nil {
		return Pagination.PaginationResponse{}, err
	}

	defer results.Close()

	var examples []Example
	for results.Next() {
		var example Example
		err := results.Scan(&example.ID, &example.Name)
		if err != nil {
			return Pagination.PaginationResponse{}, err
		}
		examples = append(examples, example)
	}

	var res Pagination.PaginationResponse
	res.Total = totalCount
	res.TotalPageCount = int(math.Ceil(float64(totalCount) / float64(filter.Limit)))
	res.NextPage = Pagination.GetNextPage("/api/v1/examples", filter.Page, filter.Limit)
	res.Page = filter.Page
	res.Limit = filter.Limit
	res.Data = examples

	return res, nil
}

func GetExample(id string) (Example, error) {
	query := "SELECT * FROM example WHERE id = " + id
	result, err := Query.RunQuery(query)
	if err != nil {
		return Example{}, err
	}
	defer result.Close()

	var example Example
	for result.Next() {
		err := result.Scan(&example.ID, &example.Name)
		if err != nil {
			return Example{}, err
		}
	}

	if len(example.Name) == 0 {
		return Example{}, errors.New("example id: " + id + " not found")
	}

	return example, nil
}

func constructQuery(filter ExampleFilter) (string, string) {
	var where string
	var order string
	var limit string

	if filter.ID != 0 {
		where = " WHERE id = " + strconv.Itoa(filter.ID)
	}

	if filter.Limit != 0 {
		if filter.Page == 0 {
			filter.Page = 1
		}

		limit = " LIMIT " + strconv.Itoa(filter.Limit) + " OFFSET " + strconv.Itoa((filter.Page-1)*filter.Limit)
	}

	if len(filter.Order) > 0 && len(filter.By) > 0 {
		order = " ORDER BY " + filter.Order + " " + filter.By
	}

	return where + order + limit, where
}
