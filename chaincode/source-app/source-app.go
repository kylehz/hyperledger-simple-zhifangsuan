package main

import (
	"encoding/json"
	"fmt"
	//  "strconv"
	//  "strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ObjChainCode struct {
}

//obj数据结构体
type ObjInfo struct {
	ObjID      string    `json:ObjID`      //物件ID
	ObjProInfo ProInfo   `json:ObjProInfo` //生产信息
	ObjIngInfo []IngInfo `json:ObjIngInfo` //配料信息
	ObjLogInfo LogInfo   `json:ObjLogInfo` //物流信息
}

type ObjAllInfo struct {
	ObjID      string    `json:ObjId`
	ObjProInfo ProInfo   `json:ObjProInfo`
	ObjIngInfo []IngInfo `json:ObjIngInfo`
	ObjLogInfo []LogInfo `json:ObjLogInfo`
}

//生产信息
type ProInfo struct {
	ObjName     string `json:ObjName`     //物件名称
	ObjSpec     string `json:ObjSpec`     //物件规格
	ObjMFGDate  string `json:ObjMFGDate`  //物件出产日期
	ObjEXPDate  string `json:ObjEXPDate`  //物件保质期
	ObjLOT      string `json:ObjLOT`      //物件批次号
	ObjQSID     string `json:ObjQSID`     //物件生产许可证编号
	ObjMFRSName string `json:ObjMFRSName` //物件生产商名称
	ObjProPrice string `json:ObjProPrice` //物件生产价格
	ObjProPlace string `json:ObjProPlace` //物件生产所在地
}
type IngInfo struct {
	IngID   string `json:IngID`   //配料ID
	IngName string `json:IngName` //配料名称
}

type LogInfo struct {
	LogDepartureTm string `json:LogDepartureTm` //出发时间
	LogArrivalTm   string `json:LogArrivalTm`   //到达时间
	LogMission     string `json:LogMission`     //处理业务(储存or运输)
	LogDeparturePl string `json:LogDeparturePl` //出发地
	LogDest        string `json:LogDest`        //目的地
	LogToSeller    string `json:LogToSeller`    //销售商
	LogStorageTm   string `json:LogStorageTm`   //存储时间
	LogMOT         string `json:LogMOT`         //运送方式
	LogCopName     string `json:LogCopName`     //物流公司名称
	LogCost        string `json:LogCost`        //费用
}

