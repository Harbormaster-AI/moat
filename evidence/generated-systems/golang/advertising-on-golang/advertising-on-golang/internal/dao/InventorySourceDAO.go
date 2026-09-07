package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InventorySourceDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInventorySource - creates a new db entry
//----------------------------------------------------------------------------
func CreateInventorySource(obj model.InventorySource)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InventorySource with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InventorySource", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInventorySource", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInventorySource - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInventorySource(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InventorySource

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InventorySource with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InventorySource using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InventorySource using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInventorySource", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInventorySource - returns all
//----------------------------------------------------------------------------
func GetAllInventorySource()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InventorySource

	//----------------------------------------------------------------------------
	// Request the ORM to find all InventorySource
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InventorySource" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InventorySource", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInventorySource", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInventorySource - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInventorySource(obj model.InventorySource)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InventorySource using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InventorySource using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInventorySource", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInventorySource - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInventorySource(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InventorySource with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInventorySource(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventorySource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InventorySource)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InventorySource using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InventorySource using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInventorySource", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Publisher on a InventorySource
//----------------------------------------------------------------------------
func AssignPublisherToInventorySource( inventorySourceId uint64, publisherId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventorySource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventorySource(inventorySourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventorySource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventorySource)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Publisher

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Publisher with a
		// matching publisherId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, publisherId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Publisher	to the InventorySource
			//----------------------------------------------------------------------------
			parentObj.Publisher = &childObj

			//----------------------------------------------------------------------------
			// save the InventorySource
			//----------------------------------------------------------------------------
			return UpdateInventorySource(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Publisher", publisherId )
			return utils.RequestResult{false, msg, "assignPublisher", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Publisher on a InventorySource
//----------------------------------------------------------------------------
func UnassignPublisherFromInventorySource(inventorySourceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventorySource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventorySource(inventorySourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventorySource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventorySource)

		//----------------------------------------------------------------------------
		// assign an empty Publisher to the Publisher
		//----------------------------------------------------------------------------
		parentObj.Publisher = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Publisher
		//----------------------------------------------------------------------------
		parentObj.PublisherId = nil;

		//----------------------------------------------------------------------------
		// save the InventorySource
		//----------------------------------------------------------------------------
		return UpdateInventorySource(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more adSlotsIds as a AdSlots to a InventorySource
//----------------------------------------------------------------------------
func AddAdSlotsToInventorySource ( inventorySourceId uint64, adSlotsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventorySource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventorySource(inventorySourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventorySource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventorySource)

		// slice the ids on comma with no spaces
		ids := strings.Split( adSlotsIds, ",")

		for _, adSlotsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdSlot

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdSlot
			// with a matching adSlotsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adSlotsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AdSlots using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdSlots").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdSlots", adSlotsId )
				return utils.RequestResult{false, msg, "unassignAdSlots", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventorySource from the gorm
		//----------------------------------------------------------------------------
		return GetInventorySource(inventorySourceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more adSlotsIds as a AdSlots from a InventorySource
//----------------------------------------------------------------------------
func RemoveAdSlotsFromInventorySource( inventorySourceId uint64, adSlotsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InventorySource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventorySource(inventorySourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventorySource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventorySource)

		// slice the ids on comma with no spaces
		ids := strings.Split( adSlotsIds, ",")

		for _, adSlotsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdSlot

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdSlot
			// with a matching adSlotsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adSlotsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AdSlotObj from the AdSlots array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdSlots").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdSlots", adSlotsId )
				return utils.RequestResult{false, msg, "removeAdSlots", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventorySource from the gorm
		//----------------------------------------------------------------------------
		return GetInventorySource(inventorySourceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dealsIds as a Deals to a InventorySource
//----------------------------------------------------------------------------
func AddDealsToInventorySource ( inventorySourceId uint64, dealsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventorySource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventorySource(inventorySourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventorySource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventorySource)

		// slice the ids on comma with no spaces
		ids := strings.Split( dealsIds, ",")

		for _, dealsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Deal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Deal
			// with a matching dealsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dealsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Deals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Deals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deals", dealsId )
				return utils.RequestResult{false, msg, "unassignDeals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventorySource from the gorm
		//----------------------------------------------------------------------------
		return GetInventorySource(inventorySourceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dealsIds as a Deals from a InventorySource
//----------------------------------------------------------------------------
func RemoveDealsFromInventorySource( inventorySourceId uint64, dealsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InventorySource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventorySource(inventorySourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventorySource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventorySource)

		// slice the ids on comma with no spaces
		ids := strings.Split( dealsIds, ",")

		for _, dealsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Deal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Deal
			// with a matching dealsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dealsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DealObj from the Deals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Deals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deals", dealsId )
				return utils.RequestResult{false, msg, "removeDeals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventorySource from the gorm
		//----------------------------------------------------------------------------
		return GetInventorySource(inventorySourceId)

	} else {
		return parentRequestResult
	}
}

