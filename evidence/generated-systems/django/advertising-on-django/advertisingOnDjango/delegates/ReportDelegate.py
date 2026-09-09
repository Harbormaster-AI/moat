from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Report import Report
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.LineItem import LineItem
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Report
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReportDelegate Declaration
#======================================================================
class ReportDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, reportId ):
		try:	
			report = Report.objects.filter(id=reportId)
			return report.first();
		except Report.DoesNotExist:
			raise ProcessingError("Report with id " + str(reportId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, report):
		for model in serializers.deserialize("json", report):
			model.save()
			return model;

	def create(self, report):
		report.save()
		return report;

	def saveFromJson(self, report):
		for model in serializers.deserialize("json", report):
			model.save()
			return report;
	
	def save(self, report):
		report.save()
		return report;
	
	def delete(self, reportId ):
		errMsg = "Failed to delete Report from db using id " + str(reportId)
		
		try:
			report = Report.objects.get(id=reportId)
			report.delete()
			return True
		except Report.DoesNotExist:
			raise ProcessingError("Report with id " + str(reportId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Report.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Report from db")
		except Exception:
			return None;
		
	def assignAdAccount( self, reportId, adAccountId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to assign element " + str(adAccountId) + " for AdAccount on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# get the AdAccount from db
			adAccount = AdAccountDelegate().get(adAccountId).first();
			
			# assign the AdAccount		
			report.adAccount = adAccount
			
			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdAccount( self, reportId ):
		errMsg = "Failed to unassign element " + str(adAccountId) + " for AdAccount on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# assign to None for unassignment
			report.adAccount = None			

			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCampaign( self, reportId, campaignId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			report.campaign = campaign
			
			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, reportId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# assign to None for unassignment
			report.campaign = None			

			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLineItem( self, reportId, lineItemId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.LineItemDelegate import LineItemDelegate

		errMsg = "Failed to assign element " + str(lineItemId) + " for LineItem on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# get the LineItem from db
			lineItem = LineItemDelegate().get(lineItemId).first();
			
			# assign the LineItem		
			report.lineItem = lineItem
			
			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except LineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : LineItem with id " + str(lineItemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLineItem( self, reportId ):
		errMsg = "Failed to unassign element " + str(lineItemId) + " for LineItem on Report"

		try:
			# get the Report from db
			report = self.get( reportId ).first()	
			
			# assign to None for unassignment
			report.lineItem = None			

			#save it
			report.save()

			# reload and return the appropriate version					
			return self.get( reportId );
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Exception:
			return None;
		
