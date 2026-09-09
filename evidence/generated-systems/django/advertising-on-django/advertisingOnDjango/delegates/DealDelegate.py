from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Deal import Deal
from advertisingOnDjango.models.Publisher import Publisher
from advertisingOnDjango.models.InventorySource import InventorySource
from advertisingOnDjango.models.Placement import Placement
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Deal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DealDelegate Declaration
#======================================================================
class DealDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dealId ):
		try:	
			deal = Deal.objects.filter(id=dealId)
			return deal.first();
		except Deal.DoesNotExist:
			raise ProcessingError("Deal with id " + str(dealId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, deal):
		for model in serializers.deserialize("json", deal):
			model.save()
			return model;

	def create(self, deal):
		deal.save()
		return deal;

	def saveFromJson(self, deal):
		for model in serializers.deserialize("json", deal):
			model.save()
			return deal;
	
	def save(self, deal):
		deal.save()
		return deal;
	
	def delete(self, dealId ):
		errMsg = "Failed to delete Deal from db using id " + str(dealId)
		
		try:
			deal = Deal.objects.get(id=dealId)
			deal.delete()
			return True
		except Deal.DoesNotExist:
			raise ProcessingError("Deal with id " + str(dealId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Deal.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Deal from db")
		except Exception:
			return None;
		
	def assignPublisher( self, dealId, publisherId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PublisherDelegate import PublisherDelegate

		errMsg = "Failed to assign element " + str(publisherId) + " for Publisher on Deal"

		try:
			# get the Deal from db
			deal = self.get( dealId ).first()	
			
			# get the Publisher from db
			publisher = PublisherDelegate().get(publisherId).first();
			
			# assign the Publisher		
			deal.publisher = publisher
			
			#save it
			deal.save()

			# reload and return the appropriate version					
			return self.get( dealId );
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPublisher( self, dealId ):
		errMsg = "Failed to unassign element " + str(publisherId) + " for Publisher on Deal"

		try:
			# get the Deal from db
			deal = self.get( dealId ).first()	
			
			# assign to None for unassignment
			deal.publisher = None			

			#save it
			deal.save()

			# reload and return the appropriate version					
			return self.get( dealId );
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except Exception:
			return None;
		
	def addInventorySources( self, dealId, inventorySourcesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InventorySourceDelegate import InventorySourceDelegate

		errMsg = "Failed to add elements " + str(inventorySourcesIds) + " for InventorySources on Deal"

		try:
			# get the Deal
			deal = self.get( dealId ).first()
				
			# split on a comma with no spaces
			idList = inventorySourcesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventorySource		
				inventorySource = InventorySourceDelegate().get(id).first();	
				# add the InventorySource
				deal.inventorySources.add(inventorySource)
				
			# save it		
			deal.save()
			
			# reload and return the appropriate version
			return self.get( dealId );
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventorySources( self, dealId, inventorySourcesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InventorySourceDelegate import InventorySourceDelegate

		errMsg = "Failed to remove elements " + str(inventorySourcesIds) + " for InventorySources on Deal"

		try:
			# get the Deal
			deal = self.get( dealId ).first()
				
			# split on a comma with no spaces
			idList = inventorySourcesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventorySource		
				inventorySource = InventorySourceDelegate().get(id).first();	
				# add the InventorySource
				deal.inventorySources.remove(inventorySource)
				
			# save it		
			deal.save()
			
			# reload and return the appropriate version
			return self.get( dealId );
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPlacements( self, dealId, placementsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

		errMsg = "Failed to add elements " + str(placementsIds) + " for Placements on Deal"

		try:
			# get the Deal
			deal = self.get( dealId ).first()
				
			# split on a comma with no spaces
			idList = placementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Placement		
				placement = PlacementDelegate().get(id).first();	
				# add the Placement
				deal.placements.add(placement)
				
			# save it		
			deal.save()
			
			# reload and return the appropriate version
			return self.get( dealId );
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlacements( self, dealId, placementsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

		errMsg = "Failed to remove elements " + str(placementsIds) + " for Placements on Deal"

		try:
			# get the Deal
			deal = self.get( dealId ).first()
				
			# split on a comma with no spaces
			idList = placementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Placement		
				placement = PlacementDelegate().get(id).first();	
				# add the Placement
				deal.placements.remove(placement)
				
			# save it		
			deal.save()
			
			# reload and return the appropriate version
			return self.get( dealId );
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
