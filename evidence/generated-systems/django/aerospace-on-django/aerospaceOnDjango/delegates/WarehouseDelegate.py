from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.Warehouse import Warehouse
from aerospaceOnDjango.models.InventoryItem import InventoryItem
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Warehouse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarehouseDelegate Declaration
#======================================================================
class WarehouseDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, warehouseId ):
		try:	
			warehouse = Warehouse.objects.filter(id=warehouseId)
			return warehouse.first();
		except Warehouse.DoesNotExist:
			raise ProcessingError("Warehouse with id " + str(warehouseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, warehouse):
		for model in serializers.deserialize("json", warehouse):
			model.save()
			return model;

	def create(self, warehouse):
		warehouse.save()
		return warehouse;

	def saveFromJson(self, warehouse):
		for model in serializers.deserialize("json", warehouse):
			model.save()
			return warehouse;
	
	def save(self, warehouse):
		warehouse.save()
		return warehouse;
	
	def delete(self, warehouseId ):
		errMsg = "Failed to delete Warehouse from db using id " + str(warehouseId)
		
		try:
			warehouse = Warehouse.objects.get(id=warehouseId)
			warehouse.delete()
			return True
		except Warehouse.DoesNotExist:
			raise ProcessingError("Warehouse with id " + str(warehouseId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Warehouse.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Warehouse from db")
		except Exception:
			return None;
		
	def addInventoryItems( self, warehouseId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				warehouse.inventoryItems.add(inventoryItem)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, warehouseId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				warehouse.inventoryItems.remove(inventoryItem)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
