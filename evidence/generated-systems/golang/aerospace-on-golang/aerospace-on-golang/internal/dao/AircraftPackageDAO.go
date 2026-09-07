package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftPackageDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraftPackage - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraftPackage(obj model.AircraftPackage)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AircraftPackage with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AircraftPackage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraftPackage", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraftPackage - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraftPackage(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AircraftPackage

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AircraftPackage with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AircraftPackage using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AircraftPackage using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraftPackage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraftPackage - returns all
//----------------------------------------------------------------------------
func GetAllAircraftPackage()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AircraftPackage

	//----------------------------------------------------------------------------
	// Request the ORM to find all AircraftPackage
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AircraftPackage" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AircraftPackage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraftPackage", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraftPackage - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraftPackage(obj model.AircraftPackage)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AircraftPackage using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AircraftPackage using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraftPackage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraftPackage - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraftPackage(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AircraftPackage with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraftPackage(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AircraftPackage)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AircraftPackage using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AircraftPackage using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraftPackage", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more optionsIds as a Options to a AircraftPackage
//----------------------------------------------------------------------------
func AddOptionsToAircraftPackage ( aircraftPackageId uint64, optionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftPackage(aircraftPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftPackage)

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
		// retrieve the modified AircraftPackage from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftPackage(aircraftPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more optionsIds as a Options from a AircraftPackage
//----------------------------------------------------------------------------
func RemoveOptionsFromAircraftPackage( aircraftPackageId uint64, optionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftPackage(aircraftPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftPackage)

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
		// retrieve the modified AircraftPackage from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftPackage(aircraftPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more variantsIds as a Variants to a AircraftPackage
//----------------------------------------------------------------------------
func AddVariantsToAircraftPackage ( aircraftPackageId uint64, variantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftPackage(aircraftPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftPackage)

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
		// retrieve the modified AircraftPackage from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftPackage(aircraftPackageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variantsIds as a Variants from a AircraftPackage
//----------------------------------------------------------------------------
func RemoveVariantsFromAircraftPackage( aircraftPackageId uint64, variantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftPackage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftPackage(aircraftPackageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftPackage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftPackage)

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
		// retrieve the modified AircraftPackage from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftPackage(aircraftPackageId)

	} else {
		return parentRequestResult
	}
}

