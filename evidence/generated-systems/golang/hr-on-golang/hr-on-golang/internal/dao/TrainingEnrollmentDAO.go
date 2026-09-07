package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TrainingEnrollmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTrainingEnrollment - creates a new db entry
//----------------------------------------------------------------------------
func CreateTrainingEnrollment(obj model.TrainingEnrollment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TrainingEnrollment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TrainingEnrollment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTrainingEnrollment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTrainingEnrollment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTrainingEnrollment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TrainingEnrollment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TrainingEnrollment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TrainingEnrollment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TrainingEnrollment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTrainingEnrollment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTrainingEnrollment - returns all
//----------------------------------------------------------------------------
func GetAllTrainingEnrollment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TrainingEnrollment

	//----------------------------------------------------------------------------
	// Request the ORM to find all TrainingEnrollment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TrainingEnrollment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TrainingEnrollment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTrainingEnrollment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTrainingEnrollment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTrainingEnrollment(obj model.TrainingEnrollment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TrainingEnrollment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TrainingEnrollment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTrainingEnrollment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTrainingEnrollment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTrainingEnrollment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TrainingEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTrainingEnrollment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TrainingEnrollment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TrainingEnrollment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TrainingEnrollment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTrainingEnrollment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Course on a TrainingEnrollment
//----------------------------------------------------------------------------
func AssignCourseToTrainingEnrollment( trainingEnrollmentId uint64, courseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TrainingEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingEnrollment(trainingEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingEnrollment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.TrainingCourse

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TrainingCourse with a
		// matching courseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, courseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Course	to the TrainingEnrollment
			//----------------------------------------------------------------------------
			parentObj.Course = &childObj

			//----------------------------------------------------------------------------
			// save the TrainingEnrollment
			//----------------------------------------------------------------------------
			return UpdateTrainingEnrollment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Course", courseId )
			return utils.RequestResult{false, msg, "assignCourse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Course on a TrainingEnrollment
//----------------------------------------------------------------------------
func UnassignCourseFromTrainingEnrollment(trainingEnrollmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingEnrollment(trainingEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingEnrollment)

		//----------------------------------------------------------------------------
		// assign an empty TrainingCourse to the Course
		//----------------------------------------------------------------------------
		parentObj.Course = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Course
		//----------------------------------------------------------------------------
		parentObj.CourseId = nil;

		//----------------------------------------------------------------------------
		// save the TrainingEnrollment
		//----------------------------------------------------------------------------
		return UpdateTrainingEnrollment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a TrainingEnrollment
//----------------------------------------------------------------------------
func AssignEmployeeToTrainingEnrollment( trainingEnrollmentId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TrainingEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingEnrollment(trainingEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingEnrollment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching employeeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, employeeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Employee	to the TrainingEnrollment
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the TrainingEnrollment
			//----------------------------------------------------------------------------
			return UpdateTrainingEnrollment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a TrainingEnrollment
//----------------------------------------------------------------------------
func UnassignEmployeeFromTrainingEnrollment(trainingEnrollmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingEnrollment(trainingEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingEnrollment)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the TrainingEnrollment
		//----------------------------------------------------------------------------
		return UpdateTrainingEnrollment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Instructor on a TrainingEnrollment
//----------------------------------------------------------------------------
func AssignInstructorToTrainingEnrollment( trainingEnrollmentId uint64, instructorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TrainingEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingEnrollment(trainingEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingEnrollment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching instructorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, instructorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Instructor	to the TrainingEnrollment
			//----------------------------------------------------------------------------
			parentObj.Instructor = &childObj

			//----------------------------------------------------------------------------
			// save the TrainingEnrollment
			//----------------------------------------------------------------------------
			return UpdateTrainingEnrollment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Instructor", instructorId )
			return utils.RequestResult{false, msg, "assignInstructor", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Instructor on a TrainingEnrollment
//----------------------------------------------------------------------------
func UnassignInstructorFromTrainingEnrollment(trainingEnrollmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingEnrollment(trainingEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingEnrollment)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Instructor
		//----------------------------------------------------------------------------
		parentObj.Instructor = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Instructor
		//----------------------------------------------------------------------------
		parentObj.InstructorId = nil;

		//----------------------------------------------------------------------------
		// save the TrainingEnrollment
		//----------------------------------------------------------------------------
		return UpdateTrainingEnrollment(parentObj)

	} else {
		return parentRequestResult
	}

}


