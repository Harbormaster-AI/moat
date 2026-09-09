from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.WishlistItem import WishlistItem
from ecommerceOnDjango.models.Wishlist import Wishlist
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model WishlistItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WishlistItemDelegate Declaration
#======================================================================
class WishlistItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, wishlistItemId ):
		try:	
			wishlistItem = WishlistItem.objects.filter(id=wishlistItemId)
			return wishlistItem.first();
		except WishlistItem.DoesNotExist:
			raise ProcessingError("WishlistItem with id " + str(wishlistItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, wishlistItem):
		for model in serializers.deserialize("json", wishlistItem):
			model.save()
			return model;

	def create(self, wishlistItem):
		wishlistItem.save()
		return wishlistItem;

	def saveFromJson(self, wishlistItem):
		for model in serializers.deserialize("json", wishlistItem):
			model.save()
			return wishlistItem;
	
	def save(self, wishlistItem):
		wishlistItem.save()
		return wishlistItem;
	
	def delete(self, wishlistItemId ):
		errMsg = "Failed to delete WishlistItem from db using id " + str(wishlistItemId)
		
		try:
			wishlistItem = WishlistItem.objects.get(id=wishlistItemId)
			wishlistItem.delete()
			return True
		except WishlistItem.DoesNotExist:
			raise ProcessingError("WishlistItem with id " + str(wishlistItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = WishlistItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all WishlistItem from db")
		except Exception:
			return None;
		
	def assignWishlist( self, wishlistItemId, wishlistId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.WishlistDelegate import WishlistDelegate

		errMsg = "Failed to assign element " + str(wishlistId) + " for Wishlist on WishlistItem"

		try:
			# get the WishlistItem from db
			wishlistItem = self.get( wishlistItemId ).first()	
			
			# get the Wishlist from db
			wishlist = WishlistDelegate().get(wishlistId).first();
			
			# assign the Wishlist		
			wishlistItem.wishlist = wishlist
			
			#save it
			wishlistItem.save()

			# reload and return the appropriate version					
			return self.get( wishlistItemId );
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem with id " + str(wishlistItemId) + " does not exist.")
		except Wishlist.DoesNotExist:
			raise ProcessingError(errMsg + " : Wishlist with id " + str(wishlistId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWishlist( self, wishlistItemId ):
		errMsg = "Failed to unassign element " + str(wishlistId) + " for Wishlist on WishlistItem"

		try:
			# get the WishlistItem from db
			wishlistItem = self.get( wishlistItemId ).first()	
			
			# assign to None for unassignment
			wishlistItem.wishlist = None			

			#save it
			wishlistItem.save()

			# reload and return the appropriate version					
			return self.get( wishlistItemId );
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem with id " + str(wishlistItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignVariant( self, wishlistItemId, variantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on WishlistItem"

		try:
			# get the WishlistItem from db
			wishlistItem = self.get( wishlistItemId ).first()	
			
			# get the ProductVariant from db
			productVariant = ProductVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			wishlistItem.variant = productVariant
			
			#save it
			wishlistItem.save()

			# reload and return the appropriate version					
			return self.get( wishlistItemId );
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem with id " + str(wishlistItemId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, wishlistItemId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on WishlistItem"

		try:
			# get the WishlistItem from db
			wishlistItem = self.get( wishlistItemId ).first()	
			
			# assign to None for unassignment
			wishlistItem.productVariant = None			

			#save it
			wishlistItem.save()

			# reload and return the appropriate version					
			return self.get( wishlistItemId );
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem with id " + str(wishlistItemId) + " does not exist.")
		except Exception:
			return None;
		
