package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FinancialInstitutionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFinancialInstitution - creates a new db entry
//----------------------------------------------------------------------------
func CreateFinancialInstitution(obj model.FinancialInstitution)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FinancialInstitution with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FinancialInstitution", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFinancialInstitution", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFinancialInstitution - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFinancialInstitution(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FinancialInstitution

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FinancialInstitution with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FinancialInstitution using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FinancialInstitution using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFinancialInstitution", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFinancialInstitution - returns all
//----------------------------------------------------------------------------
func GetAllFinancialInstitution()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FinancialInstitution

	//----------------------------------------------------------------------------
	// Request the ORM to find all FinancialInstitution
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FinancialInstitution" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FinancialInstitution", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFinancialInstitution", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFinancialInstitution - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFinancialInstitution(obj model.FinancialInstitution)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FinancialInstitution using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FinancialInstitution using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFinancialInstitution", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFinancialInstitution - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFinancialInstitution(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFinancialInstitution(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FinancialInstitution)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FinancialInstitution using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FinancialInstitution using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFinancialInstitution", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more branchesIds as a Branches to a FinancialInstitution
//----------------------------------------------------------------------------
func AddBranchesToFinancialInstitution ( financialInstitutionId uint64, branchesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( branchesIds, ",")

		for _, branchesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Branch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Branch
			// with a matching branchesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , branchesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Branches using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Branches").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branches", branchesId )
				return utils.RequestResult{false, msg, "unassignBranches", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more branchesIds as a Branches from a FinancialInstitution
//----------------------------------------------------------------------------
func RemoveBranchesFromFinancialInstitution( financialInstitutionId uint64, branchesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( branchesIds, ",")

		for _, branchesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Branch

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Branch
			// with a matching branchesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , branchesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BranchObj from the Branches array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Branches").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Branches", branchesId )
				return utils.RequestResult{false, msg, "removeBranches", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more customersIds as a Customers to a FinancialInstitution
//----------------------------------------------------------------------------
func AddCustomersToFinancialInstitution ( financialInstitutionId uint64, customersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( customersIds, ",")

		for _, customersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Customers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Customers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )
				return utils.RequestResult{false, msg, "unassignCustomers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more customersIds as a Customers from a FinancialInstitution
//----------------------------------------------------------------------------
func RemoveCustomersFromFinancialInstitution( financialInstitutionId uint64, customersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( customersIds, ",")

		for _, customersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CustomerObj from the Customers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Customers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )
				return utils.RequestResult{false, msg, "removeCustomers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more productOfferingsIds as a ProductOfferings to a FinancialInstitution
//----------------------------------------------------------------------------
func AddProductOfferingsToFinancialInstitution ( financialInstitutionId uint64, productOfferingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( productOfferingsIds, ",")

		for _, productOfferingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductOffering

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductOffering
			// with a matching productOfferingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productOfferingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProductOfferings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductOfferings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductOfferings", productOfferingsId )
				return utils.RequestResult{false, msg, "unassignProductOfferings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productOfferingsIds as a ProductOfferings from a FinancialInstitution
//----------------------------------------------------------------------------
func RemoveProductOfferingsFromFinancialInstitution( financialInstitutionId uint64, productOfferingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( productOfferingsIds, ",")

		for _, productOfferingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductOffering

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductOffering
			// with a matching productOfferingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productOfferingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProductOfferingObj from the ProductOfferings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductOfferings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductOfferings", productOfferingsId )
				return utils.RequestResult{false, msg, "removeProductOfferings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentProcessorsIds as a PaymentProcessors to a FinancialInstitution
//----------------------------------------------------------------------------
func AddPaymentProcessorsToFinancialInstitution ( financialInstitutionId uint64, paymentProcessorsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentProcessorsIds, ",")

		for _, paymentProcessorsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentProcessor

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentProcessor
			// with a matching paymentProcessorsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentProcessorsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PaymentProcessors using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentProcessors").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentProcessors", paymentProcessorsId )
				return utils.RequestResult{false, msg, "unassignPaymentProcessors", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentProcessorsIds as a PaymentProcessors from a FinancialInstitution
//----------------------------------------------------------------------------
func RemovePaymentProcessorsFromFinancialInstitution( financialInstitutionId uint64, paymentProcessorsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentProcessorsIds, ",")

		for _, paymentProcessorsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PaymentProcessor

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PaymentProcessor
			// with a matching paymentProcessorsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentProcessorsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PaymentProcessorObj from the PaymentProcessors array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PaymentProcessors").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PaymentProcessors", paymentProcessorsId )
				return utils.RequestResult{false, msg, "removePaymentProcessors", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more compliancePoliciesIds as a CompliancePolicies to a FinancialInstitution
//----------------------------------------------------------------------------
func AddCompliancePoliciesToFinancialInstitution ( financialInstitutionId uint64, compliancePoliciesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( compliancePoliciesIds, ",")

		for _, compliancePoliciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompliancePolicy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompliancePolicy
			// with a matching compliancePoliciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , compliancePoliciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CompliancePolicies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompliancePolicies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompliancePolicies", compliancePoliciesId )
				return utils.RequestResult{false, msg, "unassignCompliancePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more compliancePoliciesIds as a CompliancePolicies from a FinancialInstitution
//----------------------------------------------------------------------------
func RemoveCompliancePoliciesFromFinancialInstitution( financialInstitutionId uint64, compliancePoliciesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FinancialInstitution with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFinancialInstitution(financialInstitutionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FinancialInstitution so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FinancialInstitution)

		// slice the ids on comma with no spaces
		ids := strings.Split( compliancePoliciesIds, ",")

		for _, compliancePoliciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CompliancePolicy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CompliancePolicy
			// with a matching compliancePoliciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , compliancePoliciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CompliancePolicyObj from the CompliancePolicies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompliancePolicies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompliancePolicies", compliancePoliciesId )
				return utils.RequestResult{false, msg, "removeCompliancePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FinancialInstitution from the gorm
		//----------------------------------------------------------------------------
		return GetFinancialInstitution(financialInstitutionId)

	} else {
		return parentRequestResult
	}
}

