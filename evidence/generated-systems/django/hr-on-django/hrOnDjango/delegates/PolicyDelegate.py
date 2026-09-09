from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Policy import Policy
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.PolicyAcknowledgement import PolicyAcknowledgement
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyDelegate Declaration
#======================================================================
class PolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, policyId ):
		try:	
			policy = Policy.objects.filter(id=policyId)
			return policy.first();
		except Policy.DoesNotExist:
			raise ProcessingError("Policy with id " + str(policyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, policy):
		for model in serializers.deserialize("json", policy):
			model.save()
			return model;

	def create(self, policy):
		policy.save()
		return policy;

	def saveFromJson(self, policy):
		for model in serializers.deserialize("json", policy):
			model.save()
			return policy;
	
	def save(self, policy):
		policy.save()
		return policy;
	
	def delete(self, policyId ):
		errMsg = "Failed to delete Policy from db using id " + str(policyId)
		
		try:
			policy = Policy.objects.get(id=policyId)
			policy.delete()
			return True
		except Policy.DoesNotExist:
			raise ProcessingError("Policy with id " + str(policyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Policy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Policy from db")
		except Exception:
			return None;
		
	def assignOrganization( self, policyId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			policy.organization = organization
			
			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, policyId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# assign to None for unassignment
			policy.organization = None			

			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
		
	def addAcknowledgements( self, policyId, acknowledgementsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PolicyAcknowledgementDelegate import PolicyAcknowledgementDelegate

		errMsg = "Failed to add elements " + str(acknowledgementsIds) + " for Acknowledgements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = acknowledgementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PolicyAcknowledgement		
				policyAcknowledgement = PolicyAcknowledgementDelegate().get(id).first();	
				# add the PolicyAcknowledgement
				policy.acknowledgements.add(policyAcknowledgement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyAcknowledgement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAcknowledgements( self, policyId, acknowledgementsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PolicyAcknowledgementDelegate import PolicyAcknowledgementDelegate

		errMsg = "Failed to remove elements " + str(acknowledgementsIds) + " for Acknowledgements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = acknowledgementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PolicyAcknowledgement		
				policyAcknowledgement = PolicyAcknowledgementDelegate().get(id).first();	
				# add the PolicyAcknowledgement
				policy.acknowledgements.remove(policyAcknowledgement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyAcknowledgement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
