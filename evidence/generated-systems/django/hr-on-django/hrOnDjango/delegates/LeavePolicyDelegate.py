from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.LeavePolicy import LeavePolicy
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.LeaveRequest import LeaveRequest
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LeavePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeavePolicyDelegate Declaration
#======================================================================
class LeavePolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, leavePolicyId ):
		try:	
			leavePolicy = LeavePolicy.objects.filter(id=leavePolicyId)
			return leavePolicy.first();
		except LeavePolicy.DoesNotExist:
			raise ProcessingError("LeavePolicy with id " + str(leavePolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, leavePolicy):
		for model in serializers.deserialize("json", leavePolicy):
			model.save()
			return model;

	def create(self, leavePolicy):
		leavePolicy.save()
		return leavePolicy;

	def saveFromJson(self, leavePolicy):
		for model in serializers.deserialize("json", leavePolicy):
			model.save()
			return leavePolicy;
	
	def save(self, leavePolicy):
		leavePolicy.save()
		return leavePolicy;
	
	def delete(self, leavePolicyId ):
		errMsg = "Failed to delete LeavePolicy from db using id " + str(leavePolicyId)
		
		try:
			leavePolicy = LeavePolicy.objects.get(id=leavePolicyId)
			leavePolicy.delete()
			return True
		except LeavePolicy.DoesNotExist:
			raise ProcessingError("LeavePolicy with id " + str(leavePolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LeavePolicy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LeavePolicy from db")
		except Exception:
			return None;
		
	def assignOrganization( self, leavePolicyId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on LeavePolicy"

		try:
			# get the LeavePolicy from db
			leavePolicy = self.get( leavePolicyId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			leavePolicy.organization = organization
			
			#save it
			leavePolicy.save()

			# reload and return the appropriate version					
			return self.get( leavePolicyId );
		except LeavePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : LeavePolicy with id " + str(leavePolicyId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, leavePolicyId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on LeavePolicy"

		try:
			# get the LeavePolicy from db
			leavePolicy = self.get( leavePolicyId ).first()	
			
			# assign to None for unassignment
			leavePolicy.organization = None			

			#save it
			leavePolicy.save()

			# reload and return the appropriate version					
			return self.get( leavePolicyId );
		except LeavePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : LeavePolicy with id " + str(leavePolicyId) + " does not exist.")
		except Exception:
			return None;
		
	def addLeaveRequests( self, leavePolicyId, leaveRequestsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LeaveRequestDelegate import LeaveRequestDelegate

		errMsg = "Failed to add elements " + str(leaveRequestsIds) + " for LeaveRequests on LeavePolicy"

		try:
			# get the LeavePolicy
			leavePolicy = self.get( leavePolicyId ).first()
				
			# split on a comma with no spaces
			idList = leaveRequestsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LeaveRequest		
				leaveRequest = LeaveRequestDelegate().get(id).first();	
				# add the LeaveRequest
				leavePolicy.leaveRequests.add(leaveRequest)
				
			# save it		
			leavePolicy.save()
			
			# reload and return the appropriate version
			return self.get( leavePolicyId );
		except LeavePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : LeavePolicy with id " + str(leavePolicyId) + " does not exist.")
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLeaveRequests( self, leavePolicyId, leaveRequestsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LeaveRequestDelegate import LeaveRequestDelegate

		errMsg = "Failed to remove elements " + str(leaveRequestsIds) + " for LeaveRequests on LeavePolicy"

		try:
			# get the LeavePolicy
			leavePolicy = self.get( leavePolicyId ).first()
				
			# split on a comma with no spaces
			idList = leaveRequestsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LeaveRequest		
				leaveRequest = LeaveRequestDelegate().get(id).first();	
				# add the LeaveRequest
				leavePolicy.leaveRequests.remove(leaveRequest)
				
			# save it		
			leavePolicy.save()
			
			# reload and return the appropriate version
			return self.get( leavePolicyId );
		except LeavePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : LeavePolicy with id " + str(leavePolicyId) + " does not exist.")
		except LeaveRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : LeaveRequest does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
