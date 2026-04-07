package tui

import "strings"

// LockNumberByCriteriaJSONPlaceholder is example JSON for
// serviceprovisioning.LockNumberByCriteriaResponse (UnmarshalAndValidate).
const LockNumberByCriteriaJSONPlaceholder = `{
  "responseHeader": {
    "resourceGroupId": "123",
    "resourceOrderId": "123",
    "resultCode": "0",
    "resultDesc": "OK",
    "developerMessage": "",
    "userSys": "SYS",
    "reTransmit": "N"
  },
  "resourceItemList": [
    {
      "resourceName": "lockNumberByCriteriaPrepaid",
      "resourceItemStatus": "success",
      "errorFlag": "0",
      "resourceItemErrMessage": "",
      "specialErrHandling": {
        "suppCode": [],
        "taskKeyCondition": [],
        "taskDeveloperMessage": [""]
      },
      "requestPrepResponse": [
        { "mobileNo": "0611234567" }
      ]
    }
  ]
}`

const lockNumberByMobileJSONPlaceholder = `{
  "responseHeader": {
    "resourceGroupId": "123",
    "resourceOrderId": "123",
    "resultCode": "0",
    "resultDesc": "OK",
    "developerMessage": "",
    "userSys": "SYS",
    "reTransmit": "N"
  },
  "resourceItemList": [
    {
      "resourceName": "lockNumberByMobilePrepaid",
      "resourceItemStatus": "Success",
      "errorFlag": "1",
      "resourceItemErrMessage": "Success",
      "specialErrHandling": {
        "suppCode": [],
        "taskKeyCondition": [],
        "taskDeveloperMessage": []
      }
    }
  ]
}`

const clearNumberPreparationJSONPlaceholder = `{
  "responseHeader": {
    "resourceGroupId": "123",
    "resourceOrderId": "123",
    "resultCode": "0",
    "resultDesc": "OK",
    "developerMessage": "",
    "userSys": "SYS",
    "reTransmit": "N"
  },
  "resourceItemList": [
    {
      "resourceName": "clearNumberPreparationPrepaid",
      "resourceItemStatus": "Success",
      "errorFlag": "1",
      "resourceItemErrMessage": "Success",
      "specialErrHandling": {
        "suppCode": [],
        "taskKeyCondition": [],
        "taskDeveloperMessage": []
      }
    }
  ]
}`

const querySimInfoJSONPlaceholder = `{
  "responseHeader": {
    "resourceGroupId": "123",
    "resourceOrderId": "123",
    "resultCode": "0",
    "resultDesc": "OK",
    "developerMessage": "",
    "userSys": "SYS",
    "reTransmit": "N"
  },
  "resourceItemList": [
    {
      "resourceName": "querySimInfo",
      "resourceItemStatus": "Success",
      "errorFlag": "1",
      "resourceItemErrMessage": "Success",
      "specialErrHandling": {
        "suppCode": [],
        "taskKeyCondition": [],
        "taskDeveloperMessage": []
      },
      "simSerialNoList": [
        {
          "simSerialNo": "89550000000000000000",
          "preparationDate": "28/02/202410:07:06",
          "simSerialNoStatus": "Reserved",
          "statusDate": "28/02/202410:07:06",
          "expiryDate": "28/02/202623:59:59",
          "packageNo": "9991425266",
          "subRegion": "C301",
          "packType": "X",
          "subPackType": "K1",
          "mobileNo": "0983044861",
          "mobileNoStatus": "Reserved",
          "numberClass": "Normal",
          "numberPattern": "77",
          "luckyName": "Mor_MAN",
          "luckyType": "GoodLove",
          "qrCodeInfo": "LPA:1$example$80D88923FADA3C76656D344AF",
          "material": "1000022401"
        }
      ]
    }
  ]
}`

const requestPrepNoJSONPlaceholder = `{
  "responseHeader": {
    "resourceGroupId": "123",
    "resourceOrderId": "123",
    "resultCode": "0",
    "resultDesc": "OK",
    "developerMessage": "",
    "userSys": "SYS",
    "reTransmit": "N"
  },
  "resourceItemList": [
    {
      "resourceName": "requestPrepNoPrepaid",
      "resourceItemStatus": "Success",
      "errorFlag": "1",
      "resourceItemErrMessage": "Success",
      "specialErrHandling": {
        "suppCode": [],
        "taskKeyCondition": [],
        "taskDeveloperMessage": []
      },
      "packageRowId": "1234567890",
      "offeringName": "Offering Name",
      "offeringCode": "Offering Code",
      "prepNoFrom": "9300015000",
      "prepNoTo": "9300015009"
    }
  ]
}`

