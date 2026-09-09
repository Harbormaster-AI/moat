from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.OrderLine import OrderLine
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model OrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderLineDelegate Declaration
#======================================================================
class OrderLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, orderLineId ):
		try:	
			orderLine = OrderLine.objects.filter(id=orderLineId)
			return orderLine.first();
		except OrderLine.DoesNotExist:
			raise ProcessingError("OrderLine with id " + str(orderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, orderLine):
		for model in serializers.deserialize("json", orderLine):
			model.save()
			return model;

	def create(self, orderLine):
		orderLine.save()
		return orderLine;

	def saveFromJson(self, orderLine):
		for model in serializers.deserialize("json", orderLine):
			model.save()
			return orderLine;
	
	def save(self, orderLine):
		orderLine.save()
		return orderLine;
	
	def delete(self, orderLineId ):
		errMsg = "Failed to delete OrderLine from db using id " + str(orderLineId)
		
		try:
			orderLine = OrderLine.objects.get(id=orderLineId)
			orderLine.delete()
			return True
		except OrderLine.DoesNotExist:
			raise ProcessingError("OrderLine with id " + str(orderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = OrderLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all OrderLine from db")
		except Exception:
			return None;
		
	def assignOrder( self, orderLineId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on OrderLine"

		try:
			# get the OrderLine from db
			orderLine = self.get( orderLineId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			orderLine.order = order
			
			#save it
			orderLine.save()

			# reload and return the appropriate version					
			return self.get( orderLineId );
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, orderLineId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on OrderLine"

		try:
			# get the OrderLine from db
			orderLine = self.get( orderLineId ).first()	
			
			# assign to None for unassignment
			orderLine.order = None			

			#save it
			orderLine.save()

			# reload and return the appropriate version					
			return self.get( orderLineId );
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignVariant( self, orderLineId, variantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

		errMsg = "Failed to assign element " + str(variantId) + " for Variant on OrderLine"

		try:
			# get the OrderLine from db
			orderLine = self.get( orderLineId ).first()	
			
			# get the ProductVariant from db
			productVariant = ProductVariantDelegate().get(variantId).first();
			
			# assign the Variant		
			orderLine.variant = productVariant
			
			#save it
			orderLine.save()

			# reload and return the appropriate version					
			return self.get( orderLineId );
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except ProductVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductVariant with id " + str(variantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignVariant( self, orderLineId ):
		errMsg = "Failed to unassign element " + str(variantId) + " for Variant on OrderLine"

		try:
			# get the OrderLine from db
			orderLine = self.get( orderLineId ).first()	
			
			# assign to None for unassignment
			orderLine.productVariant = None			

			#save it
			orderLine.save()

			# reload and return the appropriate version					
			return self.get( orderLineId );
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def addAppliedPromotions( self, orderLineId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to add elements " + str(appliedPromotionsIds) + " for AppliedPromotions on OrderLine"

		try:
			# get the OrderLine
			orderLine = self.get( orderLineId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				orderLine.appliedPromotions.add(promotion)
				
			# save it		
			orderLine.save()
			
			# reload and return the appropriate version
			return self.get( orderLineId );
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAppliedPromotions( self, orderLineId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to remove elements " + str(appliedPromotionsIds) + " for AppliedPromotions on OrderLine"

		try:
			# get the OrderLine
			orderLine = self.get( orderLineId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				orderLine.appliedPromotions.remove(promotion)
				
			# save it		
			orderLine.save()
			
			# reload and return the appropriate version
			return self.get( orderLineId );
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
