package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CreativeAssetDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCreativeAsset - creates a new db entry
//----------------------------------------------------------------------------
func CreateCreativeAsset(obj model.CreativeAsset)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CreativeAsset with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CreativeAsset", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCreativeAsset", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCreativeAsset - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCreativeAsset(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CreativeAsset

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CreativeAsset with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CreativeAsset using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CreativeAsset using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCreativeAsset", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCreativeAsset - returns all
//----------------------------------------------------------------------------
func GetAllCreativeAsset()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CreativeAsset

	//----------------------------------------------------------------------------
	// Request the ORM to find all CreativeAsset
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CreativeAsset" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CreativeAsset", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCreativeAsset", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCreativeAsset - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCreativeAsset(obj model.CreativeAsset)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CreativeAsset using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CreativeAsset using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCreativeAsset", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCreativeAsset - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCreativeAsset(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCreativeAsset(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CreativeAsset)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CreativeAsset using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CreativeAsset using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCreativeAsset", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more filesIds as a Files to a CreativeAsset
//----------------------------------------------------------------------------
func AddFilesToCreativeAsset ( creativeAssetId uint64, filesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( filesIds, ",")

		for _, filesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeFile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeFile
			// with a matching filesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , filesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Files using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Files").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Files", filesId )
				return utils.RequestResult{false, msg, "unassignFiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more filesIds as a Files from a CreativeAsset
//----------------------------------------------------------------------------
func RemoveFilesFromCreativeAsset( creativeAssetId uint64, filesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( filesIds, ",")

		for _, filesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeFile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeFile
			// with a matching filesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , filesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CreativeFileObj from the Files array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Files").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Files", filesId )
				return utils.RequestResult{false, msg, "removeFiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more approvalsIds as a Approvals to a CreativeAsset
//----------------------------------------------------------------------------
func AddApprovalsToCreativeAsset ( creativeAssetId uint64, approvalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( approvalsIds, ",")

		for _, approvalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeApproval

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeApproval
			// with a matching approvalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , approvalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Approvals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Approvals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Approvals", approvalsId )
				return utils.RequestResult{false, msg, "unassignApprovals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more approvalsIds as a Approvals from a CreativeAsset
//----------------------------------------------------------------------------
func RemoveApprovalsFromCreativeAsset( creativeAssetId uint64, approvalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( approvalsIds, ",")

		for _, approvalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeApproval

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeApproval
			// with a matching approvalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , approvalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CreativeApprovalObj from the Approvals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Approvals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Approvals", approvalsId )
				return utils.RequestResult{false, msg, "removeApprovals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more variationsIds as a Variations to a CreativeAsset
//----------------------------------------------------------------------------
func AddVariationsToCreativeAsset ( creativeAssetId uint64, variationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( variationsIds, ",")

		for _, variationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeVariation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeVariation
			// with a matching variationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Variations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variations", variationsId )
				return utils.RequestResult{false, msg, "unassignVariations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variationsIds as a Variations from a CreativeAsset
//----------------------------------------------------------------------------
func RemoveVariationsFromCreativeAsset( creativeAssetId uint64, variationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( variationsIds, ",")

		for _, variationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeVariation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeVariation
			// with a matching variationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CreativeVariationObj from the Variations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variations", variationsId )
				return utils.RequestResult{false, msg, "removeVariations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more lineItemsIds as a LineItems to a CreativeAsset
//----------------------------------------------------------------------------
func AddLineItemsToCreativeAsset ( creativeAssetId uint64, lineItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineItem
			// with a matching lineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LineItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItems", lineItemsId )
				return utils.RequestResult{false, msg, "unassignLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more lineItemsIds as a LineItems from a CreativeAsset
//----------------------------------------------------------------------------
func RemoveLineItemsFromCreativeAsset( creativeAssetId uint64, lineItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CreativeAsset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCreativeAsset(creativeAssetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CreativeAsset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CreativeAsset)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineItemsIds, ",")

		for _, lineItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineItem
			// with a matching lineItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LineItemObj from the LineItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineItems", lineItemsId )
				return utils.RequestResult{false, msg, "removeLineItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CreativeAsset from the gorm
		//----------------------------------------------------------------------------
		return GetCreativeAsset(creativeAssetId)

	} else {
		return parentRequestResult
	}
}

