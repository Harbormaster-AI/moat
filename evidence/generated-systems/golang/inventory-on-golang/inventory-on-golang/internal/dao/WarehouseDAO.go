package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WarehouseDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWarehouse - creates a new db entry
//----------------------------------------------------------------------------
func CreateWarehouse(obj model.Warehouse)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Warehouse with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Warehouse", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWarehouse", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWarehouse - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWarehouse(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Warehouse

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Warehouse with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Warehouse using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Warehouse using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWarehouse", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWarehouse - returns all
//----------------------------------------------------------------------------
func GetAllWarehouse()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Warehouse

	//----------------------------------------------------------------------------
	// Request the ORM to find all Warehouse
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Warehouse" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Warehouse", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWarehouse", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWarehouse - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWarehouse(obj model.Warehouse)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Warehouse using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Warehouse using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWarehouse", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWarehouse - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWarehouse(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWarehouse(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Warehouse)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Warehouse using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Warehouse using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWarehouse", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more storageLocationsIds as a StorageLocations to a Warehouse
//----------------------------------------------------------------------------
func AddStorageLocationsToWarehouse ( warehouseId uint64, storageLocationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( storageLocationsIds, ",")

		for _, storageLocationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StorageLocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StorageLocation
			// with a matching storageLocationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , storageLocationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the StorageLocations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("StorageLocations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "StorageLocations", storageLocationsId )
				return utils.RequestResult{false, msg, "unassignStorageLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more storageLocationsIds as a StorageLocations from a Warehouse
//----------------------------------------------------------------------------
func RemoveStorageLocationsFromWarehouse( warehouseId uint64, storageLocationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( storageLocationsIds, ",")

		for _, storageLocationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StorageLocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StorageLocation
			// with a matching storageLocationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , storageLocationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove StorageLocationObj from the StorageLocations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("StorageLocations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "StorageLocations", storageLocationsId )
				return utils.RequestResult{false, msg, "removeStorageLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a Warehouse
//----------------------------------------------------------------------------
func AddInventoryItemsToWarehouse ( warehouseId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InventoryItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "unassignInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a Warehouse
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromWarehouse( warehouseId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InventoryItemObj from the InventoryItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "removeInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more inboundShipmentsIds as a InboundShipments to a Warehouse
//----------------------------------------------------------------------------
func AddInboundShipmentsToWarehouse ( warehouseId uint64, inboundShipmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( inboundShipmentsIds, ",")

		for _, inboundShipmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InboundShipment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InboundShipment
			// with a matching inboundShipmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inboundShipmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InboundShipments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InboundShipments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InboundShipments", inboundShipmentsId )
				return utils.RequestResult{false, msg, "unassignInboundShipments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inboundShipmentsIds as a InboundShipments from a Warehouse
//----------------------------------------------------------------------------
func RemoveInboundShipmentsFromWarehouse( warehouseId uint64, inboundShipmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( inboundShipmentsIds, ",")

		for _, inboundShipmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InboundShipment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InboundShipment
			// with a matching inboundShipmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inboundShipmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InboundShipmentObj from the InboundShipments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InboundShipments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InboundShipments", inboundShipmentsId )
				return utils.RequestResult{false, msg, "removeInboundShipments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more outboundAllocationsIds as a OutboundAllocations to a Warehouse
//----------------------------------------------------------------------------
func AddOutboundAllocationsToWarehouse ( warehouseId uint64, outboundAllocationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( outboundAllocationsIds, ",")

		for _, outboundAllocationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OutboundAllocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OutboundAllocation
			// with a matching outboundAllocationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , outboundAllocationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OutboundAllocations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OutboundAllocations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OutboundAllocations", outboundAllocationsId )
				return utils.RequestResult{false, msg, "unassignOutboundAllocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more outboundAllocationsIds as a OutboundAllocations from a Warehouse
//----------------------------------------------------------------------------
func RemoveOutboundAllocationsFromWarehouse( warehouseId uint64, outboundAllocationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( outboundAllocationsIds, ",")

		for _, outboundAllocationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.OutboundAllocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a OutboundAllocation
			// with a matching outboundAllocationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , outboundAllocationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OutboundAllocationObj from the OutboundAllocations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OutboundAllocations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OutboundAllocations", outboundAllocationsId )
				return utils.RequestResult{false, msg, "removeOutboundAllocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more originTransfersIds as a OriginTransfers to a Warehouse
//----------------------------------------------------------------------------
func AddOriginTransfersToWarehouse ( warehouseId uint64, originTransfersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( originTransfersIds, ",")

		for _, originTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TransferOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TransferOrder
			// with a matching originTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , originTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OriginTransfers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OriginTransfers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OriginTransfers", originTransfersId )
				return utils.RequestResult{false, msg, "unassignOriginTransfers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more originTransfersIds as a OriginTransfers from a Warehouse
//----------------------------------------------------------------------------
func RemoveOriginTransfersFromWarehouse( warehouseId uint64, originTransfersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( originTransfersIds, ",")

		for _, originTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TransferOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TransferOrder
			// with a matching originTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , originTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransferOrderObj from the OriginTransfers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OriginTransfers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OriginTransfers", originTransfersId )
				return utils.RequestResult{false, msg, "removeOriginTransfers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more destinationTransfersIds as a DestinationTransfers to a Warehouse
//----------------------------------------------------------------------------
func AddDestinationTransfersToWarehouse ( warehouseId uint64, destinationTransfersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( destinationTransfersIds, ",")

		for _, destinationTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TransferOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TransferOrder
			// with a matching destinationTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , destinationTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DestinationTransfers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DestinationTransfers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationTransfers", destinationTransfersId )
				return utils.RequestResult{false, msg, "unassignDestinationTransfers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more destinationTransfersIds as a DestinationTransfers from a Warehouse
//----------------------------------------------------------------------------
func RemoveDestinationTransfersFromWarehouse( warehouseId uint64, destinationTransfersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( destinationTransfersIds, ",")

		for _, destinationTransfersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TransferOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TransferOrder
			// with a matching destinationTransfersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , destinationTransfersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransferOrderObj from the DestinationTransfers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DestinationTransfers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationTransfers", destinationTransfersId )
				return utils.RequestResult{false, msg, "removeDestinationTransfers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more cycleCountsIds as a CycleCounts to a Warehouse
//----------------------------------------------------------------------------
func AddCycleCountsToWarehouse ( warehouseId uint64, cycleCountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( cycleCountsIds, ",")

		for _, cycleCountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CycleCount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CycleCount
			// with a matching cycleCountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cycleCountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CycleCounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CycleCounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CycleCounts", cycleCountsId )
				return utils.RequestResult{false, msg, "unassignCycleCounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more cycleCountsIds as a CycleCounts from a Warehouse
//----------------------------------------------------------------------------
func RemoveCycleCountsFromWarehouse( warehouseId uint64, cycleCountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( cycleCountsIds, ",")

		for _, cycleCountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CycleCount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CycleCount
			// with a matching cycleCountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cycleCountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CycleCountObj from the CycleCounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CycleCounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CycleCounts", cycleCountsId )
				return utils.RequestResult{false, msg, "removeCycleCounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

