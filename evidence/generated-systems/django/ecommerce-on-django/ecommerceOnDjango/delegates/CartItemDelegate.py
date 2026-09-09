from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.CartItem import CartItem
from ecommerceOnDjango.models.Cart import Cart
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CartItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartItemDelegate Declaration
#======================================================================
class CartItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, cartItemId ):
		try:	
			cartItem = CartItem.objects.filter(id=cartItemId)
			return cartItem.first();
		except CartItem.DoesNotExist:
			raise ProcessingError("CartItem with id " + str(cartItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, cartItem):
		for model in serializers.deserialize("json", cartItem):
			model.save()
			return model;

	def create(self, cartItem):
		cartItem.save()
		return cartItem;

	def saveFromJson(self, cartItem):
		for model in serializers.deserialize("json", cartItem):
			model.save()
			return cartItem;
	
	def save(self, cartItem):
		cartItem.save()
		return cartItem;
	
	def delete(self, cartItemId ):
		errMsg = "Failed to delete CartItem from db using id " + str(cartItemId)
		
		try:
			cartItem = CartItem.objects.get(id=cartItemId)
			cartItem.delete()
			return True
		except CartItem.DoesNotExist:
			raise ProcessingError("CartItem with id " + str(cartItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CartItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CartItem from db")
		except Exception:
			return None;
		
	def assignCart( self, cartItemId, cartId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CartDelegate import CartDelegate

		errMsg = "Failed to assign element " + str(cartId) + " for Cart on CartItem"

		try:
			# get the CartItem from db
			cartItem = self.get( cartItemId ).first()	
			
			# get the Cart from db
			cart = CartDelegate().get(cartId).first();
			
			# assign the Cart		
			cartItem.cart = cart
			
			#save it
			cartItem.save()

			# reload and return the appropriate version					
			return self.get( cartItemId );
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem with id " + str(cartItemId) + " does not exist.")
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCart( self, cartItemId ):
		errMsg = "Failed to unassign element " + str(cartId) + " for Cart on CartItem"

		try:
			# get the CartItem from db
			cartItem = self.get( cartItemId ).first()	
			
			# assign to None for unassignment
			cartItem.cart = None			

			#save it
			cartItem.save()

			# reload and return the appropriate version					
			return self.get( cartItemId );
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem with id " + str(cartItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignVariant( self, cartItemId, variantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on CartItem"

		try:
			# get the CartItem from db
			cartItem = self.get( cartItemId ).first()	
			
			# get the ProductVariant from db
			productVariant = ProductVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			cartItem.variant = productVariant
			
			#save it
			cartItem.save()

			# reload and return the appropriate version					
			return self.get( cartItemId );
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem with id " + str(cartItemId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, cartItemId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on CartItem"

		try:
			# get the CartItem from db
			cartItem = self.get( cartItemId ).first()	
			
			# assign to None for unassignment
			cartItem.productVariant = None			

			#save it
			cartItem.save()

			# reload and return the appropriate version					
			return self.get( cartItemId );
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem with id " + str(cartItemId) + " does not exist.")
		except Exception:
			return None;
		
	def addAppliedPromotions( self, cartItemId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to add elements " + str(appliedPromotionsIds) + " for AppliedPromotions on CartItem"

		try:
			# get the CartItem
			cartItem = self.get( cartItemId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				cartItem.appliedPromotions.add(promotion)
				
			# save it		
			cartItem.save()
			
			# reload and return the appropriate version
			return self.get( cartItemId );
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem with id " + str(cartItemId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAppliedPromotions( self, cartItemId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to remove elements " + str(appliedPromotionsIds) + " for AppliedPromotions on CartItem"

		try:
			# get the CartItem
			cartItem = self.get( cartItemId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				cartItem.appliedPromotions.remove(promotion)
				
			# save it		
			cartItem.save()
			
			# reload and return the appropriate version
			return self.get( cartItemId );
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem with id " + str(cartItemId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
