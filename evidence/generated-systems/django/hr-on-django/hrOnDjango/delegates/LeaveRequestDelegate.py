from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.LeaveRequest import LeaveRequest
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.LeavePolicy import LeavePolicy
from hrOnDjango.models.Approval import Approval
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LeaveRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeaveRequestDelegate Declaration
#======================================================================
class LeaveRequestDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, leaveRequestId ):
		try:	
			leaveRequest = LeaveRequest.objects.filter(id=leaveRequestId)
			return leaveRequest.first();
		except LeaveRequest.DoesNotExist:
			raise ProcessingError("LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, leaveRequest):
		for model in serializers.deserialize("json", leaveRequest):
			model.save()
			return model;

	def create(self, leaveRequest):
		leaveRequest.save()
		return leaveRequest;

	def saveFromJson(self, leaveRequest):
		for model in serializers.deserialize("json", leaveRequest):
			model.save()
			return leaveRequest;
	
	def save(self, leaveRequest):
		leaveRequest.save()
		return leaveRequest;
	
	def delete(self, leaveRequestId ):
		errMsg = "Failed to delete LeaveRequest from db using id " + str(leaveRequestId)
		
		try:
			leaveRequest = LeaveRequest.objects.get(id=leaveRequestId)
			leaveRequest.delete()
			return True
		except LeaveRequest.DoesNotExist:
			raise ProcessingError("LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LeaveRequest.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LeaveRequest from db")
		except Exception:
			return None;
		
	def assignEmployee( self, leaveRequestId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on LeaveRequest"

		try:
			# get the LeaveRequest from db
			leaveRequest = self.get( leaveRequestId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			leaveRequest.employee = employee
			
			#save it
			leaveRequest.save()

			# reload and return the appropriate version					
			return self.get( leaveRequestId );
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, leaveRequestId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on LeaveRequest"

		try:
			# get the LeaveRequest from db
			leaveRequest = self.get( leaveRequestId ).first()	
			
			# assign to None for unassignment
			leaveRequest.employee = None			

			#save it
			leaveRequest.save()

			# reload and return the appropriate version					
			return self.get( leaveRequestId );
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLeavePolicy( self, leaveRequestId, leavePolicyId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LeavePolicyDelegate import LeavePolicyDelegate

		errMsg = "Failed to assign element " + str(leavePolicyId) + " for LeavePolicy on LeaveRequest"

		try:
			# get the LeaveRequest from db
			leaveRequest = self.get( leaveRequestId ).first()	
			
			# get the LeavePolicy from db
			leavePolicy = LeavePolicyDelegate().get(leavePolicyId).first();
			
			# assign the LeavePolicy		
			leaveRequest.leavePolicy = leavePolicy
			
			#save it
			leaveRequest.save()

			# reload and return the appropriate version					
			return self.get( leaveRequestId );
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except LeavePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : LeavePolicy with id " + str(leavePolicyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLeavePolicy( self, leaveRequestId ):
		errMsg = "Failed to unassign element " + str(leavePolicyId) + " for LeavePolicy on LeaveRequest"

		try:
			# get the LeaveRequest from db
			leaveRequest = self.get( leaveRequestId ).first()	
			
			# assign to None for unassignment
			leaveRequest.leavePolicy = None			

			#save it
			leaveRequest.save()

			# reload and return the appropriate version					
			return self.get( leaveRequestId );
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except Exception:
			return None;
		
	def addApprovals( self, leaveRequestId, approvalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ApprovalDelegate import ApprovalDelegate

		errMsg = "Failed to add elements " + str(approvalsIds) + " for Approvals on LeaveRequest"

		try:
			# get the LeaveRequest
			leaveRequest = self.get( leaveRequestId ).first()
				
			# split on a comma with no spaces
			idList = approvalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Approval		
				approval = ApprovalDelegate().get(id).first();	
				# add the Approval
				leaveRequest.approvals.add(approval)
				
			# save it		
			leaveRequest.save()
			
			# reload and return the appropriate version
			return self.get( leaveRequestId );
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApprovals( self, leaveRequestId, approvalsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.ApprovalDelegate import ApprovalDelegate

		errMsg = "Failed to remove elements " + str(approvalsIds) + " for Approvals on LeaveRequest"

		try:
			# get the LeaveRequest
			leaveRequest = self.get( leaveRequestId ).first()
				
			# split on a comma with no spaces
			idList = approvalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Approval		
				approval = ApprovalDelegate().get(id).first();	
				# add the Approval
				leaveRequest.approvals.remove(approval)
				
			# save it		
			leaveRequest.save()
			
			# reload and return the appropriate version
			return self.get( leaveRequestId );
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest with id " + str(leaveRequestId) + " does not exist.")
		except Approval.DoesNotExist:
			raise ProcessingError(errMsg + " : Approval does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
