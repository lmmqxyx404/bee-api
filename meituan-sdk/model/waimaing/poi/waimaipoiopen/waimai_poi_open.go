package waimaipoiopen

/**
* 门店置营业
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const waimai_poi_open_url = "/waimai/poi/open"

type WaimaiPoiOpenResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type WaimaiPoiOpenRequest struct {
}



func (req *WaimaiPoiOpenRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WaimaiPoiOpenResponse, error) {
    resp, err := client.InvokeApi(waimai_poi_open_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WaimaiPoiOpenResponse
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

func (response *WaimaiPoiOpenResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

