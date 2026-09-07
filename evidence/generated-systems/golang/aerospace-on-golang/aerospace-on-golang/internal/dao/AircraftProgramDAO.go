package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftProgramDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraftProgram - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraftProgram(obj model.AircraftProgram)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AircraftProgram with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AircraftProgram", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraftProgram", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraftProgram - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraftProgram(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AircraftProgram

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AircraftProgram with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AircraftProgram using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AircraftProgram using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraftProgram", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraftProgram - returns all
//----------------------------------------------------------------------------
func GetAllAircraftProgram()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AircraftProgram

	//----------------------------------------------------------------------------
	// Request the ORM to find all AircraftProgram
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AircraftProgram" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AircraftProgram", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraftProgram", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraftProgram - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraftProgram(obj model.AircraftProgram)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AircraftProgram using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AircraftProgram using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraftProgram", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraftProgram - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraftProgram(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraftProgram(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AircraftProgram)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AircraftProgram using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AircraftProgram using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraftProgram", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Manufacturer on a AircraftProgram
//----------------------------------------------------------------------------
func AssignManufacturerToAircraftProgram( aircraftProgramId uint64, manufacturerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AerospaceManufacturer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AerospaceManufacturer with a
		// matching manufacturerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, manufacturerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Manufacturer	to the AircraftProgram
			//----------------------------------------------------------------------------
			parentObj.Manufacturer = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftProgram
			//----------------------------------------------------------------------------
			return UpdateAircraftProgram(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Manufacturer", manufacturerId )
			return utils.RequestResult{false, msg, "assignManufacturer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Manufacturer on a AircraftProgram
//----------------------------------------------------------------------------
func UnassignManufacturerFromAircraftProgram(aircraftProgramId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		//----------------------------------------------------------------------------
		// assign an empty AerospaceManufacturer to the Manufacturer
		//----------------------------------------------------------------------------
		parentObj.Manufacturer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Manufacturer
		//----------------------------------------------------------------------------
		parentObj.ManufacturerId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftProgram
		//----------------------------------------------------------------------------
		return UpdateAircraftProgram(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a TypeCertificate on a AircraftProgram
//----------------------------------------------------------------------------
func AssignTypeCertificateToAircraftProgram( aircraftProgramId uint64, typeCertificateId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.TypeCertificate

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TypeCertificate with a
		// matching typeCertificateId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, typeCertificateId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the TypeCertificate	to the AircraftProgram
			//----------------------------------------------------------------------------
			parentObj.TypeCertificate = &childObj

			//----------------------------------------------------------------------------
			// save the AircraftProgram
			//----------------------------------------------------------------------------
			return UpdateAircraftProgram(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TypeCertificate", typeCertificateId )
			return utils.RequestResult{false, msg, "assignTypeCertificate", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TypeCertificate on a AircraftProgram
//----------------------------------------------------------------------------
func UnassignTypeCertificateFromAircraftProgram(aircraftProgramId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		//----------------------------------------------------------------------------
		// assign an empty TypeCertificate to the TypeCertificate
		//----------------------------------------------------------------------------
		parentObj.TypeCertificate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TypeCertificate
		//----------------------------------------------------------------------------
		parentObj.TypeCertificateId = nil;

		//----------------------------------------------------------------------------
		// save the AircraftProgram
		//----------------------------------------------------------------------------
		return UpdateAircraftProgram(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more aircraftFamiliesIds as a AircraftFamilies to a AircraftProgram
//----------------------------------------------------------------------------
func AddAircraftFamiliesToAircraftProgram ( aircraftProgramId uint64, aircraftFamiliesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftFamiliesIds, ",")

		for _, aircraftFamiliesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftFamily

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftFamily
			// with a matching aircraftFamiliesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftFamiliesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AircraftFamilies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AircraftFamilies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftFamilies", aircraftFamiliesId )
				return utils.RequestResult{false, msg, "unassignAircraftFamilies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftProgram from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftProgram(aircraftProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more aircraftFamiliesIds as a AircraftFamilies from a AircraftProgram
//----------------------------------------------------------------------------
func RemoveAircraftFamiliesFromAircraftProgram( aircraftProgramId uint64, aircraftFamiliesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftFamiliesIds, ",")

		for _, aircraftFamiliesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftFamily

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftFamily
			// with a matching aircraftFamiliesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftFamiliesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftFamilyObj from the AircraftFamilies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AircraftFamilies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftFamilies", aircraftFamiliesId )
				return utils.RequestResult{false, msg, "removeAircraftFamilies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftProgram from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftProgram(aircraftProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more keySuppliersIds as a KeySuppliers to a AircraftProgram
//----------------------------------------------------------------------------
func AddKeySuppliersToAircraftProgram ( aircraftProgramId uint64, keySuppliersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( keySuppliersIds, ",")

		for _, keySuppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching keySuppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , keySuppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the KeySuppliers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("KeySuppliers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KeySuppliers", keySuppliersId )
				return utils.RequestResult{false, msg, "unassignKeySuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftProgram from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftProgram(aircraftProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more keySuppliersIds as a KeySuppliers from a AircraftProgram
//----------------------------------------------------------------------------
func RemoveKeySuppliersFromAircraftProgram( aircraftProgramId uint64, keySuppliersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AircraftProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraftProgram(aircraftProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AircraftProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AircraftProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( keySuppliersIds, ",")

		for _, keySuppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching keySuppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , keySuppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SupplierObj from the KeySuppliers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("KeySuppliers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KeySuppliers", keySuppliersId )
				return utils.RequestResult{false, msg, "removeKeySuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AircraftProgram from the gorm
		//----------------------------------------------------------------------------
		return GetAircraftProgram(aircraftProgramId)

	} else {
		return parentRequestResult
	}
}

