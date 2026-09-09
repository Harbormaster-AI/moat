from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.ReturnRequest import ReturnRequest
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.ReturnItem import ReturnItem
from ecommerceOnDjango.models.Refund import Refund
from ecommerceOnDjango.models.Shipment import Shipment
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ReturnRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnRequestDelegate Declaration
#======================================================================
class ReturnRequestDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, returnRequestId ):
		try:	
			returnRequest = ReturnRequest.objects.filter(id=returnRequestId)
			return returnRequest.first();
		except ReturnRequest.DoesNotExist:
			raise ProcessingError("ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, returnRequest):
		for model in serializers.deserialize("json", returnRequest):
			model.save()
			return model;

	def create(self, returnRequest):
		returnRequest.save()
		return returnRequest;

	def saveFromJson(self, returnRequest):
		for model in serializers.deserialize("json", returnRequest):
			model.save()
			return returnRequest;
	
	def save(self, returnRequest):
		returnRequest.save()
		return returnRequest;
	
	def delete(self, returnRequestId ):
		errMsg = "Failed to delete ReturnRequest from db using id " + str(returnRequestId)
		
		try:
			returnRequest = ReturnRequest.objects.get(id=returnRequestId)
			returnRequest.delete()
			return True
		except ReturnRequest.DoesNotExist:
			raise ProcessingError("ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ReturnRequest.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ReturnRequest from db")
		except Exception:
			return None;
		
	def assignOrder( self, returnRequestId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on ReturnRequest"

		try:
			# get the ReturnRequest from db
			returnRequest = self.get( returnRequestId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			returnRequest.order = order
			
			#save it
			returnRequest.save()

			# reload and return the appropriate version					
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, returnRequestId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on ReturnRequest"

		try:
			# get the ReturnRequest from db
			returnRequest = self.get( returnRequestId ).first()	
			
			# assign to None for unassignment
			returnRequest.order = None			

			#save it
			returnRequest.save()

			# reload and return the appropriate version					
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRefund( self, returnRequestId, refundId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.RefundDelegate import RefundDelegate

		errMsg = "Failed to assign element " + str(refundId) + " for Refund on ReturnRequest"

		try:
			# get the ReturnRequest from db
			returnRequest = self.get( returnRequestId ).first()	
			
			# get the Refund from db
			refund = RefundDelegate().get(refundId).first();
			
			# assign the Refund		
			returnRequest.refund = refund
			
			#save it
			returnRequest.save()

			# reload and return the appropriate version					
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund with id " + str(refundId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRefund( self, returnRequestId ):
		errMsg = "Failed to unassign element " + str(refundId) + " for Refund on ReturnRequest"

		try:
			# get the ReturnRequest from db
			returnRequest = self.get( returnRequestId ).first()	
			
			# assign to None for unassignment
			returnRequest.refund = None			

			#save it
			returnRequest.save()

			# reload and return the appropriate version					
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except Exception:
			return None;
		
	def assignShipment( self, returnRequestId, shipmentId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

		errMsg = "Failed to assign element " + str(shipmentId) + " for Shipment on ReturnRequest"

		try:
			# get the ReturnRequest from db
			returnRequest = self.get( returnRequestId ).first()	
			
			# get the Shipment from db
			shipment = ShipmentDelegate().get(shipmentId).first();
			
			# assign the Shipment		
			returnRequest.shipment = shipment
			
			#save it
			returnRequest.save()

			# reload and return the appropriate version					
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignShipment( self, returnRequestId ):
		errMsg = "Failed to unassign element " + str(shipmentId) + " for Shipment on ReturnRequest"

		try:
			# get the ReturnRequest from db
			returnRequest = self.get( returnRequestId ).first()	
			
			# assign to None for unassignment
			returnRequest.shipment = None			

			#save it
			returnRequest.save()

			# reload and return the appropriate version					
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except Exception:
			return None;
		
	def addItems( self, returnRequestId, itemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReturnItemDelegate import ReturnItemDelegate

		errMsg = "Failed to add elements " + str(itemsIds) + " for Items on ReturnRequest"

		try:
			# get the ReturnRequest
			returnRequest = self.get( returnRequestId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ReturnItem		
				returnItem = ReturnItemDelegate().get(id).first();	
				# add the ReturnItem
				returnRequest.items.add(returnItem)
				
			# save it		
			returnRequest.save()
			
			# reload and return the appropriate version
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except ReturnItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeItems( self, returnRequestId, itemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReturnItemDelegate import ReturnItemDelegate

		errMsg = "Failed to remove elements " + str(itemsIds) + " for Items on ReturnRequest"

		try:
			# get the ReturnRequest
			returnRequest = self.get( returnRequestId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ReturnItem		
				returnItem = ReturnItemDelegate().get(id).first();	
				# add the ReturnItem
				returnRequest.items.remove(returnItem)
				
			# save it		
			returnRequest.save()
			
			# reload and return the appropriate version
			return self.get( returnRequestId );
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest with id " + str(returnRequestId) + " does not exist.")
		except ReturnItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
