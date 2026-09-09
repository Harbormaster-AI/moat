from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Employee import Employee
from manufacturingOnDjango.models.WorkCenter import WorkCenter
from manufacturingOnDjango.models.ShiftAssignment import ShiftAssignment
from manufacturingOnDjango.models.CorrectiveAction import CorrectiveAction
from manufacturingOnDjango.exceptions import Exceptions

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
		
	def assignWorkCenter( self, employeeId, workCenterId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to assign element " + str(workCenterId) + " for WorkCenter on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# get the WorkCenter from db
			workCenter = WorkCenterDelegate().get(workCenterId).first();
			
			# assign the WorkCenter		
			employee.workCenter = workCenter
			
			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter with id " + str(workCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkCenter( self, employeeId ):
		errMsg = "Failed to unassign element " + str(workCenterId) + " for WorkCenter on Employee"

		try:
			# get the Employee from db
			employee = self.get( employeeId ).first()	
			
			# assign to None for unassignment
			employee.workCenter = None			

			#save it
			employee.save()

			# reload and return the appropriate version					
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
		
	def addShiftAssignments( self, employeeId, shiftAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ShiftAssignmentDelegate import ShiftAssignmentDelegate

		errMsg = "Failed to add elements " + str(shiftAssignmentsIds) + " for ShiftAssignments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = shiftAssignmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ShiftAssignment		
				shiftAssignment = ShiftAssignmentDelegate().get(id).first();	
				# add the ShiftAssignment
				employee.shiftAssignments.add(shiftAssignment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeShiftAssignments( self, employeeId, shiftAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ShiftAssignmentDelegate import ShiftAssignmentDelegate

		errMsg = "Failed to remove elements " + str(shiftAssignmentsIds) + " for ShiftAssignments on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = shiftAssignmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ShiftAssignment		
				shiftAssignment = ShiftAssignmentDelegate().get(id).first();	
				# add the ShiftAssignment
				employee.shiftAssignments.remove(shiftAssignment)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCorrectiveActions( self, employeeId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to add elements " + str(correctiveActionsIds) + " for CorrectiveActions on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				employee.correctiveActions.add(correctiveAction)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCorrectiveActions( self, employeeId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to remove elements " + str(correctiveActionsIds) + " for CorrectiveActions on Employee"

		try:
			# get the Employee
			employee = self.get( employeeId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				employee.correctiveActions.remove(correctiveAction)
				
			# save it		
			employee.save()
			
			# reload and return the appropriate version
			return self.get( employeeId );
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
