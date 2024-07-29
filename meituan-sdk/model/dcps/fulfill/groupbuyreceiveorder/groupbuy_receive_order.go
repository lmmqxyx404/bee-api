package groupbuyreceiveorder

/**
* 套餐配送-接单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const groupbuy_receive_order_url = "/dcps/fulfill/receive/order"

type GroupbuyReceiveOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpKey string `json:"opKey"`
}
type GroupbuyReceiveOrderRequest struct {
    BookingOrderId string `json:"bookingOrderId"`
    OpIp string `json:"opIp"`
    OpPlatform string `json:"opPlatform"`
    OpUuid string `json:"opUuid"`
}



func (req *GroupbuyReceiveOrderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupbuyReceiveOrderResponse, error) {
    resp, err := client.InvokeApi(groupbuy_receive_order_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupbuyReceiveOrderResponse
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

func (response *GroupbuyReceiveOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

