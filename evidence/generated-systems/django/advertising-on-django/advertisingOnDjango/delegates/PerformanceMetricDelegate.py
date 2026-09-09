from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.PerformanceMetric import PerformanceMetric
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.models.Placement import Placement
from advertisingOnDjango.models.CreativeAsset import CreativeAsset
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PerformanceMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceMetricDelegate Declaration
#======================================================================
class PerformanceMetricDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, performanceMetricId ):
		try:	
			performanceMetric = PerformanceMetric.objects.filter(id=performanceMetricId)
			return performanceMetric.first();
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError("PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, performanceMetric):
		for model in serializers.deserialize("json", performanceMetric):
			model.save()
			return model;

	def create(self, performanceMetric):
		performanceMetric.save()
		return performanceMetric;

	def saveFromJson(self, performanceMetric):
		for model in serializers.deserialize("json", performanceMetric):
			model.save()
			return performanceMetric;
	
	def save(self, performanceMetric):
		performanceMetric.save()
		return performanceMetric;
	
	def delete(self, performanceMetricId ):
		errMsg = "Failed to delete PerformanceMetric from db using id " + str(performanceMetricId)
		
		try:
			performanceMetric = PerformanceMetric.objects.get(id=performanceMetricId)
			performanceMetric.delete()
			return True
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError("PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PerformanceMetric.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PerformanceMetric from db")
		except Exception:
			return None;
		
	def assignAdAccount( self, performanceMetricId, adAccountId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to assign element " + str(adAccountId) + " for AdAccount on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# get the AdAccount from db
			adAccount = AdAccountDelegate().get(adAccountId).first();
			
			# assign the AdAccount		
			performanceMetric.adAccount = adAccount
			
			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdAccount( self, performanceMetricId ):
		errMsg = "Failed to unassign element " + str(adAccountId) + " for AdAccount on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# assign to None for unassignment
			performanceMetric.adAccount = None			

			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCampaign( self, performanceMetricId, campaignId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			performanceMetric.campaign = campaign
			
			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, performanceMetricId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# assign to None for unassignment
			performanceMetric.campaign = None			

			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLineItem( self, performanceMetricId, lineItemId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to assign element " + str(lineItemId) + " for LineItem on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# get the LineItem from db
			lineItem = LineItemDelegate().get(lineItemId).first();
			
			# assign the LineItem		
			performanceMetric.lineItem = lineItem
			
			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLineItem( self, performanceMetricId ):
		errMsg = "Failed to unassign element " + str(lineItemId) + " for LineItem on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# assign to None for unassignment
			performanceMetric.lineItem = None			

			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlacement( self, performanceMetricId, placementId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PlacementDelegate import PlacementDelegate

		errMsg = "Failed to assign element " + str(placementId) + " for Placement on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# get the Placement from db
			placement = PlacementDelegate().get(placementId).first();
			
			# assign the Placement		
			performanceMetric.placement = placement
			
			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except Placement.DoesNotExist:
			raise ProcessingError(errMsg + " : Placement with id " + str(placementId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlacement( self, performanceMetricId ):
		errMsg = "Failed to unassign element " + str(placementId) + " for Placement on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# assign to None for unassignment
			performanceMetric.placement = None			

			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCreativeAsset( self, performanceMetricId, creativeAssetId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CreativeAssetDelegate import CreativeAssetDelegate

		errMsg = "Failed to assign element " + str(creativeAssetId) + " for CreativeAsset on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# get the CreativeAsset from db
			creativeAsset = CreativeAssetDelegate().get(creativeAssetId).first();
			
			# assign the CreativeAsset		
			performanceMetric.creativeAsset = creativeAsset
			
			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except CreativeAsset.DoesNotExist:
			raise ProcessingError(errMsg + " : CreativeAsset with id " + str(creativeAssetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCreativeAsset( self, performanceMetricId ):
		errMsg = "Failed to unassign element " + str(creativeAssetId) + " for CreativeAsset on PerformanceMetric"

		try:
			# get the PerformanceMetric from db
			performanceMetric = self.get( performanceMetricId ).first()	
			
			# assign to None for unassignment
			performanceMetric.creativeAsset = None			

			#save it
			performanceMetric.save()

			# reload and return the appropriate version					
			return self.get( performanceMetricId );
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric with id " + str(performanceMetricId) + " does not exist.")
		except Exception:
			return None;
		
