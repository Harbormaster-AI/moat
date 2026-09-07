package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftModelDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraftModel - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraftModel(obj model.AircraftModel)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AircraftModel with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AircraftModel", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraftModel", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraftModel - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraftModel(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AircraftModel

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AircraftModel with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AircraftModel using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AircraftModel using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraftModel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraftModel - returns all
//----------------------------------------------------------------------------
func GetAllAircraftModel()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AircraftModel

	//----------------------------------------------------------------------------
	// Request the ORM to find all AircraftModel
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AircraftModel" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AircraftModel", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraftModel", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraftModel - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraftModel(obj model.AircraftModel)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AircraftModel using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AircraftModel using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraftModel", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraftModel - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraftModel(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AircraftModel with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraftModel(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AircraftModel)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AircraftModel using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AircraftModel using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraftModel", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Family on a AircraftModel
//----------------------------------------------------------------------------
func AssignFamilyToAircraftModel( aircraftModelId uint64, familyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftModel(aircraftModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftModel)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftFamily

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftFamily with a
		// matching familyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, familyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Family	to the AircraftModel
			//----------------------------------------------------------------------------
			parentObj.Family = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftModel
			//----------------------------------------------------------------------------
			return UpdateAircraftModel(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Family", familyId )
			return utils.RequestResult{false, msg, "assignFamily", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Family on a AircraftModel
//----------------------------------------------------------------------------
func UnassignFamilyFromAircraftModel(aircraftModelId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftModel(aircraftModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftModel)

		//----------------------------------------------------------------------------
		// assign an empty AircraftFamily to the Family
		//----------------------------------------------------------------------------
		parentObj.Family = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Family
		//----------------------------------------------------------------------------
		parentObj.FamilyId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftModel
		//----------------------------------------------------------------------------
		return UpdateAircraftModel(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more variantsIds as a Variants to a AircraftModel
//----------------------------------------------------------------------------
func AddVariantsToAircraftModel ( aircraftModelId uint64, variantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftModel(aircraftModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftModel)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Variants using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "unassignVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftModel from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftModel(aircraftModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variantsIds as a Variants from a AircraftModel
//----------------------------------------------------------------------------
func RemoveVariantsFromAircraftModel( aircraftModelId uint64, variantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftModel(aircraftModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftModel)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftVariantObj from the Variants array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "removeVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftModel from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftModel(aircraftModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more engineTypesIds as a EngineTypes to a AircraftModel
//----------------------------------------------------------------------------
func AddEngineTypesToAircraftModel ( aircraftModelId uint64, engineTypesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftModel(aircraftModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftModel)

		// slice the ids on comma with no spaces
		ids := strings.Split( engineTypesIds, ",")

		for _, engineTypesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EngineType

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EngineType
			// with a matching engineTypesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , engineTypesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EngineTypes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EngineTypes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EngineTypes", engineTypesId )
				return utils.RequestResult{false, msg, "unassignEngineTypes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftModel from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftModel(aircraftModelId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more engineTypesIds as a EngineTypes from a AircraftModel
//----------------------------------------------------------------------------
func RemoveEngineTypesFromAircraftModel( aircraftModelId uint64, engineTypesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftModel with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftModel(aircraftModelId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftModel so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftModel)

		// slice the ids on comma with no spaces
		ids := strings.Split( engineTypesIds, ",")

		for _, engineTypesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EngineType

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EngineType
			// with a matching engineTypesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , engineTypesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EngineTypeObj from the EngineTypes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EngineTypes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EngineTypes", engineTypesId )
				return utils.RequestResult{false, msg, "removeEngineTypes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftModel from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftModel(aircraftModelId)

	} else {
		return parentRequestResult
	}
}

