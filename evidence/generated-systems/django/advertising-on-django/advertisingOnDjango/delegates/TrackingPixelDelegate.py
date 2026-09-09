from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.TrackingPixel import TrackingPixel
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.Advertiser import Advertiser
from advertisingOnDjango.models.ConversionEvent import ConversionEvent
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TrackingPixel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrackingPixelDelegate Declaration
#======================================================================
class TrackingPixelDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, trackingPixelId ):
		try:	
			trackingPixel = TrackingPixel.objects.filter(id=trackingPixelId)
			return trackingPixel.first();
		except TrackingPixel.DoesNotExist:
			raise ProcessingError("TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, trackingPixel):
		for model in serializers.deserialize("json", trackingPixel):
			model.save()
			return model;

	def create(self, trackingPixel):
		trackingPixel.save()
		return trackingPixel;

	def saveFromJson(self, trackingPixel):
		for model in serializers.deserialize("json", trackingPixel):
			model.save()
			return trackingPixel;
	
	def save(self, trackingPixel):
		trackingPixel.save()
		return trackingPixel;
	
	def delete(self, trackingPixelId ):
		errMsg = "Failed to delete TrackingPixel from db using id " + str(trackingPixelId)
		
		try:
			trackingPixel = TrackingPixel.objects.get(id=trackingPixelId)
			trackingPixel.delete()
			return True
		except TrackingPixel.DoesNotExist:
			raise ProcessingError("TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TrackingPixel.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TrackingPixel from db")
		except Exception:
			return None;
		
	def assignCampaign( self, trackingPixelId, campaignId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on TrackingPixel"

		try:
			# get the TrackingPixel from db
			trackingPixel = self.get( trackingPixelId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			trackingPixel.campaign = campaign
			
			#save it
			trackingPixel.save()

			# reload and return the appropriate version					
			return self.get( trackingPixelId );
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, trackingPixelId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on TrackingPixel"

		try:
			# get the TrackingPixel from db
			trackingPixel = self.get( trackingPixelId ).first()	
			
			# assign to None for unassignment
			trackingPixel.campaign = None			

			#save it
			trackingPixel.save()

			# reload and return the appropriate version					
			return self.get( trackingPixelId );
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAdvertiser( self, trackingPixelId, advertiserId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

		errMsg = "Failed to assign element " + str(advertiserId) + " for Advertiser on TrackingPixel"

		try:
			# get the TrackingPixel from db
			trackingPixel = self.get( trackingPixelId ).first()	
			
			# get the Advertiser from db
			advertiser = AdvertiserDelegate().get(advertiserId).first();
			
			# assign the Advertiser		
			trackingPixel.advertiser = advertiser
			
			#save it
			trackingPixel.save()

			# reload and return the appropriate version					
			return self.get( trackingPixelId );
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdvertiser( self, trackingPixelId ):
		errMsg = "Failed to unassign element " + str(advertiserId) + " for Advertiser on TrackingPixel"

		try:
			# get the TrackingPixel from db
			trackingPixel = self.get( trackingPixelId ).first()	
			
			# assign to None for unassignment
			trackingPixel.advertiser = None			

			#save it
			trackingPixel.save()

			# reload and return the appropriate version					
			return self.get( trackingPixelId );
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except Exception:
			return None;
		
	def addConversionEvents( self, trackingPixelId, conversionEventsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ConversionEventDelegate import ConversionEventDelegate

		errMsg = "Failed to add elements " + str(conversionEventsIds) + " for ConversionEvents on TrackingPixel"

		try:
			# get the TrackingPixel
			trackingPixel = self.get( trackingPixelId ).first()
				
			# split on a comma with no spaces
			idList = conversionEventsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ConversionEvent		
				conversionEvent = ConversionEventDelegate().get(id).first();	
				# add the ConversionEvent
				trackingPixel.conversionEvents.add(conversionEvent)
				
			# save it		
			trackingPixel.save()
			
			# reload and return the appropriate version
			return self.get( trackingPixelId );
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeConversionEvents( self, trackingPixelId, conversionEventsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ConversionEventDelegate import ConversionEventDelegate

		errMsg = "Failed to remove elements " + str(conversionEventsIds) + " for ConversionEvents on TrackingPixel"

		try:
			# get the TrackingPixel
			trackingPixel = self.get( trackingPixelId ).first()
				
			# split on a comma with no spaces
			idList = conversionEventsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ConversionEvent		
				conversionEvent = ConversionEventDelegate().get(id).first();	
				# add the ConversionEvent
				trackingPixel.conversionEvents.remove(conversionEvent)
				
			# save it		
			trackingPixel.save()
			
			# reload and return the appropriate version
			return self.get( trackingPixelId );
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
