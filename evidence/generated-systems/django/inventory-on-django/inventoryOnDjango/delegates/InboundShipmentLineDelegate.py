from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from inventoryOnDjango.models.InboundShipmentLine import InboundShipmentLine
from inventoryOnDjango.models.InboundShipment import InboundShipment
from inventoryOnDjango.models.StockKeepingUnit import StockKeepingUnit
from inventoryOnDjango.models.Lot import Lot
from inventoryOnDjango.models.SerialNumber import SerialNumber
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InboundShipmentLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InboundShipmentLineDelegate Declaration
#======================================================================
class InboundShipmentLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inboundShipmentLineId ):
		try:	
			inboundShipmentLine = InboundShipmentLine.objects.filter(id=inboundShipmentLineId)
			return inboundShipmentLine.first();
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError("InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inboundShipmentLine):
		for model in serializers.deserialize("json", inboundShipmentLine):
			model.save()
			return model;

	def create(self, inboundShipmentLine):
		inboundShipmentLine.save()
		return inboundShipmentLine;

	def saveFromJson(self, inboundShipmentLine):
		for model in serializers.deserialize("json", inboundShipmentLine):
			model.save()
			return inboundShipmentLine;
	
	def save(self, inboundShipmentLine):
		inboundShipmentLine.save()
		return inboundShipmentLine;
	
	def delete(self, inboundShipmentLineId ):
		errMsg = "Failed to delete InboundShipmentLine from db using id " + str(inboundShipmentLineId)
		
		try:
			inboundShipmentLine = InboundShipmentLine.objects.get(id=inboundShipmentLineId)
			inboundShipmentLine.delete()
			return True
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError("InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InboundShipmentLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InboundShipmentLine from db")
		except Exception:
			return None;
		
	def assignInboundShipment( self, inboundShipmentLineId, inboundShipmentId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.InboundShipmentDelegate import InboundShipmentDelegate

		errMsg = "Failed to assign element " + str(inboundShipmentId) + " for InboundShipment on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# get the InboundShipment from db
			inboundShipment = InboundShipmentDelegate().get(inboundShipmentId).first();
			
			# assign the InboundShipment		
			inboundShipmentLine.inboundShipment = inboundShipment
			
			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except InboundShipment.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipment with id " + str(inboundShipmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInboundShipment( self, inboundShipmentLineId ):
		errMsg = "Failed to unassign element " + str(inboundShipmentId) + " for InboundShipment on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# assign to None for unassignment
			inboundShipmentLine.inboundShipment = None			

			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSku( self, inboundShipmentLineId, skuId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StockKeepingUnitDelegate import StockKeepingUnitDelegate

		errMsg = "Failed to assign element " + str(skuId) + " for Sku on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# get the StockKeepingUnit from db
			stockKeepingUnit = StockKeepingUnitDelegate().get(skuId).first();
			
			# assign the Sku		
			inboundShipmentLine.sku = stockKeepingUnit
			
			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except StockKeepingUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : StockKeepingUnit with id " + str(skuId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSku( self, inboundShipmentLineId ):
		errMsg = "Failed to unassign element " + str(skuId) + " for Sku on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# assign to None for unassignment
			inboundShipmentLine.stockKeepingUnit = None			

			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLot( self, inboundShipmentLineId, lotId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.LotDelegate import LotDelegate

		errMsg = "Failed to assign element " + str(lotId) + " for Lot on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# get the Lot from db
			lot = LotDelegate().get(lotId).first();
			
			# assign the Lot		
			inboundShipmentLine.lot = lot
			
			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except Lot.DoesNotExist:
			raise ProcessingError(errMsg + " : Lot with id " + str(lotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLot( self, inboundShipmentLineId ):
		errMsg = "Failed to unassign element " + str(lotId) + " for Lot on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# assign to None for unassignment
			inboundShipmentLine.lot = None			

			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDestinationLocation( self, inboundShipmentLineId, destinationLocationId ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

		errMsg = "Failed to assign element " + str(destinationLocationId) + " for DestinationLocation on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# get the StorageLocation from db
			storageLocation = StorageLocationDelegate().get(destinationLocationId).first();
			
			# assign the DestinationLocation		
			inboundShipmentLine.destinationLocation = storageLocation
			
			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except StorageLocation.DoesNotExist:
			raise ProcessingError(errMsg + " : StorageLocation with id " + str(destinationLocationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDestinationLocation( self, inboundShipmentLineId ):
		errMsg = "Failed to unassign element " + str(destinationLocationId) + " for DestinationLocation on InboundShipmentLine"

		try:
			# get the InboundShipmentLine from db
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()	
			
			# assign to None for unassignment
			inboundShipmentLine.storageLocation = None			

			#save it
			inboundShipmentLine.save()

			# reload and return the appropriate version					
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except Exception:
			return None;
		
	def addSerialNumbers( self, inboundShipmentLineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to add elements " + str(serialNumbersIds) + " for SerialNumbers on InboundShipmentLine"

		try:
			# get the InboundShipmentLine
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				inboundShipmentLine.serialNumbers.add(serialNumber)
				
			# save it		
			inboundShipmentLine.save()
			
			# reload and return the appropriate version
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSerialNumbers( self, inboundShipmentLineId, serialNumbersIds ):
		# lazy importing avoids circular dependencies
		from inventoryOnDjango.delegates.SerialNumberDelegate import SerialNumberDelegate

		errMsg = "Failed to remove elements " + str(serialNumbersIds) + " for SerialNumbers on InboundShipmentLine"

		try:
			# get the InboundShipmentLine
			inboundShipmentLine = self.get( inboundShipmentLineId ).first()
				
			# split on a comma with no spaces
			idList = serialNumbersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SerialNumber		
				serialNumber = SerialNumberDelegate().get(id).first();	
				# add the SerialNumber
				inboundShipmentLine.serialNumbers.remove(serialNumber)
				
			# save it		
			inboundShipmentLine.save()
			
			# reload and return the appropriate version
			return self.get( inboundShipmentLineId );
		except InboundShipmentLine.DoesNotExist:
			raise ProcessingError(errMsg + " : InboundShipmentLine with id " + str(inboundShipmentLineId) + " does not exist.")
		except SerialNumber.DoesNotExist:
			raise ProcessingError(errMsg + " : SerialNumber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
