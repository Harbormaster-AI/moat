from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Cart import Cart
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.CartItem import CartItem
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Cart
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartDelegate Declaration
#======================================================================
class CartDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, cartId ):
		try:	
			cart = Cart.objects.filter(id=cartId)
			return cart.first();
		except Cart.DoesNotExist:
			raise ProcessingError("Cart with id " + str(cartId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, cart):
		for model in serializers.deserialize("json", cart):
			model.save()
			return model;

	def create(self, cart):
		cart.save()
		return cart;

	def saveFromJson(self, cart):
		for model in serializers.deserialize("json", cart):
			model.save()
			return cart;
	
	def save(self, cart):
		cart.save()
		return cart;
	
	def delete(self, cartId ):
		errMsg = "Failed to delete Cart from db using id " + str(cartId)
		
		try:
			cart = Cart.objects.get(id=cartId)
			cart.delete()
			return True
		except Cart.DoesNotExist:
			raise ProcessingError("Cart with id " + str(cartId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Cart.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Cart from db")
		except Exception:
			return None;
		
	def assignCustomer( self, cartId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Cart"

		try:
			# get the Cart from db
			cart = self.get( cartId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			cart.customer = customer
			
			#save it
			cart.save()

			# reload and return the appropriate version					
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, cartId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Cart"

		try:
			# get the Cart from db
			cart = self.get( cartId ).first()	
			
			# assign to None for unassignment
			cart.customer = None			

			#save it
			cart.save()

			# reload and return the appropriate version					
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except Exception:
			return None;
		
	def assignChannel( self, cartId, channelId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to assign element " + str(channelId) + " for Channel on Cart"

		try:
			# get the Cart from db
			cart = self.get( cartId ).first()	
			
			# get the Channel from db
			channel = ChannelDelegate().get(channelId).first();
			
			# assign the Channel		
			cart.channel = channel
			
			#save it
			cart.save()

			# reload and return the appropriate version					
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignChannel( self, cartId ):
		errMsg = "Failed to unassign element " + str(channelId) + " for Channel on Cart"

		try:
			# get the Cart from db
			cart = self.get( cartId ).first()	
			
			# assign to None for unassignment
			cart.channel = None			

			#save it
			cart.save()

			# reload and return the appropriate version					
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except Exception:
			return None;
		
	def addItems( self, cartId, itemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CartItemDelegate import CartItemDelegate

		errMsg = "Failed to add elements " + str(itemsIds) + " for Items on Cart"

		try:
			# get the Cart
			cart = self.get( cartId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CartItem		
				cartItem = CartItemDelegate().get(id).first();	
				# add the CartItem
				cart.items.add(cartItem)
				
			# save it		
			cart.save()
			
			# reload and return the appropriate version
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeItems( self, cartId, itemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CartItemDelegate import CartItemDelegate

		errMsg = "Failed to remove elements " + str(itemsIds) + " for Items on Cart"

		try:
			# get the Cart
			cart = self.get( cartId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CartItem		
				cartItem = CartItemDelegate().get(id).first();	
				# add the CartItem
				cart.items.remove(cartItem)
				
			# save it		
			cart.save()
			
			# reload and return the appropriate version
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except CartItem.DoesNotExist:
			raise ProcessingError(errMsg + " : CartItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAppliedPromotions( self, cartId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to add elements " + str(appliedPromotionsIds) + " for AppliedPromotions on Cart"

		try:
			# get the Cart
			cart = self.get( cartId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				cart.appliedPromotions.add(promotion)
				
			# save it		
			cart.save()
			
			# reload and return the appropriate version
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAppliedPromotions( self, cartId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to remove elements " + str(appliedPromotionsIds) + " for AppliedPromotions on Cart"

		try:
			# get the Cart
			cart = self.get( cartId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				cart.appliedPromotions.remove(promotion)
				
			# save it		
			cart.save()
			
			# reload and return the appropriate version
			return self.get( cartId );
		except Cart.DoesNotExist:
			raise ProcessingError(errMsg + " : Cart with id " + str(cartId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
