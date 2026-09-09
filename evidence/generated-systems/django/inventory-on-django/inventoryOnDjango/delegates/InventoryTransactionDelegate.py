from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.InventoryTransaction import InventoryTransaction
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.models.Reservation import Reservation
from inventoryOnDjango.models.TransferOrder import TransferOrder
from inventoryOnDjango.models.StockAdjustment import StockAdjustment
from inventoryOnDjango.models.CycleCount import CycleCount
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InventoryTransaction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryTransactionDelegate Declaration
#======================================================================
class InventoryTransactionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inventoryTransactionId ):
		try:	
			inventoryTransaction = InventoryTransaction.objects.filter(id=inventoryTransactionId)
			return inventoryTransaction.first();
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError("InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inventoryTransaction):
		for model in serializers.deserialize("json", inventoryTransaction):
			model.save()
			return model;

	def create(self, inventoryTransaction):
		inventoryTransaction.save()
		return inventoryTransaction;

	def saveFromJson(self, inventoryTransaction):
		for model in serializers.deserialize("json", inventoryTransaction):
			model.save()
			return inventoryTransaction;
	
	def save(self, inventoryTransaction):
		inventoryTransaction.save()
		return inventoryTransaction;
	
	def delete(self, inventoryTransactionId ):
		errMsg = "Failed to delete InventoryTransaction from db using id " + str(inventoryTransactionId)
		
		try:
			inventoryTransaction = InventoryTransaction.objects.get(id=inventoryTransactionId)
			inventoryTransaction.delete()
			return True
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError("InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InventoryTransaction.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InventoryTransaction from db")
		except Exception:
			return None;
		
	def assignSku( self, inventoryTransactionId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			inventoryTransaction.sku = stockKeepingUnit
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.stockKeepingUnit = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarehouse( self, inventoryTransactionId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			inventoryTransaction.warehouse = warehouse
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.warehouse = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, inventoryTransactionId, locationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(locationId).first();
			
			# assign the Location		
			inventoryTransaction.location = storageLocation
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.storageLocation = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, inventoryTransactionId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			inventoryTransaction.lot = lot
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.lot = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRelatedReservation( self, inventoryTransactionId, relatedReservationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

		errMsg = "Failed to assign element " + str(relatedReservationId) + " for RelatedReservation on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the Reservation from db
			reservation = ReservationDelegate().get(relatedReservationId).first();
			
			# assign the RelatedReservation		
			inventoryTransaction.relatedReservation = reservation
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(relatedReservationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRelatedReservation( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(relatedReservationId) + " for RelatedReservation on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.reservation = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTransferOrder( self, inventoryTransactionId, transferOrderId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

		errMsg = "Failed to assign element " + str(transferOrderId) + " for TransferOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the TransferOrder from db
			transferOrder = TransferOrderDelegate().get(transferOrderId).first();
			
			# assign the TransferOrder		
			inventoryTransaction.transferOrder = transferOrder
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTransferOrder( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(transferOrderId) + " for TransferOrder on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.transferOrder = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAdjustment( self, inventoryTransactionId, adjustmentId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockAdjustmentDelegate import StockAdjustmentDelegate

		errMsg = "Failed to assign element " + str(adjustmentId) + " for Adjustment on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the StockAdjustment from db
			stockAdjustment = StockAdjustmentDelegate().get(adjustmentId).first();
			
			# assign the Adjustment		
			inventoryTransaction.adjustment = stockAdjustment
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(adjustmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdjustment( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(adjustmentId) + " for Adjustment on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.stockAdjustment = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCycleCount( self, inventoryTransactionId, cycleCountId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.CycleCountDelegate import CycleCountDelegate

		errMsg = "Failed to assign element " + str(cycleCountId) + " for CycleCount on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# get the CycleCount from db
			cycleCount = CycleCountDelegate().get(cycleCountId).first();
			
			# assign the CycleCount		
			inventoryTransaction.cycleCount = cycleCount
			
			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCycleCount( self, inventoryTransactionId ):
		errMsg = "Failed to unassign element " + str(cycleCountId) + " for CycleCount on InventoryTransaction"

		try:
			# get the InventoryTransaction from db
			inventoryTransaction = self.get( inventoryTransactionId ).first()	
			
			# assign to None for unassignment
			inventoryTransaction.cycleCount = None			

			#save it
			inventoryTransaction.save()

			# reload and return the appropriate version					
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, inventoryTransactionId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on InventoryTransaction"

		try:
			# get the InventoryTransaction
			inventoryTransaction = self.get( inventoryTransactionId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				inventoryTransaction.serialNumbers.add(serialNumber)
				
			# save it		
			inventoryTransaction.save()
			
			# reload and return the appropriate version
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, inventoryTransactionId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on InventoryTransaction"

		try:
			# get the InventoryTransaction
			inventoryTransaction = self.get( inventoryTransactionId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				inventoryTransaction.serialNumbers.remove(serialNumber)
				
			# save it		
			inventoryTransaction.save()
			
			# reload and return the appropriate version
			return self.get( inventoryTransactionId );
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction with id " + str(inventoryTransactionId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
