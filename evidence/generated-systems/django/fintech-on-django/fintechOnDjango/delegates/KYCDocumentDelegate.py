from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.KYCDocument import KYCDocument
from fintechOnDjango.models.KYCProfile import KYCProfile
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model KYCDocument
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCDocumentDelegate Declaration
#======================================================================
class KYCDocumentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, kYCDocumentId ):
		try:	
			kYCDocument = KYCDocument.objects.filter(id=kYCDocumentId)
			return kYCDocument.first();
		except KYCDocument.DoesNotExist:
			raise ProcessingError("KYCDocument with id " + str(kYCDocumentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, kYCDocument):
		for model in serializers.deserialize("json", kYCDocument):
			model.save()
			return model;

	def create(self, kYCDocument):
		kYCDocument.save()
		return kYCDocument;

	def saveFromJson(self, kYCDocument):
		for model in serializers.deserialize("json", kYCDocument):
			model.save()
			return kYCDocument;
	
	def save(self, kYCDocument):
		kYCDocument.save()
		return kYCDocument;
	
	def delete(self, kYCDocumentId ):
		errMsg = "Failed to delete KYCDocument from db using id " + str(kYCDocumentId)
		
		try:
			kYCDocument = KYCDocument.objects.get(id=kYCDocumentId)
			kYCDocument.delete()
			return True
		except KYCDocument.DoesNotExist:
			raise ProcessingError("KYCDocument with id " + str(kYCDocumentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = KYCDocument.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all KYCDocument from db")
		except Exception:
			return None;
		
	def assignKycProfile( self, kYCDocumentId, kycProfileId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.KYCProfileDelegate import KYCProfileDelegate

		errMsg = "Failed to assign element " + str(kycProfileId) + " for KycProfile on KYCDocument"

		try:
			# get the KYCDocument from db
			kYCDocument = self.get( kYCDocumentId ).first()	
			
			# get the KYCProfile from db
			kYCProfile = KYCProfileDelegate().get(kycProfileId).first();
			
			# assign the KycProfile		
			kYCDocument.kycProfile = kYCProfile
			
			#save it
			kYCDocument.save()

			# reload and return the appropriate version					
			return self.get( kYCDocumentId );
		except KYCDocument.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCDocument with id " + str(kYCDocumentId) + " does not exist.")
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kycProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignKycProfile( self, kYCDocumentId ):
		errMsg = "Failed to unassign element " + str(kycProfileId) + " for KycProfile on KYCDocument"

		try:
			# get the KYCDocument from db
			kYCDocument = self.get( kYCDocumentId ).first()	
			
			# assign to None for unassignment
			kYCDocument.kYCProfile = None			

			#save it
			kYCDocument.save()

			# reload and return the appropriate version					
			return self.get( kYCDocumentId );
		except KYCDocument.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCDocument with id " + str(kYCDocumentId) + " does not exist.")
		except Exception:
			return None;
		
