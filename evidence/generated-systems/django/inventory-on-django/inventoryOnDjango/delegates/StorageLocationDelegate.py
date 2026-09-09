from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model StorageLocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StorageLocationDelegate Declaration
#======================================================================
class StorageLocationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, storageLocationId ):
		try:	
			storageLocation = StorageLocation.objects.filter(id=storageLocationId)
			return storageLocation.first();
		except StorageLocation.DoesNotExist:
			raise ProcessingError("StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, storageLocation):
		for model in serializers.deserialize("json", storageLocation):
			model.save()
			return model;

	def create(self, storageLocation):
		storageLocation.save()
		return storageLocation;

	def saveFromJson(self, storageLocation):
		for model in serializers.deserialize("json", storageLocation):
			model.save()
			return storageLocation;
	
	def save(self, storageLocation):
		storageLocation.save()
		return storageLocation;
	
	def delete(self, storageLocationId ):
		errMsg = "Failed to delete StorageLocation from db using id " + str(storageLocationId)
		
		try:
			storageLocation = StorageLocation.objects.get(id=storageLocationId)
			storageLocation.delete()
			return True
		except StorageLocation.DoesNotExist:
			raise ProcessingError("StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = StorageLocation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all StorageLocation from db")
		except Exception:
			return None;
		
	def assignWarehouse( self, storageLocationId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on StorageLocation"

		try:
			# get the StorageLocation from db
			storageLocation = self.get( storageLocationId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			storageLocation.warehouse = warehouse
			
			#save it
			storageLocation.save()

			# reload and return the appropriate version					
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, storageLocationId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on StorageLocation"

		try:
			# get the StorageLocation from db
			storageLocation = self.get( storageLocationId ).first()	
			
			# assign to None for unassignment
			storageLocation.warehouse = None			

			#save it
			storageLocation.save()

			# reload and return the appropriate version					
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignParentLocation( self, storageLocationId, parentLocationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(parentLocationId) + " for ParentLocation on StorageLocation"

		try:
			# get the StorageLocation from db
			storageLocation = self.get( storageLocationId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(parentLocationId).first();
			
			# assign the ParentLocation		
			storageLocation.parentLocation = storageLocation
			
			#save it
			storageLocation.save()

			# reload and return the appropriate version					
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(parentLocationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignParentLocation( self, storageLocationId ):
		errMsg = "Failed to unassign element " + str(parentLocationId) + " for ParentLocation on StorageLocation"

		try:
			# get the StorageLocation from db
			storageLocation = self.get( storageLocationId ).first()	
			
			# assign to None for unassignment
			storageLocation.storageLocation = None			

			#save it
			storageLocation.save()

			# reload and return the appropriate version					
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except Exception:
			return None;
		
	def addChildLocations( self, storageLocationId, childLocationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to add elements " + str(childLocationsIds) + " for ChildLocations on StorageLocation"

		try:
			# get the StorageLocation
			storageLocation = self.get( storageLocationId ).first()
				
			# split on a comma with no spaces
			idList = childLocationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the StorageLocation		
				storageLocation = StorageLocationDelegate().get(id).first();	
				# add the StorageLocation
				storageLocation.childLocations.add(storageLocation)
				
			# save it		
			storageLocation.save()
			
			# reload and return the appropriate version
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChildLocations( self, storageLocationId, childLocationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to remove elements " + str(childLocationsIds) + " for ChildLocations on StorageLocation"

		try:
			# get the StorageLocation
			storageLocation = self.get( storageLocationId ).first()
				
			# split on a comma with no spaces
			idList = childLocationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the StorageLocation		
				storageLocation = StorageLocationDelegate().get(id).first();	
				# add the StorageLocation
				storageLocation.childLocations.remove(storageLocation)
				
			# save it		
			storageLocation.save()
			
			# reload and return the appropriate version
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInventoryItems( self, storageLocationId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on StorageLocation"

		try:
			# get the StorageLocation
			storageLocation = self.get( storageLocationId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				storageLocation.inventoryItems.add(inventoryItem)
				
			# save it		
			storageLocation.save()
			
			# reload and return the appropriate version
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, storageLocationId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on StorageLocation"

		try:
			# get the StorageLocation
			storageLocation = self.get( storageLocationId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				storageLocation.inventoryItems.remove(inventoryItem)
				
			# save it		
			storageLocation.save()
			
			# reload and return the appropriate version
			return self.get( storageLocationId );
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(storageLocationId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
