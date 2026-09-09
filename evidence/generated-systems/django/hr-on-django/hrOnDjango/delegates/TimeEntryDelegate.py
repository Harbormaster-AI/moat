from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.TimeEntry import TimeEntry
from hrOnDjango.models.Timesheet import Timesheet
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.CostCenter import CostCenter
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TimeEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeEntryDelegate Declaration
#======================================================================
class TimeEntryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, timeEntryId ):
		try:	
			timeEntry = TimeEntry.objects.filter(id=timeEntryId)
			return timeEntry.first();
		except TimeEntry.DoesNotExist:
			raise ProcessingError("TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, timeEntry):
		for model in serializers.deserialize("json", timeEntry):
			model.save()
			return model;

	def create(self, timeEntry):
		timeEntry.save()
		return timeEntry;

	def saveFromJson(self, timeEntry):
		for model in serializers.deserialize("json", timeEntry):
			model.save()
			return timeEntry;
	
	def save(self, timeEntry):
		timeEntry.save()
		return timeEntry;
	
	def delete(self, timeEntryId ):
		errMsg = "Failed to delete TimeEntry from db using id " + str(timeEntryId)
		
		try:
			timeEntry = TimeEntry.objects.get(id=timeEntryId)
			timeEntry.delete()
			return True
		except TimeEntry.DoesNotExist:
			raise ProcessingError("TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TimeEntry.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TimeEntry from db")
		except Exception:
			return None;
		
	def assignTimesheet( self, timeEntryId, timesheetId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TimesheetDelegate import TimesheetDelegate

		errMsg = "Failed to assign element " + str(timesheetId) + " for Timesheet on TimeEntry"

		try:
			# get the TimeEntry from db
			timeEntry = self.get( timeEntryId ).first()	
			
			# get the Timesheet from db
			timesheet = TimesheetDelegate().get(timesheetId).first();
			
			# assign the Timesheet		
			timeEntry.timesheet = timesheet
			
			#save it
			timeEntry.save()

			# reload and return the appropriate version					
			return self.get( timeEntryId );
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTimesheet( self, timeEntryId ):
		errMsg = "Failed to unassign element " + str(timesheetId) + " for Timesheet on TimeEntry"

		try:
			# get the TimeEntry from db
			timeEntry = self.get( timeEntryId ).first()	
			
			# assign to None for unassignment
			timeEntry.timesheet = None			

			#save it
			timeEntry.save()

			# reload and return the appropriate version					
			return self.get( timeEntryId );
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, timeEntryId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on TimeEntry"

		try:
			# get the TimeEntry from db
			timeEntry = self.get( timeEntryId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			timeEntry.employee = employee
			
			#save it
			timeEntry.save()

			# reload and return the appropriate version					
			return self.get( timeEntryId );
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, timeEntryId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on TimeEntry"

		try:
			# get the TimeEntry from db
			timeEntry = self.get( timeEntryId ).first()	
			
			# assign to None for unassignment
			timeEntry.employee = None			

			#save it
			timeEntry.save()

			# reload and return the appropriate version					
			return self.get( timeEntryId );
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCostCenter( self, timeEntryId, costCenterId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

		errMsg = "Failed to assign element " + str(costCenterId) + " for CostCenter on TimeEntry"

		try:
			# get the TimeEntry from db
			timeEntry = self.get( timeEntryId ).first()	
			
			# get the CostCenter from db
			costCenter = CostCenterDelegate().get(costCenterId).first();
			
			# assign the CostCenter		
			timeEntry.costCenter = costCenter
			
			#save it
			timeEntry.save()

			# reload and return the appropriate version					
			return self.get( timeEntryId );
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCostCenter( self, timeEntryId ):
		errMsg = "Failed to unassign element " + str(costCenterId) + " for CostCenter on TimeEntry"

		try:
			# get the TimeEntry from db
			timeEntry = self.get( timeEntryId ).first()	
			
			# assign to None for unassignment
			timeEntry.costCenter = None			

			#save it
			timeEntry.save()

			# reload and return the appropriate version					
			return self.get( timeEntryId );
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry with id " + str(timeEntryId) + " does not exist.")
		except Exception:
			return None;
		
