package tui

import (
	"encoding/json"

	"github.com/DilanHera/mockTP/internal/services/dt"
	"github.com/DilanHera/mockTP/internal/services/eos"
	"github.com/DilanHera/mockTP/internal/services/esb"
	"github.com/DilanHera/mockTP/internal/services/ids"
	"github.com/DilanHera/mockTP/internal/services/im"
	"github.com/DilanHera/mockTP/internal/services/mychannel"
	pgzinvmodel "github.com/DilanHera/mockTP/internal/services/pgzinv/model"
	"github.com/DilanHera/mockTP/internal/services/pgzinv/serviceprovisioning"
	"github.com/DilanHera/mockTP/internal/services/phx"
	"github.com/DilanHera/mockTP/internal/services/smis"
)

func EsbMockPlaceholder(apiName string) string {
	switch apiName {
	case "oauthToken":
		res, _ := json.Marshal(esb.OauthTokenResponse{
			AccessToken: "mock-token",
			TokenType:   "bearer",
			ExpiresIn:   3600,
			Error:       "remove for success case",
		})
		return string(res)
	case "createFreightOrder":
		res, _ := json.Marshal(esb.CreateFreightOrderResponse{
			FreightOrderNumber: "6200088900",
			MessageType:        "S",
			MessageDesc:        "Business document with temporary number $1 saved as business doc. 6200088900",
			MessageID:          "ECF92F6BF2EA4976B423247278CEB458",
			PartnerName:        "OPTIMUS",
			PartnerMessageID:   "20260427080253557",
		})
		return string(res)
	case "doCreation":
		res, _ := json.Marshal(esb.DOCreationResponse{
			MessageType:      "S",
			MessageDesc:      "Outbound Delivery created",
			MessageID:        "A4DA6775FDAC4415BF683366C0E3401A",
			PartnerName:      "SAP",
			PartnerMessageID: "8e864f32-6beb-4e90-bb76-0355d7e31d86_1",
			Item: []esb.DOCreationResponseItem{
				{
					MessageType:        "S",
					MessageDesc:        "Outbound Delivery created",
					InterfacerecordID:  "1",
					SDdocumentcateg:    "J",
					SDDocument:         "2001085811",
					SDItem:             "10",
					QuantitySalesUom:   "1.000",
					SalesUnit:          "PC",
					SDDocumentCategory: "J",
				},
			},
		})
		return string(res)
	case "legoUpdateOrderStatus":
		res, _ := json.Marshal(esb.LegoupdateOrderStatusResponse{
			TransactionID: "2ad5135a-8cb6-482a-9f48-610ef68cc435",
			ResultCode:    "20000",
			ResultMessage: "Success.",
			Result:        "{}",
			StatusOrder:   "03",
			OrderNo:       "AS2604261179348",
		})
		return string(res)
	case "persoSim":
		res, _ := json.Marshal(esb.PersosimResponse{
			StatusCode:        "200",
			StatusDescription: "Success",
		})
		return string(res)
	case "serialNumberExpirationDate":
		res, _ := json.Marshal(esb.SerialNumberExpirationDateResponse{
			MessageID:        "AGnu3gnT_T6pmJOnuzxDk0BX4EEg",
			PartnerName:      "SAP",
			PartnerMessageID: "0953A48139EC4CBCA97AF7B402305C6E",
			Item: []esb.SerialNumberExpirationDateResponseItem{
				{
					MessageType:    "S",
					MessageClass:   "ZSCM00",
					MessageNumber:  "000",
					MessageDesc:    "Equipment 37127803 has been updated",
					Material:       "1000000178",
					SerialNumber:   "2610506305004",
					ExpirationDate: "20270531",
					ConfirmStatus:  "X",
				},
			},
		})
		return string(res)
	default:
		return "{}"
	}
}

