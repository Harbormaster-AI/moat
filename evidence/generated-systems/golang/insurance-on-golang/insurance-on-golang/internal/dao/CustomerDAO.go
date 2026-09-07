package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
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
// adds one or more applicationsIds as a Applications to a Customer
//----------------------------------------------------------------------------
func AddApplicationsToCustomer ( customerId uint64, applicationsIds string )(utils.RequestResult) {

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
		ids := strings.Split( applicationsIds, ",")

		for _, applicationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Application

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Application
			// with a matching applicationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , applicationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Applications using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Applications").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Applications", applicationsId )
				return utils.RequestResult{false, msg, "unassignApplications", childObj}
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
// removes one or more applicationsIds as a Applications from a Customer
//----------------------------------------------------------------------------
func RemoveApplicationsFromCustomer( customerId uint64, applicationsIds string )(utils.RequestResult) {
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
		ids := strings.Split( applicationsIds, ",")

		for _, applicationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Application

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Application
			// with a matching applicationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , applicationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ApplicationObj from the Applications array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Applications").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Applications", applicationsId )
				return utils.RequestResult{false, msg, "removeApplications", childObj}
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
// adds one or more policiesIds as a Policies to a Customer
//----------------------------------------------------------------------------
func AddPoliciesToCustomer ( customerId uint64, policiesIds string )(utils.RequestResult) {

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
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Policies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "unassignPolicies", childObj}
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
// removes one or more policiesIds as a Policies from a Customer
//----------------------------------------------------------------------------
func RemovePoliciesFromCustomer( customerId uint64, policiesIds string )(utils.RequestResult) {
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
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyObj from the Policies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "removePolicies", childObj}
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
// adds one or more claimsIds as a Claims to a Customer
//----------------------------------------------------------------------------
func AddClaimsToCustomer ( customerId uint64, claimsIds string )(utils.RequestResult) {

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
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a Customer
//----------------------------------------------------------------------------
func RemoveClaimsFromCustomer( customerId uint64, claimsIds string )(utils.RequestResult) {
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
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more agentsIds as a Agents to a Customer
//----------------------------------------------------------------------------
func AddAgentsToCustomer ( customerId uint64, agentsIds string )(utils.RequestResult) {

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
		ids := strings.Split( agentsIds, ",")

		for _, agentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Agent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Agent
			// with a matching agentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , agentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Agents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Agents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agents", agentsId )
				return utils.RequestResult{false, msg, "unassignAgents", childObj}
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
// removes one or more agentsIds as a Agents from a Customer
//----------------------------------------------------------------------------
func RemoveAgentsFromCustomer( customerId uint64, agentsIds string )(utils.RequestResult) {
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
		ids := strings.Split( agentsIds, ",")

		for _, agentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Agent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Agent
			// with a matching agentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , agentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AgentObj from the Agents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Agents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Agents", agentsId )
				return utils.RequestResult{false, msg, "removeAgents", childObj}
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
// adds one or more beneficiariesIds as a Beneficiaries to a Customer
//----------------------------------------------------------------------------
func AddBeneficiariesToCustomer ( customerId uint64, beneficiariesIds string )(utils.RequestResult) {

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
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more beneficiariesIds as a Beneficiaries from a Customer
//----------------------------------------------------------------------------
func RemoveBeneficiariesFromCustomer( customerId uint64, beneficiariesIds string )(utils.RequestResult) {
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
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

