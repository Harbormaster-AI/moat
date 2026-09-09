from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.ReinsuranceAgreement import ReinsuranceAgreement
from insuranceOnDjango.models.Insurer import Insurer
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ReinsuranceAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReinsuranceAgreementDelegate Declaration
#======================================================================
class ReinsuranceAgreementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, reinsuranceAgreementId ):
		try:	
			reinsuranceAgreement = ReinsuranceAgreement.objects.filter(id=reinsuranceAgreementId)
			return reinsuranceAgreement.first();
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError("ReinsuranceAgreement with id " + str(reinsuranceAgreementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, reinsuranceAgreement):
		for model in serializers.deserialize("json", reinsuranceAgreement):
			model.save()
			return model;

	def create(self, reinsuranceAgreement):
		reinsuranceAgreement.save()
		return reinsuranceAgreement;

	def saveFromJson(self, reinsuranceAgreement):
		for model in serializers.deserialize("json", reinsuranceAgreement):
			model.save()
			return reinsuranceAgreement;
	
	def save(self, reinsuranceAgreement):
		reinsuranceAgreement.save()
		return reinsuranceAgreement;
	
	def delete(self, reinsuranceAgreementId ):
		errMsg = "Failed to delete ReinsuranceAgreement from db using id " + str(reinsuranceAgreementId)
		
		try:
			reinsuranceAgreement = ReinsuranceAgreement.objects.get(id=reinsuranceAgreementId)
			reinsuranceAgreement.delete()
			return True
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError("ReinsuranceAgreement with id " + str(reinsuranceAgreementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ReinsuranceAgreement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ReinsuranceAgreement from db")
		except Exception:
			return None;
		
	def assignInsurer( self, reinsuranceAgreementId, insurerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

		errMsg = "Failed to assign element " + str(insurerId) + " for Insurer on ReinsuranceAgreement"

		try:
			# get the ReinsuranceAgreement from db
			reinsuranceAgreement = self.get( reinsuranceAgreementId ).first()	
			
			# get the Insurer from db
			insurer = InsurerDelegate().get(insurerId).first();
			
			# assign the Insurer		
			reinsuranceAgreement.insurer = insurer
			
			#save it
			reinsuranceAgreement.save()

			# reload and return the appropriate version					
			return self.get( reinsuranceAgreementId );
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement with id " + str(reinsuranceAgreementId) + " does not exist.")
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInsurer( self, reinsuranceAgreementId ):
		errMsg = "Failed to unassign element " + str(insurerId) + " for Insurer on ReinsuranceAgreement"

		try:
			# get the ReinsuranceAgreement from db
			reinsuranceAgreement = self.get( reinsuranceAgreementId ).first()	
			
			# assign to None for unassignment
			reinsuranceAgreement.insurer = None			

			#save it
			reinsuranceAgreement.save()

			# reload and return the appropriate version					
			return self.get( reinsuranceAgreementId );
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement with id " + str(reinsuranceAgreementId) + " does not exist.")
		except Exception:
			return None;
		
	def addPolicies( self, reinsuranceAgreementId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on ReinsuranceAgreement"

		try:
			# get the ReinsuranceAgreement
			reinsuranceAgreement = self.get( reinsuranceAgreementId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				reinsuranceAgreement.policies.add(policy)
				
			# save it		
			reinsuranceAgreement.save()
			
			# reload and return the appropriate version
			return self.get( reinsuranceAgreementId );
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement with id " + str(reinsuranceAgreementId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, reinsuranceAgreementId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on ReinsuranceAgreement"

		try:
			# get the ReinsuranceAgreement
			reinsuranceAgreement = self.get( reinsuranceAgreementId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				reinsuranceAgreement.policies.remove(policy)
				
			# save it		
			reinsuranceAgreement.save()
			
			# reload and return the appropriate version
			return self.get( reinsuranceAgreementId );
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement with id " + str(reinsuranceAgreementId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
