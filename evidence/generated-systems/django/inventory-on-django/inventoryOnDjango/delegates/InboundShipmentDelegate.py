from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.InboundShipment import InboundShipment
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.models.InboundShipmentLine import InboundShipmentLine
from inventoryOnDjango.models.InventoryTransaction import InventoryTransaction
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InboundShipment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipmentDelegate Declaration
#======================================================================
class InboundShipmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inboundShipmentId ):
		try:	
			inboundShipment = InboundShipment.objects.filter(id=inboundShipmentId)
			return inboundShipment.first();
		except InboundShipment.DoesNotExist:
			raise ProcessingError("InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inboundShipment):
		for model in serializers.deserialize("json", inboundShipment):
			model.save()
			return model;

	def create(self, inboundShipment):
		inboundShipment.save()
		return inboundShipment;

	def saveFromJson(self, inboundShipment):
		for model in serializers.deserialize("json", inboundShipment):
			model.save()
			return inboundShipment;
	
	def save(self, inboundShipment):
		inboundShipment.save()
		return inboundShipment;
	
	def delete(self, inboundShipmentId ):
		errMsg = "Failed to delete InboundShipment from db using id " + str(inboundShipmentId)
		
		try:
			inboundShipment = InboundShipment.objects.get(id=inboundShipmentId)
			inboundShipment.delete()
			return True
		except InboundShipment.DoesNotExist:
			raise ProcessingError("InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InboundShipment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InboundShipment from db")
		except Exception:
			return None;
		
	def assignWarehouse( self, inboundShipmentId, warehouseId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

		errMsg = "Failed to assign element " + str(warehouseId) + " for Warehouse on InboundShipment"

		try:
			# get the InboundShipment from db
			inboundShipment = self.get( inboundShipmentId ).first()	
			
			# get the Warehouse from db
			warehouse = WarehouseDelegate().get(warehouseId).first();
			
			# assign the Warehouse		
			inboundShipment.warehouse = warehouse
			
			#save it
			inboundShipment.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentId );
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except Warehouse.DoesNotExist:
			raise ProcessingError(errMsg + " : Warehouse with id " + str(warehouseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWarehouse( self, inboundShipmentId ):
		errMsg = "Failed to unassign element " + str(warehouseId) + " for Warehouse on InboundShipment"

		try:
			# get the InboundShipment from db
			inboundShipment = self.get( inboundShipmentId ).first()	
			
			# assign to None for unassignment
			inboundShipment.warehouse = None			

			#save it
			inboundShipment.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentId );
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addLines( self, inboundShipmentId, linesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InboundShipmentLineDelegate import InboundShipmentLineDelegate

		errMsg = "Failed to add elements " + str(linesIds) + " for Lines on InboundShipment"

		try:
			# get the InboundShipment
			inboundShipment = self.get( inboundShipmentId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InboundShipmentLine		
				inboundShipmentLine = InboundShipmentLineDelegate().get(id).first();	
				# add the InboundShipmentLine
				inboundShipment.lines.add(inboundShipmentLine)
				
			# save it		
			inboundShipment.save()
			
			# reload and return the appropriate version
			return self.get( inboundShipmentId );
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLines( self, inboundShipmentId, linesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InboundShipmentLineDelegate import InboundShipmentLineDelegate

		errMsg = "Failed to remove elements " + str(linesIds) + " for Lines on InboundShipment"

		try:
			# get the InboundShipment
			inboundShipment = self.get( inboundShipmentId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InboundShipmentLine		
				inboundShipmentLine = InboundShipmentLineDelegate().get(id).first();	
				# add the InboundShipmentLine
				inboundShipment.lines.remove(inboundShipmentLine)
				
			# save it		
			inboundShipment.save()
			
			# reload and return the appropriate version
			return self.get( inboundShipmentId );
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTransactions( self, inboundShipmentId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to add elements " + str(transactionsIds) + " for Transactions on InboundShipment"

		try:
			# get the InboundShipment
			inboundShipment = self.get( inboundShipmentId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				inboundShipment.transactions.add(inventoryTransaction)
				
			# save it		
			inboundShipment.save()
			
			# reload and return the appropriate version
			return self.get( inboundShipmentId );
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTransactions( self, inboundShipmentId, transactionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryTransactionDelegate import InventoryTransactionDelegate

		errMsg = "Failed to remove elements " + str(transactionsIds) + " for Transactions on InboundShipment"

		try:
			# get the InboundShipment
			inboundShipment = self.get( inboundShipmentId ).first()
				
			# split on a comma with no spaces
			idList = transactionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryTransaction		
				inventoryTransaction = InventoryTransactionDelegate().get(id).first();	
				# add the InventoryTransaction
				inboundShipment.transactions.remove(inventoryTransaction)
				
			# save it		
			inboundShipment.save()
			
			# reload and return the appropriate version
			return self.get( inboundShipmentId );
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except InventoryTransaction.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryTransaction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
