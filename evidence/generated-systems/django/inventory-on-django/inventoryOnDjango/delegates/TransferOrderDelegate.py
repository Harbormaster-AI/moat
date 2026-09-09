from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.TransferOrder import TransferOrder
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.TransferOrderLine import TransferOrderLine
from inventoryOnDjango.models.InventoryTransaction import InventoryTransaction
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TransferOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TransferOrderDelegate Declaration
#======================================================================
class TransferOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, transferOrderId ):
		try:	
			transferOrder = TransferOrder.objects.filter(id=transferOrderId)
			return transferOrder.first();
		except TransferOrder.DoesNotExist:
			raise ProcessingError("TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, transferOrder):
		for model in serializers.deserialize("json", transferOrder):
			model.save()
			return model;

	def create(self, transferOrder):
		transferOrder.save()
		return transferOrder;

	def saveFromJson(self, transferOrder):
		for model in serializers.deserialize("json", transferOrder):
			model.save()
			return transferOrder;
	
	def save(self, transferOrder):
		transferOrder.save()
		return transferOrder;
	
	def delete(self, transferOrderId ):
		errMsg = "Failed to delete TransferOrder from db using id " + str(transferOrderId)
		
		try:
			transferOrder = TransferOrder.objects.get(id=transferOrderId)
			transferOrder.delete()
			return True
		except TransferOrder.DoesNotExist:
			raise ProcessingError("TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TransferOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TransferOrder from db")
		except Exception:
			return None;
		
	def assignOriginWarehouse( self, transferOrderId, originWarehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(originWarehouseId) + " for OriginWarehouse on TransferOrder"

		try:
			# get the TransferOrder from db
			transferOrder = self.get( transferOrderId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(originWarehouseId).first();
			
			# assign the OriginWarehouse		
			transferOrder.originWarehouse = warehouse
			
			#save it
			transferOrder.save()

			# reload and return the appropriate version					
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(originWarehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOriginWarehouse( self, transferOrderId ):
		errMsg = "Failed to unassign element " + str(originWarehouseId) + " for OriginWarehouse on TransferOrder"

		try:
			# get the TransferOrder from db
			transferOrder = self.get( transferOrderId ).first()	
			
			# assign to None for unassignment
			transferOrder.warehouse = None			

			#save it
			transferOrder.save()

			# reload and return the appropriate version					
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDestinationWarehouse( self, transferOrderId, destinationWarehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(destinationWarehouseId) + " for DestinationWarehouse on TransferOrder"

		try:
			# get the TransferOrder from db
			transferOrder = self.get( transferOrderId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(destinationWarehouseId).first();
			
			# assign the DestinationWarehouse		
			transferOrder.destinationWarehouse = warehouse
			
			#save it
			transferOrder.save()

			# reload and return the appropriate version					
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(destinationWarehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDestinationWarehouse( self, transferOrderId ):
		errMsg = "Failed to unassign element " + str(destinationWarehouseId) + " for DestinationWarehouse on TransferOrder"

		try:
			# get the TransferOrder from db
			transferOrder = self.get( transferOrderId ).first()	
			
			# assign to None for unassignment
			transferOrder.warehouse = None			

			#save it
			transferOrder.save()

			# reload and return the appropriate version					
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addLines( self, transferOrderId, linesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderLineDelegate import TransferOrderLineDelegate

		errMsg = "Failed to add elements " + str(linesIds) + " for Lines on TransferOrder"

		try:
			# get the TransferOrder
			transferOrder = self.get( transferOrderId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TransferOrderLine		
				transferOrderLine = TransferOrderLineDelegate().get(id).first();	
				# add the TransferOrderLine
				transferOrder.lines.add(transferOrderLine)
				
			# save it		
			transferOrder.save()
			
			# reload and return the appropriate version
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLines( self, transferOrderId, linesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.TransferOrderLineDelegate import TransferOrderLineDelegate

		errMsg = "Failed to remove elements " + str(linesIds) + " for Lines on TransferOrder"

		try:
			# get the TransferOrder
			transferOrder = self.get( transferOrderId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TransferOrderLine		
				transferOrderLine = TransferOrderLineDelegate().get(id).first();	
				# add the TransferOrderLine
				transferOrder.lines.remove(transferOrderLine)
				
			# save it		
			transferOrder.save()
			
			# reload and return the appropriate version
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except TransferOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrderLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTransactions( self, transferOrderId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on TransferOrder"

		try:
			# get the TransferOrder
			transferOrder = self.get( transferOrderId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				transferOrder.transactions.add(inventoryTransaction)
				
			# save it		
			transferOrder.save()
			
			# reload and return the appropriate version
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, transferOrderId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on TransferOrder"

		try:
			# get the TransferOrder
			transferOrder = self.get( transferOrderId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				transferOrder.transactions.remove(inventoryTransaction)
				
			# save it		
			transferOrder.save()
			
			# reload and return the appropriate version
			return self.get( transferOrderId );
		except TransferOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TransferOrder with id " + str(transferOrderId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
