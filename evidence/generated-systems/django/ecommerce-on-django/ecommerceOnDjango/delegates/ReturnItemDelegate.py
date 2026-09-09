from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.ReturnItem import ReturnItem
from ecommerceOnDjango.models.ReturnRequest import ReturnRequest
from ecommerceOnDjango.models.OrderLine import OrderLine
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ReturnItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnItemDelegate Declaration
#======================================================================
class ReturnItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, returnItemId ):
		try:	
			returnItem = ReturnItem.objects.filter(id=returnItemId)
			return returnItem.first();
		except ReturnItem.DoesNotExist:
			raise ProcessingError("ReturnItem with id " + str(returnItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, returnItem):
		for model in serializers.deserialize("json", returnItem):
			model.save()
			return model;

	def create(self, returnItem):
		returnItem.save()
		return returnItem;

	def saveFromJson(self, returnItem):
		for model in serializers.deserialize("json", returnItem):
			model.save()
			return returnItem;
	
	def save(self, returnItem):
		returnItem.save()
		return returnItem;
	
	def delete(self, returnItemId ):
		errMsg = "Failed to delete ReturnItem from db using id " + str(returnItemId)
		
		try:
			returnItem = ReturnItem.objects.get(id=returnItemId)
			returnItem.delete()
			return True
		except ReturnItem.DoesNotExist:
			raise ProcessingError("ReturnItem with id " + str(returnItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ReturnItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ReturnItem from db")
		except Exception:
			return None;
		
	def assignReturnRequest( self, returnItemId, returnRequestId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReturnRequestDelegate import ReturnRequestDelegate

		errMsg = "Failed to assign element " + str(returnRequestId) + " for ReturnRequest on ReturnItem"

		try:
			# get the ReturnItem from db
			returnItem = self.get( returnItemId ).first()	
			
			# get the ReturnRequest from db
			returnRequest = ReturnRequestDelegate().get(returnRequestId).first();
			
			# assign the ReturnRequest		
			returnItem.returnRequest = returnRequest
			
			#save it
			returnItem.save()

			# reload and return the appropriate version					
			return self.get( returnItemId );
		except ReturnItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnItem with id " + str(returnItemId) + " does not exist.")
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignReturnRequest( self, returnItemId ):
		errMsg = "Failed to unassign element " + str(returnRequestId) + " for ReturnRequest on ReturnItem"

		try:
			# get the ReturnItem from db
			returnItem = self.get( returnItemId ).first()	
			
			# assign to None for unassignment
			returnItem.returnRequest = None			

			#save it
			returnItem.save()

			# reload and return the appropriate version					
			return self.get( returnItemId );
		except ReturnItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnItem with id " + str(returnItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrderLine( self, returnItemId, orderLineId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

		errMsg = "Failed to assign element " + str(orderLineId) + " for OrderLine on ReturnItem"

		try:
			# get the ReturnItem from db
			returnItem = self.get( returnItemId ).first()	
			
			# get the OrderLine from db
			orderLine = OrderLineDelegate().get(orderLineId).first();
			
			# assign the OrderLine		
			returnItem.orderLine = orderLine
			
			#save it
			returnItem.save()

			# reload and return the appropriate version					
			return self.get( returnItemId );
		except ReturnItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnItem with id " + str(returnItemId) + " does not exist.")
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrderLine( self, returnItemId ):
		errMsg = "Failed to unassign element " + str(orderLineId) + " for OrderLine on ReturnItem"

		try:
			# get the ReturnItem from db
			returnItem = self.get( returnItemId ).first()	
			
			# assign to None for unassignment
			returnItem.orderLine = None			

			#save it
			returnItem.save()

			# reload and return the appropriate version					
			return self.get( returnItemId );
		except ReturnItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnItem with id " + str(returnItemId) + " does not exist.")
		except Exception:
			return None;
		
