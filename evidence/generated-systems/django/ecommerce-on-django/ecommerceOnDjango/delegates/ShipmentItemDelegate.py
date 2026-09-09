from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.ShipmentItem import ShipmentItem
from ecommerceOnDjango.models.Shipment import Shipment
from ecommerceOnDjango.models.OrderLine import OrderLine
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ShipmentItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentItemDelegate Declaration
#======================================================================
class ShipmentItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, shipmentItemId ):
		try:	
			shipmentItem = ShipmentItem.objects.filter(id=shipmentItemId)
			return shipmentItem.first();
		except ShipmentItem.DoesNotExist:
			raise ProcessingError("ShipmentItem with id " + str(shipmentItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, shipmentItem):
		for model in serializers.deserialize("json", shipmentItem):
			model.save()
			return model;

	def create(self, shipmentItem):
		shipmentItem.save()
		return shipmentItem;

	def saveFromJson(self, shipmentItem):
		for model in serializers.deserialize("json", shipmentItem):
			model.save()
			return shipmentItem;
	
	def save(self, shipmentItem):
		shipmentItem.save()
		return shipmentItem;
	
	def delete(self, shipmentItemId ):
		errMsg = "Failed to delete ShipmentItem from db using id " + str(shipmentItemId)
		
		try:
			shipmentItem = ShipmentItem.objects.get(id=shipmentItemId)
			shipmentItem.delete()
			return True
		except ShipmentItem.DoesNotExist:
			raise ProcessingError("ShipmentItem with id " + str(shipmentItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ShipmentItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ShipmentItem from db")
		except Exception:
			return None;
		
	def assignShipment( self, shipmentItemId, shipmentId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

		errMsg = "Failed to assign element " + str(shipmentId) + " for Shipment on ShipmentItem"

		try:
			# get the ShipmentItem from db
			shipmentItem = self.get( shipmentItemId ).first()	
			
			# get the Shipment from db
			shipment = ShipmentDelegate().get(shipmentId).first();
			
			# assign the Shipment		
			shipmentItem.shipment = shipment
			
			#save it
			shipmentItem.save()

			# reload and return the appropriate version					
			return self.get( shipmentItemId );
		except ShipmentItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ShipmentItem with id " + str(shipmentItemId) + " does not exist.")
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignShipment( self, shipmentItemId ):
		errMsg = "Failed to unassign element " + str(shipmentId) + " for Shipment on ShipmentItem"

		try:
			# get the ShipmentItem from db
			shipmentItem = self.get( shipmentItemId ).first()	
			
			# assign to None for unassignment
			shipmentItem.shipment = None			

			#save it
			shipmentItem.save()

			# reload and return the appropriate version					
			return self.get( shipmentItemId );
		except ShipmentItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ShipmentItem with id " + str(shipmentItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrderLine( self, shipmentItemId, orderLineId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

		errMsg = "Failed to assign element " + str(orderLineId) + " for OrderLine on ShipmentItem"

		try:
			# get the ShipmentItem from db
			shipmentItem = self.get( shipmentItemId ).first()	
			
			# get the OrderLine from db
			orderLine = OrderLineDelegate().get(orderLineId).first();
			
			# assign the OrderLine		
			shipmentItem.orderLine = orderLine
			
			#save it
			shipmentItem.save()

			# reload and return the appropriate version					
			return self.get( shipmentItemId );
		except ShipmentItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ShipmentItem with id " + str(shipmentItemId) + " does not exist.")
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine with id " + str(orderLineId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrderLine( self, shipmentItemId ):
		errMsg = "Failed to unassign element " + str(orderLineId) + " for OrderLine on ShipmentItem"

		try:
			# get the ShipmentItem from db
			shipmentItem = self.get( shipmentItemId ).first()	
			
			# assign to None for unassignment
			shipmentItem.orderLine = None			

			#save it
			shipmentItem.save()

			# reload and return the appropriate version					
			return self.get( shipmentItemId );
		except ShipmentItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ShipmentItem with id " + str(shipmentItemId) + " does not exist.")
		except Exception:
			return None;
		
