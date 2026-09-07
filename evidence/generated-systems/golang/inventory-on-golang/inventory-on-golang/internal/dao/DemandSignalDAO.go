package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DemandSignalDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDemandSignal - creates a new db entry
//----------------------------------------------------------------------------
func CreateDemandSignal(obj model.DemandSignal)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DemandSignal with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DemandSignal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDemandSignal", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDemandSignal - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDemandSignal(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DemandSignal

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DemandSignal with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DemandSignal using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DemandSignal using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDemandSignal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDemandSignal - returns all
//----------------------------------------------------------------------------
func GetAllDemandSignal()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DemandSignal

	//----------------------------------------------------------------------------
	// Request the ORM to find all DemandSignal
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DemandSignal" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DemandSignal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDemandSignal", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDemandSignal - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDemandSignal(obj model.DemandSignal)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DemandSignal using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DemandSignal using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDemandSignal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDemandSignal - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDemandSignal(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DemandSignal with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDemandSignal(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DemandSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DemandSignal)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DemandSignal using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DemandSignal using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDemandSignal", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a DemandSignal
//----------------------------------------------------------------------------
func AssignSkuToDemandSignal( demandSignalId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DemandSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDemandSignal(demandSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DemandSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DemandSignal)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.StockKeepingUnit

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StockKeepingUnit with a
		// matching skuId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, skuId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Sku	to the DemandSignal
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the DemandSignal
			//----------------------------------------------------------------------------
			return UpdateDemandSignal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a DemandSignal
//----------------------------------------------------------------------------
func UnassignSkuFromDemandSignal(demandSignalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DemandSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDemandSignal(demandSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DemandSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DemandSignal)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the DemandSignal
		//----------------------------------------------------------------------------
		return UpdateDemandSignal(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more reservationsIds as a Reservations to a DemandSignal
//----------------------------------------------------------------------------
func AddReservationsToDemandSignal ( demandSignalId uint64, reservationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DemandSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDemandSignal(demandSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DemandSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DemandSignal)

		// slice the ids on comma with no spaces
		ids := strings.Split( reservationsIds, ",")

		for _, reservationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Reservation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Reservation
			// with a matching reservationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reservations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reservations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reservations", reservationsId )
				return utils.RequestResult{false, msg, "unassignReservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DemandSignal from the gorm
		//----------------------------------------------------------------------------
		return GetDemandSignal(demandSignalId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reservationsIds as a Reservations from a DemandSignal
//----------------------------------------------------------------------------
func RemoveReservationsFromDemandSignal( demandSignalId uint64, reservationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DemandSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDemandSignal(demandSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DemandSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DemandSignal)

		// slice the ids on comma with no spaces
		ids := strings.Split( reservationsIds, ",")

		for _, reservationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Reservation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Reservation
			// with a matching reservationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReservationObj from the Reservations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reservations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reservations", reservationsId )
				return utils.RequestResult{false, msg, "removeReservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DemandSignal from the gorm
		//----------------------------------------------------------------------------
		return GetDemandSignal(demandSignalId)

	} else {
		return parentRequestResult
	}
}

