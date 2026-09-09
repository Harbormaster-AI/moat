from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.EmploymentAssignment import EmploymentAssignment
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.Position import Position
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model EmploymentAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentAssignmentDelegate Declaration
#======================================================================
class EmploymentAssignmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, employmentAssignmentId ):
		try:	
			employmentAssignment = EmploymentAssignment.objects.filter(id=employmentAssignmentId)
			return employmentAssignment.first();
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError("EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, employmentAssignment):
		for model in serializers.deserialize("json", employmentAssignment):
			model.save()
			return model;

	def create(self, employmentAssignment):
		employmentAssignment.save()
		return employmentAssignment;

	def saveFromJson(self, employmentAssignment):
		for model in serializers.deserialize("json", employmentAssignment):
			model.save()
			return employmentAssignment;
	
	def save(self, employmentAssignment):
		employmentAssignment.save()
		return employmentAssignment;
	
	def delete(self, employmentAssignmentId ):
		errMsg = "Failed to delete EmploymentAssignment from db using id " + str(employmentAssignmentId)
		
		try:
			employmentAssignment = EmploymentAssignment.objects.get(id=employmentAssignmentId)
			employmentAssignment.delete()
			return True
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError("EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = EmploymentAssignment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all EmploymentAssignment from db")
		except Exception:
			return None;
		
	def assignEmployee( self, employmentAssignmentId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on EmploymentAssignment"

		try:
			# get the EmploymentAssignment from db
			employmentAssignment = self.get( employmentAssignmentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			employmentAssignment.employee = employee
			
			#save it
			employmentAssignment.save()

			# reload and return the appropriate version					
			return self.get( employmentAssignmentId );
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, employmentAssignmentId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on EmploymentAssignment"

		try:
			# get the EmploymentAssignment from db
			employmentAssignment = self.get( employmentAssignmentId ).first()	
			
			# assign to None for unassignment
			employmentAssignment.employee = None			

			#save it
			employmentAssignment.save()

			# reload and return the appropriate version					
			return self.get( employmentAssignmentId );
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPosition( self, employmentAssignmentId, positionId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to assign element " + str(positionId) + " for Position on EmploymentAssignment"

		try:
			# get the EmploymentAssignment from db
			employmentAssignment = self.get( employmentAssignmentId ).first()	
			
			# get the Position from db
			position = PositionDelegate().get(positionId).first();
			
			# assign the Position		
			employmentAssignment.position = position
			
			#save it
			employmentAssignment.save()

			# reload and return the appropriate version					
			return self.get( employmentAssignmentId );
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPosition( self, employmentAssignmentId ):
		errMsg = "Failed to unassign element " + str(positionId) + " for Position on EmploymentAssignment"

		try:
			# get the EmploymentAssignment from db
			employmentAssignment = self.get( employmentAssignmentId ).first()	
			
			# assign to None for unassignment
			employmentAssignment.position = None			

			#save it
			employmentAssignment.save()

			# reload and return the appropriate version					
			return self.get( employmentAssignmentId );
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSupervisor( self, employmentAssignmentId, supervisorId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(supervisorId) + " for Supervisor on EmploymentAssignment"

		try:
			# get the EmploymentAssignment from db
			employmentAssignment = self.get( employmentAssignmentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(supervisorId).first();
			
			# assign the Supervisor		
			employmentAssignment.supervisor = employee
			
			#save it
			employmentAssignment.save()

			# reload and return the appropriate version					
			return self.get( employmentAssignmentId );
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(supervisorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupervisor( self, employmentAssignmentId ):
		errMsg = "Failed to unassign element " + str(supervisorId) + " for Supervisor on EmploymentAssignment"

		try:
			# get the EmploymentAssignment from db
			employmentAssignment = self.get( employmentAssignmentId ).first()	
			
			# assign to None for unassignment
			employmentAssignment.employee = None			

			#save it
			employmentAssignment.save()

			# reload and return the appropriate version					
			return self.get( employmentAssignmentId );
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment with id " + str(employmentAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
