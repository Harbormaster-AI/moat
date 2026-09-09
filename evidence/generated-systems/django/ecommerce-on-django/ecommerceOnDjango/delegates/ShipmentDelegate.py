from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Shipment import Shipment
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.ShipmentItem import ShipmentItem
from ecommerceOnDjango.models.FulfillmentCenter import FulfillmentCenter
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Shipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShipmentDelegate Declaration
#======================================================================
class ShipmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, shipmentId ):
		try:	
			shipment = Shipment.objects.filter(id=shipmentId)
			return shipment.first();
		except Shipment.DoesNotExist:
			raise ProcessingError("Shipment with id " + str(shipmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, shipment):
		for model in serializers.deserialize("json", shipment):
			model.save()
			return model;

	def create(self, shipment):
		shipment.save()
		return shipment;

	def saveFromJson(self, shipment):
		for model in serializers.deserialize("json", shipment):
			model.save()
			return shipment;
	
	def save(self, shipment):
		shipment.save()
		return shipment;
	
	def delete(self, shipmentId ):
		errMsg = "Failed to delete Shipment from db using id " + str(shipmentId)
		
		try:
			shipment = Shipment.objects.get(id=shipmentId)
			shipment.delete()
			return True
		except Shipment.DoesNotExist:
			raise ProcessingError("Shipment with id " + str(shipmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Shipment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Shipment from db")
		except Exception:
			return None;
		
	def assignOrder( self, shipmentId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Shipment"

		try:
			# get the Shipment from db
			shipment = self.get( shipmentId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			shipment.order = order
			
			#save it
			shipment.save()

			# reload and return the appropriate version					
			return self.get( shipmentId );
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, shipmentId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Shipment"

		try:
			# get the Shipment from db
			shipment = self.get( shipmentId ).first()	
			
			# assign to None for unassignment
			shipment.order = None			

			#save it
			shipment.save()

			# reload and return the appropriate version					
			return self.get( shipmentId );
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFulfillmentCenter( self, shipmentId, fulfillmentCenterId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

		errMsg = "Failed to assign element " + str(fulfillmentCenterId) + " for FulfillmentCenter on Shipment"

		try:
			# get the Shipment from db
			shipment = self.get( shipmentId ).first()	
			
			# get the FulfillmentCenter from db
			fulfillmentCenter = FulfillmentCenterDelegate().get(fulfillmentCenterId).first();
			
			# assign the FulfillmentCenter		
			shipment.fulfillmentCenter = fulfillmentCenter
			
			#save it
			shipment.save()

			# reload and return the appropriate version					
			return self.get( shipmentId );
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFulfillmentCenter( self, shipmentId ):
		errMsg = "Failed to unassign element " + str(fulfillmentCenterId) + " for FulfillmentCenter on Shipment"

		try:
			# get the Shipment from db
			shipment = self.get( shipmentId ).first()	
			
			# assign to None for unassignment
			shipment.fulfillmentCenter = None			

			#save it
			shipment.save()

			# reload and return the appropriate version					
			return self.get( shipmentId );
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addShipmentItems( self, shipmentId, shipmentItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentItemDelegate import ShipmentItemDelegate

		errMsg = "Failed to add elements " + str(shipmentItemsIds) + " for ShipmentItems on Shipment"

		try:
			# get the Shipment
			shipment = self.get( shipmentId ).first()
				
			# split on a comma with no spaces
			idList = shipmentItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ShipmentItem		
				shipmentItem = ShipmentItemDelegate().get(id).first();	
				# add the ShipmentItem
				shipment.shipmentItems.add(shipmentItem)
				
			# save it		
			shipment.save()
			
			# reload and return the appropriate version
			return self.get( shipmentId );
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except ShipmentItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ShipmentItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeShipmentItems( self, shipmentId, shipmentItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentItemDelegate import ShipmentItemDelegate

		errMsg = "Failed to remove elements " + str(shipmentItemsIds) + " for ShipmentItems on Shipment"

		try:
			# get the Shipment
			shipment = self.get( shipmentId ).first()
				
			# split on a comma with no spaces
			idList = shipmentItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ShipmentItem		
				shipmentItem = ShipmentItemDelegate().get(id).first();	
				# add the ShipmentItem
				shipment.shipmentItems.remove(shipmentItem)
				
			# save it		
			shipment.save()
			
			# reload and return the appropriate version
			return self.get( shipmentId );
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment with id " + str(shipmentId) + " does not exist.")
		except ShipmentItem.DoesNotExist:
			raise ProcessingError(errMsg + " : ShipmentItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
