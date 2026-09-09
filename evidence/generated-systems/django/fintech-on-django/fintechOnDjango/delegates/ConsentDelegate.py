from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Consent import Consent
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.APIClient import APIClient
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Consent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConsentDelegate Declaration
#======================================================================
class ConsentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, consentId ):
		try:	
			consent = Consent.objects.filter(id=consentId)
			return consent.first();
		except Consent.DoesNotExist:
			raise ProcessingError("Consent with id " + str(consentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, consent):
		for model in serializers.deserialize("json", consent):
			model.save()
			return model;

	def create(self, consent):
		consent.save()
		return consent;

	def saveFromJson(self, consent):
		for model in serializers.deserialize("json", consent):
			model.save()
			return consent;
	
	def save(self, consent):
		consent.save()
		return consent;
	
	def delete(self, consentId ):
		errMsg = "Failed to delete Consent from db using id " + str(consentId)
		
		try:
			consent = Consent.objects.get(id=consentId)
			consent.delete()
			return True
		except Consent.DoesNotExist:
			raise ProcessingError("Consent with id " + str(consentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Consent.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Consent from db")
		except Exception:
			return None;
		
	def assignCustomer( self, consentId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Consent"

		try:
			# get the Consent from db
			consent = self.get( consentId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			consent.customer = customer
			
			#save it
			consent.save()

			# reload and return the appropriate version					
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, consentId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Consent"

		try:
			# get the Consent from db
			consent = self.get( consentId ).first()	
			
			# assign to None for unassignment
			consent.customer = None			

			#save it
			consent.save()

			# reload and return the appropriate version					
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignApiClient( self, consentId, apiClientId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.APIClientDelegate import APIClientDelegate

		errMsg = "Failed to assign element " + str(apiClientId) + " for ApiClient on Consent"

		try:
			# get the Consent from db
			consent = self.get( consentId ).first()	
			
			# get the APIClient from db
			aPIClient = APIClientDelegate().get(apiClientId).first();
			
			# assign the ApiClient		
			consent.apiClient = aPIClient
			
			#save it
			consent.save()

			# reload and return the appropriate version					
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except APIClient.DoesNotExist:
			raise ProcessingError(errMsg + " : APIClient with id " + str(apiClientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignApiClient( self, consentId ):
		errMsg = "Failed to unassign element " + str(apiClientId) + " for ApiClient on Consent"

		try:
			# get the Consent from db
			consent = self.get( consentId ).first()	
			
			# assign to None for unassignment
			consent.aPIClient = None			

			#save it
			consent.save()

			# reload and return the appropriate version					
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except Exception:
			return None;
		
