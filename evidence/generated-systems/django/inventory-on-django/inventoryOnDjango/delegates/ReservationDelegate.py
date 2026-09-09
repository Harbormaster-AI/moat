from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.Reservation import Reservation
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.models.DemandSignal import DemandSignal
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Reservation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReservationDelegate Declaration
#======================================================================
class ReservationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, reservationId ):
		try:	
			reservation = Reservation.objects.filter(id=reservationId)
			return reservation.first();
		except Reservation.DoesNotExist:
			raise ProcessingError("Reservation with id " + str(reservationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, reservation):
		for model in serializers.deserialize("json", reservation):
			model.save()
			return model;

	def create(self, reservation):
		reservation.save()
		return reservation;

	def saveFromJson(self, reservation):
		for model in serializers.deserialize("json", reservation):
			model.save()
			return reservation;
	
	def save(self, reservation):
		reservation.save()
		return reservation;
	
	def delete(self, reservationId ):
		errMsg = "Failed to delete Reservation from db using id " + str(reservationId)
		
		try:
			reservation = Reservation.objects.get(id=reservationId)
			reservation.delete()
			return True
		except Reservation.DoesNotExist:
			raise ProcessingError("Reservation with id " + str(reservationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Reservation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Reservation from db")
		except Exception:
			return None;
		
	def assignSku( self, reservationId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			reservation.sku = stockKeepingUnit
			
			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, reservationId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# assign to None for unassignment
			reservation.stockKeepingUnit = None			

			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWarehouse( self, reservationId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			reservation.warehouse = warehouse
			
			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, reservationId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# assign to None for unassignment
			reservation.warehouse = None			

			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, reservationId, locationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(locationId).first();
			
			# assign the Location		
			reservation.location = storageLocation
			
			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, reservationId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# assign to None for unassignment
			reservation.storageLocation = None			

			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInventoryItem( self, reservationId, inventoryItemId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to assign element " + str(inventoryItemId) + " for InventoryItem on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# get the InventoryItem from db
			inventoryItem = InventoryItemDelegate().get(inventoryItemId).first();
			
			# assign the InventoryItem		
			reservation.inventoryItem = inventoryItem
			
			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInventoryItem( self, reservationId ):
		errMsg = "Failed to unassign element " + str(inventoryItemId) + " for InventoryItem on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# assign to None for unassignment
			reservation.inventoryItem = None			

			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, reservationId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			reservation.lot = lot
			
			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, reservationId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# assign to None for unassignment
			reservation.lot = None			

			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDemandSignal( self, reservationId, demandSignalId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.DemandSignalDelegate import DemandSignalDelegate

		errMsg = "Failed to assign element " + str(demandSignalId) + " for DemandSignal on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# get the DemandSignal from db
			demandSignal = DemandSignalDelegate().get(demandSignalId).first();
			
			# assign the DemandSignal		
			reservation.demandSignal = demandSignal
			
			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except DemandSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : DemandSignal with id " + str(demandSignalId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDemandSignal( self, reservationId ):
		errMsg = "Failed to unassign element " + str(demandSignalId) + " for DemandSignal on Reservation"

		try:
			# get the Reservation from db
			reservation = self.get( reservationId ).first()	
			
			# assign to None for unassignment
			reservation.demandSignal = None			

			#save it
			reservation.save()

			# reload and return the appropriate version					
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, reservationId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on Reservation"

		try:
			# get the Reservation
			reservation = self.get( reservationId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				reservation.serialNumbers.add(serialNumber)
				
			# save it		
			reservation.save()
			
			# reload and return the appropriate version
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, reservationId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on Reservation"

		try:
			# get the Reservation
			reservation = self.get( reservationId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				reservation.serialNumbers.remove(serialNumber)
				
			# save it		
			reservation.save()
			
			# reload and return the appropriate version
			return self.get( reservationId );
		except Reservation.DoesNotExist:
			raise ProcessingError(errMsg + " : Reservation with id " + str(reservationId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
