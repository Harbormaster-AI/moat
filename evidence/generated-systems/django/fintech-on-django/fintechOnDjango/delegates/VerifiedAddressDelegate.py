from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.VerifiedAddress import VerifiedAddress
from fintechOnDjango.models.KYCProfile import KYCProfile
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model VerifiedAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VerifiedAddressDelegate Declaration
#======================================================================
class VerifiedAddressDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, verifiedAddressId ):
		try:	
			verifiedAddress = VerifiedAddress.objects.filter(id=verifiedAddressId)
			return verifiedAddress.first();
		except VerifiedAddress.DoesNotExist:
			raise ProcessingError("VerifiedAddress with id " + str(verifiedAddressId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, verifiedAddress):
		for model in serializers.deserialize("json", verifiedAddress):
			model.save()
			return model;

	def create(self, verifiedAddress):
		verifiedAddress.save()
		return verifiedAddress;

	def saveFromJson(self, verifiedAddress):
		for model in serializers.deserialize("json", verifiedAddress):
			model.save()
			return verifiedAddress;
	
	def save(self, verifiedAddress):
		verifiedAddress.save()
		return verifiedAddress;
	
	def delete(self, verifiedAddressId ):
		errMsg = "Failed to delete VerifiedAddress from db using id " + str(verifiedAddressId)
		
		try:
			verifiedAddress = VerifiedAddress.objects.get(id=verifiedAddressId)
			verifiedAddress.delete()
			return True
		except VerifiedAddress.DoesNotExist:
			raise ProcessingError("VerifiedAddress with id " + str(verifiedAddressId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = VerifiedAddress.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all VerifiedAddress from db")
		except Exception:
			return None;
		
	def assignKycProfile( self, verifiedAddressId, kycProfileId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.KYCProfileDelegate import KYCProfileDelegate

		errMsg = "Failed to assign element " + str(kycProfileId) + " for KycProfile on VerifiedAddress"

		try:
			# get the VerifiedAddress from db
			verifiedAddress = self.get( verifiedAddressId ).first()	
			
			# get the KYCProfile from db
			kYCProfile = KYCProfileDelegate().get(kycProfileId).first();
			
			# assign the KycProfile		
			verifiedAddress.kycProfile = kYCProfile
			
			#save it
			verifiedAddress.save()

			# reload and return the appropriate version					
			return self.get( verifiedAddressId );
		except VerifiedAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : VerifiedAddress with id " + str(verifiedAddressId) + " does not exist.")
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kycProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignKycProfile( self, verifiedAddressId ):
		errMsg = "Failed to unassign element " + str(kycProfileId) + " for KycProfile on VerifiedAddress"

		try:
			# get the VerifiedAddress from db
			verifiedAddress = self.get( verifiedAddressId ).first()	
			
			# assign to None for unassignment
			verifiedAddress.kYCProfile = None			

			#save it
			verifiedAddress.save()

			# reload and return the appropriate version					
			return self.get( verifiedAddressId );
		except VerifiedAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : VerifiedAddress with id " + str(verifiedAddressId) + " does not exist.")
		except Exception:
			return None;
		
