package main

import (
	"fmt"
	"reflect"
)

// We could have used generics to pass a converter func that takes []any and returns []any
// The converter should have been implemented for each type though
// Below Works for now!!
func convertDBStructSliceToResponseStructSlice(dbSlice interface{}, responseStruct interface{}) interface{} {
	dbSliceVal := reflect.ValueOf(dbSlice)
	responseSliceType := reflect.TypeOf(responseStruct)
	fmt.Println("responseSliceType", responseSliceType)
	fmt.Println("dbSliceVal", dbSliceVal)
	responseFeed := reflect.MakeSlice(reflect.SliceOf(responseSliceType), dbSliceVal.Len(), dbSliceVal.Len())

	for i := 0; i < dbSliceVal.Len(); i++ {
		newResponseStruct := reflect.New(responseSliceType).Elem()
		dbItem := dbSliceVal.Index(i)

		for j := 0; j < responseSliceType.NumField(); j++ {
			responseField := newResponseStruct.Field(j)
			dbField := dbItem.FieldByName(responseSliceType.Field(j).Name)
			if dbField.IsValid() && responseField.CanSet() {
				responseField.Set(dbField)
			}
		}

		responseFeed.Index(i).Set(newResponseStruct)
	}

	return responseFeed.Interface()
}
