from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Termination import Termination
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.EmploymentAssignment import EmploymentAssignment
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Termination
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminationDelegate Declaration
#======================================================================
class TerminationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, terminationId ):
		try:	
			termination = Termination.objects.filter(id=terminationId)
			return termination.first();
		except Termination.DoesNotExist:
			raise ProcessingError("Termination with id " + str(terminationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, termination):
		for model in serializers.deserialize("json", termination):
			model.save()
			return model;

	def create(self, termination):
		termination.save()
		return termination;

	def saveFromJson(self, termination):
		for model in serializers.deserialize("json", termination):
			model.save()
			return termination;
	
	def save(self, termination):
		termination.save()
		return termination;
	
	def delete(self, terminationId ):
		errMsg = "Failed to delete Termination from db using id " + str(terminationId)
		
		try:
			termination = Termination.objects.get(id=terminationId)
			termination.delete()
			return True
		except Termination.DoesNotExist:
			raise ProcessingError("Termination with id " + str(terminationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Termination.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Termination from db")
		except Exception:
			return None;
		
	def assignEmployee( self, terminationId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on Termination"

		try:
			# get the Termination from db
			termination = self.get( terminationId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			termination.employee = employee
			
			#save it
			termination.save()

			# reload and return the appropriate version					
			return self.get( terminationId );
		except Termination.DoesNotExist:
			raise ProcessingError(errMsg + " : Termination with id " + str(terminationId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, terminationId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on Termination"

		try:
			# get the Termination from db
			termination = self.get( terminationId ).first()	
			
			# assign to None for unassignment
			termination.employee = None			

			#save it
			termination.save()

			# reload and return the appropriate version					
			return self.get( terminationId );
		except Termination.DoesNotExist:
			raise ProcessingError(errMsg + " : Termination with id " + str(terminationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAssignment( self, terminationId, assignmentId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentAssignmentDelegate import EmploymentAssignmentDelegate

		errMsg = "Failed to assign element " + str(assignmentId) + " for Assignment on Termination"

		try:
			# get the Termination from db
			termination = self.get( terminationId ).first()	
			
			# get the EmploymentAssignment from db
			employmentAssignment = EmploymentAssignmentDelegate().get(assignmentId).first();
			
			# assign the Assignment		
			termination.assignment = employmentAssignment
			
			#save it
			termination.save()

			# reload and return the appropriate version					
			return self.get( terminationId );
		except Termination.DoesNotExist:
			raise ProcessingError(errMsg + " : Termination with id " + str(terminationId) + " does not exist.")
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment with id " + str(assignmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAssignment( self, terminationId ):
		errMsg = "Failed to unassign element " + str(assignmentId) + " for Assignment on Termination"

		try:
			# get the Termination from db
			termination = self.get( terminationId ).first()	
			
			# assign to None for unassignment
			termination.employmentAssignment = None			

			#save it
			termination.save()

			# reload and return the appropriate version					
			return self.get( terminationId );
		except Termination.DoesNotExist:
			raise ProcessingError(errMsg + " : Termination with id " + str(terminationId) + " does not exist.")
		except Exception:
			return None;
		
