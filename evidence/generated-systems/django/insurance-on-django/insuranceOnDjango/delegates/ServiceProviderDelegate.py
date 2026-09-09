from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.ServiceProvider import ServiceProvider
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ServiceProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceProviderDelegate Declaration
#======================================================================
class ServiceProviderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, serviceProviderId ):
		try:	
			serviceProvider = ServiceProvider.objects.filter(id=serviceProviderId)
			return serviceProvider.first();
		except ServiceProvider.DoesNotExist:
			raise ProcessingError("ServiceProvider with id " + str(serviceProviderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, serviceProvider):
		for model in serializers.deserialize("json", serviceProvider):
			model.save()
			return model;

	def create(self, serviceProvider):
		serviceProvider.save()
		return serviceProvider;

	def saveFromJson(self, serviceProvider):
		for model in serializers.deserialize("json", serviceProvider):
			model.save()
			return serviceProvider;
	
	def save(self, serviceProvider):
		serviceProvider.save()
		return serviceProvider;
	
	def delete(self, serviceProviderId ):
		errMsg = "Failed to delete ServiceProvider from db using id " + str(serviceProviderId)
		
		try:
			serviceProvider = ServiceProvider.objects.get(id=serviceProviderId)
			serviceProvider.delete()
			return True
		except ServiceProvider.DoesNotExist:
			raise ProcessingError("ServiceProvider with id " + str(serviceProviderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ServiceProvider.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ServiceProvider from db")
		except Exception:
			return None;
		
	def addClaims( self, serviceProviderId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on ServiceProvider"

		try:
			# get the ServiceProvider
			serviceProvider = self.get( serviceProviderId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				serviceProvider.claims.add(claim)
				
			# save it		
			serviceProvider.save()
			
			# reload and return the appropriate version
			return self.get( serviceProviderId );
		except ServiceProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceProvider with id " + str(serviceProviderId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, serviceProviderId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on ServiceProvider"

		try:
			# get the ServiceProvider
			serviceProvider = self.get( serviceProviderId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				serviceProvider.claims.remove(claim)
				
			# save it		
			serviceProvider.save()
			
			# reload and return the appropriate version
			return self.get( serviceProviderId );
		except ServiceProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceProvider with id " + str(serviceProviderId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
