from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.APIClient import APIClient
from fintechOnDjango.models.Consent import Consent
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model APIClient
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APIClientDelegate Declaration
#======================================================================
class APIClientDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aPIClientId ):
		try:	
			aPIClient = APIClient.objects.filter(id=aPIClientId)
			return aPIClient.first();
		except APIClient.DoesNotExist:
			raise ProcessingError("APIClient with id " + str(aPIClientId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aPIClient):
		for model in serializers.deserialize("json", aPIClient):
			model.save()
			return model;

	def create(self, aPIClient):
		aPIClient.save()
		return aPIClient;

	def saveFromJson(self, aPIClient):
		for model in serializers.deserialize("json", aPIClient):
			model.save()
			return aPIClient;
	
	def save(self, aPIClient):
		aPIClient.save()
		return aPIClient;
	
	def delete(self, aPIClientId ):
		errMsg = "Failed to delete APIClient from db using id " + str(aPIClientId)
		
		try:
			aPIClient = APIClient.objects.get(id=aPIClientId)
			aPIClient.delete()
			return True
		except APIClient.DoesNotExist:
			raise ProcessingError("APIClient with id " + str(aPIClientId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = APIClient.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all APIClient from db")
		except Exception:
			return None;
		
	def addConsents( self, aPIClientId, consentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to add elements " + str(consentsIds) + " for Consents on APIClient"

		try:
			# get the APIClient
			aPIClient = self.get( aPIClientId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				aPIClient.consents.add(consent)
				
			# save it		
			aPIClient.save()
			
			# reload and return the appropriate version
			return self.get( aPIClientId );
		except APIClient.DoesNotExist:
			raise ProcessingError(errMsg + " : APIClient with id " + str(aPIClientId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeConsents( self, aPIClientId, consentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to remove elements " + str(consentsIds) + " for Consents on APIClient"

		try:
			# get the APIClient
			aPIClient = self.get( aPIClientId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				aPIClient.consents.remove(consent)
				
			# save it		
			aPIClient.save()
			
			# reload and return the appropriate version
			return self.get( aPIClientId );
		except APIClient.DoesNotExist:
			raise ProcessingError(errMsg + " : APIClient with id " + str(aPIClientId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
