from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.AudienceSegment import AudienceSegment
from advertisingOnDjango.models.DataProvider import DataProvider
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AudienceSegment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AudienceSegmentDelegate Declaration
#======================================================================
class AudienceSegmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, audienceSegmentId ):
		try:	
			audienceSegment = AudienceSegment.objects.filter(id=audienceSegmentId)
			return audienceSegment.first();
		except AudienceSegment.DoesNotExist:
			raise ProcessingError("AudienceSegment with id " + str(audienceSegmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, audienceSegment):
		for model in serializers.deserialize("json", audienceSegment):
			model.save()
			return model;

	def create(self, audienceSegment):
		audienceSegment.save()
		return audienceSegment;

	def saveFromJson(self, audienceSegment):
		for model in serializers.deserialize("json", audienceSegment):
			model.save()
			return audienceSegment;
	
	def save(self, audienceSegment):
		audienceSegment.save()
		return audienceSegment;
	
	def delete(self, audienceSegmentId ):
		errMsg = "Failed to delete AudienceSegment from db using id " + str(audienceSegmentId)
		
		try:
			audienceSegment = AudienceSegment.objects.get(id=audienceSegmentId)
			audienceSegment.delete()
			return True
		except AudienceSegment.DoesNotExist:
			raise ProcessingError("AudienceSegment with id " + str(audienceSegmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AudienceSegment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AudienceSegment from db")
		except Exception:
			return None;
		
	def assignProvider( self, audienceSegmentId, providerId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DataProviderDelegate import DataProviderDelegate

		errMsg = "Failed to assign element " + str(providerId) + " for Provider on AudienceSegment"

		try:
			# get the AudienceSegment from db
			audienceSegment = self.get( audienceSegmentId ).first()	
			
			# get the DataProvider from db
			dataProvider = DataProviderDelegate().get(providerId).first();
			
			# assign the Provider		
			audienceSegment.provider = dataProvider
			
			#save it
			audienceSegment.save()

			# reload and return the appropriate version					
			return self.get( audienceSegmentId );
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment with id " + str(audienceSegmentId) + " does not exist.")
		except DataProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProvider with id " + str(providerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProvider( self, audienceSegmentId ):
		errMsg = "Failed to unassign element " + str(providerId) + " for Provider on AudienceSegment"

		try:
			# get the AudienceSegment from db
			audienceSegment = self.get( audienceSegmentId ).first()	
			
			# assign to None for unassignment
			audienceSegment.dataProvider = None			

			#save it
			audienceSegment.save()

			# reload and return the appropriate version					
			return self.get( audienceSegmentId );
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment with id " + str(audienceSegmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addCampaigns( self, audienceSegmentId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on AudienceSegment"

		try:
			# get the AudienceSegment
			audienceSegment = self.get( audienceSegmentId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				audienceSegment.campaigns.add(campaign)
				
			# save it		
			audienceSegment.save()
			
			# reload and return the appropriate version
			return self.get( audienceSegmentId );
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment with id " + str(audienceSegmentId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, audienceSegmentId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on AudienceSegment"

		try:
			# get the AudienceSegment
			audienceSegment = self.get( audienceSegmentId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				audienceSegment.campaigns.remove(campaign)
				
			# save it		
			audienceSegment.save()
			
			# reload and return the appropriate version
			return self.get( audienceSegmentId );
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment with id " + str(audienceSegmentId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
