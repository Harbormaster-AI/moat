from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.StockAdjustmentLine import StockAdjustmentLine
from inventoryOnDjango.models.StockAdjustment import StockAdjustment
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model StockAdjustmentLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StockAdjustmentLineDelegate Declaration
#======================================================================
class StockAdjustmentLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, stockAdjustmentLineId ):
		try:	
			stockAdjustmentLine = StockAdjustmentLine.objects.filter(id=stockAdjustmentLineId)
			return stockAdjustmentLine.first();
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError("StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, stockAdjustmentLine):
		for model in serializers.deserialize("json", stockAdjustmentLine):
			model.save()
			return model;

	def create(self, stockAdjustmentLine):
		stockAdjustmentLine.save()
		return stockAdjustmentLine;

	def saveFromJson(self, stockAdjustmentLine):
		for model in serializers.deserialize("json", stockAdjustmentLine):
			model.save()
			return stockAdjustmentLine;
	
	def save(self, stockAdjustmentLine):
		stockAdjustmentLine.save()
		return stockAdjustmentLine;
	
	def delete(self, stockAdjustmentLineId ):
		errMsg = "Failed to delete StockAdjustmentLine from db using id " + str(stockAdjustmentLineId)
		
		try:
			stockAdjustmentLine = StockAdjustmentLine.objects.get(id=stockAdjustmentLineId)
			stockAdjustmentLine.delete()
			return True
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError("StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = StockAdjustmentLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all StockAdjustmentLine from db")
		except Exception:
			return None;
		
	def assignAdjustment( self, stockAdjustmentLineId, adjustmentId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockAdjustmentDelegate import StockAdjustmentDelegate

		errMsg = "Failed to assign element " + str(adjustmentId) + " for Adjustment on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# get the StockAdjustment from db
			stockAdjustment = StockAdjustmentDelegate().get(adjustmentId).first();
			
			# assign the Adjustment		
			stockAdjustmentLine.adjustment = stockAdjustment
			
			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except StockAdjustment.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustment with id " + str(adjustmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdjustment( self, stockAdjustmentLineId ):
		errMsg = "Failed to unassign element " + str(adjustmentId) + " for Adjustment on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# assign to None for unassignment
			stockAdjustmentLine.stockAdjustment = None			

			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSku( self, stockAdjustmentLineId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			stockAdjustmentLine.sku = stockKeepingUnit
			
			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, stockAdjustmentLineId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# assign to None for unassignment
			stockAdjustmentLine.stockKeepingUnit = None			

			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, stockAdjustmentLineId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			stockAdjustmentLine.lot = lot
			
			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, stockAdjustmentLineId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# assign to None for unassignment
			stockAdjustmentLine.lot = None			

			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, stockAdjustmentLineId, locationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(locationId).first();
			
			# assign the Location		
			stockAdjustmentLine.location = storageLocation
			
			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, stockAdjustmentLineId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine from db
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()	
			
			# assign to None for unassignment
			stockAdjustmentLine.storageLocation = None			

			#save it
			stockAdjustmentLine.save()

			# reload and return the appropriate version					
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, stockAdjustmentLineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				stockAdjustmentLine.serialNumbers.add(serialNumber)
				
			# save it		
			stockAdjustmentLine.save()
			
			# reload and return the appropriate version
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, stockAdjustmentLineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on StockAdjustmentLine"

		try:
			# get the StockAdjustmentLine
			stockAdjustmentLine = self.get( stockAdjustmentLineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				stockAdjustmentLine.serialNumbers.remove(serialNumber)
				
			# save it		
			stockAdjustmentLine.save()
			
			# reload and return the appropriate version
			return self.get( stockAdjustmentLineId );
		except StockAdjustmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : StockAdjustmentLine with id " + str(stockAdjustmentLineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
