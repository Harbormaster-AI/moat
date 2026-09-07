package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AdSlotDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAdSlot - creates a new db entry
//----------------------------------------------------------------------------
func CreateAdSlot(obj model.AdSlot)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AdSlot with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AdSlot", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAdSlot", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAdSlot - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAdSlot(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AdSlot

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AdSlot with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AdSlot using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AdSlot using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAdSlot", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAdSlot - returns all
//----------------------------------------------------------------------------
func GetAllAdSlot()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AdSlot

	//----------------------------------------------------------------------------
	// Request the ORM to find all AdSlot
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AdSlot" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AdSlot", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAdSlot", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAdSlot - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAdSlot(obj model.AdSlot)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AdSlot using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AdSlot using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAdSlot", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAdSlot - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAdSlot(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AdSlot with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAdSlot(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdSlot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AdSlot)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AdSlot using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AdSlot using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAdSlot", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a InventorySource on a AdSlot
//----------------------------------------------------------------------------
func AssignInventorySourceToAdSlot( adSlotId uint64, inventorySourceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AdSlot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdSlot(adSlotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdSlot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdSlot)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InventorySource

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InventorySource with a
		// matching inventorySourceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inventorySourceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InventorySource	to the AdSlot
			//----------------------------------------------------------------------------
			parentObj.InventorySource = &childObj

			//----------------------------------------------------------------------------
			// save the AdSlot
			//----------------------------------------------------------------------------
			return UpdateAdSlot(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventorySource", inventorySourceId )
			return utils.RequestResult{false, msg, "assignInventorySource", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InventorySource on a AdSlot
//----------------------------------------------------------------------------
func UnassignInventorySourceFromAdSlot(adSlotId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdSlot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdSlot(adSlotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdSlot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdSlot)

		//----------------------------------------------------------------------------
		// assign an empty InventorySource to the InventorySource
		//----------------------------------------------------------------------------
		parentObj.InventorySource = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InventorySource
		//----------------------------------------------------------------------------
		parentObj.InventorySourceId = nil;

		//----------------------------------------------------------------------------
		// save the AdSlot
		//----------------------------------------------------------------------------
		return UpdateAdSlot(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more placementsIds as a Placements to a AdSlot
//----------------------------------------------------------------------------
func AddPlacementsToAdSlot ( adSlotId uint64, placementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdSlot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdSlot(adSlotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdSlot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdSlot)

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
		// retrieve the modified AdSlot from the gorm
		//----------------------------------------------------------------------------
		return GetAdSlot(adSlotId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more placementsIds as a Placements from a AdSlot
//----------------------------------------------------------------------------
func RemovePlacementsFromAdSlot( adSlotId uint64, placementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AdSlot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdSlot(adSlotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdSlot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdSlot)

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
		// retrieve the modified AdSlot from the gorm
		//----------------------------------------------------------------------------
		return GetAdSlot(adSlotId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ratesIds as a Rates to a AdSlot
//----------------------------------------------------------------------------
func AddRatesToAdSlot ( adSlotId uint64, ratesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AdSlot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdSlot(adSlotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdSlot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdSlot)

		// slice the ids on comma with no spaces
		ids := strings.Split( ratesIds, ",")

		for _, ratesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Rate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Rate
			// with a matching ratesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ratesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Rates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Rates").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Rates", ratesId )
				return utils.RequestResult{false, msg, "unassignRates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AdSlot from the gorm
		//----------------------------------------------------------------------------
		return GetAdSlot(adSlotId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ratesIds as a Rates from a AdSlot
//----------------------------------------------------------------------------
func RemoveRatesFromAdSlot( adSlotId uint64, ratesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AdSlot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAdSlot(adSlotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AdSlot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AdSlot)

		// slice the ids on comma with no spaces
		ids := strings.Split( ratesIds, ",")

		for _, ratesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Rate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Rate
			// with a matching ratesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ratesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RateObj from the Rates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Rates").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Rates", ratesId )
				return utils.RequestResult{false, msg, "removeRates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AdSlot from the gorm
		//----------------------------------------------------------------------------
		return GetAdSlot(adSlotId)

	} else {
		return parentRequestResult
	}
}

