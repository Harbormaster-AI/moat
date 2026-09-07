package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EngineTypeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEngineType - creates a new db entry
//----------------------------------------------------------------------------
func CreateEngineType(obj model.EngineType)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a EngineType with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a EngineType", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEngineType", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEngineType - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEngineType(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.EngineType

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a EngineType with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a EngineType using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a EngineType using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEngineType", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEngineType - returns all
//----------------------------------------------------------------------------
func GetAllEngineType()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.EngineType

	//----------------------------------------------------------------------------
	// Request the ORM to find all EngineType
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all EngineType" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all EngineType", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEngineType", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEngineType - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEngineType(obj model.EngineType)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a EngineType using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a EngineType using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEngineType", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEngineType - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEngineType(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the EngineType with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEngineType(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EngineType so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.EngineType)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a EngineType using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a EngineType using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEngineType", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Supplier on a EngineType
//----------------------------------------------------------------------------
func AssignSupplierToEngineType( engineTypeId uint64, supplierId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EngineType with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEngineType(engineTypeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EngineType so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EngineType)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Supplier

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Supplier with a
		// matching supplierId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, supplierId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Supplier	to the EngineType
			//----------------------------------------------------------------------------
			parentObj.Supplier = &childObj

			//----------------------------------------------------------------------------
			// save the EngineType
			//----------------------------------------------------------------------------
			return UpdateEngineType(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Supplier", supplierId )
			return utils.RequestResult{false, msg, "assignSupplier", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Supplier on a EngineType
//----------------------------------------------------------------------------
func UnassignSupplierFromEngineType(engineTypeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EngineType with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEngineType(engineTypeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EngineType so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EngineType)

		//----------------------------------------------------------------------------
		// assign an empty Supplier to the Supplier
		//----------------------------------------------------------------------------
		parentObj.Supplier = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Supplier
		//----------------------------------------------------------------------------
		parentObj.SupplierId = nil;

		//----------------------------------------------------------------------------
		// save the EngineType
		//----------------------------------------------------------------------------
		return UpdateEngineType(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more compatibleModelsIds as a CompatibleModels to a EngineType
//----------------------------------------------------------------------------
func AddCompatibleModelsToEngineType ( engineTypeId uint64, compatibleModelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EngineType with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEngineType(engineTypeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EngineType so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EngineType)

		// slice the ids on comma with no spaces
		ids := strings.Split( compatibleModelsIds, ",")

		for _, compatibleModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftModel
			// with a matching compatibleModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , compatibleModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CompatibleModels using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompatibleModels").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompatibleModels", compatibleModelsId )
				return utils.RequestResult{false, msg, "unassignCompatibleModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified EngineType from the gorm
		//----------------------------------------------------------------------------
		return GetEngineType(engineTypeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more compatibleModelsIds as a CompatibleModels from a EngineType
//----------------------------------------------------------------------------
func RemoveCompatibleModelsFromEngineType( engineTypeId uint64, compatibleModelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the EngineType with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEngineType(engineTypeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EngineType so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EngineType)

		// slice the ids on comma with no spaces
		ids := strings.Split( compatibleModelsIds, ",")

		for _, compatibleModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftModel
			// with a matching compatibleModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , compatibleModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftModelObj from the CompatibleModels array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompatibleModels").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompatibleModels", compatibleModelsId )
				return utils.RequestResult{false, msg, "removeCompatibleModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified EngineType from the gorm
		//----------------------------------------------------------------------------
		return GetEngineType(engineTypeId)

	} else {
		return parentRequestResult
	}
}

