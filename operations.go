package db

import "log"

func (dbPlugin *DbPlugin) Select(query string, inputParams []interface{}) interface{} {
	var result interface{}
	err := dbPlugin.db.Select(result, query, inputParams)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)

	return result
}
