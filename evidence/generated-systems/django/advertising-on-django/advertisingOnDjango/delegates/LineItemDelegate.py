from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.Placement import Placement
from advertisingOnDjango.models.TargetingProfile import TargetingProfile
from advertisingOnDjango.models.Deal import Deal
from advertisingOnDjango.models.CreativeAsset import CreativeAsset
from advertisingOnDjango.models.PerformanceMetric import PerformanceMetric
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LineItemDelegate Declaration
#======================================================================
class LineItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, lineItemId ):
		try:	
			lineItem = LineItem.objects.filter(id=lineItemId)
			return lineItem.first();
		except LineItem.DoesNotExist:
			raise ProcessingError("LineItem with id " + str(lineItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, lineItem):
		for model in serializers.deserialize("json", lineItem):
			model.save()
			return model;

	def create(self, lineItem):
		lineItem.save()
		return lineItem;

	def saveFromJson(self, lineItem):
		for model in serializers.deserialize("json", lineItem):
			model.save()
			return lineItem;
	
	def save(self, lineItem):
		lineItem.save()
		return lineItem;
	
	def delete(self, lineItemId ):
		errMsg = "Failed to delete LineItem from db using id " + str(lineItemId)
		
		try:
			lineItem = LineItem.objects.get(id=lineItemId)
			lineItem.delete()
			return True
		except LineItem.DoesNotExist:
			raise ProcessingError("LineItem with id " + str(lineItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LineItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LineItem from db")
		except Exception:
			return None;
		
	def assignCampaign( self, lineItemId, campaignId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on LineItem"

		try:
			# get the LineItem from db
			lineItem = self.get( lineItemId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			lineItem.campaign = campaign
			
			#save it
			lineItem.save()

			# reload and return the appropriate version					
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, lineItemId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on LineItem"

		try:
			# get the LineItem from db
			lineItem = self.get( lineItemId ).first()	
			
			# assign to None for unassignment
			lineItem.campaign = None			

			#save it
			lineItem.save()

			# reload and return the appropriate version					
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTargetingProfile( self, lineItemId, targetingProfileId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TargetingProfileDelegate import TargetingProfileDelegate

		errMsg = "Failed to assign element " + str(targetingProfileId) + " for TargetingProfile on LineItem"

		try:
			# get the LineItem from db
			lineItem = self.get( lineItemId ).first()	
			
			# get the TargetingProfile from db
			targetingProfile = TargetingProfileDelegate().get(targetingProfileId).first();
			
			# assign the TargetingProfile		
			lineItem.targetingProfile = targetingProfile
			
			#save it
			lineItem.save()

			# reload and return the appropriate version					
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except TargetingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : TargetingProfile with id " + str(targetingProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTargetingProfile( self, lineItemId ):
		errMsg = "Failed to unassign element " + str(targetingProfileId) + " for TargetingProfile on LineItem"

		try:
			# get the LineItem from db
			lineItem = self.get( lineItemId ).first()	
			
			# assign to None for unassignment
			lineItem.targetingProfile = None			

			#save it
			lineItem.save()

			# reload and return the appropriate version					
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDeal( self, lineItemId, dealId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DealDelegate import DealDelegate

		errMsg = "Failed to assign element " + str(dealId) + " for Deal on LineItem"

		try:
			# get the LineItem from db
			lineItem = self.get( lineItemId ).first()	
			
			# get the Deal from db
			deal = DealDelegate().get(dealId).first();
			
			# assign the Deal		
			lineItem.deal = deal
			
			#save it
			lineItem.save()

			# reload and return the appropriate version					
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Deal.DoesNotExist:
			raise ProcessingError(errMsg + " : Deal with id " + str(dealId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDeal( self, lineItemId ):
		errMsg = "Failed to unassign element " + str(dealId) + " for Deal on LineItem"

		try:
			# get the LineItem from db
			lineItem = self.get( lineItemId ).first()	
			
			# assign to None for unassignment
			lineItem.deal = None			

			#save it
			lineItem.save()

			# reload and return the appropriate version					
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
		
	def addPlacements( self, lineItemId, placementsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

		errMsg = "Failed to add elements " + str(placementsIds) + " for Placements on LineItem"

		try:
			# get the LineItem
			lineItem = self.get( lineItemId ).first()
				
			# split on a comma with no spaces
			idList = placementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Placement		
				placement = PlacementDelegate().get(id).first();	
				# add the Placement
				lineItem.placements.add(placement)
				
			# save it		
			lineItem.save()
			
			# reload and return the appropriate version
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlacements( self, lineItemId, placementsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

		errMsg = "Failed to remove elements " + str(placementsIds) + " for Placements on LineItem"

		try:
			# get the LineItem
			lineItem = self.get( lineItemId ).first()
				
			# split on a comma with no spaces
			idList = placementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Placement		
				placement = PlacementDelegate().get(id).first();	
				# add the Placement
				lineItem.placements.remove(placement)
				
			# save it		
			lineItem.save()
			
			# reload and return the appropriate version
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCreatives( self, lineItemId, creativesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

		errMsg = "Failed to add elements " + str(creativesIds) + " for Creatives on LineItem"

		try:
			# get the LineItem
			lineItem = self.get( lineItemId ).first()
				
			# split on a comma with no spaces
			idList = creativesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CreativeAsset		
				creativeAsset = CreativeAssetDelegate().get(id).first();	
				# add the CreativeAsset
				lineItem.creatives.add(creativeAsset)
				
			# save it		
			lineItem.save()
			
			# reload and return the appropriate version
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCreatives( self, lineItemId, creativesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

		errMsg = "Failed to remove elements " + str(creativesIds) + " for Creatives on LineItem"

		try:
			# get the LineItem
			lineItem = self.get( lineItemId ).first()
				
			# split on a comma with no spaces
			idList = creativesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CreativeAsset		
				creativeAsset = CreativeAssetDelegate().get(id).first();	
				# add the CreativeAsset
				lineItem.creatives.remove(creativeAsset)
				
			# save it		
			lineItem.save()
			
			# reload and return the appropriate version
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPerformanceMetrics( self, lineItemId, performanceMetricsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PerformanceMetricDelegate import PerformanceMetricDelegate

		errMsg = "Failed to add elements " + str(performanceMetricsIds) + " for PerformanceMetrics on LineItem"

		try:
			# get the LineItem
			lineItem = self.get( lineItemId ).first()
				
			# split on a comma with no spaces
			idList = performanceMetricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PerformanceMetric		
				performanceMetric = PerformanceMetricDelegate().get(id).first();	
				# add the PerformanceMetric
				lineItem.performanceMetrics.add(performanceMetric)
				
			# save it		
			lineItem.save()
			
			# reload and return the appropriate version
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePerformanceMetrics( self, lineItemId, performanceMetricsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PerformanceMetricDelegate import PerformanceMetricDelegate

		errMsg = "Failed to remove elements " + str(performanceMetricsIds) + " for PerformanceMetrics on LineItem"

		try:
			# get the LineItem
			lineItem = self.get( lineItemId ).first()
				
			# split on a comma with no spaces
			idList = performanceMetricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PerformanceMetric		
				performanceMetric = PerformanceMetricDelegate().get(id).first();	
				# add the PerformanceMetric
				lineItem.performanceMetrics.remove(performanceMetric)
				
			# save it		
			lineItem.save()
			
			# reload and return the appropriate version
			return self.get( lineItemId );
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
