from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Adjuster import Adjuster
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.ServiceProvider import ServiceProvider
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Adjuster
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdjusterDelegate Declaration
#======================================================================
class AdjusterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, adjusterId ):
		try:	
			adjuster = Adjuster.objects.filter(id=adjusterId)
			return adjuster.first();
		except Adjuster.DoesNotExist:
			raise ProcessingError("Adjuster with id " + str(adjusterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, adjuster):
		for model in serializers.deserialize("json", adjuster):
			model.save()
			return model;

	def create(self, adjuster):
		adjuster.save()
		return adjuster;

	def saveFromJson(self, adjuster):
		for model in serializers.deserialize("json", adjuster):
			model.save()
			return adjuster;
	
	def save(self, adjuster):
		adjuster.save()
		return adjuster;
	
	def delete(self, adjusterId ):
		errMsg = "Failed to delete Adjuster from db using id " + str(adjusterId)
		
		try:
			adjuster = Adjuster.objects.get(id=adjusterId)
			adjuster.delete()
			return True
		except Adjuster.DoesNotExist:
			raise ProcessingError("Adjuster with id " + str(adjusterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Adjuster.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Adjuster from db")
		except Exception:
			return None;
		
	def addClaims( self, adjusterId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on Adjuster"

		try:
			# get the Adjuster
			adjuster = self.get( adjusterId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				adjuster.claims.add(claim)
				
			# save it		
			adjuster.save()
			
			# reload and return the appropriate version
			return self.get( adjusterId );
		except Adjuster.DoesNotExist:
			raise ProcessingError(errMsg + " : Adjuster with id " + str(adjusterId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, adjusterId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on Adjuster"

		try:
			# get the Adjuster
			adjuster = self.get( adjusterId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				adjuster.claims.remove(claim)
				
			# save it		
			adjuster.save()
			
			# reload and return the appropriate version
			return self.get( adjusterId );
		except Adjuster.DoesNotExist:
			raise ProcessingError(errMsg + " : Adjuster with id " + str(adjusterId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addServiceProviders( self, adjusterId, serviceProvidersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ServiceProviderDelegate import ServiceProviderDelegate

		errMsg = "Failed to add elements " + str(serviceProvidersIds) + " for ServiceProviders on Adjuster"

		try:
			# get the Adjuster
			adjuster = self.get( adjusterId ).first()
				
			# split on a comma with no spaces
			idList = serviceProvidersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ServiceProvider		
				serviceProvider = ServiceProviderDelegate().get(id).first();	
				# add the ServiceProvider
				adjuster.serviceProviders.add(serviceProvider)
				
			# save it		
			adjuster.save()
			
			# reload and return the appropriate version
			return self.get( adjusterId );
		except Adjuster.DoesNotExist:
			raise ProcessingError(errMsg + " : Adjuster with id " + str(adjusterId) + " does not exist.")
		except ServiceProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceProvider does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeServiceProviders( self, adjusterId, serviceProvidersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ServiceProviderDelegate import ServiceProviderDelegate

		errMsg = "Failed to remove elements " + str(serviceProvidersIds) + " for ServiceProviders on Adjuster"

		try:
			# get the Adjuster
			adjuster = self.get( adjusterId ).first()
				
			# split on a comma with no spaces
			idList = serviceProvidersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ServiceProvider		
				serviceProvider = ServiceProviderDelegate().get(id).first();	
				# add the ServiceProvider
				adjuster.serviceProviders.remove(serviceProvider)
				
			# save it		
			adjuster.save()
			
			# reload and return the appropriate version
			return self.get( adjusterId );
		except Adjuster.DoesNotExist:
			raise ProcessingError(errMsg + " : Adjuster with id " + str(adjusterId) + " does not exist.")
		except ServiceProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceProvider does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
