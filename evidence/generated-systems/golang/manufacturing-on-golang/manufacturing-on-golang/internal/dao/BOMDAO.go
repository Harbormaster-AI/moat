package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BOMDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBOM - creates a new db entry
//----------------------------------------------------------------------------
func CreateBOM(obj model.BOM)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BOM with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BOM", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBOM", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBOM - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBOM(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BOM

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BOM with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BOM using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BOM using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBOM", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBOM - returns all
//----------------------------------------------------------------------------
func GetAllBOM()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BOM

	//----------------------------------------------------------------------------
	// Request the ORM to find all BOM
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BOM" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BOM", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBOM", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBOM - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBOM(obj model.BOM)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BOM using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BOM using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBOM", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBOM - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBOM(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BOM with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBOM(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BOM)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BOM using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BOM using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBOM", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ParentItem on a BOM
//----------------------------------------------------------------------------
func AssignParentItemToBOM( bOMId uint64, parentItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BOM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOM(bOMId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOM)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching parentItemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, parentItemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ParentItem	to the BOM
			//----------------------------------------------------------------------------
			parentObj.ParentItem = &childObj

			//----------------------------------------------------------------------------
			// save the BOM
			//----------------------------------------------------------------------------
			return UpdateBOM(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ParentItem", parentItemId )
			return utils.RequestResult{false, msg, "assignParentItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ParentItem on a BOM
//----------------------------------------------------------------------------
func UnassignParentItemFromBOM(bOMId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BOM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOM(bOMId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOM)

		//----------------------------------------------------------------------------
		// assign an empty Item to the ParentItem
		//----------------------------------------------------------------------------
		parentObj.ParentItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ParentItem
		//----------------------------------------------------------------------------
		parentObj.ParentItemId = nil;

		//----------------------------------------------------------------------------
		// save the BOM
		//----------------------------------------------------------------------------
		return UpdateBOM(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more bomItemsIds as a BomItems to a BOM
//----------------------------------------------------------------------------
func AddBomItemsToBOM ( bOMId uint64, bomItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BOM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOM(bOMId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOM)

		// slice the ids on comma with no spaces
		ids := strings.Split( bomItemsIds, ",")

		for _, bomItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BOMItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BOMItem
			// with a matching bomItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , bomItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the BomItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BomItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BomItems", bomItemsId )
				return utils.RequestResult{false, msg, "unassignBomItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BOM from the gorm
		//----------------------------------------------------------------------------
		return GetBOM(bOMId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more bomItemsIds as a BomItems from a BOM
//----------------------------------------------------------------------------
func RemoveBomItemsFromBOM( bOMId uint64, bomItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BOM with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBOM(bOMId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BOM so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BOM)

		// slice the ids on comma with no spaces
		ids := strings.Split( bomItemsIds, ",")

		for _, bomItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BOMItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BOMItem
			// with a matching bomItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , bomItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BOMItemObj from the BomItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BomItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BomItems", bomItemsId )
				return utils.RequestResult{false, msg, "removeBomItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BOM from the gorm
		//----------------------------------------------------------------------------
		return GetBOM(bOMId)

	} else {
		return parentRequestResult
	}
}

