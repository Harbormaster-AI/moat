package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InboundShipmentLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInboundShipmentLine - creates a new db entry
//----------------------------------------------------------------------------
func CreateInboundShipmentLine(obj model.InboundShipmentLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InboundShipmentLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InboundShipmentLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInboundShipmentLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInboundShipmentLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInboundShipmentLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InboundShipmentLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InboundShipmentLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InboundShipmentLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InboundShipmentLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInboundShipmentLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInboundShipmentLine - returns all
//----------------------------------------------------------------------------
func GetAllInboundShipmentLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InboundShipmentLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all InboundShipmentLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InboundShipmentLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InboundShipmentLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInboundShipmentLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInboundShipmentLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInboundShipmentLine(obj model.InboundShipmentLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InboundShipmentLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InboundShipmentLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInboundShipmentLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInboundShipmentLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInboundShipmentLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInboundShipmentLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InboundShipmentLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InboundShipmentLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInboundShipmentLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a InboundShipment on a InboundShipmentLine
//----------------------------------------------------------------------------
func AssignInboundShipmentToInboundShipmentLine( inboundShipmentLineId uint64, inboundShipmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InboundShipment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InboundShipment with a
		// matching inboundShipmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inboundShipmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InboundShipment	to the InboundShipmentLine
			//----------------------------------------------------------------------------
			parentObj.InboundShipment = &childObj

			//----------------------------------------------------------------------------
			// save the InboundShipmentLine
			//----------------------------------------------------------------------------
			return UpdateInboundShipmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InboundShipment", inboundShipmentId )
			return utils.RequestResult{false, msg, "assignInboundShipment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InboundShipment on a InboundShipmentLine
//----------------------------------------------------------------------------
func UnassignInboundShipmentFromInboundShipmentLine(inboundShipmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// assign an empty InboundShipment to the InboundShipment
		//----------------------------------------------------------------------------
		parentObj.InboundShipment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InboundShipment
		//----------------------------------------------------------------------------
		parentObj.InboundShipmentId = nil;

		//----------------------------------------------------------------------------
		// save the InboundShipmentLine
		//----------------------------------------------------------------------------
		return UpdateInboundShipmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Sku on a InboundShipmentLine
//----------------------------------------------------------------------------
func AssignSkuToInboundShipmentLine( inboundShipmentLineId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

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
			// assign the Sku	to the InboundShipmentLine
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the InboundShipmentLine
			//----------------------------------------------------------------------------
			return UpdateInboundShipmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a InboundShipmentLine
//----------------------------------------------------------------------------
func UnassignSkuFromInboundShipmentLine(inboundShipmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the InboundShipmentLine
		//----------------------------------------------------------------------------
		return UpdateInboundShipmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a InboundShipmentLine
//----------------------------------------------------------------------------
func AssignLotToInboundShipmentLine( inboundShipmentLineId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Lot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Lot with a
		// matching lotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Lot	to the InboundShipmentLine
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the InboundShipmentLine
			//----------------------------------------------------------------------------
			return UpdateInboundShipmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a InboundShipmentLine
//----------------------------------------------------------------------------
func UnassignLotFromInboundShipmentLine(inboundShipmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the InboundShipmentLine
		//----------------------------------------------------------------------------
		return UpdateInboundShipmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationLocation on a InboundShipmentLine
//----------------------------------------------------------------------------
func AssignDestinationLocationToInboundShipmentLine( inboundShipmentLineId uint64, destinationLocationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.StorageLocation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StorageLocation with a
		// matching destinationLocationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, destinationLocationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the DestinationLocation	to the InboundShipmentLine
			//----------------------------------------------------------------------------
			parentObj.DestinationLocation = &childObj

			//----------------------------------------------------------------------------
			// save the InboundShipmentLine
			//----------------------------------------------------------------------------
			return UpdateInboundShipmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationLocation", destinationLocationId )
			return utils.RequestResult{false, msg, "assignDestinationLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationLocation on a InboundShipmentLine
//----------------------------------------------------------------------------
func UnassignDestinationLocationFromInboundShipmentLine(inboundShipmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the DestinationLocation
		//----------------------------------------------------------------------------
		parentObj.DestinationLocation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationLocation
		//----------------------------------------------------------------------------
		parentObj.DestinationLocationId = nil;

		//----------------------------------------------------------------------------
		// save the InboundShipmentLine
		//----------------------------------------------------------------------------
		return UpdateInboundShipmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a InboundShipmentLine
//----------------------------------------------------------------------------
func AddSerialNumbersToInboundShipmentLine ( inboundShipmentLineId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		// slice the ids on comma with no spaces
		ids := strings.Split( serialNumbersIds, ",")

		for _, serialNumbersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SerialNumber

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SerialNumber
			// with a matching serialNumbersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serialNumbersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SerialNumbers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SerialNumbers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SerialNumbers", serialNumbersId )
				return utils.RequestResult{false, msg, "unassignSerialNumbers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InboundShipmentLine from the gorm
		//----------------------------------------------------------------------------
		return GetInboundShipmentLine(inboundShipmentLineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a InboundShipmentLine
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromInboundShipmentLine( inboundShipmentLineId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InboundShipmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipmentLine(inboundShipmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipmentLine)

		// slice the ids on comma with no spaces
		ids := strings.Split( serialNumbersIds, ",")

		for _, serialNumbersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SerialNumber

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SerialNumber
			// with a matching serialNumbersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serialNumbersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SerialNumberObj from the SerialNumbers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SerialNumbers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SerialNumbers", serialNumbersId )
				return utils.RequestResult{false, msg, "removeSerialNumbers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InboundShipmentLine from the gorm
		//----------------------------------------------------------------------------
		return GetInboundShipmentLine(inboundShipmentLineId)

	} else {
		return parentRequestResult
	}
}

