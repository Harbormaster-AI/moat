from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Approval import Approval
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.Timesheet import Timesheet
from hrOnDjango.models.LeaveRequest import LeaveRequest
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Approval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ApprovalDelegate Declaration
#======================================================================
class ApprovalDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, approvalId ):
		try:	
			approval = Approval.objects.filter(id=approvalId)
			return approval.first();
		except Approval.DoesNotExist:
			raise ProcessingError("Approval with id " + str(approvalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, approval):
		for model in serializers.deserialize("json", approval):
			model.save()
			return model;

	def create(self, approval):
		approval.save()
		return approval;

	def saveFromJson(self, approval):
		for model in serializers.deserialize("json", approval):
			model.save()
			return approval;
	
	def save(self, approval):
		approval.save()
		return approval;
	
	def delete(self, approvalId ):
		errMsg = "Failed to delete Approval from db using id " + str(approvalId)
		
		try:
			approval = Approval.objects.get(id=approvalId)
			approval.delete()
			return True
		except Approval.DoesNotExist:
			raise ProcessingError("Approval with id " + str(approvalId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Approval.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Approval from db")
		except Exception:
			return None;
		
	def assignApprover( self, approvalId, approverId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(approverId) + " for Approver on Approval"

		try:
			# get the Approval from db
			approval = self.get( approvalId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(approverId).first();
			
			# assign the Approver		
			approval.approver = employee
			
			#save it
			approval.save()

			# reload and return the appropriate version					
			return self.get( approvalId );
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval with id " + str(approvalId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(approverId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignApprover( self, approvalId ):
		errMsg = "Failed to unassign element " + str(approverId) + " for Approver on Approval"

		try:
			# get the Approval from db
			approval = self.get( approvalId ).first()	
			
			# assign to None for unassignment
			approval.employee = None			

			#save it
			approval.save()

			# reload and return the appropriate version					
			return self.get( approvalId );
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval with id " + str(approvalId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTimesheet( self, approvalId, timesheetId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TimesheetDelegate import TimesheetDelegate

		errMsg = "Failed to assign element " + str(timesheetId) + " for Timesheet on Approval"

		try:
			# get the Approval from db
			approval = self.get( approvalId ).first()	
			
			# get the Timesheet from db
			timesheet = TimesheetDelegate().get(timesheetId).first();
			
			# assign the Timesheet		
			approval.timesheet = timesheet
			
			#save it
			approval.save()

			# reload and return the appropriate version					
			return self.get( approvalId );
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval with id " + str(approvalId) + " does not exist.")
		except Timesheet.DoesNotExist:
			raise ProcessingError(errMsg + " : Timesheet with id " + str(timesheetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTimesheet( self, approvalId ):
		errMsg = "Failed to unassign element " + str(timesheetId) + " for Timesheet on Approval"

		try:
			# get the Approval from db
			approval = self.get( approvalId ).first()	
			
			# assign to None for unassignment
			approval.timesheet = None			

			#save it
			approval.save()

			# reload and return the appropriate version					
			return self.get( approvalId );
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval with id " + str(approvalId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLeaveRequest( self, approvalId, leaveRequestId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LeaveRequestDelegate import LeaveRequestDelegate

		errMsg = "Failed to assign element " + str(leaveRequestId) + " for LeaveRequest on Approval"

		try:
			# get the Approval from db
			approval = self.get( approvalId ).first()	
			
			# get the LeaveRequest from db
			leaveRequest = LeaveRequestDelegate().get(leaveRequestId).first();
			
			# assign the LeaveRequest		
			approval.leaveRequest = leaveRequest
			
			#save it
			approval.save()

			# reload and return the appropriate version					
			return self.get( approvalId );
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval with id " + str(approvalId) + " does not exist.")
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLeaveRequest( self, approvalId ):
		errMsg = "Failed to unassign element " + str(leaveRequestId) + " for LeaveRequest on Approval"

		try:
			# get the Approval from db
			approval = self.get( approvalId ).first()	
			
			# assign to None for unassignment
			approval.leaveRequest = None			

			#save it
			approval.save()

			# reload and return the appropriate version					
			return self.get( approvalId );
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval with id " + str(approvalId) + " does not exist.")
		except Exception:
			return None;
		
