package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProductionLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProductionLine - creates a new db entry
//----------------------------------------------------------------------------
func CreateProductionLine(obj model.ProductionLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ProductionLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ProductionLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProductionLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProductionLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProductionLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ProductionLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ProductionLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ProductionLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ProductionLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProductionLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProductionLine - returns all
//----------------------------------------------------------------------------
func GetAllProductionLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ProductionLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all ProductionLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ProductionLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ProductionLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProductionLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProductionLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProductionLine(obj model.ProductionLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ProductionLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ProductionLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProductionLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProductionLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProductionLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ProductionLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProductionLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ProductionLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ProductionLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ProductionLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProductionLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Plant on a ProductionLine
//----------------------------------------------------------------------------
func AssignPlantToProductionLine( productionLineId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProductionLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionLine(productionLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionLine)

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
			// assign the Plant	to the ProductionLine
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the ProductionLine
			//----------------------------------------------------------------------------
			return UpdateProductionLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a ProductionLine
//----------------------------------------------------------------------------
func UnassignPlantFromProductionLine(productionLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionLine(productionLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionLine)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the ProductionLine
		//----------------------------------------------------------------------------
		return UpdateProductionLine(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more workCentersIds as a WorkCenters to a ProductionLine
//----------------------------------------------------------------------------
func AddWorkCentersToProductionLine ( productionLineId uint64, workCentersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionLine(productionLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionLine)

		// slice the ids on comma with no spaces
		ids := strings.Split( workCentersIds, ",")

		for _, workCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkCenter
			// with a matching workCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the WorkCenters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkCenters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenters", workCentersId )
				return utils.RequestResult{false, msg, "unassignWorkCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ProductionLine from the gorm
		//----------------------------------------------------------------------------
		return GetProductionLine(productionLineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workCentersIds as a WorkCenters from a ProductionLine
//----------------------------------------------------------------------------
func RemoveWorkCentersFromProductionLine( productionLineId uint64, workCentersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ProductionLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionLine(productionLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionLine)

		// slice the ids on comma with no spaces
		ids := strings.Split( workCentersIds, ",")

		for _, workCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkCenter
			// with a matching workCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WorkCenterObj from the WorkCenters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkCenters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenters", workCentersId )
				return utils.RequestResult{false, msg, "removeWorkCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ProductionLine from the gorm
		//----------------------------------------------------------------------------
		return GetProductionLine(productionLineId)

	} else {
		return parentRequestResult
	}
}

