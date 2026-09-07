package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CabinLayoutDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCabinLayout - creates a new db entry
//----------------------------------------------------------------------------
func CreateCabinLayout(obj model.CabinLayout)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a CabinLayout with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CabinLayout", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCabinLayout", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCabinLayout - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCabinLayout(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CabinLayout

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CabinLayout with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CabinLayout using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CabinLayout using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCabinLayout", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCabinLayout - returns all
//----------------------------------------------------------------------------
func GetAllCabinLayout()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CabinLayout

	//----------------------------------------------------------------------------
	// Request the ORM to find all CabinLayout
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CabinLayout" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CabinLayout", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCabinLayout", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCabinLayout - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCabinLayout(obj model.CabinLayout)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a CabinLayout using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CabinLayout using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCabinLayout", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCabinLayout - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCabinLayout(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CabinLayout with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCabinLayout(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CabinLayout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CabinLayout)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CabinLayout using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CabinLayout using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCabinLayout", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Variant on a CabinLayout
//----------------------------------------------------------------------------
func AssignVariantToCabinLayout( cabinLayoutId uint64, variantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CabinLayout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCabinLayout(cabinLayoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CabinLayout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CabinLayout)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftVariant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftVariant with a
		// matching variantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, variantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Variant	to the CabinLayout
			//----------------------------------------------------------------------------
			parentObj.Variant = &childObj

			//----------------------------------------------------------------------------
			// save the CabinLayout
			//----------------------------------------------------------------------------
			return UpdateCabinLayout(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variant", variantId )
			return utils.RequestResult{false, msg, "assignVariant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Variant on a CabinLayout
//----------------------------------------------------------------------------
func UnassignVariantFromCabinLayout(cabinLayoutId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CabinLayout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCabinLayout(cabinLayoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CabinLayout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CabinLayout)

		//----------------------------------------------------------------------------
		// assign an empty AircraftVariant to the Variant
		//----------------------------------------------------------------------------
		parentObj.Variant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Variant
		//----------------------------------------------------------------------------
		parentObj.VariantId = nil;

		//----------------------------------------------------------------------------
		// save the CabinLayout
		//----------------------------------------------------------------------------
		return UpdateCabinLayout(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more aircraftIds as a Aircraft to a CabinLayout
//----------------------------------------------------------------------------
func AddAircraftToCabinLayout ( cabinLayoutId uint64, aircraftIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CabinLayout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCabinLayout(cabinLayoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CabinLayout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CabinLayout)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftIds, ",")

		for _, aircraftId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Aircraft

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Aircraft
			// with a matching aircraftId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Aircraft using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Aircraft").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Aircraft", aircraftId )
				return utils.RequestResult{false, msg, "unassignAircraft", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CabinLayout from the gorm
		//----------------------------------------------------------------------------
		return GetCabinLayout(cabinLayoutId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more aircraftIds as a Aircraft from a CabinLayout
//----------------------------------------------------------------------------
func RemoveAircraftFromCabinLayout( cabinLayoutId uint64, aircraftIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CabinLayout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCabinLayout(cabinLayoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CabinLayout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CabinLayout)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftIds, ",")

		for _, aircraftId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Aircraft

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Aircraft
			// with a matching aircraftId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftObj from the Aircraft array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Aircraft").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Aircraft", aircraftId )
				return utils.RequestResult{false, msg, "removeAircraft", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CabinLayout from the gorm
		//----------------------------------------------------------------------------
		return GetCabinLayout(cabinLayoutId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more optionsIds as a Options to a CabinLayout
//----------------------------------------------------------------------------
func AddOptionsToCabinLayout ( cabinLayoutId uint64, optionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CabinLayout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCabinLayout(cabinLayoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CabinLayout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CabinLayout)

		// slice the ids on comma with no spaces
		ids := strings.Split( optionsIds, ",")

		for _, optionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftOption

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftOption
			// with a matching optionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , optionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Options using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Options").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Options", optionsId )
				return utils.RequestResult{false, msg, "unassignOptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CabinLayout from the gorm
		//----------------------------------------------------------------------------
		return GetCabinLayout(cabinLayoutId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more optionsIds as a Options from a CabinLayout
//----------------------------------------------------------------------------
func RemoveOptionsFromCabinLayout( cabinLayoutId uint64, optionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CabinLayout with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCabinLayout(cabinLayoutId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CabinLayout so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CabinLayout)

		// slice the ids on comma with no spaces
		ids := strings.Split( optionsIds, ",")

		for _, optionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftOption

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftOption
			// with a matching optionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , optionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftOptionObj from the Options array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Options").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Options", optionsId )
				return utils.RequestResult{false, msg, "removeOptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CabinLayout from the gorm
		//----------------------------------------------------------------------------
		return GetCabinLayout(cabinLayoutId)

	} else {
		return parentRequestResult
	}
}

