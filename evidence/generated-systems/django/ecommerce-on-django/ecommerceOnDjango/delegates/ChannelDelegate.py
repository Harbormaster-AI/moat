from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Catalog import Catalog
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.models.ShippingMethod import ShippingMethod
from ecommerceOnDjango.models.PaymentProvider import PaymentProvider
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Channel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChannelDelegate Declaration
#======================================================================
class ChannelDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, channelId ):
		try:	
			channel = Channel.objects.filter(id=channelId)
			return channel.first();
		except Channel.DoesNotExist:
			raise ProcessingError("Channel with id " + str(channelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, channel):
		for model in serializers.deserialize("json", channel):
			model.save()
			return model;

	def create(self, channel):
		channel.save()
		return channel;

	def saveFromJson(self, channel):
		for model in serializers.deserialize("json", channel):
			model.save()
			return channel;
	
	def save(self, channel):
		channel.save()
		return channel;
	
	def delete(self, channelId ):
		errMsg = "Failed to delete Channel from db using id " + str(channelId)
		
		try:
			channel = Channel.objects.get(id=channelId)
			channel.delete()
			return True
		except Channel.DoesNotExist:
			raise ProcessingError("Channel with id " + str(channelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Channel.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Channel from db")
		except Exception:
			return None;
		
	def assignMerchant( self, channelId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Channel"

		try:
			# get the Channel from db
			channel = self.get( channelId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			channel.merchant = merchant
			
			#save it
			channel.save()

			# reload and return the appropriate version					
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, channelId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Channel"

		try:
			# get the Channel from db
			channel = self.get( channelId ).first()	
			
			# assign to None for unassignment
			channel.merchant = None			

			#save it
			channel.save()

			# reload and return the appropriate version					
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Exception:
			return None;
		
	def addCatalogs( self, channelId, catalogsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CatalogDelegate import CatalogDelegate

		errMsg = "Failed to add elements " + str(catalogsIds) + " for Catalogs on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = catalogsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Catalog		
				catalog = CatalogDelegate().get(id).first();	
				# add the Catalog
				channel.catalogs.add(catalog)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Catalog.DoesNotExist:
			raise ProcessingError(errMsg + " : Catalog does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCatalogs( self, channelId, catalogsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CatalogDelegate import CatalogDelegate

		errMsg = "Failed to remove elements " + str(catalogsIds) + " for Catalogs on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = catalogsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Catalog		
				catalog = CatalogDelegate().get(id).first();	
				# add the Catalog
				channel.catalogs.remove(catalog)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Catalog.DoesNotExist:
			raise ProcessingError(errMsg + " : Catalog does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPromotions( self, channelId, promotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to add elements " + str(promotionsIds) + " for Promotions on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = promotionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				channel.promotions.add(promotion)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePromotions( self, channelId, promotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to remove elements " + str(promotionsIds) + " for Promotions on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = promotionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				channel.promotions.remove(promotion)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addShippingMethods( self, channelId, shippingMethodsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShippingMethodDelegate import ShippingMethodDelegate

		errMsg = "Failed to add elements " + str(shippingMethodsIds) + " for ShippingMethods on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = shippingMethodsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ShippingMethod		
				shippingMethod = ShippingMethodDelegate().get(id).first();	
				# add the ShippingMethod
				channel.shippingMethods.add(shippingMethod)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeShippingMethods( self, channelId, shippingMethodsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShippingMethodDelegate import ShippingMethodDelegate

		errMsg = "Failed to remove elements " + str(shippingMethodsIds) + " for ShippingMethods on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = shippingMethodsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ShippingMethod		
				shippingMethod = ShippingMethodDelegate().get(id).first();	
				# add the ShippingMethod
				channel.shippingMethods.remove(shippingMethod)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except ShippingMethod.DoesNotExist:
			raise ProcessingError(errMsg + " : ShippingMethod does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPaymentProviders( self, channelId, paymentProvidersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

		errMsg = "Failed to add elements " + str(paymentProvidersIds) + " for PaymentProviders on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = paymentProvidersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentProvider		
				paymentProvider = PaymentProviderDelegate().get(id).first();	
				# add the PaymentProvider
				channel.paymentProviders.add(paymentProvider)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePaymentProviders( self, channelId, paymentProvidersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

		errMsg = "Failed to remove elements " + str(paymentProvidersIds) + " for PaymentProviders on Channel"

		try:
			# get the Channel
			channel = self.get( channelId ).first()
				
			# split on a comma with no spaces
			idList = paymentProvidersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentProvider		
				paymentProvider = PaymentProviderDelegate().get(id).first();	
				# add the PaymentProvider
				channel.paymentProviders.remove(paymentProvider)
				
			# save it		
			channel.save()
			
			# reload and return the appropriate version
			return self.get( channelId );
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
