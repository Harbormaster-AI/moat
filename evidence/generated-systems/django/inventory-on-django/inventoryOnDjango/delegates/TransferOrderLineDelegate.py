from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.TransferOrderLine import TransferOrderLine
from inventoryOnDjango.models.TransferOrder import TransferOrder
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TransferOrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderLineDelegate Declaration
#======================================================================
class TransferOrderLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, transferOrderLineId ):
		try:	
			transferOrderLine = TransferOrderLine.objects.filter(id=transferOrderLineId)
			return transferOrderLine.first();
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError("TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, transferOrderLine):
		for model in serializers.deserialize("json", transferOrderLine):
			model.save()
			return model;

	def create(self, transferOrderLine):
		transferOrderLine.save()
		return transferOrderLine;

	def saveFromJson(self, transferOrderLine):
		for model in serializers.deserialize("json", transferOrderLine):
			model.save()
			return transferOrderLine;
	
	def save(self, transferOrderLine):
		transferOrderLine.save()
		return transferOrderLine;
	
	def delete(self, transferOrderLineId ):
		errMsg = "Failed to delete TransferOrderLine from db using id " + str(transferOrderLineId)
		
		try:
			transferOrderLine = TransferOrderLine.objects.get(id=transferOrderLineId)
			transferOrderLine.delete()
			return True
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError("TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TransferOrderLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TransferOrderLine from db")
		except Exception:
			return None;
		
	def assignTransferOrder( self, transferOrderLineId, transferOrderId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderDelegate import TransferOrderDelegate

		errMsg = "Failed to assign element " + str(transferOrderId) + " for TransferOrder on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# get the TransferOrder from db
			transferOrder = TransferOrderDelegate().get(transferOrderId).first();
			
			# assign the TransferOrder		
			transferOrderLine.transferOrder = transferOrder
			
			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTransferOrder( self, transferOrderLineId ):
		errMsg = "Failed to unassign element " + str(transferOrderId) + " for TransferOrder on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# assign to None for unassignment
			transferOrderLine.transferOrder = None			

			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSku( self, transferOrderLineId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			transferOrderLine.sku = stockKeepingUnit
			
			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, transferOrderLineId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# assign to None for unassignment
			transferOrderLine.stockKeepingUnit = None			

			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, transferOrderLineId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			transferOrderLine.lot = lot
			
			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, transferOrderLineId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# assign to None for unassignment
			transferOrderLine.lot = None			

			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFromLocation( self, transferOrderLineId, fromLocationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(fromLocationId) + " for FromLocation on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(fromLocationId).first();
			
			# assign the FromLocation		
			transferOrderLine.fromLocation = storageLocation
			
			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(fromLocationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFromLocation( self, transferOrderLineId ):
		errMsg = "Failed to unassign element " + str(fromLocationId) + " for FromLocation on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# assign to None for unassignment
			transferOrderLine.storageLocation = None			

			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignToLocation( self, transferOrderLineId, toLocationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(toLocationId) + " for ToLocation on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(toLocationId).first();
			
			# assign the ToLocation		
			transferOrderLine.toLocation = storageLocation
			
			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(toLocationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignToLocation( self, transferOrderLineId ):
		errMsg = "Failed to unassign element " + str(toLocationId) + " for ToLocation on TransferOrderLine"

		try:
			# get the TransferOrderLine from db
			transferOrderLine = self.get( transferOrderLineId ).first()	
			
			# assign to None for unassignment
			transferOrderLine.storageLocation = None			

			#save it
			transferOrderLine.save()

			# reload and return the appropriate version					
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, transferOrderLineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on TransferOrderLine"

		try:
			# get the TransferOrderLine
			transferOrderLine = self.get( transferOrderLineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				transferOrderLine.serialNumbers.add(serialNumber)
				
			# save it		
			transferOrderLine.save()
			
			# reload and return the appropriate version
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, transferOrderLineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on TransferOrderLine"

		try:
			# get the TransferOrderLine
			transferOrderLine = self.get( transferOrderLineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				transferOrderLine.serialNumbers.remove(serialNumber)
				
			# save it		
			transferOrderLine.save()
			
			# reload and return the appropriate version
			return self.get( transferOrderLineId );
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine with id " + str(transferOrderLineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
