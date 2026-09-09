from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Timesheet import Timesheet
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.TimeEntry import TimeEntry
from hrOnDjango.models.Approval import Approval
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Timesheet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimesheetDelegate Declaration
#======================================================================
class TimesheetDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, timesheetId ):
		try:	
			timesheet = Timesheet.objects.filter(id=timesheetId)
			return timesheet.first();
		except Timesheet.DoesNotExist:
			raise ProcessingError("Timesheet with id " + str(timesheetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, timesheet):
		for model in serializers.deserialize("json", timesheet):
			model.save()
			return model;

	def create(self, timesheet):
		timesheet.save()
		return timesheet;

	def saveFromJson(self, timesheet):
		for model in serializers.deserialize("json", timesheet):
			model.save()
			return timesheet;
	
	def save(self, timesheet):
		timesheet.save()
		return timesheet;
	
	def delete(self, timesheetId ):
		errMsg = "Failed to delete Timesheet from db using id " + str(timesheetId)
		
		try:
			timesheet = Timesheet.objects.get(id=timesheetId)
			timesheet.delete()
			return True
		except Timesheet.DoesNotExist:
			raise ProcessingError("Timesheet with id " + str(timesheetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Timesheet.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Timesheet from db")
		except Exception:
			return None;
		
	def assignEmployee( self, timesheetId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on Timesheet"

		try:
			# get the Timesheet from db
			timesheet = self.get( timesheetId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			timesheet.employee = employee
			
			#save it
			timesheet.save()

			# reload and return the appropriate version					
			return self.get( timesheetId );
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, timesheetId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on Timesheet"

		try:
			# get the Timesheet from db
			timesheet = self.get( timesheetId ).first()	
			
			# assign to None for unassignment
			timesheet.employee = None			

			#save it
			timesheet.save()

			# reload and return the appropriate version					
			return self.get( timesheetId );
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except Exception:
			return None;
		
	def addTimeEntries( self, timesheetId, timeEntriesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TimeEntryDelegate import TimeEntryDelegate

		errMsg = "Failed to add elements " + str(timeEntriesIds) + " for TimeEntries on Timesheet"

		try:
			# get the Timesheet
			timesheet = self.get( timesheetId ).first()
				
			# split on a comma with no spaces
			idList = timeEntriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TimeEntry		
				timeEntry = TimeEntryDelegate().get(id).first();	
				# add the TimeEntry
				timesheet.timeEntries.add(timeEntry)
				
			# save it		
			timesheet.save()
			
			# reload and return the appropriate version
			return self.get( timesheetId );
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTimeEntries( self, timesheetId, timeEntriesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TimeEntryDelegate import TimeEntryDelegate

		errMsg = "Failed to remove elements " + str(timeEntriesIds) + " for TimeEntries on Timesheet"

		try:
			# get the Timesheet
			timesheet = self.get( timesheetId ).first()
				
			# split on a comma with no spaces
			idList = timeEntriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TimeEntry		
				timeEntry = TimeEntryDelegate().get(id).first();	
				# add the TimeEntry
				timesheet.timeEntries.remove(timeEntry)
				
			# save it		
			timesheet.save()
			
			# reload and return the appropriate version
			return self.get( timesheetId );
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except TimeEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeEntry does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addApprovals( self, timesheetId, approvalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ApprovalDelegate import ApprovalDelegate

		errMsg = "Failed to add elements " + str(approvalsIds) + " for Approvals on Timesheet"

		try:
			# get the Timesheet
			timesheet = self.get( timesheetId ).first()
				
			# split on a comma with no spaces
			idList = approvalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Approval		
				approval = ApprovalDelegate().get(id).first();	
				# add the Approval
				timesheet.approvals.add(approval)
				
			# save it		
			timesheet.save()
			
			# reload and return the appropriate version
			return self.get( timesheetId );
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApprovals( self, timesheetId, approvalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ApprovalDelegate import ApprovalDelegate

		errMsg = "Failed to remove elements " + str(approvalsIds) + " for Approvals on Timesheet"

		try:
			# get the Timesheet
			timesheet = self.get( timesheetId ).first()
				
			# split on a comma with no spaces
			idList = approvalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Approval		
				approval = ApprovalDelegate().get(id).first();	
				# add the Approval
				timesheet.approvals.remove(approval)
				
			# save it		
			timesheet.save()
			
			# reload and return the appropriate version
			return self.get( timesheetId );
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
