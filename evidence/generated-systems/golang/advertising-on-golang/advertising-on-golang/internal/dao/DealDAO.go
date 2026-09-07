package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DealDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDeal - creates a new db entry
//----------------------------------------------------------------------------
func CreateDeal(obj model.Deal)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Deal with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Deal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDeal", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDeal - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDeal(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Deal

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Deal with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Deal using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Deal using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDeal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDeal - returns all
//----------------------------------------------------------------------------
func GetAllDeal()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Deal

	//----------------------------------------------------------------------------
	// Request the ORM to find all Deal
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Deal" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Deal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDeal", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDeal - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDeal(obj model.Deal)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Deal using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Deal using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDeal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDeal - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDeal(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Deal with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDeal(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Deal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Deal)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Deal using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Deal using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDeal", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Publisher on a Deal
//----------------------------------------------------------------------------
func AssignPublisherToDeal( dealId uint64, publisherId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Deal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeal(dealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Deal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Deal)

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
			// assign the Publisher	to the Deal
			//----------------------------------------------------------------------------
			parentObj.Publisher = &childObj

			//----------------------------------------------------------------------------
			// save the Deal
			//----------------------------------------------------------------------------
			return UpdateDeal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Publisher", publisherId )
			return utils.RequestResult{false, msg, "assignPublisher", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Publisher on a Deal
//----------------------------------------------------------------------------
func UnassignPublisherFromDeal(dealId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Deal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeal(dealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Deal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Deal)

		//----------------------------------------------------------------------------
		// assign an empty Publisher to the Publisher
		//----------------------------------------------------------------------------
		parentObj.Publisher = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Publisher
		//----------------------------------------------------------------------------
		parentObj.PublisherId = nil;

		//----------------------------------------------------------------------------
		// save the Deal
		//----------------------------------------------------------------------------
		return UpdateDeal(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more inventorySourcesIds as a InventorySources to a Deal
//----------------------------------------------------------------------------
func AddInventorySourcesToDeal ( dealId uint64, inventorySourcesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Deal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeal(dealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Deal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Deal)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventorySourcesIds, ",")

		for _, inventorySourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventorySource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventorySource
			// with a matching inventorySourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventorySourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InventorySources using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventorySources").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventorySources", inventorySourcesId )
				return utils.RequestResult{false, msg, "unassignInventorySources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Deal from the gorm
		//----------------------------------------------------------------------------
		return GetDeal(dealId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventorySourcesIds as a InventorySources from a Deal
//----------------------------------------------------------------------------
func RemoveInventorySourcesFromDeal( dealId uint64, inventorySourcesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Deal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeal(dealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Deal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Deal)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventorySourcesIds, ",")

		for _, inventorySourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventorySource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventorySource
			// with a matching inventorySourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventorySourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InventorySourceObj from the InventorySources array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventorySources").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventorySources", inventorySourcesId )
				return utils.RequestResult{false, msg, "removeInventorySources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Deal from the gorm
		//----------------------------------------------------------------------------
		return GetDeal(dealId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more placementsIds as a Placements to a Deal
//----------------------------------------------------------------------------
func AddPlacementsToDeal ( dealId uint64, placementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Deal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeal(dealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Deal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Deal)

		// slice the ids on comma with no spaces
		ids := strings.Split( placementsIds, ",")

		for _, placementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Placement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Placement
			// with a matching placementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , placementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Placements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Placements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Placements", placementsId )
				return utils.RequestResult{false, msg, "unassignPlacements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Deal from the gorm
		//----------------------------------------------------------------------------
		return GetDeal(dealId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more placementsIds as a Placements from a Deal
//----------------------------------------------------------------------------
func RemovePlacementsFromDeal( dealId uint64, placementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Deal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDeal(dealId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Deal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Deal)

		// slice the ids on comma with no spaces
		ids := strings.Split( placementsIds, ",")

		for _, placementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Placement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Placement
			// with a matching placementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , placementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PlacementObj from the Placements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Placements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Placements", placementsId )
				return utils.RequestResult{false, msg, "removePlacements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Deal from the gorm
		//----------------------------------------------------------------------------
		return GetDeal(dealId)

	} else {
		return parentRequestResult
	}
}

