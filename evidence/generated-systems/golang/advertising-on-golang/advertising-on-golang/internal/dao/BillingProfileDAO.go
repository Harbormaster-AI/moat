package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BillingProfileDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBillingProfile - creates a new db entry
//----------------------------------------------------------------------------
func CreateBillingProfile(obj model.BillingProfile)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BillingProfile with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BillingProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBillingProfile", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBillingProfile - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBillingProfile(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BillingProfile

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BillingProfile with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BillingProfile using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BillingProfile using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBillingProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBillingProfile - returns all
//----------------------------------------------------------------------------
func GetAllBillingProfile()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BillingProfile

	//----------------------------------------------------------------------------
	// Request the ORM to find all BillingProfile
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BillingProfile" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BillingProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBillingProfile", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBillingProfile - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBillingProfile(obj model.BillingProfile)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BillingProfile using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BillingProfile using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBillingProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBillingProfile - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBillingProfile(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BillingProfile with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBillingProfile(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BillingProfile)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BillingProfile using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BillingProfile using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBillingProfile", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Advertiser on a BillingProfile
//----------------------------------------------------------------------------
func AssignAdvertiserToBillingProfile( billingProfileId uint64, advertiserId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BillingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingProfile(billingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingProfile)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Advertiser

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Advertiser with a
		// matching advertiserId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, advertiserId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Advertiser	to the BillingProfile
			//----------------------------------------------------------------------------
			parentObj.Advertiser = &childObj

			//----------------------------------------------------------------------------
			// save the BillingProfile
			//----------------------------------------------------------------------------
			return UpdateBillingProfile(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Advertiser", advertiserId )
			return utils.RequestResult{false, msg, "assignAdvertiser", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Advertiser on a BillingProfile
//----------------------------------------------------------------------------
func UnassignAdvertiserFromBillingProfile(billingProfileId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BillingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingProfile(billingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingProfile)

		//----------------------------------------------------------------------------
		// assign an empty Advertiser to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.Advertiser = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Advertiser
		//----------------------------------------------------------------------------
		parentObj.AdvertiserId = nil;

		//----------------------------------------------------------------------------
		// save the BillingProfile
		//----------------------------------------------------------------------------
		return UpdateBillingProfile(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more paymentMethodsIds as a PaymentMethods to a BillingProfile
//----------------------------------------------------------------------------
func AddPaymentMethodsToBillingProfile ( billingProfileId uint64, paymentMethodsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BillingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingProfile(billingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentMethodsIds, ",")

		for _, paymentMethodsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentMethod

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentMethod
			// with a matching paymentMethodsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentMethodsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PaymentMethods using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentMethods").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentMethods", paymentMethodsId )
				return utils.RequestResult{false, msg, "unassignPaymentMethods", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetBillingProfile(billingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentMethodsIds as a PaymentMethods from a BillingProfile
//----------------------------------------------------------------------------
func RemovePaymentMethodsFromBillingProfile( billingProfileId uint64, paymentMethodsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BillingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingProfile(billingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentMethodsIds, ",")

		for _, paymentMethodsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentMethod

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentMethod
			// with a matching paymentMethodsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentMethodsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentMethodObj from the PaymentMethods array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentMethods").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentMethods", paymentMethodsId )
				return utils.RequestResult{false, msg, "removePaymentMethods", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetBillingProfile(billingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more adAccountsIds as a AdAccounts to a BillingProfile
//----------------------------------------------------------------------------
func AddAdAccountsToBillingProfile ( billingProfileId uint64, adAccountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BillingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingProfile(billingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( adAccountsIds, ",")

		for _, adAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdAccount
			// with a matching adAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AdAccounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdAccounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccounts", adAccountsId )
				return utils.RequestResult{false, msg, "unassignAdAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetBillingProfile(billingProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more adAccountsIds as a AdAccounts from a BillingProfile
//----------------------------------------------------------------------------
func RemoveAdAccountsFromBillingProfile( billingProfileId uint64, adAccountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BillingProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBillingProfile(billingProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BillingProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BillingProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( adAccountsIds, ",")

		for _, adAccountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AdAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AdAccount
			// with a matching adAccountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , adAccountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AdAccountObj from the AdAccounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AdAccounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdAccounts", adAccountsId )
				return utils.RequestResult{false, msg, "removeAdAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BillingProfile from the gorm
		//----------------------------------------------------------------------------
		return GetBillingProfile(billingProfileId)

	} else {
		return parentRequestResult
	}
}

