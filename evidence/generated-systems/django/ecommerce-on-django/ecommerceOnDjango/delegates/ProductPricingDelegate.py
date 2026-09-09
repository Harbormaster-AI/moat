from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.ProductPricing import ProductPricing
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProductPricing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductPricingDelegate Declaration
#======================================================================
class ProductPricingDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productPricingId ):
		try:	
			productPricing = ProductPricing.objects.filter(id=productPricingId)
			return productPricing.first();
		except ProductPricing.DoesNotExist:
			raise ProcessingError("ProductPricing with id " + str(productPricingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, productPricing):
		for model in serializers.deserialize("json", productPricing):
			model.save()
			return model;

	def create(self, productPricing):
		productPricing.save()
		return productPricing;

	def saveFromJson(self, productPricing):
		for model in serializers.deserialize("json", productPricing):
			model.save()
			return productPricing;
	
	def save(self, productPricing):
		productPricing.save()
		return productPricing;
	
	def delete(self, productPricingId ):
		errMsg = "Failed to delete ProductPricing from db using id " + str(productPricingId)
		
		try:
			productPricing = ProductPricing.objects.get(id=productPricingId)
			productPricing.delete()
			return True
		except ProductPricing.DoesNotExist:
			raise ProcessingError("ProductPricing with id " + str(productPricingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProductPricing.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProductPricing from db")
		except Exception:
			return None;
		
	def assignVariant( self, productPricingId, variantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on ProductPricing"

		try:
			# get the ProductPricing from db
			productPricing = self.get( productPricingId ).first()	
			
			# get the ProductVariant from db
			productVariant = ProductVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			productPricing.variant = productVariant
			
			#save it
			productPricing.save()

			# reload and return the appropriate version					
			return self.get( productPricingId );
		except ProductPricing.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductPricing with id " + str(productPricingId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, productPricingId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on ProductPricing"

		try:
			# get the ProductPricing from db
			productPricing = self.get( productPricingId ).first()	
			
			# assign to None for unassignment
			productPricing.productVariant = None			

			#save it
			productPricing.save()

			# reload and return the appropriate version					
			return self.get( productPricingId );
		except ProductPricing.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductPricing with id " + str(productPricingId) + " does not exist.")
		except Exception:
			return None;
		
	def assignChannel( self, productPricingId, channelId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to assign element " + str(channelId) + " for Channel on ProductPricing"

		try:
			# get the ProductPricing from db
			productPricing = self.get( productPricingId ).first()	
			
			# get the Channel from db
			channel = ChannelDelegate().get(channelId).first();
			
			# assign the Channel		
			productPricing.channel = channel
			
			#save it
			productPricing.save()

			# reload and return the appropriate version					
			return self.get( productPricingId );
		except ProductPricing.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductPricing with id " + str(productPricingId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignChannel( self, productPricingId ):
		errMsg = "Failed to unassign element " + str(channelId) + " for Channel on ProductPricing"

		try:
			# get the ProductPricing from db
			productPricing = self.get( productPricingId ).first()	
			
			# assign to None for unassignment
			productPricing.channel = None			

			#save it
			productPricing.save()

			# reload and return the appropriate version					
			return self.get( productPricingId );
		except ProductPricing.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductPricing with id " + str(productPricingId) + " does not exist.")
		except Exception:
			return None;
		