func (a *ObjChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (a *ObjChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "addProInfo" {
		return a.addProInfo(stub, args)
	} else if fn == "addIngInfo" {
		return a.addIngInfo(stub, args)
	} else if fn == "getObjInfo" {
		return a.getObjInfo(stub, args)
	} else if fn == "addLogInfo" {
		return a.addLogInfo(stub, args)
	} else if fn == "getProInfo" {
		return a.getProInfo(stub, args)
	} else if fn == "getLogInfo" {
		return a.getLogInfo(stub, args)
	} else if fn == "getIngInfo" {
		return a.getIngInfo(stub, args)
	} else if fn == "getLogInfo_l" {
		return a.getLogInfo_l(stub, args)
	}

	return shim.Error("Recevied unkown function invocation")
}

func (a *ObjChainCode) addProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var ObjInfos ObjInfo

	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments.")
	}
	ObjInfos.ObjID = args[0]
	if ObjInfos.ObjID == "" {
		return shim.Error("ObjID can not be empty.")
	}

	ObjInfos.ObjProInfo.ObjName = args[1]
	ObjInfos.ObjProInfo.ObjSpec = args[2]
	ObjInfos.ObjProInfo.ObjMFGDate = args[3]
	ObjInfos.ObjProInfo.ObjEXPDate = args[4]
	ObjInfos.ObjProInfo.ObjLOT = args[5]
	ObjInfos.ObjProInfo.ObjQSID = args[6]
	ObjInfos.ObjProInfo.ObjMFRSName = args[7]
	ObjInfos.ObjProInfo.ObjProPrice = args[8]
	ObjInfos.ObjProInfo.ObjProPlace = args[9]
	ProInfosJSONasBytes, err := json.Marshal(ObjInfos)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ObjInfos.ObjID, ProInfosJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (a *ObjChainCode) addIngInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var ObjInfos ObjInfo
	var IngInfoitem IngInfo

	if (len(args)-1)%2 != 0 || len(args) == 1 {
		return shim.Error("Incorrect number of arguments")
	}

	ObjID := args[0]
	for i := 1; i < len(args); {
		IngInfoitem.IngID = args[i]
		IngInfoitem.IngName = args[i+1]
		ObjInfos.ObjIngInfo = append(ObjInfos.ObjIngInfo, IngInfoitem)
		i = i + 2
	}

	ObjInfos.ObjID = ObjID
	/*  ObjInfos.ObjIngInfo = objIngInfo*/
	IngInfoJsonAsBytes, err := json.Marshal(ObjInfos)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ObjInfos.ObjID, IngInfoJsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (a *ObjChainCode) addLogInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	var ObjInfos ObjInfo

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments.")
	}
	ObjInfos.ObjID = args[0]
	if ObjInfos.ObjID == "" {
		return shim.Error("ObjID can not be empty.")
	}
	ObjInfos.ObjLogInfo.LogDepartureTm = args[1]
	ObjInfos.ObjLogInfo.LogArrivalTm = args[2]
	ObjInfos.ObjLogInfo.LogMission = args[3]
	ObjInfos.ObjLogInfo.LogDeparturePl = args[4]
	ObjInfos.ObjLogInfo.LogDest = args[5]
	ObjInfos.ObjLogInfo.LogToSeller = args[6]
	ObjInfos.ObjLogInfo.LogStorageTm = args[7]
	ObjInfos.ObjLogInfo.LogMOT = args[8]
	ObjInfos.ObjLogInfo.LogCopName = args[9]
	ObjInfos.ObjLogInfo.LogCost = args[10]

	LogInfosJSONasBytes, err := json.Marshal(ObjInfos)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(ObjInfos.ObjID, LogInfosJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (a *ObjChainCode) getObjInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	ObjID := args[0]
	resultsIterator, err := stub.GetHistoryForKey(ObjID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var objAllinfo ObjAllInfo

	for resultsIterator.HasNext() {
		var ObjInfos ObjInfo
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &ObjInfos)
		if ObjInfos.ObjProInfo.ObjName != "" {
			objAllinfo.ObjProInfo = ObjInfos.ObjProInfo
		} else if ObjInfos.ObjIngInfo != nil {
			objAllinfo.ObjIngInfo = ObjInfos.ObjIngInfo
		} else if ObjInfos.ObjLogInfo.LogMission != "" {
			objAllinfo.ObjLogInfo = append(objAllinfo.ObjLogInfo, ObjInfos.ObjLogInfo)
		}

	}

	jsonsAsBytes, err := json.Marshal(objAllinfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonsAsBytes)
}

func (a *ObjChainCode) getProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	ObjID := args[0]
	resultsIterator, err := stub.GetHistoryForKey(ObjID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var objProInfo ProInfo

	for resultsIterator.HasNext() {
		var ObjInfos ObjInfo
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &ObjInfos)
		if ObjInfos.ObjProInfo.ObjName != "" {
			objProInfo = ObjInfos.ObjProInfo
			continue
		}
	}
	jsonsAsBytes, err := json.Marshal(objProInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func (a *ObjChainCode) getIngInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}
	ObjID := args[0]
	resultsIterator, err := stub.GetHistoryForKey(ObjID)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var objIngInfo []IngInfo
	for resultsIterator.HasNext() {
		var ObjInfos ObjInfo
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &ObjInfos)
		if ObjInfos.ObjIngInfo != nil {
			objIngInfo = ObjInfos.ObjIngInfo
			continue
		}
	}
	jsonsAsBytes, err := json.Marshal(objIngInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func (a *ObjChainCode) getLogInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var LogInfos []LogInfo

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}

	ObjID := args[0]
	resultsIterator, err := stub.GetHistoryForKey(ObjID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		var ObjInfos ObjInfo
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &ObjInfos)
		if ObjInfos.ObjLogInfo.LogMission != "" {
			LogInfos = append(LogInfos, ObjInfos.ObjLogInfo)
		}
	}
	jsonsAsBytes, err := json.Marshal(LogInfos)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func (a *ObjChainCode) getLogInfo_l(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var Loginfo LogInfo

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}

	ObjID := args[0]
	resultsIterator, err := stub.GetHistoryForKey(ObjID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		var ObjInfos ObjInfo
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		json.Unmarshal(response.Value, &ObjInfos)
		if ObjInfos.ObjLogInfo.LogMission != "" {
			Loginfo = ObjInfos.ObjLogInfo
			continue
		}
	}
	jsonsAsBytes, err := json.Marshal(Loginfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonsAsBytes)
}

func main() {
	err := shim.Start(new(ObjChainCode))
	if err != nil {
		fmt.Printf("Error starting Obj chaincode: %s ", err)
	}
}
