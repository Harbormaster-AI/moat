from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.CycleCountEntry import CycleCountEntry
from inventoryOnDjango.models.CycleCount import CycleCount
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CycleCountEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCountEntryDelegate Declaration
#======================================================================
class CycleCountEntryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, cycleCountEntryId ):
		try:	
			cycleCountEntry = CycleCountEntry.objects.filter(id=cycleCountEntryId)
			return cycleCountEntry.first();
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError("CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, cycleCountEntry):
		for model in serializers.deserialize("json", cycleCountEntry):
			model.save()
			return model;

	def create(self, cycleCountEntry):
		cycleCountEntry.save()
		return cycleCountEntry;

	def saveFromJson(self, cycleCountEntry):
		for model in serializers.deserialize("json", cycleCountEntry):
			model.save()
			return cycleCountEntry;
	
	def save(self, cycleCountEntry):
		cycleCountEntry.save()
		return cycleCountEntry;
	
	def delete(self, cycleCountEntryId ):
		errMsg = "Failed to delete CycleCountEntry from db using id " + str(cycleCountEntryId)
		
		try:
			cycleCountEntry = CycleCountEntry.objects.get(id=cycleCountEntryId)
			cycleCountEntry.delete()
			return True
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError("CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CycleCountEntry.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CycleCountEntry from db")
		except Exception:
			return None;
		
	def assignCycleCount( self, cycleCountEntryId, cycleCountId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.CycleCountDelegate import CycleCountDelegate

		errMsg = "Failed to assign element " + str(cycleCountId) + " for CycleCount on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# get the CycleCount from db
			cycleCount = CycleCountDelegate().get(cycleCountId).first();
			
			# assign the CycleCount		
			cycleCountEntry.cycleCount = cycleCount
			
			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except CycleCount.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCount with id " + str(cycleCountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCycleCount( self, cycleCountEntryId ):
		errMsg = "Failed to unassign element " + str(cycleCountId) + " for CycleCount on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# assign to None for unassignment
			cycleCountEntry.cycleCount = None			

			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSku( self, cycleCountEntryId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			cycleCountEntry.sku = stockKeepingUnit
			
			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, cycleCountEntryId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# assign to None for unassignment
			cycleCountEntry.stockKeepingUnit = None			

			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, cycleCountEntryId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			cycleCountEntry.lot = lot
			
			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, cycleCountEntryId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# assign to None for unassignment
			cycleCountEntry.lot = None			

			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, cycleCountEntryId, locationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(locationId).first();
			
			# assign the Location		
			cycleCountEntry.location = storageLocation
			
			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, cycleCountEntryId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on CycleCountEntry"

		try:
			# get the CycleCountEntry from db
			cycleCountEntry = self.get( cycleCountEntryId ).first()	
			
			# assign to None for unassignment
			cycleCountEntry.storageLocation = None			

			#save it
			cycleCountEntry.save()

			# reload and return the appropriate version					
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, cycleCountEntryId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on CycleCountEntry"

		try:
			# get the CycleCountEntry
			cycleCountEntry = self.get( cycleCountEntryId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				cycleCountEntry.serialNumbers.add(serialNumber)
				
			# save it		
			cycleCountEntry.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, cycleCountEntryId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on CycleCountEntry"

		try:
			# get the CycleCountEntry
			cycleCountEntry = self.get( cycleCountEntryId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				cycleCountEntry.serialNumbers.remove(serialNumber)
				
			# save it		
			cycleCountEntry.save()
			
			# reload and return the appropriate version
			return self.get( cycleCountEntryId );
		except CycleCountEntry.DoesNotExist:
			raise ProcessingError(errMsg + " : CycleCountEntry with id " + str(cycleCountEntryId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
