package canteen

import "github.com/ipp-ementa/iped/model/canteen"

// This file contains model views representation for GET functionalities of canteens collection

// GetAvailableCanteensModelView is the model view representation
// for the available canteens functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/canteens.md#available-canteens
type GetAvailableCanteensModelView []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetDetailedCanteenInformationModelView is the model view representation
// for the detailed canteen information functionality
// See more info at: https://github.com/ipp-ementa/iped-documentation/blob/master/documentation/rest_api/canteens.md#detailed-canteen-information
type GetDetailedCanteenInformationModelView struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ToGetAvailableCanteensModelView creates a GetAvailableCanteensModelView using a slice of canteen models
func ToGetAvailableCanteensModelView(canteens []canteen.Canteen) GetAvailableCanteensModelView {
	modelview := make(GetAvailableCanteensModelView, len(canteens))

	for index, canteen := range canteens {
		element := &modelview[index]
		element.ID = int(canteen.ID)
		element.Name = canteen.Name
	}

	return modelview
}

// ToGetDetailedCanteenInformationModelView creates a GetDetailedCanteenInformationModelView using a canteen model
func ToGetDetailedCanteenInformationModelView(canteen canteen.Canteen) GetDetailedCanteenInformationModelView {

	modelview := GetDetailedCanteenInformationModelView{ID: int(canteen.ID), Name: canteen.Name}

	return modelview
}
