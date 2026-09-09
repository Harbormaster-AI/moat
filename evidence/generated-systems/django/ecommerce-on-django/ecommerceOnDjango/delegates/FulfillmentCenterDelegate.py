from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.FulfillmentCenter import FulfillmentCenter
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.InventoryItem import InventoryItem
from ecommerceOnDjango.models.Shipment import Shipment
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FulfillmentCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FulfillmentCenterDelegate Declaration
#======================================================================
class FulfillmentCenterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, fulfillmentCenterId ):
		try:	
			fulfillmentCenter = FulfillmentCenter.objects.filter(id=fulfillmentCenterId)
			return fulfillmentCenter.first();
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError("FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, fulfillmentCenter):
		for model in serializers.deserialize("json", fulfillmentCenter):
			model.save()
			return model;

	def create(self, fulfillmentCenter):
		fulfillmentCenter.save()
		return fulfillmentCenter;

	def saveFromJson(self, fulfillmentCenter):
		for model in serializers.deserialize("json", fulfillmentCenter):
			model.save()
			return fulfillmentCenter;
	
	def save(self, fulfillmentCenter):
		fulfillmentCenter.save()
		return fulfillmentCenter;
	
	def delete(self, fulfillmentCenterId ):
		errMsg = "Failed to delete FulfillmentCenter from db using id " + str(fulfillmentCenterId)
		
		try:
			fulfillmentCenter = FulfillmentCenter.objects.get(id=fulfillmentCenterId)
			fulfillmentCenter.delete()
			return True
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError("FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FulfillmentCenter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FulfillmentCenter from db")
		except Exception:
			return None;
		
	def assignMerchant( self, fulfillmentCenterId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on FulfillmentCenter"

		try:
			# get the FulfillmentCenter from db
			fulfillmentCenter = self.get( fulfillmentCenterId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			fulfillmentCenter.merchant = merchant
			
			#save it
			fulfillmentCenter.save()

			# reload and return the appropriate version					
			return self.get( fulfillmentCenterId );
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, fulfillmentCenterId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on FulfillmentCenter"

		try:
			# get the FulfillmentCenter from db
			fulfillmentCenter = self.get( fulfillmentCenterId ).first()	
			
			# assign to None for unassignment
			fulfillmentCenter.merchant = None			

			#save it
			fulfillmentCenter.save()

			# reload and return the appropriate version					
			return self.get( fulfillmentCenterId );
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except Exception:
			return None;
		
	def addInventoryItems( self, fulfillmentCenterId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on FulfillmentCenter"

		try:
			# get the FulfillmentCenter
			fulfillmentCenter = self.get( fulfillmentCenterId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				fulfillmentCenter.inventoryItems.add(inventoryItem)
				
			# save it		
			fulfillmentCenter.save()
			
			# reload and return the appropriate version
			return self.get( fulfillmentCenterId );
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, fulfillmentCenterId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on FulfillmentCenter"

		try:
			# get the FulfillmentCenter
			fulfillmentCenter = self.get( fulfillmentCenterId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				fulfillmentCenter.inventoryItems.remove(inventoryItem)
				
			# save it		
			fulfillmentCenter.save()
			
			# reload and return the appropriate version
			return self.get( fulfillmentCenterId );
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addShipments( self, fulfillmentCenterId, shipmentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

		errMsg = "Failed to add elements " + str(shipmentsIds) + " for Shipments on FulfillmentCenter"

		try:
			# get the FulfillmentCenter
			fulfillmentCenter = self.get( fulfillmentCenterId ).first()
				
			# split on a comma with no spaces
			idList = shipmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Shipment		
				shipment = ShipmentDelegate().get(id).first();	
				# add the Shipment
				fulfillmentCenter.shipments.add(shipment)
				
			# save it		
			fulfillmentCenter.save()
			
			# reload and return the appropriate version
			return self.get( fulfillmentCenterId );
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeShipments( self, fulfillmentCenterId, shipmentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

		errMsg = "Failed to remove elements " + str(shipmentsIds) + " for Shipments on FulfillmentCenter"

		try:
			# get the FulfillmentCenter
			fulfillmentCenter = self.get( fulfillmentCenterId ).first()
				
			# split on a comma with no spaces
			idList = shipmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Shipment		
				shipment = ShipmentDelegate().get(id).first();	
				# add the Shipment
				fulfillmentCenter.shipments.remove(shipment)
				
			# save it		
			fulfillmentCenter.save()
			
			# reload and return the appropriate version
			return self.get( fulfillmentCenterId );
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter with id " + str(fulfillmentCenterId) + " does not exist.")
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
