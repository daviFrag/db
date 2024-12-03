package db

import "log"

func (dbPlugin *DbPlugin) Select(query string, inputParams map[string]interface{}) interface{} {
	rows, err := dbPlugin.db.NamedQuery(query, inputParams)
	if err != nil {
		log.Fatalln(err)
	}
	counter := 0
	for rows.Next() {
		log.Println(counter)
		counter += 1
	}

	log.Println(rows)

	return rows
}
