package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BusinessUnitDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBusinessUnit - creates a new db entry
//----------------------------------------------------------------------------
func CreateBusinessUnit(obj model.BusinessUnit)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BusinessUnit with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BusinessUnit", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBusinessUnit", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBusinessUnit - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBusinessUnit(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BusinessUnit

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BusinessUnit with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BusinessUnit using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BusinessUnit using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBusinessUnit", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBusinessUnit - returns all
//----------------------------------------------------------------------------
func GetAllBusinessUnit()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BusinessUnit

	//----------------------------------------------------------------------------
	// Request the ORM to find all BusinessUnit
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BusinessUnit" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BusinessUnit", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBusinessUnit", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBusinessUnit - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBusinessUnit(obj model.BusinessUnit)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BusinessUnit using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BusinessUnit using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBusinessUnit", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBusinessUnit - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBusinessUnit(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BusinessUnit with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBusinessUnit(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BusinessUnit)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BusinessUnit using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BusinessUnit using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBusinessUnit", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Enterprise on a BusinessUnit
//----------------------------------------------------------------------------
func AssignEnterpriseToBusinessUnit( businessUnitId uint64, enterpriseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BusinessUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessUnit(businessUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessUnit)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Enterprise

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Enterprise with a
		// matching enterpriseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, enterpriseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Enterprise	to the BusinessUnit
			//----------------------------------------------------------------------------
			parentObj.Enterprise = &childObj

			//----------------------------------------------------------------------------
			// save the BusinessUnit
			//----------------------------------------------------------------------------
			return UpdateBusinessUnit(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enterprise", enterpriseId )
			return utils.RequestResult{false, msg, "assignEnterprise", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Enterprise on a BusinessUnit
//----------------------------------------------------------------------------
func UnassignEnterpriseFromBusinessUnit(businessUnitId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessUnit(businessUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessUnit)

		//----------------------------------------------------------------------------
		// assign an empty Enterprise to the Enterprise
		//----------------------------------------------------------------------------
		parentObj.Enterprise = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Enterprise
		//----------------------------------------------------------------------------
		parentObj.EnterpriseId = nil;

		//----------------------------------------------------------------------------
		// save the BusinessUnit
		//----------------------------------------------------------------------------
		return UpdateBusinessUnit(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more itemsIds as a Items to a BusinessUnit
//----------------------------------------------------------------------------
func AddItemsToBusinessUnit ( businessUnitId uint64, itemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessUnit(businessUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Item

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Item
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Items using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "unassignItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BusinessUnit from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessUnit(businessUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more itemsIds as a Items from a BusinessUnit
//----------------------------------------------------------------------------
func RemoveItemsFromBusinessUnit( businessUnitId uint64, itemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BusinessUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessUnit(businessUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Item

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Item
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ItemObj from the Items array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "removeItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BusinessUnit from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessUnit(businessUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more plantsIds as a Plants to a BusinessUnit
//----------------------------------------------------------------------------
func AddPlantsToBusinessUnit ( businessUnitId uint64, plantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BusinessUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessUnit(businessUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessUnit)

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
		// retrieve the modified BusinessUnit from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessUnit(businessUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more plantsIds as a Plants from a BusinessUnit
//----------------------------------------------------------------------------
func RemovePlantsFromBusinessUnit( businessUnitId uint64, plantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BusinessUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBusinessUnit(businessUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BusinessUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BusinessUnit)

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
		// retrieve the modified BusinessUnit from the gorm
		//----------------------------------------------------------------------------
		return GetBusinessUnit(businessUnitId)

	} else {
		return parentRequestResult
	}
}

