package example

import (
	"errors"
	"ginexample/model/Query"
	"strconv"
)

type Example struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ExampleFilter struct {
	ID int `json:"id"`
}

func HandleGetExampleParams(params map[string][]string) (ExampleFilter, error) {
	var exampleFilter ExampleFilter

	if len(params["id"]) > 0 {
		id, err := strconv.Atoi(params["id"][0])
		if err != nil {
			return ExampleFilter{}, err
		}
		exampleFilter.ID = id
	}

	return exampleFilter, nil
}

func GetExamples(filter ExampleFilter) ([]Example, error) {
	query := constructQuery(filter)
	results, err := Query.RunQuery(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var examples []Example
	for results.Next() {
		var example Example
		err := results.Scan(&example.ID, &example.Name)
		if err != nil {
			return nil, err
		}
		examples = append(examples, example)
	}

	return examples, nil
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

func constructQuery(filter ExampleFilter) string {
	sql := "SELECT * FROM example"
	if filter.ID != 0 {
		sql += " WHERE id = " + strconv.Itoa(filter.ID)
	}
	return sql
}
