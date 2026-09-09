from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.InventorySource import InventorySource
from advertisingOnDjango.models.Publisher import Publisher
from advertisingOnDjango.models.AdSlot import AdSlot
from advertisingOnDjango.models.Deal import Deal
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InventorySource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventorySourceDelegate Declaration
#======================================================================
class InventorySourceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inventorySourceId ):
		try:	
			inventorySource = InventorySource.objects.filter(id=inventorySourceId)
			return inventorySource.first();
		except InventorySource.DoesNotExist:
			raise ProcessingError("InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inventorySource):
		for model in serializers.deserialize("json", inventorySource):
			model.save()
			return model;

	def create(self, inventorySource):
		inventorySource.save()
		return inventorySource;

	def saveFromJson(self, inventorySource):
		for model in serializers.deserialize("json", inventorySource):
			model.save()
			return inventorySource;
	
	def save(self, inventorySource):
		inventorySource.save()
		return inventorySource;
	
	def delete(self, inventorySourceId ):
		errMsg = "Failed to delete InventorySource from db using id " + str(inventorySourceId)
		
		try:
			inventorySource = InventorySource.objects.get(id=inventorySourceId)
			inventorySource.delete()
			return True
		except InventorySource.DoesNotExist:
			raise ProcessingError("InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InventorySource.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InventorySource from db")
		except Exception:
			return None;
		
	def assignPublisher( self, inventorySourceId, publisherId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PublisherDelegate import PublisherDelegate

		errMsg = "Failed to assign element " + str(publisherId) + " for Publisher on InventorySource"

		try:
			# get the InventorySource from db
			inventorySource = self.get( inventorySourceId ).first()	
			
			# get the Publisher from db
			publisher = PublisherDelegate().get(publisherId).first();
			
			# assign the Publisher		
			inventorySource.publisher = publisher
			
			#save it
			inventorySource.save()

			# reload and return the appropriate version					
			return self.get( inventorySourceId );
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPublisher( self, inventorySourceId ):
		errMsg = "Failed to unassign element " + str(publisherId) + " for Publisher on InventorySource"

		try:
			# get the InventorySource from db
			inventorySource = self.get( inventorySourceId ).first()	
			
			# assign to None for unassignment
			inventorySource.publisher = None			

			#save it
			inventorySource.save()

			# reload and return the appropriate version					
			return self.get( inventorySourceId );
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except Exception:
			return None;
		
	def addAdSlots( self, inventorySourceId, adSlotsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdSlotDelegate import AdSlotDelegate

		errMsg = "Failed to add elements " + str(adSlotsIds) + " for AdSlots on InventorySource"

		try:
			# get the InventorySource
			inventorySource = self.get( inventorySourceId ).first()
				
			# split on a comma with no spaces
			idList = adSlotsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AdSlot		
				adSlot = AdSlotDelegate().get(id).first();	
				# add the AdSlot
				inventorySource.adSlots.add(adSlot)
				
			# save it		
			inventorySource.save()
			
			# reload and return the appropriate version
			return self.get( inventorySourceId );
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAdSlots( self, inventorySourceId, adSlotsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdSlotDelegate import AdSlotDelegate

		errMsg = "Failed to remove elements " + str(adSlotsIds) + " for AdSlots on InventorySource"

		try:
			# get the InventorySource
			inventorySource = self.get( inventorySourceId ).first()
				
			# split on a comma with no spaces
			idList = adSlotsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AdSlot		
				adSlot = AdSlotDelegate().get(id).first();	
				# add the AdSlot
				inventorySource.adSlots.remove(adSlot)
				
			# save it		
			inventorySource.save()
			
			# reload and return the appropriate version
			return self.get( inventorySourceId );
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except AdSlot.DoesNotExist:
			raise ProcessingError(errMsg + " : AdSlot does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDeals( self, inventorySourceId, dealsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DealDelegate import DealDelegate

		errMsg = "Failed to add elements " + str(dealsIds) + " for Deals on InventorySource"

		try:
			# get the InventorySource
			inventorySource = self.get( inventorySourceId ).first()
				
			# split on a comma with no spaces
			idList = dealsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Deal		
				deal = DealDelegate().get(id).first();	
				# add the Deal
				inventorySource.deals.add(deal)
				
			# save it		
			inventorySource.save()
			
			# reload and return the appropriate version
			return self.get( inventorySourceId );
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDeals( self, inventorySourceId, dealsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DealDelegate import DealDelegate

		errMsg = "Failed to remove elements " + str(dealsIds) + " for Deals on InventorySource"

		try:
			# get the InventorySource
			inventorySource = self.get( inventorySourceId ).first()
				
			# split on a comma with no spaces
			idList = dealsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Deal		
				deal = DealDelegate().get(id).first();	
				# add the Deal
				inventorySource.deals.remove(deal)
				
			# save it		
			inventorySource.save()
			
			# reload and return the appropriate version
			return self.get( inventorySourceId );
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource with id " + str(inventorySourceId) + " does not exist.")
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
