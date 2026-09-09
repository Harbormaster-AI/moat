from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Wishlist import Wishlist
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.WishlistItem import WishlistItem
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Wishlist
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WishlistDelegate Declaration
#======================================================================
class WishlistDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, wishlistId ):
		try:	
			wishlist = Wishlist.objects.filter(id=wishlistId)
			return wishlist.first();
		except Wishlist.DoesNotExist:
			raise ProcessingError("Wishlist with id " + str(wishlistId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, wishlist):
		for model in serializers.deserialize("json", wishlist):
			model.save()
			return model;

	def create(self, wishlist):
		wishlist.save()
		return wishlist;

	def saveFromJson(self, wishlist):
		for model in serializers.deserialize("json", wishlist):
			model.save()
			return wishlist;
	
	def save(self, wishlist):
		wishlist.save()
		return wishlist;
	
	def delete(self, wishlistId ):
		errMsg = "Failed to delete Wishlist from db using id " + str(wishlistId)
		
		try:
			wishlist = Wishlist.objects.get(id=wishlistId)
			wishlist.delete()
			return True
		except Wishlist.DoesNotExist:
			raise ProcessingError("Wishlist with id " + str(wishlistId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Wishlist.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Wishlist from db")
		except Exception:
			return None;
		
	def assignCustomer( self, wishlistId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Wishlist"

		try:
			# get the Wishlist from db
			wishlist = self.get( wishlistId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			wishlist.customer = customer
			
			#save it
			wishlist.save()

			# reload and return the appropriate version					
			return self.get( wishlistId );
		except Wishlist.DoesNotExist:
			raise ProcessingError(errMsg + " : Wishlist with id " + str(wishlistId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, wishlistId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Wishlist"

		try:
			# get the Wishlist from db
			wishlist = self.get( wishlistId ).first()	
			
			# assign to None for unassignment
			wishlist.customer = None			

			#save it
			wishlist.save()

			# reload and return the appropriate version					
			return self.get( wishlistId );
		except Wishlist.DoesNotExist:
			raise ProcessingError(errMsg + " : Wishlist with id " + str(wishlistId) + " does not exist.")
		except Exception:
			return None;
		
	def addItems( self, wishlistId, itemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.WishlistItemDelegate import WishlistItemDelegate

		errMsg = "Failed to add elements " + str(itemsIds) + " for Items on Wishlist"

		try:
			# get the Wishlist
			wishlist = self.get( wishlistId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WishlistItem		
				wishlistItem = WishlistItemDelegate().get(id).first();	
				# add the WishlistItem
				wishlist.items.add(wishlistItem)
				
			# save it		
			wishlist.save()
			
			# reload and return the appropriate version
			return self.get( wishlistId );
		except Wishlist.DoesNotExist:
			raise ProcessingError(errMsg + " : Wishlist with id " + str(wishlistId) + " does not exist.")
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeItems( self, wishlistId, itemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.WishlistItemDelegate import WishlistItemDelegate

		errMsg = "Failed to remove elements " + str(itemsIds) + " for Items on Wishlist"

		try:
			# get the Wishlist
			wishlist = self.get( wishlistId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WishlistItem		
				wishlistItem = WishlistItemDelegate().get(id).first();	
				# add the WishlistItem
				wishlist.items.remove(wishlistItem)
				
			# save it		
			wishlist.save()
			
			# reload and return the appropriate version
			return self.get( wishlistId );
		except Wishlist.DoesNotExist:
			raise ProcessingError(errMsg + " : Wishlist with id " + str(wishlistId) + " does not exist.")
		except WishlistItem.DoesNotExist:
			raise ProcessingError(errMsg + " : WishlistItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
