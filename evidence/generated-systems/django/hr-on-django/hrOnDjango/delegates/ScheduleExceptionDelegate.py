from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.ScheduleException import ScheduleException
from hrOnDjango.models.WorkSchedule import WorkSchedule
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ScheduleException
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScheduleExceptionDelegate Declaration
#======================================================================
class ScheduleExceptionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, scheduleExceptionId ):
		try:	
			scheduleException = ScheduleException.objects.filter(id=scheduleExceptionId)
			return scheduleException.first();
		except ScheduleException.DoesNotExist:
			raise ProcessingError("ScheduleException with id " + str(scheduleExceptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, scheduleException):
		for model in serializers.deserialize("json", scheduleException):
			model.save()
			return model;

	def create(self, scheduleException):
		scheduleException.save()
		return scheduleException;

	def saveFromJson(self, scheduleException):
		for model in serializers.deserialize("json", scheduleException):
			model.save()
			return scheduleException;
	
	def save(self, scheduleException):
		scheduleException.save()
		return scheduleException;
	
	def delete(self, scheduleExceptionId ):
		errMsg = "Failed to delete ScheduleException from db using id " + str(scheduleExceptionId)
		
		try:
			scheduleException = ScheduleException.objects.get(id=scheduleExceptionId)
			scheduleException.delete()
			return True
		except ScheduleException.DoesNotExist:
			raise ProcessingError("ScheduleException with id " + str(scheduleExceptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ScheduleException.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ScheduleException from db")
		except Exception:
			return None;
		
	def assignWorkSchedule( self, scheduleExceptionId, workScheduleId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.WorkScheduleDelegate import WorkScheduleDelegate

		errMsg = "Failed to assign element " + str(workScheduleId) + " for WorkSchedule on ScheduleException"

		try:
			# get the ScheduleException from db
			scheduleException = self.get( scheduleExceptionId ).first()	
			
			# get the WorkSchedule from db
			workSchedule = WorkScheduleDelegate().get(workScheduleId).first();
			
			# assign the WorkSchedule		
			scheduleException.workSchedule = workSchedule
			
			#save it
			scheduleException.save()

			# reload and return the appropriate version					
			return self.get( scheduleExceptionId );
		except ScheduleException.DoesNotExist:
			raise ProcessingError(errMsg + " : ScheduleException with id " + str(scheduleExceptionId) + " does not exist.")
		except WorkSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkSchedule with id " + str(workScheduleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkSchedule( self, scheduleExceptionId ):
		errMsg = "Failed to unassign element " + str(workScheduleId) + " for WorkSchedule on ScheduleException"

		try:
			# get the ScheduleException from db
			scheduleException = self.get( scheduleExceptionId ).first()	
			
			# assign to None for unassignment
			scheduleException.workSchedule = None			

			#save it
			scheduleException.save()

			# reload and return the appropriate version					
			return self.get( scheduleExceptionId );
		except ScheduleException.DoesNotExist:
			raise ProcessingError(errMsg + " : ScheduleException with id " + str(scheduleExceptionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, scheduleExceptionId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on ScheduleException"

		try:
			# get the ScheduleException from db
			scheduleException = self.get( scheduleExceptionId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			scheduleException.employee = employee
			
			#save it
			scheduleException.save()

			# reload and return the appropriate version					
			return self.get( scheduleExceptionId );
		except ScheduleException.DoesNotExist:
			raise ProcessingError(errMsg + " : ScheduleException with id " + str(scheduleExceptionId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, scheduleExceptionId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on ScheduleException"

		try:
			# get the ScheduleException from db
			scheduleException = self.get( scheduleExceptionId ).first()	
			
			# assign to None for unassignment
			scheduleException.employee = None			

			#save it
			scheduleException.save()

			# reload and return the appropriate version					
			return self.get( scheduleExceptionId );
		except ScheduleException.DoesNotExist:
			raise ProcessingError(errMsg + " : ScheduleException with id " + str(scheduleExceptionId) + " does not exist.")
		except Exception:
			return None;
		
