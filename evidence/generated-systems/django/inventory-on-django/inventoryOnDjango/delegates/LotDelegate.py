from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Lot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LotDelegate Declaration
#======================================================================
class LotDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, lotId ):
		try:	
			lot = Lot.objects.filter(id=lotId)
			return lot.first();
		except Lot.DoesNotExist:
			raise ProcessingError("Lot with id " + str(lotId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, lot):
		for model in serializers.deserialize("json", lot):
			model.save()
			return model;

	def create(self, lot):
		lot.save()
		return lot;

	def saveFromJson(self, lot):
		for model in serializers.deserialize("json", lot):
			model.save()
			return lot;
	
	def save(self, lot):
		lot.save()
		return lot;
	
	def delete(self, lotId ):
		errMsg = "Failed to delete Lot from db using id " + str(lotId)
		
		try:
			lot = Lot.objects.get(id=lotId)
			lot.delete()
			return True
		except Lot.DoesNotExist:
			raise ProcessingError("Lot with id " + str(lotId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Lot.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Lot from db")
		except Exception:
			return None;
		
	def assignSku( self, lotId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on Lot"

		try:
			# get the Lot from db
			lot = self.get( lotId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			lot.sku = stockKeepingUnit
			
			#save it
			lot.save()

			# reload and return the appropriate version					
			return self.get( lotId );
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, lotId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on Lot"

		try:
			# get the Lot from db
			lot = self.get( lotId ).first()	
			
			# assign to None for unassignment
			lot.stockKeepingUnit = None			

			#save it
			lot.save()

			# reload and return the appropriate version					
			return self.get( lotId );
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
		
	def addInventoryItems( self, lotId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on Lot"

		try:
			# get the Lot
			lot = self.get( lotId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				lot.inventoryItems.add(inventoryItem)
				
			# save it		
			lot.save()
			
			# reload and return the appropriate version
			return self.get( lotId );
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, lotId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on Lot"

		try:
			# get the Lot
			lot = self.get( lotId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				lot.inventoryItems.remove(inventoryItem)
				
			# save it		
			lot.save()
			
			# reload and return the appropriate version
			return self.get( lotId );
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
