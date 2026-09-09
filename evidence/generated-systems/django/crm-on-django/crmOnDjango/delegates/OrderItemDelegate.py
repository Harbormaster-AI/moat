from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.OrderItem import OrderItem
from crmOnDjango.models.Order import Order
from crmOnDjango.models.Product import Product
from crmOnDjango.models.PriceBookEntry import PriceBookEntry
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model OrderItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderItemDelegate Declaration
#======================================================================
class OrderItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, orderItemId ):
		try:	
			orderItem = OrderItem.objects.filter(id=orderItemId)
			return orderItem.first();
		except OrderItem.DoesNotExist:
			raise ProcessingError("OrderItem with id " + str(orderItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, orderItem):
		for model in serializers.deserialize("json", orderItem):
			model.save()
			return model;

	def create(self, orderItem):
		orderItem.save()
		return orderItem;

	def saveFromJson(self, orderItem):
		for model in serializers.deserialize("json", orderItem):
			model.save()
			return orderItem;
	
	def save(self, orderItem):
		orderItem.save()
		return orderItem;
	
	def delete(self, orderItemId ):
		errMsg = "Failed to delete OrderItem from db using id " + str(orderItemId)
		
		try:
			orderItem = OrderItem.objects.get(id=orderItemId)
			orderItem.delete()
			return True
		except OrderItem.DoesNotExist:
			raise ProcessingError("OrderItem with id " + str(orderItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = OrderItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all OrderItem from db")
		except Exception:
			return None;
		
	def assignOrder( self, orderItemId, orderId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on OrderItem"

		try:
			# get the OrderItem from db
			orderItem = self.get( orderItemId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			orderItem.order = order
			
			#save it
			orderItem.save()

			# reload and return the appropriate version					
			return self.get( orderItemId );
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem with id " + str(orderItemId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, orderItemId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on OrderItem"

		try:
			# get the OrderItem from db
			orderItem = self.get( orderItemId ).first()	
			
			# assign to None for unassignment
			orderItem.order = None			

			#save it
			orderItem.save()

			# reload and return the appropriate version					
			return self.get( orderItemId );
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem with id " + str(orderItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProduct( self, orderItemId, productId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to assign element " + str(productId) + " for Product on OrderItem"

		try:
			# get the OrderItem from db
			orderItem = self.get( orderItemId ).first()	
			
			# get the Product from db
			product = ProductDelegate().get(productId).first();
			
			# assign the Product		
			orderItem.product = product
			
			#save it
			orderItem.save()

			# reload and return the appropriate version					
			return self.get( orderItemId );
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem with id " + str(orderItemId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product with id " + str(productId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProduct( self, orderItemId ):
		errMsg = "Failed to unassign element " + str(productId) + " for Product on OrderItem"

		try:
			# get the OrderItem from db
			orderItem = self.get( orderItemId ).first()	
			
			# assign to None for unassignment
			orderItem.product = None			

			#save it
			orderItem.save()

			# reload and return the appropriate version					
			return self.get( orderItemId );
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem with id " + str(orderItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPriceBookEntry( self, orderItemId, priceBookEntryId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

		errMsg = "Failed to assign element " + str(priceBookEntryId) + " for PriceBookEntry on OrderItem"

		try:
			# get the OrderItem from db
			orderItem = self.get( orderItemId ).first()	
			
			# get the PriceBookEntry from db
			priceBookEntry = PriceBookEntryDelegate().get(priceBookEntryId).first();
			
			# assign the PriceBookEntry		
			orderItem.priceBookEntry = priceBookEntry
			
			#save it
			orderItem.save()

			# reload and return the appropriate version					
			return self.get( orderItemId );
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem with id " + str(orderItemId) + " does not exist.")
		except PriceBookEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBookEntry with id " + str(priceBookEntryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPriceBookEntry( self, orderItemId ):
		errMsg = "Failed to unassign element " + str(priceBookEntryId) + " for PriceBookEntry on OrderItem"

		try:
			# get the OrderItem from db
			orderItem = self.get( orderItemId ).first()	
			
			# assign to None for unassignment
			orderItem.priceBookEntry = None			

			#save it
			orderItem.save()

			# reload and return the appropriate version					
			return self.get( orderItemId );
		except OrderItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderItem with id " + str(orderItemId) + " does not exist.")
		except Exception:
			return None;
		
