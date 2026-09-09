from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Endorsement import Endorsement
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Endorsement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EndorsementDelegate Declaration
#======================================================================
class EndorsementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, endorsementId ):
		try:	
			endorsement = Endorsement.objects.filter(id=endorsementId)
			return endorsement.first();
		except Endorsement.DoesNotExist:
			raise ProcessingError("Endorsement with id " + str(endorsementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, endorsement):
		for model in serializers.deserialize("json", endorsement):
			model.save()
			return model;

	def create(self, endorsement):
		endorsement.save()
		return endorsement;

	def saveFromJson(self, endorsement):
		for model in serializers.deserialize("json", endorsement):
			model.save()
			return endorsement;
	
	def save(self, endorsement):
		endorsement.save()
		return endorsement;
	
	def delete(self, endorsementId ):
		errMsg = "Failed to delete Endorsement from db using id " + str(endorsementId)
		
		try:
			endorsement = Endorsement.objects.get(id=endorsementId)
			endorsement.delete()
			return True
		except Endorsement.DoesNotExist:
			raise ProcessingError("Endorsement with id " + str(endorsementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Endorsement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Endorsement from db")
		except Exception:
			return None;
		
	def assignPolicy( self, endorsementId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Endorsement"

		try:
			# get the Endorsement from db
			endorsement = self.get( endorsementId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			endorsement.policy = policy
			
			#save it
			endorsement.save()

			# reload and return the appropriate version					
			return self.get( endorsementId );
		except Endorsement.DoesNotExist:
			raise ProcessingError(errMsg + " : Endorsement with id " + str(endorsementId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, endorsementId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Endorsement"

		try:
			# get the Endorsement from db
			endorsement = self.get( endorsementId ).first()	
			
			# assign to None for unassignment
			endorsement.policy = None			

			#save it
			endorsement.save()

			# reload and return the appropriate version					
			return self.get( endorsementId );
		except Endorsement.DoesNotExist:
			raise ProcessingError(errMsg + " : Endorsement with id " + str(endorsementId) + " does not exist.")
		except Exception:
			return None;
		