func ImMockPlaceholder(apiName string) string {
	switch apiName {
	case "sendSimSerialNo":
		res, _ := json.Marshal(im.SendSimSerialNoResponse{
			StatusDescription: "success",
			OrderNo:           "mock-order-no",
			TrackingNo:        "mock-tracking-no",
			MobileNo:          "0999999999",
			Imsi:              "mock-imsi",
			TransactionId:     "mock-transaction-id",
			StatusCode:        "200",
			SerialNo:          "mock-serial-no",
		})
		return string(res)
	default:
		return "{}"
	}
}

func EosMockPlaceholder(apiName string) string {
	switch apiName {
	case "updateSimSerialNo":
		res, _ := json.Marshal(eos.UpdateSimSerialNoResponse{
			StatusCode:        "200",
			StatusDescription: "Success",
		})
		return string(res)
	default:
		return "{}"
	}
}

func IdsMockPlaceholder(apiName string) string {
	switch apiName {
	case "userInfo":
		res, _ := json.Marshal(ids.UserInfoResponse{
			Sub:               "EMPLOYEELDAP",
			Pincode:           "00066026",
			Firstname:         "Chomnipha",
			Mobile:            "0992457344",
			Groups:            "Internal/everyone",
			Section:           "Supply Chain - E-Commerce Fulfillment",
			PreferredUsername: "Chomnipha Wetchiyo",
			Title:             "Administrative Support staff",
			Consent:           "Y",
			Lastname:          "Wetchiyo",
			LocationCode:      "Location_WDS",
			Name:              "Chomnipha Wetchiyo",
			Company:           "AWN",
			PhoneNumber:       "0992457344",
			Department:        "Supply Chain - WholeSale Fulfillment",
			FamilyName:        "Wetchiyo",
			Email:             "chomnipw@ais.co.th",
			Username:          "chomnipw",
		})
		return string(res)
	default:
		return "{}"
	}
}

func PhxMockPlaceholder(apiName string) string {
	switch apiName {
	case "requestESIM":
		res, _ := json.Marshal(phx.RequestESIMResponse{
			ResultCode: "20000",
			ResultDesc: "Success",
			ResultData: phx.ResultData{
				NewSimItem: phx.NewSimItem{
					Imsi:       "1234567890",
					QRCodeInfo: "1234567890",
					RegionCode: "1234567890",
					SerialNo:   "1234567890",
				},
			},
		})
		return string(res)
	case "newRegistration":
		res, _ := json.Marshal(phx.NewRegistrationResponse{
			ResultCode: "20000",
			ResultDesc: "Success",
		})
		return string(res)
	case "encryptLib":
		res, _ := json.Marshal(phx.EncryptLibResponse{
			ResultCode:       "20000",
			ResultDesc:       "Success",
			DeveloperMessage: "Success",
			ResultData: phx.ResultDataEncryptLib{
				Key:   "perso.encryptLib",
				Value: "mock-encrypted-value",
			},
		})
		return string(res)
	case "checkPerso":
		res, _ := json.Marshal(phx.CheckPersoResponse{
			ResultCode:       "20000",
			DeveloperMessage: "",
			ResultDesc:       "Success",
		})
		return string(res)
	case "productProvisioning":
		res, _ := json.Marshal(phx.ProductProvisioningResponse{
			ResultCode: "20000",
			ResultDesc: "Success",
		})
		return string(res)
	default:
		return "{}"
	}
}

func SmisMockPlaceholder(apiName string) string {
	switch apiName {
	case "updateSerial":
		res, _ := json.Marshal(smis.UpdateSerialResponse{
			ResponseCode:    "0000",
			ResponseMessage: "Success",
		})
		return string(res)
	default:
		return "{}"
	}
}

