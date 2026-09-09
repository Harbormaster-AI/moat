from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.SubrogationRecovery import SubrogationRecovery
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.Exposure import Exposure
from insuranceOnDjango.models.ThirdParty import ThirdParty
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SubrogationRecovery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubrogationRecoveryDelegate Declaration
#======================================================================
class SubrogationRecoveryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, subrogationRecoveryId ):
		try:	
			subrogationRecovery = SubrogationRecovery.objects.filter(id=subrogationRecoveryId)
			return subrogationRecovery.first();
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError("SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, subrogationRecovery):
		for model in serializers.deserialize("json", subrogationRecovery):
			model.save()
			return model;

	def create(self, subrogationRecovery):
		subrogationRecovery.save()
		return subrogationRecovery;

	def saveFromJson(self, subrogationRecovery):
		for model in serializers.deserialize("json", subrogationRecovery):
			model.save()
			return subrogationRecovery;
	
	def save(self, subrogationRecovery):
		subrogationRecovery.save()
		return subrogationRecovery;
	
	def delete(self, subrogationRecoveryId ):
		errMsg = "Failed to delete SubrogationRecovery from db using id " + str(subrogationRecoveryId)
		
		try:
			subrogationRecovery = SubrogationRecovery.objects.get(id=subrogationRecoveryId)
			subrogationRecovery.delete()
			return True
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError("SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SubrogationRecovery.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SubrogationRecovery from db")
		except Exception:
			return None;
		
	def assignClaim( self, subrogationRecoveryId, claimId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to assign element " + str(claimId) + " for Claim on SubrogationRecovery"

		try:
			# get the SubrogationRecovery from db
			subrogationRecovery = self.get( subrogationRecoveryId ).first()	
			
			# get the Claim from db
			claim = ClaimDelegate().get(claimId).first();
			
			# assign the Claim		
			subrogationRecovery.claim = claim
			
			#save it
			subrogationRecovery.save()

			# reload and return the appropriate version					
			return self.get( subrogationRecoveryId );
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClaim( self, subrogationRecoveryId ):
		errMsg = "Failed to unassign element " + str(claimId) + " for Claim on SubrogationRecovery"

		try:
			# get the SubrogationRecovery from db
			subrogationRecovery = self.get( subrogationRecoveryId ).first()	
			
			# assign to None for unassignment
			subrogationRecovery.claim = None			

			#save it
			subrogationRecovery.save()

			# reload and return the appropriate version					
			return self.get( subrogationRecoveryId );
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignExposure( self, subrogationRecoveryId, exposureId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ExposureDelegate import ExposureDelegate

		errMsg = "Failed to assign element " + str(exposureId) + " for Exposure on SubrogationRecovery"

		try:
			# get the SubrogationRecovery from db
			subrogationRecovery = self.get( subrogationRecoveryId ).first()	
			
			# get the Exposure from db
			exposure = ExposureDelegate().get(exposureId).first();
			
			# assign the Exposure		
			subrogationRecovery.exposure = exposure
			
			#save it
			subrogationRecovery.save()

			# reload and return the appropriate version					
			return self.get( subrogationRecoveryId );
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignExposure( self, subrogationRecoveryId ):
		errMsg = "Failed to unassign element " + str(exposureId) + " for Exposure on SubrogationRecovery"

		try:
			# get the SubrogationRecovery from db
			subrogationRecovery = self.get( subrogationRecoveryId ).first()	
			
			# assign to None for unassignment
			subrogationRecovery.exposure = None			

			#save it
			subrogationRecovery.save()

			# reload and return the appropriate version					
			return self.get( subrogationRecoveryId );
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCounterparty( self, subrogationRecoveryId, counterpartyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to assign element " + str(counterpartyId) + " for Counterparty on SubrogationRecovery"

		try:
			# get the SubrogationRecovery from db
			subrogationRecovery = self.get( subrogationRecoveryId ).first()	
			
			# get the ThirdParty from db
			thirdParty = ThirdPartyDelegate().get(counterpartyId).first();
			
			# assign the Counterparty		
			subrogationRecovery.counterparty = thirdParty
			
			#save it
			subrogationRecovery.save()

			# reload and return the appropriate version					
			return self.get( subrogationRecoveryId );
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(counterpartyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCounterparty( self, subrogationRecoveryId ):
		errMsg = "Failed to unassign element " + str(counterpartyId) + " for Counterparty on SubrogationRecovery"

		try:
			# get the SubrogationRecovery from db
			subrogationRecovery = self.get( subrogationRecoveryId ).first()	
			
			# assign to None for unassignment
			subrogationRecovery.thirdParty = None			

			#save it
			subrogationRecovery.save()

			# reload and return the appropriate version					
			return self.get( subrogationRecoveryId );
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery with id " + str(subrogationRecoveryId) + " does not exist.")
		except Exception:
			return None;
		
