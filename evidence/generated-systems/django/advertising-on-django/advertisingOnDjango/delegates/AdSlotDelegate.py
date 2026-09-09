from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.AdSlot import AdSlot
from advertisingOnDjango.models.InventorySource import InventorySource
from advertisingOnDjango.models.Placement import Placement
from advertisingOnDjango.models.Rate import Rate
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AdSlot
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdSlotDelegate Declaration
#======================================================================
class AdSlotDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, adSlotId ):
		try:	
			adSlot = AdSlot.objects.filter(id=adSlotId)
			return adSlot.first();
		except AdSlot.DoesNotExist:
			raise ProcessingError("AdSlot with id " + str(adSlotId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, adSlot):
		for model in serializers.deserialize("json", adSlot):
			model.save()
			return model;

	def create(self, adSlot):
		adSlot.save()
		return adSlot;

	def saveFromJson(self, adSlot):
		for model in serializers.deserialize("json", adSlot):
			model.save()
			return adSlot;
	
	def save(self, adSlot):
		adSlot.save()
		return adSlot;
	
	def delete(self, adSlotId ):
		errMsg = "Failed to delete AdSlot from db using id " + str(adSlotId)
		
		try:
			adSlot = AdSlot.objects.get(id=adSlotId)
			adSlot.delete()
			return True
		except AdSlot.DoesNotExist:
			raise ProcessingError("AdSlot with id " + str(adSlotId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AdSlot.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AdSlot from db")
		except Exception:
			return None;
		
	def assignInventorySource( self, adSlotId, inventorySourceId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InventorySourceDelegate import InventorySourceDelegate

		errMsg = "Failed to assign element " + str(inventorySourceId) + " for InventorySource on AdSlot"

		try:
			# get the AdSlot from db
			adSlot = self.get( adSlotId ).first()	
			
			# get the InventorySource from db
			inventorySource = InventorySourceDelegate().get(inventorySourceId).first();
			
			# assign the InventorySource		
			adSlot.inventorySource = inventorySource
			
			#save it
			adSlot.save()

			# reload and return the appropriate version					
			return self.get( adSlotId );
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInventorySource( self, adSlotId ):
		errMsg = "Failed to unassign element " + str(inventorySourceId) + " for InventorySource on AdSlot"

		try:
			# get the AdSlot from db
			adSlot = self.get( adSlotId ).first()	
			
			# assign to None for unassignment
			adSlot.inventorySource = None			

			#save it
			adSlot.save()

			# reload and return the appropriate version					
			return self.get( adSlotId );
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except Exception:
			return None;
		
	def addPlacements( self, adSlotId, placementsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

		errMsg = "Failed to add elements " + str(placementsIds) + " for Placements on AdSlot"

		try:
			# get the AdSlot
			adSlot = self.get( adSlotId ).first()
				
			# split on a comma with no spaces
			idList = placementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Placement		
				placement = PlacementDelegate().get(id).first();	
				# add the Placement
				adSlot.placements.add(placement)
				
			# save it		
			adSlot.save()
			
			# reload and return the appropriate version
			return self.get( adSlotId );
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlacements( self, adSlotId, placementsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

		errMsg = "Failed to remove elements " + str(placementsIds) + " for Placements on AdSlot"

		try:
			# get the AdSlot
			adSlot = self.get( adSlotId ).first()
				
			# split on a comma with no spaces
			idList = placementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Placement		
				placement = PlacementDelegate().get(id).first();	
				# add the Placement
				adSlot.placements.remove(placement)
				
			# save it		
			adSlot.save()
			
			# reload and return the appropriate version
			return self.get( adSlotId );
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRates( self, adSlotId, ratesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.RateDelegate import RateDelegate

		errMsg = "Failed to add elements " + str(ratesIds) + " for Rates on AdSlot"

		try:
			# get the AdSlot
			adSlot = self.get( adSlotId ).first()
				
			# split on a comma with no spaces
			idList = ratesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Rate		
				rate = RateDelegate().get(id).first();	
				# add the Rate
				adSlot.rates.add(rate)
				
			# save it		
			adSlot.save()
			
			# reload and return the appropriate version
			return self.get( adSlotId );
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRates( self, adSlotId, ratesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.RateDelegate import RateDelegate

		errMsg = "Failed to remove elements " + str(ratesIds) + " for Rates on AdSlot"

		try:
			# get the AdSlot
			adSlot = self.get( adSlotId ).first()
				
			# split on a comma with no spaces
			idList = ratesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Rate		
				rate = RateDelegate().get(id).first();	
				# add the Rate
				adSlot.rates.remove(rate)
				
			# save it		
			adSlot.save()
			
			# reload and return the appropriate version
			return self.get( adSlotId );
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot with id " + str(adSlotId) + " does not exist.")
		except Rate.DoesNotExist:
			raise ProcessingError(errMsg + " : Rate does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
