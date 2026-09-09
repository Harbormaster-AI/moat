from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.ConversionEvent import ConversionEvent
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.models.TrackingPixel import TrackingPixel
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ConversionEvent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConversionEventDelegate Declaration
#======================================================================
class ConversionEventDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, conversionEventId ):
		try:	
			conversionEvent = ConversionEvent.objects.filter(id=conversionEventId)
			return conversionEvent.first();
		except ConversionEvent.DoesNotExist:
			raise ProcessingError("ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, conversionEvent):
		for model in serializers.deserialize("json", conversionEvent):
			model.save()
			return model;

	def create(self, conversionEvent):
		conversionEvent.save()
		return conversionEvent;

	def saveFromJson(self, conversionEvent):
		for model in serializers.deserialize("json", conversionEvent):
			model.save()
			return conversionEvent;
	
	def save(self, conversionEvent):
		conversionEvent.save()
		return conversionEvent;
	
	def delete(self, conversionEventId ):
		errMsg = "Failed to delete ConversionEvent from db using id " + str(conversionEventId)
		
		try:
			conversionEvent = ConversionEvent.objects.get(id=conversionEventId)
			conversionEvent.delete()
			return True
		except ConversionEvent.DoesNotExist:
			raise ProcessingError("ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ConversionEvent.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ConversionEvent from db")
		except Exception:
			return None;
		
	def assignCampaign( self, conversionEventId, campaignId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on ConversionEvent"

		try:
			# get the ConversionEvent from db
			conversionEvent = self.get( conversionEventId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			conversionEvent.campaign = campaign
			
			#save it
			conversionEvent.save()

			# reload and return the appropriate version					
			return self.get( conversionEventId );
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, conversionEventId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on ConversionEvent"

		try:
			# get the ConversionEvent from db
			conversionEvent = self.get( conversionEventId ).first()	
			
			# assign to None for unassignment
			conversionEvent.campaign = None			

			#save it
			conversionEvent.save()

			# reload and return the appropriate version					
			return self.get( conversionEventId );
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLineItem( self, conversionEventId, lineItemId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to assign element " + str(lineItemId) + " for LineItem on ConversionEvent"

		try:
			# get the ConversionEvent from db
			conversionEvent = self.get( conversionEventId ).first()	
			
			# get the LineItem from db
			lineItem = LineItemDelegate().get(lineItemId).first();
			
			# assign the LineItem		
			conversionEvent.lineItem = lineItem
			
			#save it
			conversionEvent.save()

			# reload and return the appropriate version					
			return self.get( conversionEventId );
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLineItem( self, conversionEventId ):
		errMsg = "Failed to unassign element " + str(lineItemId) + " for LineItem on ConversionEvent"

		try:
			# get the ConversionEvent from db
			conversionEvent = self.get( conversionEventId ).first()	
			
			# assign to None for unassignment
			conversionEvent.lineItem = None			

			#save it
			conversionEvent.save()

			# reload and return the appropriate version					
			return self.get( conversionEventId );
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTrackingPixel( self, conversionEventId, trackingPixelId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TrackingPixelDelegate import TrackingPixelDelegate

		errMsg = "Failed to assign element " + str(trackingPixelId) + " for TrackingPixel on ConversionEvent"

		try:
			# get the ConversionEvent from db
			conversionEvent = self.get( conversionEventId ).first()	
			
			# get the TrackingPixel from db
			trackingPixel = TrackingPixelDelegate().get(trackingPixelId).first();
			
			# assign the TrackingPixel		
			conversionEvent.trackingPixel = trackingPixel
			
			#save it
			conversionEvent.save()

			# reload and return the appropriate version					
			return self.get( conversionEventId );
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel with id " + str(trackingPixelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTrackingPixel( self, conversionEventId ):
		errMsg = "Failed to unassign element " + str(trackingPixelId) + " for TrackingPixel on ConversionEvent"

		try:
			# get the ConversionEvent from db
			conversionEvent = self.get( conversionEventId ).first()	
			
			# assign to None for unassignment
			conversionEvent.trackingPixel = None			

			#save it
			conversionEvent.save()

			# reload and return the appropriate version					
			return self.get( conversionEventId );
		except ConversionEvent.DoesNotExist:
			raise ProcessingError(errMsg + " : ConversionEvent with id " + str(conversionEventId) + " does not exist.")
		except Exception:
			return None;
		
