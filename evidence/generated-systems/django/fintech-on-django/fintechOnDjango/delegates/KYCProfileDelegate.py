from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.KYCProfile import KYCProfile
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.KYCDocument import KYCDocument
from fintechOnDjango.models.Screening import Screening
from fintechOnDjango.models.VerifiedAddress import VerifiedAddress
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model KYCProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCProfileDelegate Declaration
#======================================================================
class KYCProfileDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, kYCProfileId ):
		try:	
			kYCProfile = KYCProfile.objects.filter(id=kYCProfileId)
			return kYCProfile.first();
		except KYCProfile.DoesNotExist:
			raise ProcessingError("KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, kYCProfile):
		for model in serializers.deserialize("json", kYCProfile):
			model.save()
			return model;

	def create(self, kYCProfile):
		kYCProfile.save()
		return kYCProfile;

	def saveFromJson(self, kYCProfile):
		for model in serializers.deserialize("json", kYCProfile):
			model.save()
			return kYCProfile;
	
	def save(self, kYCProfile):
		kYCProfile.save()
		return kYCProfile;
	
	def delete(self, kYCProfileId ):
		errMsg = "Failed to delete KYCProfile from db using id " + str(kYCProfileId)
		
		try:
			kYCProfile = KYCProfile.objects.get(id=kYCProfileId)
			kYCProfile.delete()
			return True
		except KYCProfile.DoesNotExist:
			raise ProcessingError("KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = KYCProfile.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all KYCProfile from db")
		except Exception:
			return None;
		
	def assignCustomer( self, kYCProfileId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on KYCProfile"

		try:
			# get the KYCProfile from db
			kYCProfile = self.get( kYCProfileId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			kYCProfile.customer = customer
			
			#save it
			kYCProfile.save()

			# reload and return the appropriate version					
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, kYCProfileId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on KYCProfile"

		try:
			# get the KYCProfile from db
			kYCProfile = self.get( kYCProfileId ).first()	
			
			# assign to None for unassignment
			kYCProfile.customer = None			

			#save it
			kYCProfile.save()

			# reload and return the appropriate version					
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except Exception:
			return None;
		
	def addDocuments( self, kYCProfileId, documentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.KYCDocumentDelegate import KYCDocumentDelegate

		errMsg = "Failed to add elements " + str(documentsIds) + " for Documents on KYCProfile"

		try:
			# get the KYCProfile
			kYCProfile = self.get( kYCProfileId ).first()
				
			# split on a comma with no spaces
			idList = documentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the KYCDocument		
				kYCDocument = KYCDocumentDelegate().get(id).first();	
				# add the KYCDocument
				kYCProfile.documents.add(kYCDocument)
				
			# save it		
			kYCProfile.save()
			
			# reload and return the appropriate version
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except KYCDocument.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCDocument does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDocuments( self, kYCProfileId, documentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.KYCDocumentDelegate import KYCDocumentDelegate

		errMsg = "Failed to remove elements " + str(documentsIds) + " for Documents on KYCProfile"

		try:
			# get the KYCProfile
			kYCProfile = self.get( kYCProfileId ).first()
				
			# split on a comma with no spaces
			idList = documentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the KYCDocument		
				kYCDocument = KYCDocumentDelegate().get(id).first();	
				# add the KYCDocument
				kYCProfile.documents.remove(kYCDocument)
				
			# save it		
			kYCProfile.save()
			
			# reload and return the appropriate version
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except KYCDocument.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCDocument does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addScreenings( self, kYCProfileId, screeningsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

		errMsg = "Failed to add elements " + str(screeningsIds) + " for Screenings on KYCProfile"

		try:
			# get the KYCProfile
			kYCProfile = self.get( kYCProfileId ).first()
				
			# split on a comma with no spaces
			idList = screeningsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Screening		
				screening = ScreeningDelegate().get(id).first();	
				# add the Screening
				kYCProfile.screenings.add(screening)
				
			# save it		
			kYCProfile.save()
			
			# reload and return the appropriate version
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeScreenings( self, kYCProfileId, screeningsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ScreeningDelegate import ScreeningDelegate

		errMsg = "Failed to remove elements " + str(screeningsIds) + " for Screenings on KYCProfile"

		try:
			# get the KYCProfile
			kYCProfile = self.get( kYCProfileId ).first()
				
			# split on a comma with no spaces
			idList = screeningsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Screening		
				screening = ScreeningDelegate().get(id).first();	
				# add the Screening
				kYCProfile.screenings.remove(screening)
				
			# save it		
			kYCProfile.save()
			
			# reload and return the appropriate version
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAddresses( self, kYCProfileId, addressesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.VerifiedAddressDelegate import VerifiedAddressDelegate

		errMsg = "Failed to add elements " + str(addressesIds) + " for Addresses on KYCProfile"

		try:
			# get the KYCProfile
			kYCProfile = self.get( kYCProfileId ).first()
				
			# split on a comma with no spaces
			idList = addressesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the VerifiedAddress		
				verifiedAddress = VerifiedAddressDelegate().get(id).first();	
				# add the VerifiedAddress
				kYCProfile.addresses.add(verifiedAddress)
				
			# save it		
			kYCProfile.save()
			
			# reload and return the appropriate version
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except VerifiedAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : VerifiedAddress does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAddresses( self, kYCProfileId, addressesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.VerifiedAddressDelegate import VerifiedAddressDelegate

		errMsg = "Failed to remove elements " + str(addressesIds) + " for Addresses on KYCProfile"

		try:
			# get the KYCProfile
			kYCProfile = self.get( kYCProfileId ).first()
				
			# split on a comma with no spaces
			idList = addressesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the VerifiedAddress		
				verifiedAddress = VerifiedAddressDelegate().get(id).first();	
				# add the VerifiedAddress
				kYCProfile.addresses.remove(verifiedAddress)
				
			# save it		
			kYCProfile.save()
			
			# reload and return the appropriate version
			return self.get( kYCProfileId );
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kYCProfileId) + " does not exist.")
		except VerifiedAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : VerifiedAddress does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
