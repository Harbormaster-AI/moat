from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Warehouse import Warehouse
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.Location import Location
from manufacturingOnDjango.models.InventoryItem import InventoryItem
from manufacturingOnDjango.exceptions import Exceptions

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
		
	def assignPlant( self, warehouseId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on Warehouse"

		try:
			# get the Warehouse from db
			warehouse = self.get( warehouseId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			warehouse.plant = plant
			
			#save it
			warehouse.save()

			# reload and return the appropriate version					
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, warehouseId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on Warehouse"

		try:
			# get the Warehouse from db
			warehouse = self.get( warehouseId ).first()	
			
			# assign to None for unassignment
			warehouse.plant = None			

			#save it
			warehouse.save()

			# reload and return the appropriate version					
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
		
	def addLocations( self, warehouseId, locationsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to add elements " + str(locationsIds) + " for Locations on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = locationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Location		
				location = LocationDelegate().get(id).first();	
				# add the Location
				warehouse.locations.add(location)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLocations( self, warehouseId, locationsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to remove elements " + str(locationsIds) + " for Locations on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = locationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Location		
				location = LocationDelegate().get(id).first();	
				# add the Location
				warehouse.locations.remove(location)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInventoryItems( self, warehouseId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

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
		from manufacturingOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

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
		
