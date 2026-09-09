from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.GovernanceBody import GovernanceBody
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.RoleAssignment import RoleAssignment
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model GovernanceBody
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GovernanceBodyDelegate Declaration
#======================================================================
class GovernanceBodyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, governanceBodyId ):
		try:	
			governanceBody = GovernanceBody.objects.filter(id=governanceBodyId)
			return governanceBody.first();
		except GovernanceBody.DoesNotExist:
			raise ProcessingError("GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, governanceBody):
		for model in serializers.deserialize("json", governanceBody):
			model.save()
			return model;

	def create(self, governanceBody):
		governanceBody.save()
		return governanceBody;

	def saveFromJson(self, governanceBody):
		for model in serializers.deserialize("json", governanceBody):
			model.save()
			return governanceBody;
	
	def save(self, governanceBody):
		governanceBody.save()
		return governanceBody;
	
	def delete(self, governanceBodyId ):
		errMsg = "Failed to delete GovernanceBody from db using id " + str(governanceBodyId)
		
		try:
			governanceBody = GovernanceBody.objects.get(id=governanceBodyId)
			governanceBody.delete()
			return True
		except GovernanceBody.DoesNotExist:
			raise ProcessingError("GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = GovernanceBody.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all GovernanceBody from db")
		except Exception:
			return None;
		
	def assignOrganization( self, governanceBodyId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on GovernanceBody"

		try:
			# get the GovernanceBody from db
			governanceBody = self.get( governanceBodyId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			governanceBody.organization = organization
			
			#save it
			governanceBody.save()

			# reload and return the appropriate version					
			return self.get( governanceBodyId );
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, governanceBodyId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on GovernanceBody"

		try:
			# get the GovernanceBody from db
			governanceBody = self.get( governanceBodyId ).first()	
			
			# assign to None for unassignment
			governanceBody.organization = None			

			#save it
			governanceBody.save()

			# reload and return the appropriate version					
			return self.get( governanceBodyId );
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except Exception:
			return None;
		
	def addRoleAssignments( self, governanceBodyId, roleAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

		errMsg = "Failed to add elements " + str(roleAssignmentsIds) + " for RoleAssignments on GovernanceBody"

		try:
			# get the GovernanceBody
			governanceBody = self.get( governanceBodyId ).first()
				
			# split on a comma with no spaces
			idList = roleAssignmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RoleAssignment		
				roleAssignment = RoleAssignmentDelegate().get(id).first();	
				# add the RoleAssignment
				governanceBody.roleAssignments.add(roleAssignment)
				
			# save it		
			governanceBody.save()
			
			# reload and return the appropriate version
			return self.get( governanceBodyId );
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRoleAssignments( self, governanceBodyId, roleAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

		errMsg = "Failed to remove elements " + str(roleAssignmentsIds) + " for RoleAssignments on GovernanceBody"

		try:
			# get the GovernanceBody
			governanceBody = self.get( governanceBodyId ).first()
				
			# split on a comma with no spaces
			idList = roleAssignmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RoleAssignment		
				roleAssignment = RoleAssignmentDelegate().get(id).first();	
				# add the RoleAssignment
				governanceBody.roleAssignments.remove(roleAssignment)
				
			# save it		
			governanceBody.save()
			
			# reload and return the appropriate version
			return self.get( governanceBodyId );
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPolicies( self, governanceBodyId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on GovernanceBody"

		try:
			# get the GovernanceBody
			governanceBody = self.get( governanceBodyId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				governanceBody.policies.add(policy)
				
			# save it		
			governanceBody.save()
			
			# reload and return the appropriate version
			return self.get( governanceBodyId );
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, governanceBodyId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on GovernanceBody"

		try:
			# get the GovernanceBody
			governanceBody = self.get( governanceBodyId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				governanceBody.policies.remove(policy)
				
			# save it		
			governanceBody.save()
			
			# reload and return the appropriate version
			return self.get( governanceBodyId );
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
