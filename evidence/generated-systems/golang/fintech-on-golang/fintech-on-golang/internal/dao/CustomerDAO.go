package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CustomerDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCustomer - creates a new db entry
//----------------------------------------------------------------------------
func CreateCustomer(obj model.Customer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Customer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Customer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCustomer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCustomer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCustomer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Customer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Customer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Customer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Customer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCustomer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCustomer - returns all
//----------------------------------------------------------------------------
func GetAllCustomer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Customer

	//----------------------------------------------------------------------------
	// Request the ORM to find all Customer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Customer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Customer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCustomer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCustomer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCustomer(obj model.Customer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Customer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Customer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCustomer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCustomer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCustomer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCustomer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Customer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Customer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Customer using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCustomer", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Institution on a Customer
//----------------------------------------------------------------------------
func AssignInstitutionToCustomer( customerId uint64, institutionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FinancialInstitution

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FinancialInstitution with a
		// matching institutionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, institutionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Institution	to the Customer
			//----------------------------------------------------------------------------
			parentObj.Institution = &childObj

			//----------------------------------------------------------------------------
			// save the Customer
			//----------------------------------------------------------------------------
			return UpdateCustomer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Institution", institutionId )
			return utils.RequestResult{false, msg, "assignInstitution", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Institution on a Customer
//----------------------------------------------------------------------------
func UnassignInstitutionFromCustomer(customerId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		//----------------------------------------------------------------------------
		// assign an empty FinancialInstitution to the Institution
		//----------------------------------------------------------------------------
		parentObj.Institution = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Institution
		//----------------------------------------------------------------------------
		parentObj.InstitutionId = nil;

		//----------------------------------------------------------------------------
		// save the Customer
		//----------------------------------------------------------------------------
		return UpdateCustomer(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a Customer
//----------------------------------------------------------------------------
func AddAccountsToCustomer ( customerId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Accounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a Customer
//----------------------------------------------------------------------------
func RemoveAccountsFromCustomer( customerId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Account

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Account
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccountObj from the Accounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more walletsIds as a Wallets to a Customer
//----------------------------------------------------------------------------
func AddWalletsToCustomer ( customerId uint64, walletsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( walletsIds, ",")

		for _, walletsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Wallet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Wallet
			// with a matching walletsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , walletsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Wallets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Wallets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Wallets", walletsId )
				return utils.RequestResult{false, msg, "unassignWallets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more walletsIds as a Wallets from a Customer
//----------------------------------------------------------------------------
func RemoveWalletsFromCustomer( customerId uint64, walletsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( walletsIds, ",")

		for _, walletsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Wallet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Wallet
			// with a matching walletsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , walletsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WalletObj from the Wallets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Wallets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Wallets", walletsId )
				return utils.RequestResult{false, msg, "removeWallets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more cardsIds as a Cards to a Customer
//----------------------------------------------------------------------------
func AddCardsToCustomer ( customerId uint64, cardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( cardsIds, ",")

		for _, cardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentCard
			// with a matching cardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Cards using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Cards").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cards", cardsId )
				return utils.RequestResult{false, msg, "unassignCards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more cardsIds as a Cards from a Customer
//----------------------------------------------------------------------------
func RemoveCardsFromCustomer( customerId uint64, cardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( cardsIds, ",")

		for _, cardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentCard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentCard
			// with a matching cardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , cardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentCardObj from the Cards array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Cards").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Cards", cardsId )
				return utils.RequestResult{false, msg, "removeCards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more kycProfilesIds as a KycProfiles to a Customer
//----------------------------------------------------------------------------
func AddKycProfilesToCustomer ( customerId uint64, kycProfilesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( kycProfilesIds, ",")

		for _, kycProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KYCProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KYCProfile
			// with a matching kycProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , kycProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the KycProfiles using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("KycProfiles").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfiles", kycProfilesId )
				return utils.RequestResult{false, msg, "unassignKycProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more kycProfilesIds as a KycProfiles from a Customer
//----------------------------------------------------------------------------
func RemoveKycProfilesFromCustomer( customerId uint64, kycProfilesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( kycProfilesIds, ",")

		for _, kycProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.KYCProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a KYCProfile
			// with a matching kycProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , kycProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove KYCProfileObj from the KycProfiles array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("KycProfiles").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfiles", kycProfilesId )
				return utils.RequestResult{false, msg, "removeKycProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a Customer
//----------------------------------------------------------------------------
func AddConsentsToCustomer ( customerId uint64, consentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Consents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "unassignConsents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more consentsIds as a Consents from a Customer
//----------------------------------------------------------------------------
func RemoveConsentsFromCustomer( customerId uint64, consentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConsentObj from the Consents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "removeConsents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more agreementsIds as a Agreements to a Customer
//----------------------------------------------------------------------------
func AddAgreementsToCustomer ( customerId uint64, agreementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( agreementsIds, ",")

		for _, agreementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Agreement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Agreement
			// with a matching agreementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , agreementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Agreements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Agreements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agreements", agreementsId )
				return utils.RequestResult{false, msg, "unassignAgreements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more agreementsIds as a Agreements from a Customer
//----------------------------------------------------------------------------
func RemoveAgreementsFromCustomer( customerId uint64, agreementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( agreementsIds, ",")

		for _, agreementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Agreement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Agreement
			// with a matching agreementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , agreementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AgreementObj from the Agreements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Agreements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agreements", agreementsId )
				return utils.RequestResult{false, msg, "removeAgreements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more loanApplicationsIds as a LoanApplications to a Customer
//----------------------------------------------------------------------------
func AddLoanApplicationsToCustomer ( customerId uint64, loanApplicationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( loanApplicationsIds, ",")

		for _, loanApplicationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanApplication

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanApplication
			// with a matching loanApplicationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , loanApplicationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LoanApplications using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LoanApplications").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanApplications", loanApplicationsId )
				return utils.RequestResult{false, msg, "unassignLoanApplications", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more loanApplicationsIds as a LoanApplications from a Customer
//----------------------------------------------------------------------------
func RemoveLoanApplicationsFromCustomer( customerId uint64, loanApplicationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( loanApplicationsIds, ",")

		for _, loanApplicationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanApplication

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanApplication
			// with a matching loanApplicationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , loanApplicationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LoanApplicationObj from the LoanApplications array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LoanApplications").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LoanApplications", loanApplicationsId )
				return utils.RequestResult{false, msg, "removeLoanApplications", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more loansIds as a Loans to a Customer
//----------------------------------------------------------------------------
func AddLoansToCustomer ( customerId uint64, loansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( loansIds, ",")

		for _, loansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Loan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Loan
			// with a matching loansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , loansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Loans using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Loans").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Loans", loansId )
				return utils.RequestResult{false, msg, "unassignLoans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more loansIds as a Loans from a Customer
//----------------------------------------------------------------------------
func RemoveLoansFromCustomer( customerId uint64, loansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( loansIds, ",")

		for _, loansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Loan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Loan
			// with a matching loansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , loansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LoanObj from the Loans array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Loans").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Loans", loansId )
				return utils.RequestResult{false, msg, "removeLoans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more portfoliosIds as a Portfolios to a Customer
//----------------------------------------------------------------------------
func AddPortfoliosToCustomer ( customerId uint64, portfoliosIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( portfoliosIds, ",")

		for _, portfoliosId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InvestmentPortfolio

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InvestmentPortfolio
			// with a matching portfoliosId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , portfoliosId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Portfolios using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Portfolios").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Portfolios", portfoliosId )
				return utils.RequestResult{false, msg, "unassignPortfolios", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more portfoliosIds as a Portfolios from a Customer
//----------------------------------------------------------------------------
func RemovePortfoliosFromCustomer( customerId uint64, portfoliosIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( portfoliosIds, ",")

		for _, portfoliosId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InvestmentPortfolio

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InvestmentPortfolio
			// with a matching portfoliosId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , portfoliosId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InvestmentPortfolioObj from the Portfolios array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Portfolios").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Portfolios", portfoliosId )
				return utils.RequestResult{false, msg, "removePortfolios", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more disputesIds as a Disputes to a Customer
//----------------------------------------------------------------------------
func AddDisputesToCustomer ( customerId uint64, disputesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Disputes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Disputes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "unassignDisputes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more disputesIds as a Disputes from a Customer
//----------------------------------------------------------------------------
func RemoveDisputesFromCustomer( customerId uint64, disputesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( disputesIds, ",")

		for _, disputesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dispute

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dispute
			// with a matching disputesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , disputesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DisputeObj from the Disputes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Disputes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Disputes", disputesId )
				return utils.RequestResult{false, msg, "removeDisputes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

