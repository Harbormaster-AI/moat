package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftOptionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraftOption - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraftOption(obj model.AircraftOption)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AircraftOption with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AircraftOption", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraftOption", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraftOption - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraftOption(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AircraftOption

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AircraftOption with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AircraftOption using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AircraftOption using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraftOption", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraftOption - returns all
//----------------------------------------------------------------------------
func GetAllAircraftOption()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AircraftOption

	//----------------------------------------------------------------------------
	// Request the ORM to find all AircraftOption
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AircraftOption" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AircraftOption", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraftOption", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraftOption - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraftOption(obj model.AircraftOption)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AircraftOption using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AircraftOption using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraftOption", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraftOption - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraftOption(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AircraftOption with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraftOption(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOption so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AircraftOption)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AircraftOption using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AircraftOption using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraftOption", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more variantsIds as a Variants to a AircraftOption
//----------------------------------------------------------------------------
func AddVariantsToAircraftOption ( aircraftOptionId uint64, variantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftOption with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOption(aircraftOptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOption so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOption)

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
		// retrieve the modified AircraftOption from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftOption(aircraftOptionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variantsIds as a Variants from a AircraftOption
//----------------------------------------------------------------------------
func RemoveVariantsFromAircraftOption( aircraftOptionId uint64, variantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftOption with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOption(aircraftOptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOption so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOption)

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
		// retrieve the modified AircraftOption from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftOption(aircraftOptionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more packagesIds as a Packages to a AircraftOption
//----------------------------------------------------------------------------
func AddPackagesToAircraftOption ( aircraftOptionId uint64, packagesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftOption with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOption(aircraftOptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOption so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOption)

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
		// retrieve the modified AircraftOption from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftOption(aircraftOptionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more packagesIds as a Packages from a AircraftOption
//----------------------------------------------------------------------------
func RemovePackagesFromAircraftOption( aircraftOptionId uint64, packagesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftOption with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftOption(aircraftOptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftOption so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftOption)

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
		// retrieve the modified AircraftOption from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftOption(aircraftOptionId)

	} else {
		return parentRequestResult
	}
}

