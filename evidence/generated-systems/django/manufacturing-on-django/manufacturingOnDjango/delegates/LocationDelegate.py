from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Location import Location
from manufacturingOnDjango.models.Warehouse import Warehouse
from manufacturingOnDjango.models.InventoryItem import InventoryItem
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Location
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LocationDelegate Declaration
#======================================================================
class LocationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, locationId ):
		try:	
			location = Location.objects.filter(id=locationId)
			return location.first();
		except Location.DoesNotExist:
			raise ProcessingError("Location with id " + str(locationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, location):
		for model in serializers.deserialize("json", location):
			model.save()
			return model;

	def create(self, location):
		location.save()
		return location;

	def saveFromJson(self, location):
		for model in serializers.deserialize("json", location):
			model.save()
			return location;
	
	def save(self, location):
		location.save()
		return location;
	
	def delete(self, locationId ):
		errMsg = "Failed to delete Location from db using id " + str(locationId)
		
		try:
			location = Location.objects.get(id=locationId)
			location.delete()
			return True
		except Location.DoesNotExist:
			raise ProcessingError("Location with id " + str(locationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Location.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Location from db")
		except Exception:
			return None;
		
	def assignWarehouse( self, locationId, warehouseId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on Location"

		try:
			# get the Location from db
			location = self.get( locationId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			location.warehouse = warehouse
			
			#save it
			location.save()

			# reload and return the appropriate version					
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, locationId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on Location"

		try:
			# get the Location from db
			location = self.get( locationId ).first()	
			
			# assign to None for unassignment
			location.warehouse = None			

			#save it
			location.save()

			# reload and return the appropriate version					
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
		
	def addInventoryItems( self, locationId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				location.inventoryItems.add(inventoryItem)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, locationId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				location.inventoryItems.remove(inventoryItem)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