func DtMockPlaceholder(apiName string) string {
	switch apiName {
	case "listOrderNoByDono":
		res, _ := json.Marshal(dt.ListOrderNoByDonoResponse{
			ResultCode:        "20000",
			ResultDescription: "Success",
			ResultStatus:      "S",
			ResultObj: []dt.ListOrderNo{
				{ListOrderNo: []string{"1234567890"}},
			},
		})
		return string(res)
	case "pickingDocument":
		res, _ := json.Marshal(dt.PickingDocumentResponse{
			ResultCode: "20000",
			ResultDesc: "Success picking document list",
			Status:     "S",
		})
		return string(res)
	case "queryPrint":
		res, _ := json.Marshal(dt.QueryPrintResponse{
			ResultCode:        "20000",
			ResultDescription: "Success",
			DeveloperMessage:  "Success",
			DataList:          []dt.DataList{},
		})
		return string(res)
	case "queryStockImeiMyStore":
		res, _ := json.Marshal(dt.QueryStockImeiMyStoreResponse{
			ResultCode:        "20000",
			ResultDescription: "Success",
			DeveloperMessage:  "Success",
			ListProduct: []dt.ListProduct{
				{
					LocationCode:   "4289",
					Company:        "WDS",
					ProductType:    "SIM",
					ProductSubtype: "PREPAID",
					Brand:          "OTC",
					Model:          "OTC110",
					UnitName:       "PC",
					MatCode:        "SIMSERVICEOPTIMUS",
					SubStockCode:   "OLS",
					SerialNo:       "000001920041336559",
					SapDescription: "SIM OTC 110B mock",
					VatType:        "Y",
					MatType:        "Serial",
				},
			},
		})
		return string(res)
	case "reprintReceiptForm":
		res, _ := json.Marshal(dt.ReprintReceiptFormResponse{
			ResultCode:        "20000",
			ResultDescription: "Success",
			DeveloperMessage:  "Success",
			Data:              "base64-encoded-pdf-data",
		})
		return string(res)
	case "updateSimSerialPerso":
		res, _ := json.Marshal(dt.UpdateSimSerialPersoResponse{
			ResultCode:        "20000",
			ResultDescription: "Success",
			Status:            "S",
		})
		return string(res)
	case "authenticate":
		res, _ := json.Marshal(dt.AuthenticateResponse{
			StatusCode: "200",
			Message:    "success authenticate",
			Token:      "mock-jwt-token",
		})
		return string(res)
	default:
		return "{}"
	}
}

func MyChannelMockPlaceholder(apiName string) string {
	switch apiName {
	case "simSerialNo":
		res, _ := json.Marshal(mychannel.SimSerialNoResponse{
			ResultCode:        "20000",
			ResultDescription: "Success",
			DeveloperMessage:  "Success",
			Data: mychannel.SimSerialNoData{
				IsSuccess: true,
			},
		})
		return string(res)
	default:
		return "{}"
	}
}

