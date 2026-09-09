from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.CarrierService import CarrierService
from ecommerceOnDjango.models.ShippingMethod import ShippingMethod
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CarrierService
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarrierServiceDelegate Declaration
#======================================================================
class CarrierServiceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, carrierServiceId ):
		try:	
			carrierService = CarrierService.objects.filter(id=carrierServiceId)
			return carrierService.first();
		except CarrierService.DoesNotExist:
			raise ProcessingError("CarrierService with id " + str(carrierServiceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, carrierService):
		for model in serializers.deserialize("json", carrierService):
			model.save()
			return model;

	def create(self, carrierService):
		carrierService.save()
		return carrierService;

	def saveFromJson(self, carrierService):
		for model in serializers.deserialize("json", carrierService):
			model.save()
			return carrierService;
	
	def save(self, carrierService):
		carrierService.save()
		return carrierService;
	
	def delete(self, carrierServiceId ):
		errMsg = "Failed to delete CarrierService from db using id " + str(carrierServiceId)
		
		try:
			carrierService = CarrierService.objects.get(id=carrierServiceId)
			carrierService.delete()
			return True
		except CarrierService.DoesNotExist:
			raise ProcessingError("CarrierService with id " + str(carrierServiceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CarrierService.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CarrierService from db")
		except Exception:
			return None;
		
	def addShippingMethods( self, carrierServiceId, shippingMethodsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShippingMethodDelegate import ShippingMethodDelegate

		errMsg = "Failed to add elements " + str(shippingMethodsIds) + " for ShippingMethods on CarrierService"

		try:
			# get the CarrierService
			carrierService = self.get( carrierServiceId ).first()
				
			# split on a comma with no spaces
			idList = shippingMethodsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ShippingMethod		
				shippingMethod = ShippingMethodDelegate().get(id).first();	
				# add the ShippingMethod
				carrierService.shippingMethods.add(shippingMethod)
				
			# save it		
			carrierService.save()
			
			# reload and return the appropriate version
			return self.get( carrierServiceId );
		except CarrierService.DoesNotExist:
			raise ProcessingError(errMsg + " : CarrierService with id " + str(carrierServiceId) + " does not exist.")
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeShippingMethods( self, carrierServiceId, shippingMethodsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShippingMethodDelegate import ShippingMethodDelegate

		errMsg = "Failed to remove elements " + str(shippingMethodsIds) + " for ShippingMethods on CarrierService"

		try:
			# get the CarrierService
			carrierService = self.get( carrierServiceId ).first()
				
			# split on a comma with no spaces
			idList = shippingMethodsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ShippingMethod		
				shippingMethod = ShippingMethodDelegate().get(id).first();	
				# add the ShippingMethod
				carrierService.shippingMethods.remove(shippingMethod)
				
			# save it		
			carrierService.save()
			
			# reload and return the appropriate version
			return self.get( carrierServiceId );
		except CarrierService.DoesNotExist:
			raise ProcessingError(errMsg + " : CarrierService with id " + str(carrierServiceId) + " does not exist.")
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
