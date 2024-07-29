package dailycpmdata

/**
* cpm离线分天数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const daily_cpm_data_url = "/ad/report/getDailyCpmData"

type DailyCpmDataResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 对象列表
    */
    Data []LaunchInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type DailyCpmDataRequest struct {
    /**
    * 对应查询元素的id列表,单次最多支持50个
    */
    Ids []int64 `json:"ids"`
    /**
    * 0账户，1计划，2品牌，3推广 4 创意
    */
    Dimension int32 `json:"dimension"`
    /**
    * 子账户
    */
    AccountId int32 `json:"accountId"`
    /**
    * 查询开始时间
    */
    BeginTime string `json:"beginTime"`
    /**
    * 查询截止时间
    */
    EndTime string `json:"endTime"`
}
type LaunchInfo struct {
    /**
    * 日期
    */
    Date string `json:"date"`
    /**
    * 查询元素id
    */
    Id int32 `json:"id"`
    /**
    * 查询元素名称
    */
    Name string `json:"name"`
    /**
    * 花费
    */
    Charge float64 `json:"charge"`
    /**
    * 曝光
    */
    Imp int32 `json:"imp"`
    /**
    * 点击
    */
    Click int32 `json:"click"`
}



func (req *DailyCpmDataRequest) DoInvoke(client mtclient.MeituanClient) (*DailyCpmDataResponse, error) {
    resp, err := client.InvokeApi(daily_cpm_data_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response DailyCpmDataResponse
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

func (response *DailyCpmDataResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

