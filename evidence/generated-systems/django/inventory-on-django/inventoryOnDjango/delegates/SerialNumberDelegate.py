from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SerialNumber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SerialNumberDelegate Declaration
#======================================================================
class SerialNumberDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, serialNumberId ):
		try:	
			serialNumber = SerialNumber.objects.filter(id=serialNumberId)
			return serialNumber.first();
		except SerialNumber.DoesNotExist:
			raise ProcessingError("SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, serialNumber):
		for model in serializers.deserialize("json", serialNumber):
			model.save()
			return model;

	def create(self, serialNumber):
		serialNumber.save()
		return serialNumber;

	def saveFromJson(self, serialNumber):
		for model in serializers.deserialize("json", serialNumber):
			model.save()
			return serialNumber;
	
	def save(self, serialNumber):
		serialNumber.save()
		return serialNumber;
	
	def delete(self, serialNumberId ):
		errMsg = "Failed to delete SerialNumber from db using id " + str(serialNumberId)
		
		try:
			serialNumber = SerialNumber.objects.get(id=serialNumberId)
			serialNumber.delete()
			return True
		except SerialNumber.DoesNotExist:
			raise ProcessingError("SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SerialNumber.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SerialNumber from db")
		except Exception:
			return None;
		
	def assignSku( self, serialNumberId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on SerialNumber"

		try:
			# get the SerialNumber from db
			serialNumber = self.get( serialNumberId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			serialNumber.sku = stockKeepingUnit
			
			#save it
			serialNumber.save()

			# reload and return the appropriate version					
			return self.get( serialNumberId );
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, serialNumberId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on SerialNumber"

		try:
			# get the SerialNumber from db
			serialNumber = self.get( serialNumberId ).first()	
			
			# assign to None for unassignment
			serialNumber.stockKeepingUnit = None			

			#save it
			serialNumber.save()

			# reload and return the appropriate version					
			return self.get( serialNumberId );
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCurrentInventoryItem( self, serialNumberId, currentInventoryItemId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to assign element " + str(currentInventoryItemId) + " for CurrentInventoryItem on SerialNumber"

		try:
			# get the SerialNumber from db
			serialNumber = self.get( serialNumberId ).first()	
			
			# get the InventoryItem from db
			inventoryItem = InventoryItemDelegate().get(currentInventoryItemId).first();
			
			# assign the CurrentInventoryItem		
			serialNumber.currentInventoryItem = inventoryItem
			
			#save it
			serialNumber.save()

			# reload and return the appropriate version					
			return self.get( serialNumberId );
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(currentInventoryItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCurrentInventoryItem( self, serialNumberId ):
		errMsg = "Failed to unassign element " + str(currentInventoryItemId) + " for CurrentInventoryItem on SerialNumber"

		try:
			# get the SerialNumber from db
			serialNumber = self.get( serialNumberId ).first()	
			
			# assign to None for unassignment
			serialNumber.inventoryItem = None			

			#save it
			serialNumber.save()

			# reload and return the appropriate version					
			return self.get( serialNumberId );
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, serialNumberId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on SerialNumber"

		try:
			# get the SerialNumber from db
			serialNumber = self.get( serialNumberId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			serialNumber.lot = lot
			
			#save it
			serialNumber.save()

			# reload and return the appropriate version					
			return self.get( serialNumberId );
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, serialNumberId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on SerialNumber"

		try:
			# get the SerialNumber from db
			serialNumber = self.get( serialNumberId ).first()	
			
			# assign to None for unassignment
			serialNumber.lot = None			

			#save it
			serialNumber.save()

			# reload and return the appropriate version					
			return self.get( serialNumberId );
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber with id " + str(serialNumberId) + " does not exist.")
		except Exception:
			return None;
		
