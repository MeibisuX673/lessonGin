package queryService

import (
	"fmt"
	"github.com/MeibisuX673/lessonGin/app/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type QueryService struct {
}

func New() *QueryService {
	return &QueryService{}
}

func (qs *QueryService) GetQueries(c *gin.Context) *model.Query {

	var query model.Query = model.Query{}

	queries := c.Request.URL.Query()

	replacer := strings.NewReplacer("[", " ", "]", " ")

	fmt.Println(queries)

	query.Page = 1
	value, ok := queries["page"]
	if ok {
		page, err := strconv.Atoi(value[0])
		if err == nil {
			query.Page = uint(page)
		}
	}

	query.Limit = 5
	value, ok = queries["limit"]
	if ok {
		limit, err := strconv.Atoi(value[0])
		if err == nil {
			query.Limit = uint(limit)
		}
	}

	for key, value := range queries {

		filter := ""
		order := ""
		rangeFilter := ""

		queryData := strings.Fields(replacer.Replace(key))
		fmt.Println(queryData)
		fmt.Println(value)

		switch queryData[0] {
		case "filter":
			switch queryData[2] {
			case "exact":
				filter = fmt.Sprintf("%s = %s", queryData[1], value[0])
			case "partial":
				filter = fmt.Sprintf("%s LIKE \"%%%s%%\"", queryData[1], value[0])
			}
		case "order":
			strategy := strings.ToLower(value[0])
			order = fmt.Sprintf("%s %s", queryData[1], strategy)
		case "range":
			switch queryData[2] {
			case "gt":
				rangeFilter = fmt.Sprintf("%s > %s", queryData[1], value[0])
			case "lt":
				rangeFilter = fmt.Sprintf("%s < %s", queryData[1], value[0])
			}

		}

		if len(filter) != 0 {
			query.Filters = append(query.Filters, filter)
		}
		if len(order) != 0 {
			query.Orders = append(query.Orders, order)
		}
		if len(rangeFilter) != 0 {
			query.RangeFilters = append(query.RangeFilters, rangeFilter)
		}

	}

	return &query
}
