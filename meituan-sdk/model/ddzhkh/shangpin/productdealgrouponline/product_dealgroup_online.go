package productdealgrouponline

/**
* 上架团单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const product_dealgroup_online_url = "/ddzhkh/shangpin/dealgroup/online"

type Response struct {
    /**
    * 团单id
    */
    DealGroupId int32 `json:"dealGroupId"`
}
type ProductDealgroupOnlineRequest struct {
    /**
    *  DealGroupDataDTO的json序列化后字符串形式 
    */
    Data string `json:"data"`
    /**
    *  来源，体检 : physical_examination 
    */
    Source string `json:"source"`
    OpPoiIds []string `json:"opPoiIds"`
    /**
    *  团单ID 
    */
    DealGroupId int64 `json:"dealGroupId"`
}
type ProductDealgroupOnlineResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Response `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *ProductDealgroupOnlineRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ProductDealgroupOnlineResponse, error) {
    resp, err := client.InvokeApi(product_dealgroup_online_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ProductDealgroupOnlineResponse
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

func (response *ProductDealgroupOnlineResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

