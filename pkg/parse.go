package pkg

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"io/ioutil"
)

func Parm(c *gin.Context,key string)string{
	return c.Param(key)
}

func GetToken(c *gin.Context)string{
	return c.Request.Header.Get("Authorization")
}

func Parse(c *gin.Context,params ...interface{}) error {
	if len(params)%2 != 0 {
		return errors.New("params count must be even")
	}
	result, _:= ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%s\n", result)
	var m map[string]interface{}
	json.Unmarshal(result, &m)
	for i := 0; i < len(params); i += 2 {
		key := ToString(params[i])
		if v, ok := m[key]; ok {
			var e error
			switch ref := params[i+1].(type) {
			case *string:
				*ref = ToString(v)
			case *float64:
				*ref, e = ToFloat64(v)
			case *int:
				*ref, e = ToInt(v)
			case *int8:
				*ref, e = ToInt8(v)
			case *int16:
				*ref, e = ToInt16(v)
			case *int32:
				*ref, e = ToInt32(v)
			case *int64:
				*ref, e = ToInt64(v)
			case *uint:
				*ref, e = ToUint(v)
			case *uint8:
				*ref, e = ToUint8(v)
			case *uint16:
				*ref, e = ToUint16(v)
			case *uint32:
				*ref, e = ToUint32(v)
			case *uint64:
				*ref, e = ToUint64(v)
			case *bool:
				*ref, e = ToBool(v)
			case *[]string:
				*ref, e = ToStringSlice(v)
			case *[]int64:
				*ref, e = ToInt64Slice(v)
			case *[]uint32:
				*ref, e = ToUint32Slice(v)
			case *map[string]interface{}:
				switch m := v.(type) {
				case map[string]interface{}:
					*ref = m
				default:
					e = errors.New("value is not map[string]iterface{}")
				}
			case *[]interface{}:
				switch m := v.(type) {
				case []interface{}:
					*ref = m
				default:
					e = errors.New("value is not []iterface{}")
				}
			case *interface{}:
				*ref = v
			default:
				return errors.New(fmt.Sprintf("unknown type %v ", reflect.TypeOf(ref)))
			}
			if e != nil {
				return errors.New(fmt.Sprintf("parse [%v] error:%v", key, e.Error()))
			}
		} else {
			return errors.New(fmt.Sprintf("%v not provided", key))
		}
	}
	return nil
}
