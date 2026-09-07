package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProductionOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProductionOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateProductionOrder(obj model.ProductionOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ProductionOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ProductionOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProductionOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProductionOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProductionOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ProductionOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ProductionOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ProductionOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ProductionOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProductionOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProductionOrder - returns all
//----------------------------------------------------------------------------
func GetAllProductionOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ProductionOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all ProductionOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ProductionOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ProductionOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProductionOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProductionOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProductionOrder(obj model.ProductionOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ProductionOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ProductionOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProductionOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProductionOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProductionOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ProductionOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProductionOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ProductionOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ProductionOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ProductionOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProductionOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Variant on a ProductionOrder
//----------------------------------------------------------------------------
func AssignVariantToProductionOrder( productionOrderId uint64, variantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProductionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionOrder(productionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftVariant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftVariant with a
		// matching variantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, variantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Variant	to the ProductionOrder
			//----------------------------------------------------------------------------
			parentObj.Variant = &childObj

			//----------------------------------------------------------------------------
			// save the ProductionOrder
			//----------------------------------------------------------------------------
			return UpdateProductionOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variant", variantId )
			return utils.RequestResult{false, msg, "assignVariant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Variant on a ProductionOrder
//----------------------------------------------------------------------------
func UnassignVariantFromProductionOrder(productionOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionOrder(productionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionOrder)

		//----------------------------------------------------------------------------
		// assign an empty AircraftVariant to the Variant
		//----------------------------------------------------------------------------
		parentObj.Variant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Variant
		//----------------------------------------------------------------------------
		parentObj.VariantId = nil;

		//----------------------------------------------------------------------------
		// save the ProductionOrder
		//----------------------------------------------------------------------------
		return UpdateProductionOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Plant on a ProductionOrder
//----------------------------------------------------------------------------
func AssignPlantToProductionOrder( productionOrderId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProductionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionOrder(productionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Plant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Plant with a
		// matching plantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, plantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plant	to the ProductionOrder
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the ProductionOrder
			//----------------------------------------------------------------------------
			return UpdateProductionOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a ProductionOrder
//----------------------------------------------------------------------------
func UnassignPlantFromProductionOrder(productionOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionOrder(productionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionOrder)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the ProductionOrder
		//----------------------------------------------------------------------------
		return UpdateProductionOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AircraftOrder on a ProductionOrder
//----------------------------------------------------------------------------
func AssignAircraftOrderToProductionOrder( productionOrderId uint64, aircraftOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProductionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionOrder(productionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftOrder with a
		// matching aircraftOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, aircraftOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AircraftOrder	to the ProductionOrder
			//----------------------------------------------------------------------------
			parentObj.AircraftOrder = &childObj

			//----------------------------------------------------------------------------
			// save the ProductionOrder
			//----------------------------------------------------------------------------
			return UpdateProductionOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftOrder", aircraftOrderId )
			return utils.RequestResult{false, msg, "assignAircraftOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AircraftOrder on a ProductionOrder
//----------------------------------------------------------------------------
func UnassignAircraftOrderFromProductionOrder(productionOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionOrder(productionOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionOrder)

		//----------------------------------------------------------------------------
		// assign an empty AircraftOrder to the AircraftOrder
		//----------------------------------------------------------------------------
		parentObj.AircraftOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AircraftOrder
		//----------------------------------------------------------------------------
		parentObj.AircraftOrderId = nil;

		//----------------------------------------------------------------------------
		// save the ProductionOrder
		//----------------------------------------------------------------------------
		return UpdateProductionOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


