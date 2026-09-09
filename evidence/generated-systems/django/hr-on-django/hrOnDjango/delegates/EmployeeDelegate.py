from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.Department import Department
from hrOnDjango.models.Location import Location
from hrOnDjango.models.CostCenter import CostCenter
from hrOnDjango.models.EmploymentAssignment import EmploymentAssignment
from hrOnDjango.models.EmploymentContract import EmploymentContract
from hrOnDjango.models.BenefitEnrollment import BenefitEnrollment
from hrOnDjango.models.Timesheet import Timesheet
from hrOnDjango.models.LeaveRequest import LeaveRequest
from hrOnDjango.models.PerformanceReview import PerformanceReview
from hrOnDjango.models.TrainingEnrollment import TrainingEnrollment
from hrOnDjango.models.WorkAuthorization import WorkAuthorization
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Employee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmployeeDelegate Declaration
#======================================================================
class EmployeeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, employeeId ):
		try:	
			employee = Employee.objects.filter(id=employeeId)
			return employee.first();
		except Employee.DoesNotExist:
			raise ProcessingError("Employee with id " + str(employeeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, employee):
		for model in serializers.deserialize("json", employee):
			model.save()
			return model;

	def create(self, employee):
		employee.save()
		return employee;

	def saveFromJson(self, employee):
		for model in serializers.deserialize("json", employee):
			model.save()
			return employee;
	
	def save(self, employee):
		employee.save()
		return employee;
	
	def delete(self, employeeId ):
		errMsg = "Failed to delete Employee from db using id " + str(employeeId)
		
		try:
			employee = Employee.objects.get(id=employeeId)
			employee.delete()
			return True
		except Employee.DoesNotExist:
			raise ProcessingError("Employee with id " + str(employeeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Employee.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Employee from db")
		except Exception:
			return None;
		
	def assignManager( self, employeeId, managerId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(managerId) + " for Manager on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(managerId).first();
			
			# assign the Manager		
			employee.manager = employee
			
			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(managerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignManager( self, employeeId ):
		errMsg = "Failed to unassign element " + str(managerId) + " for Manager on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# assign to None for unassignment
			employee.employee = None			

			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDepartment( self, employeeId, departmentId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to assign element " + str(departmentId) + " for Department on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# get the Department from db
			department = DepartmentDelegate().get(departmentId).first();
			
			# assign the Department		
			employee.department = department
			
			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDepartment( self, employeeId ):
		errMsg = "Failed to unassign element " + str(departmentId) + " for Department on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# assign to None for unassignment
			employee.department = None			

			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPrimaryLocation( self, employeeId, primaryLocationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to assign element " + str(primaryLocationId) + " for PrimaryLocation on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# get the Location from db
			location = LocationDelegate().get(primaryLocationId).first();
			
			# assign the PrimaryLocation		
			employee.primaryLocation = location
			
			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(primaryLocationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPrimaryLocation( self, employeeId ):
		errMsg = "Failed to unassign element " + str(primaryLocationId) + " for PrimaryLocation on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# assign to None for unassignment
			employee.location = None			

			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCostCenter( self, employeeId, costCenterId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

		errMsg = "Failed to assign element " + str(costCenterId) + " for CostCenter on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# get the CostCenter from db
			costCenter = CostCenterDelegate().get(costCenterId).first();
			
			# assign the CostCenter		
			employee.costCenter = costCenter
			
			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCostCenter( self, employeeId ):
		errMsg = "Failed to unassign element " + str(costCenterId) + " for CostCenter on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# assign to None for unassignment
			employee.costCenter = None			

			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
		
	def addDirectReports( self, employeeId, directReportsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to add elements " + str(directReportsIds) + " for DirectReports on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = directReportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				employee.directReports.add(employee)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDirectReports( self, employeeId, directReportsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to remove elements " + str(directReportsIds) + " for DirectReports on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = directReportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				employee.directReports.remove(employee)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmploymentAssignments( self, employeeId, employmentAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentAssignmentDelegate import EmploymentAssignmentDelegate

		errMsg = "Failed to add elements " + str(employmentAssignmentsIds) + " for EmploymentAssignments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = employmentAssignmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmploymentAssignment		
				employmentAssignment = EmploymentAssignmentDelegate().get(id).first();	
				# add the EmploymentAssignment
				employee.employmentAssignments.add(employmentAssignment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmploymentAssignments( self, employeeId, employmentAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentAssignmentDelegate import EmploymentAssignmentDelegate

		errMsg = "Failed to remove elements " + str(employmentAssignmentsIds) + " for EmploymentAssignments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = employmentAssignmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmploymentAssignment		
				employmentAssignment = EmploymentAssignmentDelegate().get(id).first();	
				# add the EmploymentAssignment
				employee.employmentAssignments.remove(employmentAssignment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContracts( self, employeeId, contractsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmploymentContract		
				employmentContract = EmploymentContractDelegate().get(id).first();	
				# add the EmploymentContract
				employee.contracts.add(employmentContract)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, employeeId, contractsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentContractDelegate import EmploymentContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmploymentContract		
				employmentContract = EmploymentContractDelegate().get(id).first();	
				# add the EmploymentContract
				employee.contracts.remove(employmentContract)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except EmploymentContract.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentContract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBenefitEnrollments( self, employeeId, benefitEnrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitEnrollmentDelegate import BenefitEnrollmentDelegate

		errMsg = "Failed to add elements " + str(benefitEnrollmentsIds) + " for BenefitEnrollments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = benefitEnrollmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BenefitEnrollment		
				benefitEnrollment = BenefitEnrollmentDelegate().get(id).first();	
				# add the BenefitEnrollment
				employee.benefitEnrollments.add(benefitEnrollment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBenefitEnrollments( self, employeeId, benefitEnrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitEnrollmentDelegate import BenefitEnrollmentDelegate

		errMsg = "Failed to remove elements " + str(benefitEnrollmentsIds) + " for BenefitEnrollments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = benefitEnrollmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BenefitEnrollment		
				benefitEnrollment = BenefitEnrollmentDelegate().get(id).first();	
				# add the BenefitEnrollment
				employee.benefitEnrollments.remove(benefitEnrollment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except BenefitEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitEnrollment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTimesheets( self, employeeId, timesheetsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TimesheetDelegate import TimesheetDelegate

		errMsg = "Failed to add elements " + str(timesheetsIds) + " for Timesheets on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = timesheetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Timesheet		
				timesheet = TimesheetDelegate().get(id).first();	
				# add the Timesheet
				employee.timesheets.add(timesheet)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTimesheets( self, employeeId, timesheetsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TimesheetDelegate import TimesheetDelegate

		errMsg = "Failed to remove elements " + str(timesheetsIds) + " for Timesheets on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = timesheetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Timesheet		
				timesheet = TimesheetDelegate().get(id).first();	
				# add the Timesheet
				employee.timesheets.remove(timesheet)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLeaveRequests( self, employeeId, leaveRequestsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LeaveRequestDelegate import LeaveRequestDelegate

		errMsg = "Failed to add elements " + str(leaveRequestsIds) + " for LeaveRequests on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = leaveRequestsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LeaveRequest		
				leaveRequest = LeaveRequestDelegate().get(id).first();	
				# add the LeaveRequest
				employee.leaveRequests.add(leaveRequest)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLeaveRequests( self, employeeId, leaveRequestsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LeaveRequestDelegate import LeaveRequestDelegate

		errMsg = "Failed to remove elements " + str(leaveRequestsIds) + " for LeaveRequests on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = leaveRequestsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LeaveRequest		
				leaveRequest = LeaveRequestDelegate().get(id).first();	
				# add the LeaveRequest
				employee.leaveRequests.remove(leaveRequest)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPerformanceReviews( self, employeeId, performanceReviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PerformanceReviewDelegate import PerformanceReviewDelegate

		errMsg = "Failed to add elements " + str(performanceReviewsIds) + " for PerformanceReviews on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = performanceReviewsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PerformanceReview		
				performanceReview = PerformanceReviewDelegate().get(id).first();	
				# add the PerformanceReview
				employee.performanceReviews.add(performanceReview)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePerformanceReviews( self, employeeId, performanceReviewsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PerformanceReviewDelegate import PerformanceReviewDelegate

		errMsg = "Failed to remove elements " + str(performanceReviewsIds) + " for PerformanceReviews on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = performanceReviewsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PerformanceReview		
				performanceReview = PerformanceReviewDelegate().get(id).first();	
				# add the PerformanceReview
				employee.performanceReviews.remove(performanceReview)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except PerformanceReview.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceReview does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTrainingEnrollments( self, employeeId, trainingEnrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingEnrollmentDelegate import TrainingEnrollmentDelegate

		errMsg = "Failed to add elements " + str(trainingEnrollmentsIds) + " for TrainingEnrollments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = trainingEnrollmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrainingEnrollment		
				trainingEnrollment = TrainingEnrollmentDelegate().get(id).first();	
				# add the TrainingEnrollment
				employee.trainingEnrollments.add(trainingEnrollment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrainingEnrollments( self, employeeId, trainingEnrollmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingEnrollmentDelegate import TrainingEnrollmentDelegate

		errMsg = "Failed to remove elements " + str(trainingEnrollmentsIds) + " for TrainingEnrollments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = trainingEnrollmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrainingEnrollment		
				trainingEnrollment = TrainingEnrollmentDelegate().get(id).first();	
				# add the TrainingEnrollment
				employee.trainingEnrollments.remove(trainingEnrollment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWorkAuthorizations( self, employeeId, workAuthorizationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.WorkAuthorizationDelegate import WorkAuthorizationDelegate

		errMsg = "Failed to add elements " + str(workAuthorizationsIds) + " for WorkAuthorizations on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = workAuthorizationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WorkAuthorization		
				workAuthorization = WorkAuthorizationDelegate().get(id).first();	
				# add the WorkAuthorization
				employee.workAuthorizations.add(workAuthorization)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkAuthorization does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkAuthorizations( self, employeeId, workAuthorizationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.WorkAuthorizationDelegate import WorkAuthorizationDelegate

		errMsg = "Failed to remove elements " + str(workAuthorizationsIds) + " for WorkAuthorizations on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = workAuthorizationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WorkAuthorization		
				workAuthorization = WorkAuthorizationDelegate().get(id).first();	
				# add the WorkAuthorization
				employee.workAuthorizations.remove(workAuthorization)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except WorkAuthorization.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkAuthorization does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
