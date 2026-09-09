from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.StockAdjustment import StockAdjustment
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.StockAdjustmentLine import StockAdjustmentLine
from inventoryOnDjango.models.InventoryTransaction import InventoryTransaction
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model StockAdjustment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockAdjustmentDelegate Declaration
#======================================================================
class StockAdjustmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, stockAdjustmentId ):
		try:	
			stockAdjustment = StockAdjustment.objects.filter(id=stockAdjustmentId)
			return stockAdjustment.first();
		except StockAdjustment.DoesNotExist:
			raise ProcessingError("StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, stockAdjustment):
		for model in serializers.deserialize("json", stockAdjustment):
			model.save()
			return model;

	def create(self, stockAdjustment):
		stockAdjustment.save()
		return stockAdjustment;

	def saveFromJson(self, stockAdjustment):
		for model in serializers.deserialize("json", stockAdjustment):
			model.save()
			return stockAdjustment;
	
	def save(self, stockAdjustment):
		stockAdjustment.save()
		return stockAdjustment;
	
	def delete(self, stockAdjustmentId ):
		errMsg = "Failed to delete StockAdjustment from db using id " + str(stockAdjustmentId)
		
		try:
			stockAdjustment = StockAdjustment.objects.get(id=stockAdjustmentId)
			stockAdjustment.delete()
			return True
		except StockAdjustment.DoesNotExist:
			raise ProcessingError("StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = StockAdjustment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all StockAdjustment from db")
		except Exception:
			return None;
		
	def assignWarehouse( self, stockAdjustmentId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on StockAdjustment"

		try:
			# get the StockAdjustment from db
			stockAdjustment = self.get( stockAdjustmentId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			stockAdjustment.warehouse = warehouse
			
			#save it
			stockAdjustment.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentId );
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, stockAdjustmentId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on StockAdjustment"

		try:
			# get the StockAdjustment from db
			stockAdjustment = self.get( stockAdjustmentId ).first()	
			
			# assign to None for unassignment
			stockAdjustment.warehouse = None			

			#save it
			stockAdjustment.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentId );
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addLines( self, stockAdjustmentId, linesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockAdjustmentLineDelegate import StockAdjustmentLineDelegate

		errMsg = "Failed to add elements " + str(linesIds) + " for Lines on StockAdjustment"

		try:
			# get the StockAdjustment
			stockAdjustment = self.get( stockAdjustmentId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the StockAdjustmentLine		
				stockAdjustmentLine = StockAdjustmentLineDelegate().get(id).first();	
				# add the StockAdjustmentLine
				stockAdjustment.lines.add(stockAdjustmentLine)
				
			# save it		
			stockAdjustment.save()
			
			# reload and return the appropriate version
			return self.get( stockAdjustmentId );
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLines( self, stockAdjustmentId, linesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockAdjustmentLineDelegate import StockAdjustmentLineDelegate

		errMsg = "Failed to remove elements " + str(linesIds) + " for Lines on StockAdjustment"

		try:
			# get the StockAdjustment
			stockAdjustment = self.get( stockAdjustmentId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the StockAdjustmentLine		
				stockAdjustmentLine = StockAdjustmentLineDelegate().get(id).first();	
				# add the StockAdjustmentLine
				stockAdjustment.lines.remove(stockAdjustmentLine)
				
			# save it		
			stockAdjustment.save()
			
			# reload and return the appropriate version
			return self.get( stockAdjustmentId );
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTransactions( self, stockAdjustmentId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on StockAdjustment"

		try:
			# get the StockAdjustment
			stockAdjustment = self.get( stockAdjustmentId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				stockAdjustment.transactions.add(inventoryTransaction)
				
			# save it		
			stockAdjustment.save()
			
			# reload and return the appropriate version
			return self.get( stockAdjustmentId );
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, stockAdjustmentId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on StockAdjustment"

		try:
			# get the StockAdjustment
			stockAdjustment = self.get( stockAdjustmentId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				stockAdjustment.transactions.remove(inventoryTransaction)
				
			# save it		
			stockAdjustment.save()
			
			# reload and return the appropriate version
			return self.get( stockAdjustmentId );
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(stockAdjustmentId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
