from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.models.InventoryTransaction import InventoryTransaction
from inventoryOnDjango.models.Reservation import Reservation
from inventoryOnDjango.exceptions import Exceptions

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
		
	def assignSku( self, inventoryItemId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			inventoryItem.sku = stockKeepingUnit
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.stockKeepingUnit = None			

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
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

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
		
	def assignLocation( self, inventoryItemId, locationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(locationId).first();
			
			# assign the Location		
			inventoryItem.location = storageLocation
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.storageLocation = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, inventoryItemId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			inventoryItem.lot = lot
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.lot = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, inventoryItemId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on InventoryItem"

		try:
			# get the InventoryItem
			inventoryItem = self.get( inventoryItemId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				inventoryItem.serialNumbers.add(serialNumber)
				
			# save it		
			inventoryItem.save()
			
			# reload and return the appropriate version
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, inventoryItemId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on InventoryItem"

		try:
			# get the InventoryItem
			inventoryItem = self.get( inventoryItemId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				inventoryItem.serialNumbers.remove(serialNumber)
				
			# save it		
			inventoryItem.save()
			
			# reload and return the appropriate version
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTransactions( self, inventoryItemId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on InventoryItem"

		try:
			# get the InventoryItem
			inventoryItem = self.get( inventoryItemId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				inventoryItem.transactions.add(inventoryTransaction)
				
			# save it		
			inventoryItem.save()
			
			# reload and return the appropriate version
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, inventoryItemId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on InventoryItem"

		try:
			# get the InventoryItem
			inventoryItem = self.get( inventoryItemId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				inventoryItem.transactions.remove(inventoryTransaction)
				
			# save it		
			inventoryItem.save()
			
			# reload and return the appropriate version
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReservations( self, inventoryItemId, reservationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

		errMsg = "Failed to add elements " + str(reservationsIds) + " for Reservations on InventoryItem"

		try:
			# get the InventoryItem
			inventoryItem = self.get( inventoryItemId ).first()
				
			# split on a comma with no spaces
			idList = reservationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Reservation		
				reservation = ReservationDelegate().get(id).first();	
				# add the Reservation
				inventoryItem.reservations.add(reservation)
				
			# save it		
			inventoryItem.save()
			
			# reload and return the appropriate version
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReservations( self, inventoryItemId, reservationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

		errMsg = "Failed to remove elements " + str(reservationsIds) + " for Reservations on InventoryItem"

		try:
			# get the InventoryItem
			inventoryItem = self.get( inventoryItemId ).first()
				
			# split on a comma with no spaces
			idList = reservationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Reservation		
				reservation = ReservationDelegate().get(id).first();	
				# add the Reservation
				inventoryItem.reservations.remove(reservation)
				
			# save it		
			inventoryItem.save()
			
			# reload and return the appropriate version
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
