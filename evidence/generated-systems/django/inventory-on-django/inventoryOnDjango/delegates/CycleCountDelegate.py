from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.CycleCount import CycleCount
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.CycleCountEntry import CycleCountEntry
from inventoryOnDjango.models.InventoryTransaction import InventoryTransaction
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CycleCount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCountDelegate Declaration
#======================================================================
class CycleCountDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, cycleCountId ):
		try:	
			cycleCount = CycleCount.objects.filter(id=cycleCountId)
			return cycleCount.first();
		except CycleCount.DoesNotExist:
			raise ProcessingError("CycleCount with id " + str(cycleCountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, cycleCount):
		for model in serializers.deserialize("json", cycleCount):
			model.save()
			return model;

	def create(self, cycleCount):
		cycleCount.save()
		return cycleCount;

	def saveFromJson(self, cycleCount):
		for model in serializers.deserialize("json", cycleCount):
			model.save()
			return cycleCount;
	
	def save(self, cycleCount):
		cycleCount.save()
		return cycleCount;
	
	def delete(self, cycleCountId ):
		errMsg = "Failed to delete CycleCount from db using id " + str(cycleCountId)
		
		try:
			cycleCount = CycleCount.objects.get(id=cycleCountId)
			cycleCount.delete()
			return True
		except CycleCount.DoesNotExist:
			raise ProcessingError("CycleCount with id " + str(cycleCountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CycleCount.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CycleCount from db")
		except Exception:
			return None;
		
	def assignWarehouse( self, cycleCountId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on CycleCount"

		try:
			# get the CycleCount from db
			cycleCount = self.get( cycleCountId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			cycleCount.warehouse = warehouse
			
			#save it
			cycleCount.save()

			# reload and return the appropriate version					
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, cycleCountId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on CycleCount"

		try:
			# get the CycleCount from db
			cycleCount = self.get( cycleCountId ).first()	
			
			# assign to None for unassignment
			cycleCount.warehouse = None			

			#save it
			cycleCount.save()

			# reload and return the appropriate version					
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except Exception:
			return None;
		
	def addLocations( self, cycleCountId, locationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to add elements " + str(locationsIds) + " for Locations on CycleCount"

		try:
			# get the CycleCount
			cycleCount = self.get( cycleCountId ).first()
				
			# split on a comma with no spaces
			idList = locationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the StorageLocation		
				storageLocation = StorageLocationDelegate().get(id).first();	
				# add the StorageLocation
				cycleCount.locations.add(storageLocation)
				
			# save it		
			cycleCount.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLocations( self, cycleCountId, locationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to remove elements " + str(locationsIds) + " for Locations on CycleCount"

		try:
			# get the CycleCount
			cycleCount = self.get( cycleCountId ).first()
				
			# split on a comma with no spaces
			idList = locationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the StorageLocation		
				storageLocation = StorageLocationDelegate().get(id).first();	
				# add the StorageLocation
				cycleCount.locations.remove(storageLocation)
				
			# save it		
			cycleCount.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEntries( self, cycleCountId, entriesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.CycleCountEntryDelegate import CycleCountEntryDelegate

		errMsg = "Failed to add elements " + str(entriesIds) + " for Entries on CycleCount"

		try:
			# get the CycleCount
			cycleCount = self.get( cycleCountId ).first()
				
			# split on a comma with no spaces
			idList = entriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CycleCountEntry		
				cycleCountEntry = CycleCountEntryDelegate().get(id).first();	
				# add the CycleCountEntry
				cycleCount.entries.add(cycleCountEntry)
				
			# save it		
			cycleCount.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEntries( self, cycleCountId, entriesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.CycleCountEntryDelegate import CycleCountEntryDelegate

		errMsg = "Failed to remove elements " + str(entriesIds) + " for Entries on CycleCount"

		try:
			# get the CycleCount
			cycleCount = self.get( cycleCountId ).first()
				
			# split on a comma with no spaces
			idList = entriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CycleCountEntry		
				cycleCountEntry = CycleCountEntryDelegate().get(id).first();	
				# add the CycleCountEntry
				cycleCount.entries.remove(cycleCountEntry)
				
			# save it		
			cycleCount.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTransactions( self, cycleCountId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on CycleCount"

		try:
			# get the CycleCount
			cycleCount = self.get( cycleCountId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				cycleCount.transactions.add(inventoryTransaction)
				
			# save it		
			cycleCount.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, cycleCountId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on CycleCount"

		try:
			# get the CycleCount
			cycleCount = self.get( cycleCountId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				cycleCount.transactions.remove(inventoryTransaction)
				
			# save it		
			cycleCount.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountId );
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
