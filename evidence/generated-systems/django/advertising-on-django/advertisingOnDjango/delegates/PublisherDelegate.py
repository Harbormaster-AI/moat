from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Publisher import Publisher
from advertisingOnDjango.models.InventorySource import InventorySource
from advertisingOnDjango.models.Deal import Deal
from advertisingOnDjango.models.CreativeApproval import CreativeApproval
from advertisingOnDjango.models.InsertionOrder import InsertionOrder
from advertisingOnDjango.models.RateCard import RateCard
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Publisher
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PublisherDelegate Declaration
#======================================================================
class PublisherDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, publisherId ):
		try:	
			publisher = Publisher.objects.filter(id=publisherId)
			return publisher.first();
		except Publisher.DoesNotExist:
			raise ProcessingError("Publisher with id " + str(publisherId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, publisher):
		for model in serializers.deserialize("json", publisher):
			model.save()
			return model;

	def create(self, publisher):
		publisher.save()
		return publisher;

	def saveFromJson(self, publisher):
		for model in serializers.deserialize("json", publisher):
			model.save()
			return publisher;
	
	def save(self, publisher):
		publisher.save()
		return publisher;
	
	def delete(self, publisherId ):
		errMsg = "Failed to delete Publisher from db using id " + str(publisherId)
		
		try:
			publisher = Publisher.objects.get(id=publisherId)
			publisher.delete()
			return True
		except Publisher.DoesNotExist:
			raise ProcessingError("Publisher with id " + str(publisherId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Publisher.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Publisher from db")
		except Exception:
			return None;
		
	def addInventorySources( self, publisherId, inventorySourcesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InventorySourceDelegate import InventorySourceDelegate

		errMsg = "Failed to add elements " + str(inventorySourcesIds) + " for InventorySources on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = inventorySourcesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventorySource		
				inventorySource = InventorySourceDelegate().get(id).first();	
				# add the InventorySource
				publisher.inventorySources.add(inventorySource)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventorySources( self, publisherId, inventorySourcesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InventorySourceDelegate import InventorySourceDelegate

		errMsg = "Failed to remove elements " + str(inventorySourcesIds) + " for InventorySources on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = inventorySourcesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventorySource		
				inventorySource = InventorySourceDelegate().get(id).first();	
				# add the InventorySource
				publisher.inventorySources.remove(inventorySource)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except InventorySource.DoesNotExist:
			raise ProcessingError(errMsg + " : InventorySource does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDeals( self, publisherId, dealsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DealDelegate import DealDelegate

		errMsg = "Failed to add elements " + str(dealsIds) + " for Deals on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = dealsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Deal		
				deal = DealDelegate().get(id).first();	
				# add the Deal
				publisher.deals.add(deal)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDeals( self, publisherId, dealsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DealDelegate import DealDelegate

		errMsg = "Failed to remove elements " + str(dealsIds) + " for Deals on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = dealsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Deal		
				deal = DealDelegate().get(id).first();	
				# add the Deal
				publisher.deals.remove(deal)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCreativeApprovals( self, publisherId, creativeApprovalsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeApprovalDelegate import CreativeApprovalDelegate

		errMsg = "Failed to add elements " + str(creativeApprovalsIds) + " for CreativeApprovals on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = creativeApprovalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CreativeApproval		
				creativeApproval = CreativeApprovalDelegate().get(id).first();	
				# add the CreativeApproval
				publisher.creativeApprovals.add(creativeApproval)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCreativeApprovals( self, publisherId, creativeApprovalsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeApprovalDelegate import CreativeApprovalDelegate

		errMsg = "Failed to remove elements " + str(creativeApprovalsIds) + " for CreativeApprovals on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = creativeApprovalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CreativeApproval		
				creativeApproval = CreativeApprovalDelegate().get(id).first();	
				# add the CreativeApproval
				publisher.creativeApprovals.remove(creativeApproval)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except CreativeApproval.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeApproval does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInsertionOrders( self, publisherId, insertionOrdersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InsertionOrderDelegate import InsertionOrderDelegate

		errMsg = "Failed to add elements " + str(insertionOrdersIds) + " for InsertionOrders on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = insertionOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InsertionOrder		
				insertionOrder = InsertionOrderDelegate().get(id).first();	
				# add the InsertionOrder
				publisher.insertionOrders.add(insertionOrder)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInsertionOrders( self, publisherId, insertionOrdersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InsertionOrderDelegate import InsertionOrderDelegate

		errMsg = "Failed to remove elements " + str(insertionOrdersIds) + " for InsertionOrders on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = insertionOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InsertionOrder		
				insertionOrder = InsertionOrderDelegate().get(id).first();	
				# add the InsertionOrder
				publisher.insertionOrders.remove(insertionOrder)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRateCards( self, publisherId, rateCardsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.RateCardDelegate import RateCardDelegate

		errMsg = "Failed to add elements " + str(rateCardsIds) + " for RateCards on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = rateCardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RateCard		
				rateCard = RateCardDelegate().get(id).first();	
				# add the RateCard
				publisher.rateCards.add(rateCard)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except RateCard.DoesNotExist:
			raise ProcessingError(errMsg + " : RateCard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRateCards( self, publisherId, rateCardsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.RateCardDelegate import RateCardDelegate

		errMsg = "Failed to remove elements " + str(rateCardsIds) + " for RateCards on Publisher"

		try:
			# get the Publisher
			publisher = self.get( publisherId ).first()
				
			# split on a comma with no spaces
			idList = rateCardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RateCard		
				rateCard = RateCardDelegate().get(id).first();	
				# add the RateCard
				publisher.rateCards.remove(rateCard)
				
			# save it		
			publisher.save()
			
			# reload and return the appropriate version
			return self.get( publisherId );
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except RateCard.DoesNotExist:
			raise ProcessingError(errMsg + " : RateCard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
