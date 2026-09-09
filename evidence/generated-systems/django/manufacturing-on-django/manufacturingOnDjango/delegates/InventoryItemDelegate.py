from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.InventoryItem import InventoryItem
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.Location import Location
from manufacturingOnDjango.exceptions import Exceptions

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
		
	def assignItem( self, inventoryItemId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			inventoryItem.item = item
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.item = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, inventoryItemId, locationId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the Location from db
			location = LocationDelegate().get(locationId).first();
			
			# assign the Location		
			inventoryItem.location = location
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.location = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
