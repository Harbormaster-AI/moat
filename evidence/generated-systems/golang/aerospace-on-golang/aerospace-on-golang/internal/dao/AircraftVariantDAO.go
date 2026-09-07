package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftVariantDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraftVariant - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraftVariant(obj model.AircraftVariant)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AircraftVariant with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AircraftVariant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraftVariant", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraftVariant - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraftVariant(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AircraftVariant

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AircraftVariant with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AircraftVariant using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AircraftVariant using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraftVariant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraftVariant - returns all
//----------------------------------------------------------------------------
func GetAllAircraftVariant()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AircraftVariant

	//----------------------------------------------------------------------------
	// Request the ORM to find all AircraftVariant
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AircraftVariant" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AircraftVariant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraftVariant", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraftVariant - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraftVariant(obj model.AircraftVariant)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AircraftVariant using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AircraftVariant using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraftVariant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraftVariant - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraftVariant(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraftVariant(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AircraftVariant using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AircraftVariant using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraftVariant", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Model_ on a AircraftVariant
//----------------------------------------------------------------------------
func AssignModel_ToAircraftVariant( aircraftVariantId uint64, model_Id uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftModel

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftModel with a
		// matching model_Id
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, model_Id).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Model_	to the AircraftVariant
			//----------------------------------------------------------------------------
			parentObj.Model_ = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftVariant
			//----------------------------------------------------------------------------
			return UpdateAircraftVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Model_", model_Id )
			return utils.RequestResult{false, msg, "assignModel_", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Model_ on a AircraftVariant
//----------------------------------------------------------------------------
func UnassignModel_FromAircraftVariant(aircraftVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// assign an empty AircraftModel to the Model_
		//----------------------------------------------------------------------------
		parentObj.Model_ = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Model_
		//----------------------------------------------------------------------------
		parentObj.Model_Id = nil;

		//----------------------------------------------------------------------------
		// save the AircraftVariant
		//----------------------------------------------------------------------------
		return UpdateAircraftVariant(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a EngineType on a AircraftVariant
//----------------------------------------------------------------------------
func AssignEngineTypeToAircraftVariant( aircraftVariantId uint64, engineTypeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.EngineType

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a EngineType with a
		// matching engineTypeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, engineTypeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the EngineType	to the AircraftVariant
			//----------------------------------------------------------------------------
			parentObj.EngineType = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftVariant
			//----------------------------------------------------------------------------
			return UpdateAircraftVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EngineType", engineTypeId )
			return utils.RequestResult{false, msg, "assignEngineType", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a EngineType on a AircraftVariant
//----------------------------------------------------------------------------
func UnassignEngineTypeFromAircraftVariant(aircraftVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// assign an empty EngineType to the EngineType
		//----------------------------------------------------------------------------
		parentObj.EngineType = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the EngineType
		//----------------------------------------------------------------------------
		parentObj.EngineTypeId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftVariant
		//----------------------------------------------------------------------------
		return UpdateAircraftVariant(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AvionicsSuite on a AircraftVariant
//----------------------------------------------------------------------------
func AssignAvionicsSuiteToAircraftVariant( aircraftVariantId uint64, avionicsSuiteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AvionicsSuite

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AvionicsSuite with a
		// matching avionicsSuiteId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, avionicsSuiteId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AvionicsSuite	to the AircraftVariant
			//----------------------------------------------------------------------------
			parentObj.AvionicsSuite = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftVariant
			//----------------------------------------------------------------------------
			return UpdateAircraftVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AvionicsSuite", avionicsSuiteId )
			return utils.RequestResult{false, msg, "assignAvionicsSuite", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AvionicsSuite on a AircraftVariant
//----------------------------------------------------------------------------
func UnassignAvionicsSuiteFromAircraftVariant(aircraftVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// assign an empty AvionicsSuite to the AvionicsSuite
		//----------------------------------------------------------------------------
		parentObj.AvionicsSuite = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AvionicsSuite
		//----------------------------------------------------------------------------
		parentObj.AvionicsSuiteId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftVariant
		//----------------------------------------------------------------------------
		return UpdateAircraftVariant(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Apu on a AircraftVariant
//----------------------------------------------------------------------------
func AssignApuToAircraftVariant( aircraftVariantId uint64, apuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.APU

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a APU with a
		// matching apuId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, apuId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Apu	to the AircraftVariant
			//----------------------------------------------------------------------------
			parentObj.Apu = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftVariant
			//----------------------------------------------------------------------------
			return UpdateAircraftVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Apu", apuId )
			return utils.RequestResult{false, msg, "assignApu", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Apu on a AircraftVariant
//----------------------------------------------------------------------------
func UnassignApuFromAircraftVariant(aircraftVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// assign an empty APU to the Apu
		//----------------------------------------------------------------------------
		parentObj.Apu = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Apu
		//----------------------------------------------------------------------------
		parentObj.ApuId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftVariant
		//----------------------------------------------------------------------------
		return UpdateAircraftVariant(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LandingGear on a AircraftVariant
//----------------------------------------------------------------------------
func AssignLandingGearToAircraftVariant( aircraftVariantId uint64, landingGearId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LandingGear

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LandingGear with a
		// matching landingGearId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, landingGearId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LandingGear	to the AircraftVariant
			//----------------------------------------------------------------------------
			parentObj.LandingGear = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftVariant
			//----------------------------------------------------------------------------
			return UpdateAircraftVariant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LandingGear", landingGearId )
			return utils.RequestResult{false, msg, "assignLandingGear", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LandingGear on a AircraftVariant
//----------------------------------------------------------------------------
func UnassignLandingGearFromAircraftVariant(aircraftVariantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		//----------------------------------------------------------------------------
		// assign an empty LandingGear to the LandingGear
		//----------------------------------------------------------------------------
		parentObj.LandingGear = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LandingGear
		//----------------------------------------------------------------------------
		parentObj.LandingGearId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftVariant
		//----------------------------------------------------------------------------
		return UpdateAircraftVariant(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more cabinLayoutsIds as a CabinLayouts to a AircraftVariant
//----------------------------------------------------------------------------
func AddCabinLayoutsToAircraftVariant ( aircraftVariantId uint64, cabinLayoutsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		// slice the ids on comma with no spaces
		ids := strings.Split( cabinLayoutsIds, ",")

		for _, cabinLayoutsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CabinLayout

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CabinLayout
			// with a matching cabinLayoutsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cabinLayoutsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CabinLayouts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CabinLayouts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CabinLayouts", cabinLayoutsId )
				return utils.RequestResult{false, msg, "unassignCabinLayouts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftVariant from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftVariant(aircraftVariantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more cabinLayoutsIds as a CabinLayouts from a AircraftVariant
//----------------------------------------------------------------------------
func RemoveCabinLayoutsFromAircraftVariant( aircraftVariantId uint64, cabinLayoutsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		// slice the ids on comma with no spaces
		ids := strings.Split( cabinLayoutsIds, ",")

		for _, cabinLayoutsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CabinLayout

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CabinLayout
			// with a matching cabinLayoutsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cabinLayoutsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CabinLayoutObj from the CabinLayouts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CabinLayouts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CabinLayouts", cabinLayoutsId )
				return utils.RequestResult{false, msg, "removeCabinLayouts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftVariant from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftVariant(aircraftVariantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more optionsIds as a Options to a AircraftVariant
//----------------------------------------------------------------------------
func AddOptionsToAircraftVariant ( aircraftVariantId uint64, optionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

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
		// retrieve the modified AircraftVariant from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftVariant(aircraftVariantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more optionsIds as a Options from a AircraftVariant
//----------------------------------------------------------------------------
func RemoveOptionsFromAircraftVariant( aircraftVariantId uint64, optionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

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
		// retrieve the modified AircraftVariant from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftVariant(aircraftVariantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more packagesIds as a Packages to a AircraftVariant
//----------------------------------------------------------------------------
func AddPackagesToAircraftVariant ( aircraftVariantId uint64, packagesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		// slice the ids on comma with no spaces
		ids := strings.Split( packagesIds, ",")

		for _, packagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftPackage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftPackage
			// with a matching packagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , packagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Packages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Packages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Packages", packagesId )
				return utils.RequestResult{false, msg, "unassignPackages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftVariant from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftVariant(aircraftVariantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more packagesIds as a Packages from a AircraftVariant
//----------------------------------------------------------------------------
func RemovePackagesFromAircraftVariant( aircraftVariantId uint64, packagesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftVariant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftVariant(aircraftVariantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftVariant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftVariant)

		// slice the ids on comma with no spaces
		ids := strings.Split( packagesIds, ",")

		for _, packagesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftPackage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftPackage
			// with a matching packagesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , packagesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftPackageObj from the Packages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Packages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Packages", packagesId )
				return utils.RequestResult{false, msg, "removePackages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftVariant from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftVariant(aircraftVariantId)

	} else {
		return parentRequestResult
	}
}

