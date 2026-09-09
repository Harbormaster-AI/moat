from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.models.InboundShipment import InboundShipment
from inventoryOnDjango.models.OutboundAllocation import OutboundAllocation
from inventoryOnDjango.models.TransferOrder import TransferOrder
from inventoryOnDjango.models.CycleCount import CycleCount
from inventoryOnDjango.exceptions import Exceptions

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
		
	def addStorageLocations( self, warehouseId, storageLocationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to add elements " + str(storageLocationsIds) + " for StorageLocations on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = storageLocationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the StorageLocation		
				storageLocation = StorageLocationDelegate().get(id).first();	
				# add the StorageLocation
				warehouse.storageLocations.add(storageLocation)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeStorageLocations( self, warehouseId, storageLocationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to remove elements " + str(storageLocationsIds) + " for StorageLocations on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = storageLocationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the StorageLocation		
				storageLocation = StorageLocationDelegate().get(id).first();	
				# add the StorageLocation
				warehouse.storageLocations.remove(storageLocation)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInventoryItems( self, warehouseId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

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
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

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
		
	def addInboundShipments( self, warehouseId, inboundShipmentsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InboundShipmentDelegate import InboundShipmentDelegate

		errMsg = "Failed to add elements " + str(inboundShipmentsIds) + " for InboundShipments on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = inboundShipmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InboundShipment		
				inboundShipment = InboundShipmentDelegate().get(id).first();	
				# add the InboundShipment
				warehouse.inboundShipments.add(inboundShipment)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInboundShipments( self, warehouseId, inboundShipmentsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InboundShipmentDelegate import InboundShipmentDelegate

		errMsg = "Failed to remove elements " + str(inboundShipmentsIds) + " for InboundShipments on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = inboundShipmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InboundShipment		
				inboundShipment = InboundShipmentDelegate().get(id).first();	
				# add the InboundShipment
				warehouse.inboundShipments.remove(inboundShipment)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOutboundAllocations( self, warehouseId, outboundAllocationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.OutboundAllocationDelegate import OutboundAllocationDelegate

		errMsg = "Failed to add elements " + str(outboundAllocationsIds) + " for OutboundAllocations on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = outboundAllocationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OutboundAllocation		
				outboundAllocation = OutboundAllocationDelegate().get(id).first();	
				# add the OutboundAllocation
				warehouse.outboundAllocations.add(outboundAllocation)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOutboundAllocations( self, warehouseId, outboundAllocationsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.OutboundAllocationDelegate import OutboundAllocationDelegate

		errMsg = "Failed to remove elements " + str(outboundAllocationsIds) + " for OutboundAllocations on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = outboundAllocationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OutboundAllocation		
				outboundAllocation = OutboundAllocationDelegate().get(id).first();	
				# add the OutboundAllocation
				warehouse.outboundAllocations.remove(outboundAllocation)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except OutboundAllocation.DoesNotExist:
			raise ProcessingError(errMsg + " : OutboundAllocation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOriginTransfers( self, warehouseId, originTransfersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

		errMsg = "Failed to add elements " + str(originTransfersIds) + " for OriginTransfers on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = originTransfersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TransferOrder		
				transferOrder = TransferOrderDelegate().get(id).first();	
				# add the TransferOrder
				warehouse.originTransfers.add(transferOrder)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOriginTransfers( self, warehouseId, originTransfersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

		errMsg = "Failed to remove elements " + str(originTransfersIds) + " for OriginTransfers on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = originTransfersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TransferOrder		
				transferOrder = TransferOrderDelegate().get(id).first();	
				# add the TransferOrder
				warehouse.originTransfers.remove(transferOrder)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDestinationTransfers( self, warehouseId, destinationTransfersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

		errMsg = "Failed to add elements " + str(destinationTransfersIds) + " for DestinationTransfers on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = destinationTransfersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TransferOrder		
				transferOrder = TransferOrderDelegate().get(id).first();	
				# add the TransferOrder
				warehouse.destinationTransfers.add(transferOrder)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDestinationTransfers( self, warehouseId, destinationTransfersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

		errMsg = "Failed to remove elements " + str(destinationTransfersIds) + " for DestinationTransfers on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = destinationTransfersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TransferOrder		
				transferOrder = TransferOrderDelegate().get(id).first();	
				# add the TransferOrder
				warehouse.destinationTransfers.remove(transferOrder)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCycleCounts( self, warehouseId, cycleCountsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.CycleCountDelegate import CycleCountDelegate

		errMsg = "Failed to add elements " + str(cycleCountsIds) + " for CycleCounts on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = cycleCountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CycleCount		
				cycleCount = CycleCountDelegate().get(id).first();	
				# add the CycleCount
				warehouse.cycleCounts.add(cycleCount)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCycleCounts( self, warehouseId, cycleCountsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.CycleCountDelegate import CycleCountDelegate

		errMsg = "Failed to remove elements " + str(cycleCountsIds) + " for CycleCounts on Warehouse"

		try:
			# get the Warehouse
			warehouse = self.get( warehouseId ).first()
				
			# split on a comma with no spaces
			idList = cycleCountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CycleCount		
				cycleCount = CycleCountDelegate().get(id).first();	
				# add the CycleCount
				warehouse.cycleCounts.remove(cycleCount)
				
			# save it		
			warehouse.save()
			
			# reload and return the appropriate version
			return self.get( warehouseId );
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