func PgzinvMockPlaceholder(resourceName string) string {
	header := pgzinvmodel.ResponseHeader{
		ResourceGroupId:  "mock-resource-group-id",
		ResourceOrderId:  "DBSIPGSA001G-PGZINV-202303171437060271",
		ReTransmit:       "0",
		UserSys:          "mock-user-sys",
		DeveloperMessage: "",
		ResultCode:       "20000",
		ResultDesc:       "Success",
	}
	baseItem := pgzinvmodel.ResourceItemListBase{
		ResourceName:           resourceName,
		ResourceItemStatus:     "Success",
		ErrorFlag:              "1",
		ResourceItemErrMessage: "Success",
		SpecialErrHandling: pgzinvmodel.SpecialErrHandling{
			SuppCode:             []string{},
			TaskKeyCondition:     []string{},
			TaskDeveloperMessage: []string{},
		},
	}

	switch resourceName {
	case "lockNumberByCriteriaPrepaid", "lockNumberByCriteriaPostpaid":
		res, _ := json.Marshal(serviceprovisioning.LockNumberByCriteriaResponse{
			ResponseHeader: header,
			ResourceItemList: []serviceprovisioning.LockNumberByCriteriaResponseItem{
				{
					ResourceItemListBase: baseItem,
					Key:                 "1234567",
					RequestPrepResponse: []serviceprovisioning.RequestPrepResponseItem{
						{MobileNo: "0610000001"},
					},
				},
			},
		})
		return string(res)
	case "lockNumberByMobilePrepaid", "lockNumberByMobilePostpaid":
		res, _ := json.Marshal(serviceprovisioning.LockNumberByMobileResponse{
			ResponseHeader:   header,
			ResourceItemList: []pgzinvmodel.ResourceItemListBase{baseItem},
		})
		return string(res)
	case "clearNumberPreparationPrepaid", "clearNumberPreparationPostpaid":
		res, _ := json.Marshal(serviceprovisioning.ClearNumberPreparationResponse{
			ResponseHeader:   header,
			ResourceItemList: []pgzinvmodel.ResourceItemListBase{baseItem},
		})
		return string(res)
	case "querySimInfo":
		res, _ := json.Marshal(serviceprovisioning.QuerySimInfoResponse{
			ResponseHeader: header,
			ResourceItemList: []serviceprovisioning.QuerySimInfoResponseItem{
				{
					ResourceItemListBase: baseItem,
					SimSerialNoList: []serviceprovisioning.SimSerialNoListItem{
						{
							SimSerialNo:       "mock-serial-no",
							PreparationDate:   "28/02/202410:07:06",
							SimSerialNoStatus: "Reserved",
							StatusDate:        "28/02/202410:07:06",
							ExpiryDate:        "28/02/202623:59:59",
							PackageNo:         "9991425266",
							SubRegion:         "C301",
							PackType:          "X",
							SubPackType:       "K1",
							MobileNo:          "0983044861",
							MobileNoStatus:    "Reserved",
							NumberClass:       "Normal",
							NumberPattern:     "77",
							LuckyName:         "Mor_MAN",
							LuckyType:         "GoodLove",
							QRCodeInfo:        "LPA:1$secsmsminiapp.eastcompeace.com$80D88923FADA3C76656D344AF",
							Material:          "1000022401",
						},
					},
				},
			},
		})
		return string(res)
	case "requestPrepNoPrepaid", "requestPrepNoPostpaid":
		res, _ := json.Marshal(serviceprovisioning.RequestPrepNoResponse{
			ResponseHeader: header,
			ResourceItemList: []serviceprovisioning.RequestPrepNoResponseItem{
				{
					ResourceItemListBase: baseItem,
					PackageRowId:         "1234567890",
					OfferingName:         "Offering Name",
					OfferingCode:         "Offering Code",
					PrepNoFrom:           "9300015000",
					PrepNoTo:             "9300015009",
				},
			},
		})
		return string(res)
	case "confirmPreparationPrepaid", "confirmPreparationPostpaid":
		res, _ := json.Marshal(serviceprovisioning.ConfirmPreparationResponse{
			ResponseHeader: header,
			ResourceItemList: []serviceprovisioning.ConfirmPreparationResponseItem{
				{
					ResourceItemListBase: baseItem,
					ConfirmPrepResponse: []serviceprovisioning.ConfirmPrepResponseItem{
						{
							SimSerialNo:    "mock-serial-no",
							MobileNo:       "0610000001",
							PrepNo:         "9300015000",
							ExpiryDate:     "31/06/2026",
							RegionCode:     "X330",
							ClassifyCode:   "N",
							PatternNo:      "51",
							NumberStatusTo: "B",
							SimType:        "Service",
							Package:        "mock-package",
							PackageRowId:   "1234567890",
							LuckyName:      "Mor_AIS",
							LuckyType:      "Good Money & Love",
							QRCodeInfo:     "LPA:1$secsmsminiapp.eastcompeace.com$80D88923FADA3C76656D344AF",
						},
					},
				},
			},
		})
		return string(res)
	default:
		return "{}"
	}
}
