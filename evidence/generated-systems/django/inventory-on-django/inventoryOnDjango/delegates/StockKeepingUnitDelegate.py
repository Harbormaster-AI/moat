from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.models.UoMConversion import UoMConversion
from inventoryOnDjango.models.ReplenishmentPolicy import ReplenishmentPolicy
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model StockKeepingUnit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockKeepingUnitDelegate Declaration
#======================================================================
class StockKeepingUnitDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, stockKeepingUnitId ):
		try:	
			stockKeepingUnit = StockKeepingUnit.objects.filter(id=stockKeepingUnitId)
			return stockKeepingUnit.first();
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError("StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, stockKeepingUnit):
		for model in serializers.deserialize("json", stockKeepingUnit):
			model.save()
			return model;

	def create(self, stockKeepingUnit):
		stockKeepingUnit.save()
		return stockKeepingUnit;

	def saveFromJson(self, stockKeepingUnit):
		for model in serializers.deserialize("json", stockKeepingUnit):
			model.save()
			return stockKeepingUnit;
	
	def save(self, stockKeepingUnit):
		stockKeepingUnit.save()
		return stockKeepingUnit;
	
	def delete(self, stockKeepingUnitId ):
		errMsg = "Failed to delete StockKeepingUnit from db using id " + str(stockKeepingUnitId)
		
		try:
			stockKeepingUnit = StockKeepingUnit.objects.get(id=stockKeepingUnitId)
			stockKeepingUnit.delete()
			return True
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError("StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = StockKeepingUnit.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all StockKeepingUnit from db")
		except Exception:
			return None;
		
	def addInventoryItems( self, stockKeepingUnitId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				stockKeepingUnit.inventoryItems.add(inventoryItem)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, stockKeepingUnitId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				stockKeepingUnit.inventoryItems.remove(inventoryItem)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addUomConversions( self, stockKeepingUnitId, uomConversionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.UoMConversionDelegate import UoMConversionDelegate

		errMsg = "Failed to add elements " + str(uomConversionsIds) + " for UomConversions on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = uomConversionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the UoMConversion		
				uoMConversion = UoMConversionDelegate().get(id).first();	
				# add the UoMConversion
				stockKeepingUnit.uomConversions.add(uoMConversion)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except UoMConversion.DoesNotExist:
			raise ProcessingError(errMsg + " : UoMConversion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUomConversions( self, stockKeepingUnitId, uomConversionsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.UoMConversionDelegate import UoMConversionDelegate

		errMsg = "Failed to remove elements " + str(uomConversionsIds) + " for UomConversions on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = uomConversionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the UoMConversion		
				uoMConversion = UoMConversionDelegate().get(id).first();	
				# add the UoMConversion
				stockKeepingUnit.uomConversions.remove(uoMConversion)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except UoMConversion.DoesNotExist:
			raise ProcessingError(errMsg + " : UoMConversion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReplenishmentPolicies( self, stockKeepingUnitId, replenishmentPoliciesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReplenishmentPolicyDelegate import ReplenishmentPolicyDelegate

		errMsg = "Failed to add elements " + str(replenishmentPoliciesIds) + " for ReplenishmentPolicies on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = replenishmentPoliciesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ReplenishmentPolicy		
				replenishmentPolicy = ReplenishmentPolicyDelegate().get(id).first();	
				# add the ReplenishmentPolicy
				stockKeepingUnit.replenishmentPolicies.add(replenishmentPolicy)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReplenishmentPolicies( self, stockKeepingUnitId, replenishmentPoliciesIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.ReplenishmentPolicyDelegate import ReplenishmentPolicyDelegate

		errMsg = "Failed to remove elements " + str(replenishmentPoliciesIds) + " for ReplenishmentPolicies on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = replenishmentPoliciesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ReplenishmentPolicy		
				replenishmentPolicy = ReplenishmentPolicyDelegate().get(id).first();	
				# add the ReplenishmentPolicy
				stockKeepingUnit.replenishmentPolicies.remove(replenishmentPolicy)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except ReplenishmentPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : ReplenishmentPolicy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLots( self, stockKeepingUnitId, lotsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to add elements " + str(lotsIds) + " for Lots on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = lotsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Lot		
				lot = LotDelegate().get(id).first();	
				# add the Lot
				stockKeepingUnit.lots.add(lot)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLots( self, stockKeepingUnitId, lotsIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to remove elements " + str(lotsIds) + " for Lots on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = lotsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Lot		
				lot = LotDelegate().get(id).first();	
				# add the Lot
				stockKeepingUnit.lots.remove(lot)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSerialNumbers( self, stockKeepingUnitId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				stockKeepingUnit.serialNumbers.add(serialNumber)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, stockKeepingUnitId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on StockKeepingUnit"

		try:
			# get the StockKeepingUnit
			stockKeepingUnit = self.get( stockKeepingUnitId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				stockKeepingUnit.serialNumbers.remove(serialNumber)
				
			# save it		
			stockKeepingUnit.save()
			
			# reload and return the appropriate version
			return self.get( stockKeepingUnitId );
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(stockKeepingUnitId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
