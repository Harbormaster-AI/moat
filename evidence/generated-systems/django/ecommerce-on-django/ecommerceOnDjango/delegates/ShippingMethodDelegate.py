from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.ShippingMethod import ShippingMethod
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.CarrierService import CarrierService
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ShippingMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShippingMethodDelegate Declaration
#======================================================================
class ShippingMethodDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, shippingMethodId ):
		try:	
			shippingMethod = ShippingMethod.objects.filter(id=shippingMethodId)
			return shippingMethod.first();
		except ShippingMethod.DoesNotExist:
			raise ProcessingError("ShippingMethod with id " + str(shippingMethodId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, shippingMethod):
		for model in serializers.deserialize("json", shippingMethod):
			model.save()
			return model;

	def create(self, shippingMethod):
		shippingMethod.save()
		return shippingMethod;

	def saveFromJson(self, shippingMethod):
		for model in serializers.deserialize("json", shippingMethod):
			model.save()
			return shippingMethod;
	
	def save(self, shippingMethod):
		shippingMethod.save()
		return shippingMethod;
	
	def delete(self, shippingMethodId ):
		errMsg = "Failed to delete ShippingMethod from db using id " + str(shippingMethodId)
		
		try:
			shippingMethod = ShippingMethod.objects.get(id=shippingMethodId)
			shippingMethod.delete()
			return True
		except ShippingMethod.DoesNotExist:
			raise ProcessingError("ShippingMethod with id " + str(shippingMethodId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ShippingMethod.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ShippingMethod from db")
		except Exception:
			return None;
		
	def assignCarrierService( self, shippingMethodId, carrierServiceId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CarrierServiceDelegate import CarrierServiceDelegate

		errMsg = "Failed to assign element " + str(carrierServiceId) + " for CarrierService on ShippingMethod"

		try:
			# get the ShippingMethod from db
			shippingMethod = self.get( shippingMethodId ).first()	
			
			# get the CarrierService from db
			carrierService = CarrierServiceDelegate().get(carrierServiceId).first();
			
			# assign the CarrierService		
			shippingMethod.carrierService = carrierService
			
			#save it
			shippingMethod.save()

			# reload and return the appropriate version					
			return self.get( shippingMethodId );
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod with id " + str(shippingMethodId) + " does not exist.")
		except CarrierService.DoesNotExist:
			raise ProcessingError(errMsg + " : CarrierService with id " + str(carrierServiceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCarrierService( self, shippingMethodId ):
		errMsg = "Failed to unassign element " + str(carrierServiceId) + " for CarrierService on ShippingMethod"

		try:
			# get the ShippingMethod from db
			shippingMethod = self.get( shippingMethodId ).first()	
			
			# assign to None for unassignment
			shippingMethod.carrierService = None			

			#save it
			shippingMethod.save()

			# reload and return the appropriate version					
			return self.get( shippingMethodId );
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod with id " + str(shippingMethodId) + " does not exist.")
		except Exception:
			return None;
		
	def addChannels( self, shippingMethodId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to add elements " + str(channelsIds) + " for Channels on ShippingMethod"

		try:
			# get the ShippingMethod
			shippingMethod = self.get( shippingMethodId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				shippingMethod.channels.add(channel)
				
			# save it		
			shippingMethod.save()
			
			# reload and return the appropriate version
			return self.get( shippingMethodId );
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod with id " + str(shippingMethodId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChannels( self, shippingMethodId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to remove elements " + str(channelsIds) + " for Channels on ShippingMethod"

		try:
			# get the ShippingMethod
			shippingMethod = self.get( shippingMethodId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				shippingMethod.channels.remove(channel)
				
			# save it		
			shippingMethod.save()
			
			# reload and return the appropriate version
			return self.get( shippingMethodId );
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod with id " + str(shippingMethodId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
