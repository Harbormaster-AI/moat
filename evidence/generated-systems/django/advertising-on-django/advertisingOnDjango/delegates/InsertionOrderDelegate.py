from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.InsertionOrder import InsertionOrder
from advertisingOnDjango.models.Advertiser import Advertiser
from advertisingOnDjango.models.Agency import Agency
from advertisingOnDjango.models.Publisher import Publisher
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InsertionOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsertionOrderDelegate Declaration
#======================================================================
class InsertionOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, insertionOrderId ):
		try:	
			insertionOrder = InsertionOrder.objects.filter(id=insertionOrderId)
			return insertionOrder.first();
		except InsertionOrder.DoesNotExist:
			raise ProcessingError("InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, insertionOrder):
		for model in serializers.deserialize("json", insertionOrder):
			model.save()
			return model;

	def create(self, insertionOrder):
		insertionOrder.save()
		return insertionOrder;

	def saveFromJson(self, insertionOrder):
		for model in serializers.deserialize("json", insertionOrder):
			model.save()
			return insertionOrder;
	
	def save(self, insertionOrder):
		insertionOrder.save()
		return insertionOrder;
	
	def delete(self, insertionOrderId ):
		errMsg = "Failed to delete InsertionOrder from db using id " + str(insertionOrderId)
		
		try:
			insertionOrder = InsertionOrder.objects.get(id=insertionOrderId)
			insertionOrder.delete()
			return True
		except InsertionOrder.DoesNotExist:
			raise ProcessingError("InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InsertionOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InsertionOrder from db")
		except Exception:
			return None;
		
	def assignAdvertiser( self, insertionOrderId, advertiserId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

		errMsg = "Failed to assign element " + str(advertiserId) + " for Advertiser on InsertionOrder"

		try:
			# get the InsertionOrder from db
			insertionOrder = self.get( insertionOrderId ).first()	
			
			# get the Advertiser from db
			advertiser = AdvertiserDelegate().get(advertiserId).first();
			
			# assign the Advertiser		
			insertionOrder.advertiser = advertiser
			
			#save it
			insertionOrder.save()

			# reload and return the appropriate version					
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdvertiser( self, insertionOrderId ):
		errMsg = "Failed to unassign element " + str(advertiserId) + " for Advertiser on InsertionOrder"

		try:
			# get the InsertionOrder from db
			insertionOrder = self.get( insertionOrderId ).first()	
			
			# assign to None for unassignment
			insertionOrder.advertiser = None			

			#save it
			insertionOrder.save()

			# reload and return the appropriate version					
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAgency( self, insertionOrderId, agencyId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AgencyDelegate import AgencyDelegate

		errMsg = "Failed to assign element " + str(agencyId) + " for Agency on InsertionOrder"

		try:
			# get the InsertionOrder from db
			insertionOrder = self.get( insertionOrderId ).first()	
			
			# get the Agency from db
			agency = AgencyDelegate().get(agencyId).first();
			
			# assign the Agency		
			insertionOrder.agency = agency
			
			#save it
			insertionOrder.save()

			# reload and return the appropriate version					
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAgency( self, insertionOrderId ):
		errMsg = "Failed to unassign element " + str(agencyId) + " for Agency on InsertionOrder"

		try:
			# get the InsertionOrder from db
			insertionOrder = self.get( insertionOrderId ).first()	
			
			# assign to None for unassignment
			insertionOrder.agency = None			

			#save it
			insertionOrder.save()

			# reload and return the appropriate version					
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPublisher( self, insertionOrderId, publisherId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PublisherDelegate import PublisherDelegate

		errMsg = "Failed to assign element " + str(publisherId) + " for Publisher on InsertionOrder"

		try:
			# get the InsertionOrder from db
			insertionOrder = self.get( insertionOrderId ).first()	
			
			# get the Publisher from db
			publisher = PublisherDelegate().get(publisherId).first();
			
			# assign the Publisher		
			insertionOrder.publisher = publisher
			
			#save it
			insertionOrder.save()

			# reload and return the appropriate version					
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Publisher.DoesNotExist:
			raise ProcessingError(errMsg + " : Publisher with id " + str(publisherId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPublisher( self, insertionOrderId ):
		errMsg = "Failed to unassign element " + str(publisherId) + " for Publisher on InsertionOrder"

		try:
			# get the InsertionOrder from db
			insertionOrder = self.get( insertionOrderId ).first()	
			
			# assign to None for unassignment
			insertionOrder.publisher = None			

			#save it
			insertionOrder.save()

			# reload and return the appropriate version					
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addCampaigns( self, insertionOrderId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on InsertionOrder"

		try:
			# get the InsertionOrder
			insertionOrder = self.get( insertionOrderId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				insertionOrder.campaigns.add(campaign)
				
			# save it		
			insertionOrder.save()
			
			# reload and return the appropriate version
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, insertionOrderId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on InsertionOrder"

		try:
			# get the InsertionOrder
			insertionOrder = self.get( insertionOrderId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				insertionOrder.campaigns.remove(campaign)
				
			# save it		
			insertionOrder.save()
			
			# reload and return the appropriate version
			return self.get( insertionOrderId );
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
