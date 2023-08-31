package dao

import (
	"fmt"
	"strings"

	aligraph "github.com/aliyun/aliyun-igraph-go-sdk"
	"github.com/aliyun/aliyun-pai-featurestore-go-sdk/v2/constants"
	"github.com/aliyun/aliyun-pai-featurestore-go-sdk/v2/datasource/igraph"
	"github.com/aliyun/aliyun-pai-featurestore-go-sdk/v2/utils"
)

type FeatureViewIGraphDao struct {
	igraphClient    *aligraph.Client
	group           string
	label           string
	primaryKeyField string
	eventTimeField  string
	ttl             int
	fieldMap        map[string]string
	fieldTypeMap    map[string]constants.FSType
	reverseFieldMap map[string]string
}

func NewFeatureViewIGraphDao(config DaoConfig) *FeatureViewIGraphDao {
	dao := FeatureViewIGraphDao{
		group:           config.GroupName,
		label:           config.LabelName,
		primaryKeyField: config.PrimaryKeyField,
		eventTimeField:  config.EventTimeField,
		ttl:             config.TTL,
		fieldMap:        config.FieldMap, // igraph name => feature view schema name mapping
		fieldTypeMap:    config.FieldTypeMap,
		reverseFieldMap: make(map[string]string, len(config.FieldMap)), // revserse fieldMap kv, feature view schema name => igraph name mapping
	}
	client, err := igraph.GetGraphClient(config.IGraphName)
	if err != nil {
		return nil
	}

	dao.igraphClient = client.GraphClient
	for k, v := range dao.fieldMap {
		dao.reverseFieldMap[v] = k
	}
	return &dao
}
func (d *FeatureViewIGraphDao) GetFeatures(keys []interface{}, selectFields []string) ([]map[string]interface{}, error) {
	var pkeys []string
	for _, key := range keys {
		if pkey := utils.ToString(key, ""); pkey != "" {
			pkeys = append(pkeys, pkey)
		}
	}
	selector := make([]string, 0, len(selectFields))
	for _, field := range selectFields {
		selector = append(selector, fmt.Sprintf("\"%s\"", d.reverseFieldMap[field]))
	}

	var queryString string
	if len(d.fieldMap) == len(selectFields) {
		queryString = fmt.Sprintf("g(\"%s\").V(\"%s\").hasLabel(\"%s\")", d.group, strings.Join(pkeys, ";"), d.label)
	} else {
		queryString = fmt.Sprintf("g(\"%s\").V(\"%s\").hasLabel(\"%s\").fields(%s)", d.group, strings.Join(pkeys, ";"), d.label, strings.Join(selector, ","))
	}

	request := aligraph.ReadRequest{
		QueryString: queryString,
	}
	resp, err := d.igraphClient.Read(&request)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0, len(keys))
	for _, resultData := range resp.Result {
		for _, data := range resultData.Data {
			properties := make(map[string]interface{}, len(data))

			for field, value := range data {
				if field == "label" {
					continue
				}

				switch d.fieldTypeMap[field] {
				case constants.FS_DOUBLE, constants.FS_FLOAT:
					properties[d.fieldMap[field]] = utils.ToFloat(value, 0)
				case constants.FS_INT32, constants.FS_INT64:
					properties[d.fieldMap[field]] = utils.ToInt(value, -1024)
				default:
					properties[d.fieldMap[field]] = value
				}
			}

			result = append(result, properties)
		}
	}

	return result, nil
}
