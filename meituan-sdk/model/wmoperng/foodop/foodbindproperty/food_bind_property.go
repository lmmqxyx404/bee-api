package foodbindproperty

/**
* 绑定菜品属性
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_bind_property_url = "/wmoper/ng/foodop/food/bind/property"

type FoodBindPropertyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodBindPropertyRequest struct {
    /**
    *  菜品属性的json数据 
    */
    FoodProperty string `json:"food_property"`
}



func (req *FoodBindPropertyRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodBindPropertyResponse, error) {
    resp, err := client.InvokeApi(food_bind_property_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodBindPropertyResponse
    err = json.Unmarshal(resp, &response)
    if err != nil {
        return nil, err
    }

    if response.IsSuccess() {
        return &response, nil
    } else {
        return &response, &api_error.ApiError{Code: response.Code, Msg: response.Msg, TraceId: response.TraceId}
    }
}

func (response *FoodBindPropertyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

