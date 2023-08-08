package helper

import (
	"github.com/MeibisuX673/lessonGin/app/model"
	"gorm.io/gorm"
)

func ConfigurationDbQuery(db *gorm.DB, query model.Query) {

	if query.Filters != nil {
		for _, value := range query.Filters {
			db.Where(value)
		}
	}
	if query.RangeFilters != nil {
		for _, value := range query.RangeFilters {
			db.Where(value)
		}
	}

	if query.Orders != nil {
		for _, value := range query.Orders {
			db.Order(value)
		}
	}

	db.Offset(int(query.Page*query.Limit - query.Limit))

	db.Limit(int(query.Limit))

}
