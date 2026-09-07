package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AerospaceManufacturerDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAerospaceManufacturer - creates a new db entry
//----------------------------------------------------------------------------
func CreateAerospaceManufacturer(obj model.AerospaceManufacturer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AerospaceManufacturer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AerospaceManufacturer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAerospaceManufacturer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAerospaceManufacturer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAerospaceManufacturer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AerospaceManufacturer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AerospaceManufacturer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AerospaceManufacturer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AerospaceManufacturer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAerospaceManufacturer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAerospaceManufacturer - returns all
//----------------------------------------------------------------------------
func GetAllAerospaceManufacturer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AerospaceManufacturer

	//----------------------------------------------------------------------------
	// Request the ORM to find all AerospaceManufacturer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AerospaceManufacturer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AerospaceManufacturer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAerospaceManufacturer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAerospaceManufacturer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAerospaceManufacturer(obj model.AerospaceManufacturer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AerospaceManufacturer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AerospaceManufacturer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAerospaceManufacturer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAerospaceManufacturer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAerospaceManufacturer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAerospaceManufacturer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AerospaceManufacturer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AerospaceManufacturer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AerospaceManufacturer using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAerospaceManufacturer", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more programsIds as a Programs to a AerospaceManufacturer
//----------------------------------------------------------------------------
func AddProgramsToAerospaceManufacturer ( aerospaceManufacturerId uint64, programsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( programsIds, ",")

		for _, programsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftProgram

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftProgram
			// with a matching programsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , programsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Programs using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Programs").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Programs", programsId )
				return utils.RequestResult{false, msg, "unassignPrograms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more programsIds as a Programs from a AerospaceManufacturer
//----------------------------------------------------------------------------
func RemoveProgramsFromAerospaceManufacturer( aerospaceManufacturerId uint64, programsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( programsIds, ",")

		for _, programsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftProgram

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftProgram
			// with a matching programsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , programsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftProgramObj from the Programs array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Programs").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Programs", programsId )
				return utils.RequestResult{false, msg, "removePrograms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more plantsIds as a Plants to a AerospaceManufacturer
//----------------------------------------------------------------------------
func AddPlantsToAerospaceManufacturer ( aerospaceManufacturerId uint64, plantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( plantsIds, ",")

		for _, plantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Plant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Plant
			// with a matching plantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Plants using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Plants").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plants", plantsId )
				return utils.RequestResult{false, msg, "unassignPlants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more plantsIds as a Plants from a AerospaceManufacturer
//----------------------------------------------------------------------------
func RemovePlantsFromAerospaceManufacturer( aerospaceManufacturerId uint64, plantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( plantsIds, ",")

		for _, plantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Plant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Plant
			// with a matching plantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PlantObj from the Plants array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Plants").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plants", plantsId )
				return utils.RequestResult{false, msg, "removePlants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more suppliersIds as a Suppliers to a AerospaceManufacturer
//----------------------------------------------------------------------------
func AddSuppliersToAerospaceManufacturer ( aerospaceManufacturerId uint64, suppliersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching suppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , suppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Suppliers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Suppliers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Suppliers", suppliersId )
				return utils.RequestResult{false, msg, "unassignSuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more suppliersIds as a Suppliers from a AerospaceManufacturer
//----------------------------------------------------------------------------
func RemoveSuppliersFromAerospaceManufacturer( aerospaceManufacturerId uint64, suppliersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching suppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , suppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SupplierObj from the Suppliers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Suppliers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Suppliers", suppliersId )
				return utils.RequestResult{false, msg, "removeSuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more productionCertificatesIds as a ProductionCertificates to a AerospaceManufacturer
//----------------------------------------------------------------------------
func AddProductionCertificatesToAerospaceManufacturer ( aerospaceManufacturerId uint64, productionCertificatesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionCertificatesIds, ",")

		for _, productionCertificatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionCertificate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionCertificate
			// with a matching productionCertificatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionCertificatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProductionCertificates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionCertificates").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionCertificates", productionCertificatesId )
				return utils.RequestResult{false, msg, "unassignProductionCertificates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productionCertificatesIds as a ProductionCertificates from a AerospaceManufacturer
//----------------------------------------------------------------------------
func RemoveProductionCertificatesFromAerospaceManufacturer( aerospaceManufacturerId uint64, productionCertificatesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AerospaceManufacturer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAerospaceManufacturer(aerospaceManufacturerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AerospaceManufacturer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AerospaceManufacturer)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionCertificatesIds, ",")

		for _, productionCertificatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionCertificate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionCertificate
			// with a matching productionCertificatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionCertificatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProductionCertificateObj from the ProductionCertificates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionCertificates").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionCertificates", productionCertificatesId )
				return utils.RequestResult{false, msg, "removeProductionCertificates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AerospaceManufacturer from the gorm
		//----------------------------------------------------------------------------
		return GetAerospaceManufacturer(aerospaceManufacturerId)

	} else {
		return parentRequestResult
	}
}

