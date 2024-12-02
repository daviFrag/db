package db

import "log"

func (dbPlugin *DbPlugin) Select(query string, inputParams []interface{}) interface{} {
	var result interface{}
	dbPlugin.db.Select(result, query, inputParams)

	log.Println(result)

	return result
}