const confirmPreparationJSONPlaceholder = `{
  "responseHeader": {
    "resourceGroupId": "123",
    "resourceOrderId": "123",
    "resultCode": "0",
    "resultDesc": "OK",
    "developerMessage": "",
    "userSys": "SYS",
    "reTransmit": "N"
  },
  "resourceItemList": [
    {
      "resourceName": "confirmPreparationPrepaid",
      "resourceItemStatus": "Success",
      "errorFlag": "1",
      "resourceItemErrMessage": "Success",
      "specialErrHandling": {
        "suppCode": [],
        "taskKeyCondition": [],
        "taskDeveloperMessage": []
      },
      "confirmPrepResponse": [
        {
          "simSerialNo": "89550000000000000000",
          "mobileNo": "0983044861",
          "prepNo": "9300015000",
          "expiryDate": "31/06/2026",
          "regionCode": "C301",
          "classifyCode": "N",
          "patternNo": "51",
          "numberStatusTo": "B",
          "simType": "USIM",
          "package": "PKG01",
          "packageRowId": "1234567890",
          "luckyName": "Mor_AIS",
          "luckyType": "Good Money & Love",
          "qrCodeInfo": "LPA:1$secsmsminiapp.eastcompeace.com$80D88923FADA3C76656D344AF"
        }
      ]
    }
  ]
}`

// ServiceProvisioningMockPlaceholder returns example response JSON for setting a mock
// from the TUI for the given serviceProvisioning resourceName.
func ServiceProvisioningMockPlaceholder(resourceName string) string {
	switch resourceName {
	case "lockNumberByCriteriaPrepaid":
		return LockNumberByCriteriaJSONPlaceholder
	case "lockNumberByCriteriaPostpaid":
		return strings.Replace(LockNumberByCriteriaJSONPlaceholder, "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid", 1)
	case "lockNumberByMobilePrepaid":
		return lockNumberByMobileJSONPlaceholder
	case "lockNumberByMobilePostpaid":
		return strings.Replace(lockNumberByMobileJSONPlaceholder, "lockNumberByMobilePrepaid", "lockNumberByMobilePostpaid", 1)
	case "clearNumberPreparationPrepaid":
		return clearNumberPreparationJSONPlaceholder
	case "clearNumberPreparationPostpaid":
		return strings.Replace(clearNumberPreparationJSONPlaceholder, "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid", 1)
	case "querySimInfo":
		return querySimInfoJSONPlaceholder
	case "requestPrepNoPrepaid":
		return requestPrepNoJSONPlaceholder
	case "requestPrepNoPostpaid":
		return strings.Replace(requestPrepNoJSONPlaceholder, "requestPrepNoPrepaid", "requestPrepNoPostpaid", 1)
	case "confirmPreparationPrepaid":
		return confirmPreparationJSONPlaceholder
	case "confirmPreparationPostpaid":
		return strings.Replace(confirmPreparationJSONPlaceholder, "confirmPreparationPrepaid", "confirmPreparationPostpaid", 1)
	default:
		return "{}"
	}
}

const phxRequestESIMResponsePlaceholder = `{
  "resultCode": "20000",
  "resultDesc": "Success",
  "resultData": {
    "newSimItem": {
      "imsi": "1234567890",
      "qrCodeInfo": "LPA:1$example$80D88923FADA3C76656D344AF",
      "regionCode": "C301",
      "serialNo": "89550000000000000000"
    }
  }
}`

const phxNewRegistrationResponsePlaceholder = `{
  "resultCode": "20000",
  "resultDesc": "Success"
}`

// PHXMockPlaceholder returns example response JSON for the given PHX API name (see PHXApis).
func PHXMockPlaceholder(apiName string) string {
	switch apiName {
	case "requestESIM":
		return phxRequestESIMResponsePlaceholder
	case "newRegistration":
		return phxNewRegistrationResponsePlaceholder
	default:
		return "{}"
	}
}
