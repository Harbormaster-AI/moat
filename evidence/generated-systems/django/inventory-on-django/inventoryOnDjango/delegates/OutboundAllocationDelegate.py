from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.OutboundAllocation import OutboundAllocation
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.models.Reservation import Reservation
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model OutboundAllocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OutboundAllocationDelegate Declaration
#======================================================================
class OutboundAllocationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, outboundAllocationId ):
		try:	
			outboundAllocation = OutboundAllocation.objects.filter(id=outboundAllocationId)
			return outboundAllocation.first();
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError("OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, outboundAllocation):
		for model in serializers.deserialize("json", outboundAllocation):
			model.save()
			return model;

	def create(self, outboundAllocation):
		outboundAllocation.save()
		return outboundAllocation;

	def saveFromJson(self, outboundAllocation):
		for model in serializers.deserialize("json", outboundAllocation):
			model.save()
			return outboundAllocation;
	
	def save(self, outboundAllocation):
		outboundAllocation.save()
		return outboundAllocation;
	
	def delete(self, outboundAllocationId ):
		errMsg = "Failed to delete OutboundAllocation from db using id " + str(outboundAllocationId)
		
		try:
			outboundAllocation = OutboundAllocation.objects.get(id=outboundAllocationId)
			outboundAllocation.delete()
			return True
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError("OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = OutboundAllocation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all OutboundAllocation from db")
		except Exception:
			return None;
		
	def assignWarehouse( self, outboundAllocationId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			outboundAllocation.warehouse = warehouse
			
			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, outboundAllocationId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# assign to None for unassignment
			outboundAllocation.warehouse = None			

			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSku( self, outboundAllocationId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			outboundAllocation.sku = stockKeepingUnit
			
			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, outboundAllocationId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# assign to None for unassignment
			outboundAllocation.stockKeepingUnit = None			

			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInventoryItem( self, outboundAllocationId, inventoryItemId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to assign element " + str(inventoryItemId) + " for InventoryItem on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# get the InventoryItem from db
			inventoryItem = InventoryItemDelegate().get(inventoryItemId).first();
			
			# assign the InventoryItem		
			outboundAllocation.inventoryItem = inventoryItem
			
			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInventoryItem( self, outboundAllocationId ):
		errMsg = "Failed to unassign element " + str(inventoryItemId) + " for InventoryItem on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# assign to None for unassignment
			outboundAllocation.inventoryItem = None			

			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignReservation( self, outboundAllocationId, reservationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

		errMsg = "Failed to assign element " + str(reservationId) + " for Reservation on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# get the Reservation from db
			reservation = ReservationDelegate().get(reservationId).first();
			
			# assign the Reservation		
			outboundAllocation.reservation = reservation
			
			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignReservation( self, outboundAllocationId ):
		errMsg = "Failed to unassign element " + str(reservationId) + " for Reservation on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# assign to None for unassignment
			outboundAllocation.reservation = None			

			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, outboundAllocationId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			outboundAllocation.lot = lot
			
			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, outboundAllocationId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# assign to None for unassignment
			outboundAllocation.lot = None			

			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSourceLocation( self, outboundAllocationId, sourceLocationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(sourceLocationId) + " for SourceLocation on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(sourceLocationId).first();
			
			# assign the SourceLocation		
			outboundAllocation.sourceLocation = storageLocation
			
			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(sourceLocationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSourceLocation( self, outboundAllocationId ):
		errMsg = "Failed to unassign element " + str(sourceLocationId) + " for SourceLocation on OutboundAllocation"

		try:
			# get the OutboundAllocation from db
			outboundAllocation = self.get( outboundAllocationId ).first()	
			
			# assign to None for unassignment
			outboundAllocation.storageLocation = None			

			#save it
			outboundAllocation.save()

			# reload and return the appropriate version					
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, outboundAllocationId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on OutboundAllocation"

		try:
			# get the OutboundAllocation
			outboundAllocation = self.get( outboundAllocationId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				outboundAllocation.serialNumbers.add(serialNumber)
				
			# save it		
			outboundAllocation.save()
			
			# reload and return the appropriate version
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, outboundAllocationId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on OutboundAllocation"

		try:
			# get the OutboundAllocation
			outboundAllocation = self.get( outboundAllocationId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				outboundAllocation.serialNumbers.remove(serialNumber)
				
			# save it		
			outboundAllocation.save()
			
			# reload and return the appropriate version
			return self.get( outboundAllocationId );
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation with id " + str(outboundAllocationId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
