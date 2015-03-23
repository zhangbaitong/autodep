package common

type Response struct {
    Method string
    Code int
    Messgae string
    Data string
}

type RequestData struct{
        Version         string
        ServerIP        string
        Port             int
        Method         string
        Params         map[string]interface{}
}

