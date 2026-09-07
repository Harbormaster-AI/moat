package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BOMItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBOMItem - creates a new db entry
//----------------------------------------------------------------------------
func CreateBOMItem(obj model.BOMItem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BOMItem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BOMItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBOMItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBOMItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBOMItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BOMItem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BOMItem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BOMItem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BOMItem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBOMItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBOMItem - returns all
//----------------------------------------------------------------------------
func GetAllBOMItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BOMItem

	//----------------------------------------------------------------------------
	// Request the ORM to find all BOMItem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BOMItem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BOMItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBOMItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBOMItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBOMItem(obj model.BOMItem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BOMItem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BOMItem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBOMItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBOMItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBOMItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BOMItem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBOMItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOMItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BOMItem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BOMItem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BOMItem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBOMItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Bom on a BOMItem
//----------------------------------------------------------------------------
func AssignBomToBOMItem( bOMItemId uint64, bomId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BOMItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOMItem(bOMItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOMItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOMItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BOM

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BOM with a
		// matching bomId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, bomId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Bom	to the BOMItem
			//----------------------------------------------------------------------------
			parentObj.Bom = &childObj

			//----------------------------------------------------------------------------
			// save the BOMItem
			//----------------------------------------------------------------------------
			return UpdateBOMItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bom", bomId )
			return utils.RequestResult{false, msg, "assignBom", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bom on a BOMItem
//----------------------------------------------------------------------------
func UnassignBomFromBOMItem(bOMItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BOMItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOMItem(bOMItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOMItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOMItem)

		//----------------------------------------------------------------------------
		// assign an empty BOM to the Bom
		//----------------------------------------------------------------------------
		parentObj.Bom = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bom
		//----------------------------------------------------------------------------
		parentObj.BomId = nil;

		//----------------------------------------------------------------------------
		// save the BOMItem
		//----------------------------------------------------------------------------
		return UpdateBOMItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Component on a BOMItem
//----------------------------------------------------------------------------
func AssignComponentToBOMItem( bOMItemId uint64, componentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BOMItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOMItem(bOMItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOMItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOMItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching componentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, componentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Component	to the BOMItem
			//----------------------------------------------------------------------------
			parentObj.Component = &childObj

			//----------------------------------------------------------------------------
			// save the BOMItem
			//----------------------------------------------------------------------------
			return UpdateBOMItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Component", componentId )
			return utils.RequestResult{false, msg, "assignComponent", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Component on a BOMItem
//----------------------------------------------------------------------------
func UnassignComponentFromBOMItem(bOMItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BOMItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOMItem(bOMItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOMItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOMItem)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Component
		//----------------------------------------------------------------------------
		parentObj.Component = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Component
		//----------------------------------------------------------------------------
		parentObj.ComponentId = nil;

		//----------------------------------------------------------------------------
		// save the BOMItem
		//----------------------------------------------------------------------------
		return UpdateBOMItem(parentObj)

	} else {
		return parentRequestResult
	}

}


