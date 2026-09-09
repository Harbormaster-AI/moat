from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.Quarantine import Quarantine
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Quarantine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuarantineDelegate Declaration
#======================================================================
class QuarantineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, quarantineId ):
		try:	
			quarantine = Quarantine.objects.filter(id=quarantineId)
			return quarantine.first();
		except Quarantine.DoesNotExist:
			raise ProcessingError("Quarantine with id " + str(quarantineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, quarantine):
		for model in serializers.deserialize("json", quarantine):
			model.save()
			return model;

	def create(self, quarantine):
		quarantine.save()
		return quarantine;

	def saveFromJson(self, quarantine):
		for model in serializers.deserialize("json", quarantine):
			model.save()
			return quarantine;
	
	def save(self, quarantine):
		quarantine.save()
		return quarantine;
	
	def delete(self, quarantineId ):
		errMsg = "Failed to delete Quarantine from db using id " + str(quarantineId)
		
		try:
			quarantine = Quarantine.objects.get(id=quarantineId)
			quarantine.delete()
			return True
		except Quarantine.DoesNotExist:
			raise ProcessingError("Quarantine with id " + str(quarantineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Quarantine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Quarantine from db")
		except Exception:
			return None;
		
	def assignWarehouse( self, quarantineId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on Quarantine"

		try:
			# get the Quarantine from db
			quarantine = self.get( quarantineId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			quarantine.warehouse = warehouse
			
			#save it
			quarantine.save()

			# reload and return the appropriate version					
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, quarantineId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on Quarantine"

		try:
			# get the Quarantine from db
			quarantine = self.get( quarantineId ).first()	
			
			# assign to None for unassignment
			quarantine.warehouse = None			

			#save it
			quarantine.save()

			# reload and return the appropriate version					
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, quarantineId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on Quarantine"

		try:
			# get the Quarantine from db
			quarantine = self.get( quarantineId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			quarantine.lot = lot
			
			#save it
			quarantine.save()

			# reload and return the appropriate version					
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, quarantineId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on Quarantine"

		try:
			# get the Quarantine from db
			quarantine = self.get( quarantineId ).first()	
			
			# assign to None for unassignment
			quarantine.lot = None			

			#save it
			quarantine.save()

			# reload and return the appropriate version					
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except Exception:
			return None;
		
	def addItems( self, quarantineId, itemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(itemsIds) + " for Items on Quarantine"

		try:
			# get the Quarantine
			quarantine = self.get( quarantineId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				quarantine.items.add(inventoryItem)
				
			# save it		
			quarantine.save()
			
			# reload and return the appropriate version
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeItems( self, quarantineId, itemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(itemsIds) + " for Items on Quarantine"

		try:
			# get the Quarantine
			quarantine = self.get( quarantineId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				quarantine.items.remove(inventoryItem)
				
			# save it		
			quarantine.save()
			
			# reload and return the appropriate version
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSerialNumbers( self, quarantineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on Quarantine"

		try:
			# get the Quarantine
			quarantine = self.get( quarantineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				quarantine.serialNumbers.add(serialNumber)
				
			# save it		
			quarantine.save()
			
			# reload and return the appropriate version
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, quarantineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on Quarantine"

		try:
			# get the Quarantine
			quarantine = self.get( quarantineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				quarantine.serialNumbers.remove(serialNumber)
				
			# save it		
			quarantine.save()
			
			# reload and return the appropriate version
			return self.get( quarantineId );
		except Quarantine.DoesNotExist:
			raise ProcessingError(errMsg + " : Quarantine with id " + str(quarantineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
