from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.ShiftAssignment import ShiftAssignment
from manufacturingOnDjango.models.Shift import Shift
from manufacturingOnDjango.models.Employee import Employee
from manufacturingOnDjango.models.WorkCenter import WorkCenter
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ShiftAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftAssignmentDelegate Declaration
#======================================================================
class ShiftAssignmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, shiftAssignmentId ):
		try:	
			shiftAssignment = ShiftAssignment.objects.filter(id=shiftAssignmentId)
			return shiftAssignment.first();
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError("ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, shiftAssignment):
		for model in serializers.deserialize("json", shiftAssignment):
			model.save()
			return model;

	def create(self, shiftAssignment):
		shiftAssignment.save()
		return shiftAssignment;

	def saveFromJson(self, shiftAssignment):
		for model in serializers.deserialize("json", shiftAssignment):
			model.save()
			return shiftAssignment;
	
	def save(self, shiftAssignment):
		shiftAssignment.save()
		return shiftAssignment;
	
	def delete(self, shiftAssignmentId ):
		errMsg = "Failed to delete ShiftAssignment from db using id " + str(shiftAssignmentId)
		
		try:
			shiftAssignment = ShiftAssignment.objects.get(id=shiftAssignmentId)
			shiftAssignment.delete()
			return True
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError("ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ShiftAssignment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ShiftAssignment from db")
		except Exception:
			return None;
		
	def assignShift( self, shiftAssignmentId, shiftId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ShiftDelegate import ShiftDelegate

		errMsg = "Failed to assign element " + str(shiftId) + " for Shift on ShiftAssignment"

		try:
			# get the ShiftAssignment from db
			shiftAssignment = self.get( shiftAssignmentId ).first()	
			
			# get the Shift from db
			shift = ShiftDelegate().get(shiftId).first();
			
			# assign the Shift		
			shiftAssignment.shift = shift
			
			#save it
			shiftAssignment.save()

			# reload and return the appropriate version					
			return self.get( shiftAssignmentId );
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except Shift.DoesNotExist:
			raise ProcessingError(errMsg + " : Shift with id " + str(shiftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignShift( self, shiftAssignmentId ):
		errMsg = "Failed to unassign element " + str(shiftId) + " for Shift on ShiftAssignment"

		try:
			# get the ShiftAssignment from db
			shiftAssignment = self.get( shiftAssignmentId ).first()	
			
			# assign to None for unassignment
			shiftAssignment.shift = None			

			#save it
			shiftAssignment.save()

			# reload and return the appropriate version					
			return self.get( shiftAssignmentId );
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, shiftAssignmentId, employeeId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on ShiftAssignment"

		try:
			# get the ShiftAssignment from db
			shiftAssignment = self.get( shiftAssignmentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			shiftAssignment.employee = employee
			
			#save it
			shiftAssignment.save()

			# reload and return the appropriate version					
			return self.get( shiftAssignmentId );
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, shiftAssignmentId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on ShiftAssignment"

		try:
			# get the ShiftAssignment from db
			shiftAssignment = self.get( shiftAssignmentId ).first()	
			
			# assign to None for unassignment
			shiftAssignment.employee = None			

			#save it
			shiftAssignment.save()

			# reload and return the appropriate version					
			return self.get( shiftAssignmentId );
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkCenter( self, shiftAssignmentId, workCenterId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

		errMsg = "Failed to assign element " + str(workCenterId) + " for WorkCenter on ShiftAssignment"

		try:
			# get the ShiftAssignment from db
			shiftAssignment = self.get( shiftAssignmentId ).first()	
			
			# get the WorkCenter from db
			workCenter = WorkCenterDelegate().get(workCenterId).first();
			
			# assign the WorkCenter		
			shiftAssignment.workCenter = workCenter
			
			#save it
			shiftAssignment.save()

			# reload and return the appropriate version					
			return self.get( shiftAssignmentId );
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except WorkCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkCenter with id " + str(workCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkCenter( self, shiftAssignmentId ):
		errMsg = "Failed to unassign element " + str(workCenterId) + " for WorkCenter on ShiftAssignment"

		try:
			# get the ShiftAssignment from db
			shiftAssignment = self.get( shiftAssignmentId ).first()	
			
			# assign to None for unassignment
			shiftAssignment.workCenter = None			

			#save it
			shiftAssignment.save()

			# reload and return the appropriate version					
			return self.get( shiftAssignmentId );
		except ShiftAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : ShiftAssignment with id " + str(shiftAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
