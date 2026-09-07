package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreatePolicy(obj model.Policy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Policy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Policy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Policy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Policy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Policy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Policy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPolicy - returns all
//----------------------------------------------------------------------------
func GetAllPolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Policy

	//----------------------------------------------------------------------------
	// Request the ORM to find all Policy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Policy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Policy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePolicy(obj model.Policy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Policy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Policy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Policy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Policy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePolicy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Insurer on a Policy
//----------------------------------------------------------------------------
func AssignInsurerToPolicy( policyId uint64, insurerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Insurer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Insurer with a
		// matching insurerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, insurerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Insurer	to the Policy
			//----------------------------------------------------------------------------
			parentObj.Insurer = &childObj

			//----------------------------------------------------------------------------
			// save the Policy
			//----------------------------------------------------------------------------
			return UpdatePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Insurer", insurerId )
			return utils.RequestResult{false, msg, "assignInsurer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Insurer on a Policy
//----------------------------------------------------------------------------
func UnassignInsurerFromPolicy(policyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// assign an empty Insurer to the Insurer
		//----------------------------------------------------------------------------
		parentObj.Insurer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Insurer
		//----------------------------------------------------------------------------
		parentObj.InsurerId = nil;

		//----------------------------------------------------------------------------
		// save the Policy
		//----------------------------------------------------------------------------
		return UpdatePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Customer on a Policy
//----------------------------------------------------------------------------
func AssignCustomerToPolicy( policyId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching customerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, customerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Customer	to the Policy
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Policy
			//----------------------------------------------------------------------------
			return UpdatePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Policy
//----------------------------------------------------------------------------
func UnassignCustomerFromPolicy(policyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Policy
		//----------------------------------------------------------------------------
		return UpdatePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a Policy
//----------------------------------------------------------------------------
func AssignProductToPolicy( policyId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsuranceProduct

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsuranceProduct with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the Policy
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the Policy
			//----------------------------------------------------------------------------
			return UpdatePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a Policy
//----------------------------------------------------------------------------
func UnassignProductFromPolicy(policyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// assign an empty InsuranceProduct to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the Policy
		//----------------------------------------------------------------------------
		return UpdatePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Agent on a Policy
//----------------------------------------------------------------------------
func AssignAgentToPolicy( policyId uint64, agentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Agent

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Agent with a
		// matching agentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, agentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Agent	to the Policy
			//----------------------------------------------------------------------------
			parentObj.Agent = &childObj

			//----------------------------------------------------------------------------
			// save the Policy
			//----------------------------------------------------------------------------
			return UpdatePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agent", agentId )
			return utils.RequestResult{false, msg, "assignAgent", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Agent on a Policy
//----------------------------------------------------------------------------
func UnassignAgentFromPolicy(policyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// assign an empty Agent to the Agent
		//----------------------------------------------------------------------------
		parentObj.Agent = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Agent
		//----------------------------------------------------------------------------
		parentObj.AgentId = nil;

		//----------------------------------------------------------------------------
		// save the Policy
		//----------------------------------------------------------------------------
		return UpdatePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a BillingAccount on a Policy
//----------------------------------------------------------------------------
func AssignBillingAccountToPolicy( policyId uint64, billingAccountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BillingAccount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BillingAccount with a
		// matching billingAccountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, billingAccountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BillingAccount	to the Policy
			//----------------------------------------------------------------------------
			parentObj.BillingAccount = &childObj

			//----------------------------------------------------------------------------
			// save the Policy
			//----------------------------------------------------------------------------
			return UpdatePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BillingAccount", billingAccountId )
			return utils.RequestResult{false, msg, "assignBillingAccount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BillingAccount on a Policy
//----------------------------------------------------------------------------
func UnassignBillingAccountFromPolicy(policyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// assign an empty BillingAccount to the BillingAccount
		//----------------------------------------------------------------------------
		parentObj.BillingAccount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BillingAccount
		//----------------------------------------------------------------------------
		parentObj.BillingAccountId = nil;

		//----------------------------------------------------------------------------
		// save the Policy
		//----------------------------------------------------------------------------
		return UpdatePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more coveragesIds as a Coverages to a Policy
//----------------------------------------------------------------------------
func AddCoveragesToPolicy ( policyId uint64, coveragesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PolicyCoverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PolicyCoverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Coverages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "unassignCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more coveragesIds as a Coverages from a Policy
//----------------------------------------------------------------------------
func RemoveCoveragesFromPolicy( policyId uint64, coveragesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PolicyCoverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PolicyCoverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyCoverageObj from the Coverages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "removeCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more insuredObjectsIds as a InsuredObjects to a Policy
//----------------------------------------------------------------------------
func AddInsuredObjectsToPolicy ( policyId uint64, insuredObjectsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( insuredObjectsIds, ",")

		for _, insuredObjectsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuredObject

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuredObject
			// with a matching insuredObjectsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insuredObjectsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InsuredObjects using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsuredObjects").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsuredObjects", insuredObjectsId )
				return utils.RequestResult{false, msg, "unassignInsuredObjects", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more insuredObjectsIds as a InsuredObjects from a Policy
//----------------------------------------------------------------------------
func RemoveInsuredObjectsFromPolicy( policyId uint64, insuredObjectsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( insuredObjectsIds, ",")

		for _, insuredObjectsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuredObject

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuredObject
			// with a matching insuredObjectsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insuredObjectsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsuredObjectObj from the InsuredObjects array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsuredObjects").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsuredObjects", insuredObjectsId )
				return utils.RequestResult{false, msg, "removeInsuredObjects", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more endorsementsIds as a Endorsements to a Policy
//----------------------------------------------------------------------------
func AddEndorsementsToPolicy ( policyId uint64, endorsementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( endorsementsIds, ",")

		for _, endorsementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Endorsement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Endorsement
			// with a matching endorsementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , endorsementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Endorsements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Endorsements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Endorsements", endorsementsId )
				return utils.RequestResult{false, msg, "unassignEndorsements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more endorsementsIds as a Endorsements from a Policy
//----------------------------------------------------------------------------
func RemoveEndorsementsFromPolicy( policyId uint64, endorsementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( endorsementsIds, ",")

		for _, endorsementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Endorsement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Endorsement
			// with a matching endorsementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , endorsementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EndorsementObj from the Endorsements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Endorsements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Endorsements", endorsementsId )
				return utils.RequestResult{false, msg, "removeEndorsements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more beneficiariesIds as a Beneficiaries to a Policy
//----------------------------------------------------------------------------
func AddBeneficiariesToPolicy ( policyId uint64, beneficiariesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( beneficiariesIds, ",")

		for _, beneficiariesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Beneficiary

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Beneficiary
			// with a matching beneficiariesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , beneficiariesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Beneficiaries using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Beneficiaries").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Beneficiaries", beneficiariesId )
				return utils.RequestResult{false, msg, "unassignBeneficiaries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more beneficiariesIds as a Beneficiaries from a Policy
//----------------------------------------------------------------------------
func RemoveBeneficiariesFromPolicy( policyId uint64, beneficiariesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( beneficiariesIds, ",")

		for _, beneficiariesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Beneficiary

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Beneficiary
			// with a matching beneficiariesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , beneficiariesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BeneficiaryObj from the Beneficiaries array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Beneficiaries").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Beneficiaries", beneficiariesId )
				return utils.RequestResult{false, msg, "removeBeneficiaries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more claimsIds as a Claims to a Policy
//----------------------------------------------------------------------------
func AddClaimsToPolicy ( policyId uint64, claimsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Claims using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "unassignClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a Policy
//----------------------------------------------------------------------------
func RemoveClaimsFromPolicy( policyId uint64, claimsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimObj from the Claims array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "removeClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reinsuranceAgreementsIds as a ReinsuranceAgreements to a Policy
//----------------------------------------------------------------------------
func AddReinsuranceAgreementsToPolicy ( policyId uint64, reinsuranceAgreementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( reinsuranceAgreementsIds, ",")

		for _, reinsuranceAgreementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ReinsuranceAgreement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ReinsuranceAgreement
			// with a matching reinsuranceAgreementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reinsuranceAgreementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ReinsuranceAgreements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ReinsuranceAgreements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ReinsuranceAgreements", reinsuranceAgreementsId )
				return utils.RequestResult{false, msg, "unassignReinsuranceAgreements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reinsuranceAgreementsIds as a ReinsuranceAgreements from a Policy
//----------------------------------------------------------------------------
func RemoveReinsuranceAgreementsFromPolicy( policyId uint64, reinsuranceAgreementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( reinsuranceAgreementsIds, ",")

		for _, reinsuranceAgreementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ReinsuranceAgreement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ReinsuranceAgreement
			// with a matching reinsuranceAgreementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reinsuranceAgreementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReinsuranceAgreementObj from the ReinsuranceAgreements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ReinsuranceAgreements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ReinsuranceAgreements", reinsuranceAgreementsId )
				return utils.RequestResult{false, msg, "removeReinsuranceAgreements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

