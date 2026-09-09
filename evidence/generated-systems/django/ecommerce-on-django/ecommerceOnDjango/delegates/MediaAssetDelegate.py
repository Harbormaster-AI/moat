from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.MediaAsset import MediaAsset
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MediaAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MediaAssetDelegate Declaration
#======================================================================
class MediaAssetDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, mediaAssetId ):
		try:	
			mediaAsset = MediaAsset.objects.filter(id=mediaAssetId)
			return mediaAsset.first();
		except MediaAsset.DoesNotExist:
			raise ProcessingError("MediaAsset with id " + str(mediaAssetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, mediaAsset):
		for model in serializers.deserialize("json", mediaAsset):
			model.save()
			return model;

	def create(self, mediaAsset):
		mediaAsset.save()
		return mediaAsset;

	def saveFromJson(self, mediaAsset):
		for model in serializers.deserialize("json", mediaAsset):
			model.save()
			return mediaAsset;
	
	def save(self, mediaAsset):
		mediaAsset.save()
		return mediaAsset;
	
	def delete(self, mediaAssetId ):
		errMsg = "Failed to delete MediaAsset from db using id " + str(mediaAssetId)
		
		try:
			mediaAsset = MediaAsset.objects.get(id=mediaAssetId)
			mediaAsset.delete()
			return True
		except MediaAsset.DoesNotExist:
			raise ProcessingError("MediaAsset with id " + str(mediaAssetId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MediaAsset.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MediaAsset from db")
		except Exception:
			return None;
		
	def assignProduct( self, mediaAssetId, productId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on MediaAsset"

		try:
			# get the MediaAsset from db
			mediaAsset = self.get( mediaAssetId ).first()	
			
			# get the Product from db
			product = ProductDelegate().get(productId).first();
			
			# assign the Product		
			mediaAsset.product = product
			
			#save it
			mediaAsset.save()

			# reload and return the appropriate version					
			return self.get( mediaAssetId );
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset with id " + str(mediaAssetId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, mediaAssetId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on MediaAsset"

		try:
			# get the MediaAsset from db
			mediaAsset = self.get( mediaAssetId ).first()	
			
			# assign to None for unassignment
			mediaAsset.product = None			

			#save it
			mediaAsset.save()

			# reload and return the appropriate version					
			return self.get( mediaAssetId );
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset with id " + str(mediaAssetId) + " does not exist.")
		except Exception:
			return None;
		
	def assignVariant( self, mediaAssetId, variantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on MediaAsset"

		try:
			# get the MediaAsset from db
			mediaAsset = self.get( mediaAssetId ).first()	
			
			# get the ProductVariant from db
			productVariant = ProductVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			mediaAsset.variant = productVariant
			
			#save it
			mediaAsset.save()

			# reload and return the appropriate version					
			return self.get( mediaAssetId );
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset with id " + str(mediaAssetId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, mediaAssetId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on MediaAsset"

		try:
			# get the MediaAsset from db
			mediaAsset = self.get( mediaAssetId ).first()	
			
			# assign to None for unassignment
			mediaAsset.productVariant = None			

			#save it
			mediaAsset.save()

			# reload and return the appropriate version					
			return self.get( mediaAssetId );
		except MediaAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : MediaAsset with id " + str(mediaAssetId) + " does not exist.")
		except Exception:
			return None;
		
