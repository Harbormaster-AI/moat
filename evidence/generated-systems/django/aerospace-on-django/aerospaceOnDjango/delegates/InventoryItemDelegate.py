from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.InventoryItem import InventoryItem
from aerospaceOnDjango.models.Component_ import Component_
from aerospaceOnDjango.models.Warehouse import Warehouse
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItemDelegate Declaration
#======================================================================
class InventoryItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inventoryItemId ):
		try:	
			inventoryItem = InventoryItem.objects.filter(id=inventoryItemId)
			return inventoryItem.first();
		except InventoryItem.DoesNotExist:
			raise ProcessingError("InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inventoryItem):
		for model in serializers.deserialize("json", inventoryItem):
			model.save()
			return model;

	def create(self, inventoryItem):
		inventoryItem.save()
		return inventoryItem;

	def saveFromJson(self, inventoryItem):
		for model in serializers.deserialize("json", inventoryItem):
			model.save()
			return inventoryItem;
	
	def save(self, inventoryItem):
		inventoryItem.save()
		return inventoryItem;
	
	def delete(self, inventoryItemId ):
		errMsg = "Failed to delete InventoryItem from db using id " + str(inventoryItemId)
		
		try:
			inventoryItem = InventoryItem.objects.get(id=inventoryItemId)
			inventoryItem.delete()
			return True
		except InventoryItem.DoesNotExist:
			raise ProcessingError("InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InventoryItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InventoryItem from db")
		except Exception:
			return None;
		
	def assignComponent( self, inventoryItemId, componentId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.Component_Delegate import Component_Delegate

		errMsg = "Failed to assign element " + str(componentId) + " for Component on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the Component_ from db
			component_ = Component_Delegate().get(componentId).first();
			
			# assign the Component		
			inventoryItem.component = component_
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Component_.DoesNotExist:
			raise ProcessingError(errMsg + " : Component_ with id " + str(componentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignComponent( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(componentId) + " for Component on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.component_ = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarehouse( self, inventoryItemId, warehouseId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			inventoryItem.warehouse = warehouse
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.warehouse = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
