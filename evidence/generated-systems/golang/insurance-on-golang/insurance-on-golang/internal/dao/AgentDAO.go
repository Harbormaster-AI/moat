package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AgentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAgent - creates a new db entry
//----------------------------------------------------------------------------
func CreateAgent(obj model.Agent)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Agent with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Agent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAgent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAgent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAgent(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Agent

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Agent with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Agent using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Agent using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAgent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAgent - returns all
//----------------------------------------------------------------------------
func GetAllAgent()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Agent

	//----------------------------------------------------------------------------
	// Request the ORM to find all Agent
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Agent" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Agent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAgent", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAgent - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAgent(obj model.Agent)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Agent using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Agent using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAgent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAgent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAgent(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Agent with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAgent(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Agent)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Agent using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Agent using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAgent", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Distributor on a Agent
//----------------------------------------------------------------------------
func AssignDistributorToAgent( agentId uint64, distributorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Agent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgent(agentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Distributor

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Distributor with a
		// matching distributorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, distributorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Distributor	to the Agent
			//----------------------------------------------------------------------------
			parentObj.Distributor = &childObj

			//----------------------------------------------------------------------------
			// save the Agent
			//----------------------------------------------------------------------------
			return UpdateAgent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Distributor", distributorId )
			return utils.RequestResult{false, msg, "assignDistributor", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Distributor on a Agent
//----------------------------------------------------------------------------
func UnassignDistributorFromAgent(agentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgent(agentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agent)

		//----------------------------------------------------------------------------
		// assign an empty Distributor to the Distributor
		//----------------------------------------------------------------------------
		parentObj.Distributor = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Distributor
		//----------------------------------------------------------------------------
		parentObj.DistributorId = nil;

		//----------------------------------------------------------------------------
		// save the Agent
		//----------------------------------------------------------------------------
		return UpdateAgent(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a Agent
//----------------------------------------------------------------------------
func AddPoliciesToAgent ( agentId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgent(agentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agent)

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
		// retrieve the modified Agent from the gorm
		//----------------------------------------------------------------------------
		return GetAgent(agentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a Agent
//----------------------------------------------------------------------------
func RemovePoliciesFromAgent( agentId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Agent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgent(agentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agent)

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
		// retrieve the modified Agent from the gorm
		//----------------------------------------------------------------------------
		return GetAgent(agentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more customersIds as a Customers to a Agent
//----------------------------------------------------------------------------
func AddCustomersToAgent ( agentId uint64, customersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgent(agentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agent)

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
		// retrieve the modified Agent from the gorm
		//----------------------------------------------------------------------------
		return GetAgent(agentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more customersIds as a Customers from a Agent
//----------------------------------------------------------------------------
func RemoveCustomersFromAgent( agentId uint64, customersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Agent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgent(agentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agent)

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
		// retrieve the modified Agent from the gorm
		//----------------------------------------------------------------------------
		return GetAgent(agentId)

	} else {
		return parentRequestResult
	}
}

