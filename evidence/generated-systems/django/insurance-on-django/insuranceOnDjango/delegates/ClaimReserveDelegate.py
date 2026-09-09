from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.ClaimReserve import ClaimReserve
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.Exposure import Exposure
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ClaimReserve
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimReserveDelegate Declaration
#======================================================================
class ClaimReserveDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, claimReserveId ):
		try:	
			claimReserve = ClaimReserve.objects.filter(id=claimReserveId)
			return claimReserve.first();
		except ClaimReserve.DoesNotExist:
			raise ProcessingError("ClaimReserve with id " + str(claimReserveId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, claimReserve):
		for model in serializers.deserialize("json", claimReserve):
			model.save()
			return model;

	def create(self, claimReserve):
		claimReserve.save()
		return claimReserve;

	def saveFromJson(self, claimReserve):
		for model in serializers.deserialize("json", claimReserve):
			model.save()
			return claimReserve;
	
	def save(self, claimReserve):
		claimReserve.save()
		return claimReserve;
	
	def delete(self, claimReserveId ):
		errMsg = "Failed to delete ClaimReserve from db using id " + str(claimReserveId)
		
		try:
			claimReserve = ClaimReserve.objects.get(id=claimReserveId)
			claimReserve.delete()
			return True
		except ClaimReserve.DoesNotExist:
			raise ProcessingError("ClaimReserve with id " + str(claimReserveId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ClaimReserve.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ClaimReserve from db")
		except Exception:
			return None;
		
	def assignClaim( self, claimReserveId, claimId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to assign element " + str(claimId) + " for Claim on ClaimReserve"

		try:
			# get the ClaimReserve from db
			claimReserve = self.get( claimReserveId ).first()	
			
			# get the Claim from db
			claim = ClaimDelegate().get(claimId).first();
			
			# assign the Claim		
			claimReserve.claim = claim
			
			#save it
			claimReserve.save()

			# reload and return the appropriate version					
			return self.get( claimReserveId );
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve with id " + str(claimReserveId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClaim( self, claimReserveId ):
		errMsg = "Failed to unassign element " + str(claimId) + " for Claim on ClaimReserve"

		try:
			# get the ClaimReserve from db
			claimReserve = self.get( claimReserveId ).first()	
			
			# assign to None for unassignment
			claimReserve.claim = None			

			#save it
			claimReserve.save()

			# reload and return the appropriate version					
			return self.get( claimReserveId );
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve with id " + str(claimReserveId) + " does not exist.")
		except Exception:
			return None;
		
	def assignExposure( self, claimReserveId, exposureId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ExposureDelegate import ExposureDelegate

		errMsg = "Failed to assign element " + str(exposureId) + " for Exposure on ClaimReserve"

		try:
			# get the ClaimReserve from db
			claimReserve = self.get( claimReserveId ).first()	
			
			# get the Exposure from db
			exposure = ExposureDelegate().get(exposureId).first();
			
			# assign the Exposure		
			claimReserve.exposure = exposure
			
			#save it
			claimReserve.save()

			# reload and return the appropriate version					
			return self.get( claimReserveId );
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve with id " + str(claimReserveId) + " does not exist.")
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignExposure( self, claimReserveId ):
		errMsg = "Failed to unassign element " + str(exposureId) + " for Exposure on ClaimReserve"

		try:
			# get the ClaimReserve from db
			claimReserve = self.get( claimReserveId ).first()	
			
			# assign to None for unassignment
			claimReserve.exposure = None			

			#save it
			claimReserve.save()

			# reload and return the appropriate version					
			return self.get( claimReserveId );
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve with id " + str(claimReserveId) + " does not exist.")
		except Exception:
			return None;
		
