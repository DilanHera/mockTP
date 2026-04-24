package phx

import "fmt"

type EncryptLibRequest struct {
	Key          string `json:"key" validate:"required"`
	SourceSystem string `json:"sourceSystem" validate:"required"`
}

type EncryptLibResponse struct {
	ResultCode       string               `json:"resultCode"`
	ResultDesc       string               `json:"resultDesc"`
	DeveloperMessage string               `json:"developerMessage"`
	ResultData       ResultDataEncryptLib `json:"resultData"`
}

type ResultDataEncryptLib struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (p *phx) EncryptLib(input *EncryptLibRequest) (*EncryptLibResponse, error) {
	result := p.GetApiInfo("encryptLib")
	if result.State == "C" {
		if UserEncryptLib != nil {
			return UserEncryptLib, nil
		}
		return nil, fmt.Errorf("no custom response set for encryptLib")
	}

	if result.State == "E" {
		return &EncryptLibResponse{
			ResultCode:       "50000",
			ResultDesc:       "Key Not Found",
			DeveloperMessage: "Key Not Found",
		}, nil
	}

	return &EncryptLibResponse{
		ResultCode:       "20000",
		ResultDesc:       "Success",
		DeveloperMessage: "Success",
		ResultData: ResultDataEncryptLib{
			Key:   "perso.encryptLib",
			Value: "35982ac107921d3d4cf6eee12ea3075aed384e6946f68f21582d8c427f3a097eda37673465bb5679178fd87a41d4705dfd3eb927a0cf12d1f1c28ea9ad936c5268eaa40192ea04f8c2475b66cfe9ad04e79ef1508a7eca3caa1ab39aca46996d08d9b82b951a4e075ca0ff1da0ae6e4185fbe1cde1d047676c2613f31d1ccf350c0c702eb6069059318262b401b7b4b00c9d1088a041f775ceb7778f3bcc9d841c048868f3472579d70d5e598423758dde699321b0c0dd0ead4f8d5a458541137428e7198d245d67f4cdca4228e0c3cbdc377f6573122f5c5361e63ba5976771e57d1d16ddc1845d1a0add6e1fd589fd2d9623c002f9a5ed798dcbcf2a0429f20128fa56803d666fc1be549b57210f3499c4794fde6133df33ec266237e6cda34d5475806dab2b09e850e0957ca301783bd9b65f1b587123202436ae103038d2ef601c806ea3b952e37dd97d243fcf5d95ce468abf4901320580aa28fd0db120cc7081a22b598570a3d88d8e8b3934bab537788ee327beaa3df4acaafec53eb67c6f933950546305d9d0ed8ae7671a43a72da59c26064438e65d830f31c70cfdc58436e8f16adb6f03e43b2b74bee284c47c92c2334c28212c72db2c65bb0cd9eec5601d43e3e094521ec5203eac3c37ca00b75f5573e4d765144526963360bd330a5c69cbaa8477923b6f3eee67320a4dff239d229c8892634e51f7424547744f01730592f94ff1a0bf3ad7509499e69c537854a0e0c4c316cb0bcbe731db8898a37ff066ff2feb635307b27789a97950b11e7b097dfad3b4c54d28bdac7ef9de886778e9c5e288e49584378dbdbcc14c42387c210b48116ca356035cc161ad6a104aaf88e6f2be43257d93324d03ff6eda5a894aa1911f2671256afd200895a7c29563d14ac17d37fa66c598c8873df63e0ef18868921788124c09714947dcda5b17d1cfcf6f6a66066c51c1df0f1edfa36f41081881bfd7166fa67b14fea0fe681dce0a435bc850175601e9cfe2c77c9ea2cf1f5e3ab7d3b0b244c560d2fbbe27ad493cb2b60390c1245dfe4ee9f5286a717d97830a646bda7a1e33d0a92a2ebc7515b007c87190e39ad3b345117065e15d372da4543249b4cc7818cc7cd5c01a9a39b5961d72c89fe0ee31e93237a3bf1b87f56f8b2fdea210c1fdbd7709a2151600fb38b58b1aaf827dd004655e231170222ce8467c1211af48955334fdb795ece5cddfdea60b4595abcafadafd03f3fa61173a3f424241b4ccc77798efeeeededb4d85aa2fb0178883d8a83111d4c37075520e43d7",
		},
	}, nil
}
