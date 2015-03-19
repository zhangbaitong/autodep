package common

type ACTION_V1 struct {
	Action string
                Result string
}

type RequestData struct
{
        Version         string
        ServerIP        string
        Port             int
        Method         string
        Params         interface{} 
}