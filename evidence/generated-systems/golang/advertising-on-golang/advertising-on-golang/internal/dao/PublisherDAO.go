package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PublisherDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePublisher - creates a new db entry
//----------------------------------------------------------------------------
func CreatePublisher(obj model.Publisher)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Publisher with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Publisher", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePublisher", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPublisher - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPublisher(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Publisher

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Publisher with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Publisher using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Publisher using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPublisher", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPublisher - returns all
//----------------------------------------------------------------------------
func GetAllPublisher()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Publisher

	//----------------------------------------------------------------------------
	// Request the ORM to find all Publisher
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Publisher" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Publisher", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPublisher", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePublisher - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePublisher(obj model.Publisher)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Publisher using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Publisher using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePublisher", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePublisher - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePublisher(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPublisher(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Publisher)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Publisher using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Publisher using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePublisher", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more inventorySourcesIds as a InventorySources to a Publisher
//----------------------------------------------------------------------------
func AddInventorySourcesToPublisher ( publisherId uint64, inventorySourcesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

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
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventorySourcesIds as a InventorySources from a Publisher
//----------------------------------------------------------------------------
func RemoveInventorySourcesFromPublisher( publisherId uint64, inventorySourcesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

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
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dealsIds as a Deals to a Publisher
//----------------------------------------------------------------------------
func AddDealsToPublisher ( publisherId uint64, dealsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( dealsIds, ",")

		for _, dealsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Deal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Deal
			// with a matching dealsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dealsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Deals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Deals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deals", dealsId )
				return utils.RequestResult{false, msg, "unassignDeals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dealsIds as a Deals from a Publisher
//----------------------------------------------------------------------------
func RemoveDealsFromPublisher( publisherId uint64, dealsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( dealsIds, ",")

		for _, dealsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Deal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Deal
			// with a matching dealsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dealsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DealObj from the Deals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Deals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deals", dealsId )
				return utils.RequestResult{false, msg, "removeDeals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more creativeApprovalsIds as a CreativeApprovals to a Publisher
//----------------------------------------------------------------------------
func AddCreativeApprovalsToPublisher ( publisherId uint64, creativeApprovalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( creativeApprovalsIds, ",")

		for _, creativeApprovalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeApproval

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeApproval
			// with a matching creativeApprovalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , creativeApprovalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CreativeApprovals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CreativeApprovals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CreativeApprovals", creativeApprovalsId )
				return utils.RequestResult{false, msg, "unassignCreativeApprovals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more creativeApprovalsIds as a CreativeApprovals from a Publisher
//----------------------------------------------------------------------------
func RemoveCreativeApprovalsFromPublisher( publisherId uint64, creativeApprovalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( creativeApprovalsIds, ",")

		for _, creativeApprovalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CreativeApproval

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CreativeApproval
			// with a matching creativeApprovalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , creativeApprovalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CreativeApprovalObj from the CreativeApprovals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CreativeApprovals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CreativeApprovals", creativeApprovalsId )
				return utils.RequestResult{false, msg, "removeCreativeApprovals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more insertionOrdersIds as a InsertionOrders to a Publisher
//----------------------------------------------------------------------------
func AddInsertionOrdersToPublisher ( publisherId uint64, insertionOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( insertionOrdersIds, ",")

		for _, insertionOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsertionOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsertionOrder
			// with a matching insertionOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insertionOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InsertionOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsertionOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsertionOrders", insertionOrdersId )
				return utils.RequestResult{false, msg, "unassignInsertionOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more insertionOrdersIds as a InsertionOrders from a Publisher
//----------------------------------------------------------------------------
func RemoveInsertionOrdersFromPublisher( publisherId uint64, insertionOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( insertionOrdersIds, ",")

		for _, insertionOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsertionOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsertionOrder
			// with a matching insertionOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insertionOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsertionOrderObj from the InsertionOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsertionOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsertionOrders", insertionOrdersId )
				return utils.RequestResult{false, msg, "removeInsertionOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more rateCardsIds as a RateCards to a Publisher
//----------------------------------------------------------------------------
func AddRateCardsToPublisher ( publisherId uint64, rateCardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( rateCardsIds, ",")

		for _, rateCardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RateCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RateCard
			// with a matching rateCardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , rateCardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RateCards using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RateCards").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RateCards", rateCardsId )
				return utils.RequestResult{false, msg, "unassignRateCards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more rateCardsIds as a RateCards from a Publisher
//----------------------------------------------------------------------------
func RemoveRateCardsFromPublisher( publisherId uint64, rateCardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Publisher with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPublisher(publisherId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Publisher so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Publisher)

		// slice the ids on comma with no spaces
		ids := strings.Split( rateCardsIds, ",")

		for _, rateCardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RateCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RateCard
			// with a matching rateCardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , rateCardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RateCardObj from the RateCards array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RateCards").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RateCards", rateCardsId )
				return utils.RequestResult{false, msg, "removeRateCards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Publisher from the gorm
		//----------------------------------------------------------------------------
		return GetPublisher(publisherId)

	} else {
		return parentRequestResult
	}
}

