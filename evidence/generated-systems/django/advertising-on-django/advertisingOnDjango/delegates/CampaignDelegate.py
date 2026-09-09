from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.models.KPI import KPI
from advertisingOnDjango.models.TrackingPixel import TrackingPixel
from advertisingOnDjango.models.AudienceSegment import AudienceSegment
from advertisingOnDjango.models.Report import Report
from advertisingOnDjango.models.InsertionOrder import InsertionOrder
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Campaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignDelegate Declaration
#======================================================================
class CampaignDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, campaignId ):
		try:	
			campaign = Campaign.objects.filter(id=campaignId)
			return campaign.first();
		except Campaign.DoesNotExist:
			raise ProcessingError("Campaign with id " + str(campaignId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, campaign):
		for model in serializers.deserialize("json", campaign):
			model.save()
			return model;

	def create(self, campaign):
		campaign.save()
		return campaign;

	def saveFromJson(self, campaign):
		for model in serializers.deserialize("json", campaign):
			model.save()
			return campaign;
	
	def save(self, campaign):
		campaign.save()
		return campaign;
	
	def delete(self, campaignId ):
		errMsg = "Failed to delete Campaign from db using id " + str(campaignId)
		
		try:
			campaign = Campaign.objects.get(id=campaignId)
			campaign.delete()
			return True
		except Campaign.DoesNotExist:
			raise ProcessingError("Campaign with id " + str(campaignId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Campaign.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Campaign from db")
		except Exception:
			return None;
		
	def assignAdAccount( self, campaignId, adAccountId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to assign element " + str(adAccountId) + " for AdAccount on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# get the AdAccount from db
			adAccount = AdAccountDelegate().get(adAccountId).first();
			
			# assign the AdAccount		
			campaign.adAccount = adAccount
			
			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdAccount( self, campaignId ):
		errMsg = "Failed to unassign element " + str(adAccountId) + " for AdAccount on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# assign to None for unassignment
			campaign.adAccount = None			

			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInsertionOrder( self, campaignId, insertionOrderId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InsertionOrderDelegate import InsertionOrderDelegate

		errMsg = "Failed to assign element " + str(insertionOrderId) + " for InsertionOrder on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# get the InsertionOrder from db
			insertionOrder = InsertionOrderDelegate().get(insertionOrderId).first();
			
			# assign the InsertionOrder		
			campaign.insertionOrder = insertionOrder
			
			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder with id " + str(insertionOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInsertionOrder( self, campaignId ):
		errMsg = "Failed to unassign element " + str(insertionOrderId) + " for InsertionOrder on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# assign to None for unassignment
			campaign.insertionOrder = None			

			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
		
	def addLineItems( self, campaignId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to add elements " + str(lineItemsIds) + " for LineItems on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LineItem		
				lineItem = LineItemDelegate().get(id).first();	
				# add the LineItem
				campaign.lineItems.add(lineItem)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLineItems( self, campaignId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to remove elements " + str(lineItemsIds) + " for LineItems on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LineItem		
				lineItem = LineItemDelegate().get(id).first();	
				# add the LineItem
				campaign.lineItems.remove(lineItem)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addKpis( self, campaignId, kpisIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.KPIDelegate import KPIDelegate

		errMsg = "Failed to add elements " + str(kpisIds) + " for Kpis on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = kpisIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the KPI		
				kPI = KPIDelegate().get(id).first();	
				# add the KPI
				campaign.kpis.add(kPI)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except KPI.DoesNotExist:
			raise ProcessingError(errMsg + " : KPI does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeKpis( self, campaignId, kpisIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.KPIDelegate import KPIDelegate

		errMsg = "Failed to remove elements " + str(kpisIds) + " for Kpis on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = kpisIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the KPI		
				kPI = KPIDelegate().get(id).first();	
				# add the KPI
				campaign.kpis.remove(kPI)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except KPI.DoesNotExist:
			raise ProcessingError(errMsg + " : KPI does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTrackingPixels( self, campaignId, trackingPixelsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TrackingPixelDelegate import TrackingPixelDelegate

		errMsg = "Failed to add elements " + str(trackingPixelsIds) + " for TrackingPixels on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = trackingPixelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrackingPixel		
				trackingPixel = TrackingPixelDelegate().get(id).first();	
				# add the TrackingPixel
				campaign.trackingPixels.add(trackingPixel)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrackingPixels( self, campaignId, trackingPixelsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TrackingPixelDelegate import TrackingPixelDelegate

		errMsg = "Failed to remove elements " + str(trackingPixelsIds) + " for TrackingPixels on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = trackingPixelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrackingPixel		
				trackingPixel = TrackingPixelDelegate().get(id).first();	
				# add the TrackingPixel
				campaign.trackingPixels.remove(trackingPixel)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAudiences( self, campaignId, audiencesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

		errMsg = "Failed to add elements " + str(audiencesIds) + " for Audiences on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = audiencesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AudienceSegment		
				audienceSegment = AudienceSegmentDelegate().get(id).first();	
				# add the AudienceSegment
				campaign.audiences.add(audienceSegment)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAudiences( self, campaignId, audiencesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

		errMsg = "Failed to remove elements " + str(audiencesIds) + " for Audiences on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = audiencesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AudienceSegment		
				audienceSegment = AudienceSegmentDelegate().get(id).first();	
				# add the AudienceSegment
				campaign.audiences.remove(audienceSegment)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReports( self, campaignId, reportsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				campaign.reports.add(report)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, campaignId, reportsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				campaign.reports.remove(report)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
