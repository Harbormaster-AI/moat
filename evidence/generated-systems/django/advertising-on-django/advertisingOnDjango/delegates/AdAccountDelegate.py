from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.models.Advertiser import Advertiser
from advertisingOnDjango.models.User import User
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.BillingProfile import BillingProfile
from advertisingOnDjango.models.DSP import DSP
from advertisingOnDjango.models.PerformanceMetric import PerformanceMetric
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AdAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdAccountDelegate Declaration
#======================================================================
class AdAccountDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, adAccountId ):
		try:	
			adAccount = AdAccount.objects.filter(id=adAccountId)
			return adAccount.first();
		except AdAccount.DoesNotExist:
			raise ProcessingError("AdAccount with id " + str(adAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, adAccount):
		for model in serializers.deserialize("json", adAccount):
			model.save()
			return model;

	def create(self, adAccount):
		adAccount.save()
		return adAccount;

	def saveFromJson(self, adAccount):
		for model in serializers.deserialize("json", adAccount):
			model.save()
			return adAccount;
	
	def save(self, adAccount):
		adAccount.save()
		return adAccount;
	
	def delete(self, adAccountId ):
		errMsg = "Failed to delete AdAccount from db using id " + str(adAccountId)
		
		try:
			adAccount = AdAccount.objects.get(id=adAccountId)
			adAccount.delete()
			return True
		except AdAccount.DoesNotExist:
			raise ProcessingError("AdAccount with id " + str(adAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AdAccount.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AdAccount from db")
		except Exception:
			return None;
		
	def assignAdvertiser( self, adAccountId, advertiserId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

		errMsg = "Failed to assign element " + str(advertiserId) + " for Advertiser on AdAccount"

		try:
			# get the AdAccount from db
			adAccount = self.get( adAccountId ).first()	
			
			# get the Advertiser from db
			advertiser = AdvertiserDelegate().get(advertiserId).first();
			
			# assign the Advertiser		
			adAccount.advertiser = advertiser
			
			#save it
			adAccount.save()

			# reload and return the appropriate version					
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdvertiser( self, adAccountId ):
		errMsg = "Failed to unassign element " + str(advertiserId) + " for Advertiser on AdAccount"

		try:
			# get the AdAccount from db
			adAccount = self.get( adAccountId ).first()	
			
			# assign to None for unassignment
			adAccount.advertiser = None			

			#save it
			adAccount.save()

			# reload and return the appropriate version					
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Exception:
			return None;
		
	def assignBillingProfile( self, adAccountId, billingProfileId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.BillingProfileDelegate import BillingProfileDelegate

		errMsg = "Failed to assign element " + str(billingProfileId) + " for BillingProfile on AdAccount"

		try:
			# get the AdAccount from db
			adAccount = self.get( adAccountId ).first()	
			
			# get the BillingProfile from db
			billingProfile = BillingProfileDelegate().get(billingProfileId).first();
			
			# assign the BillingProfile		
			adAccount.billingProfile = billingProfile
			
			#save it
			adAccount.save()

			# reload and return the appropriate version					
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile with id " + str(billingProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBillingProfile( self, adAccountId ):
		errMsg = "Failed to unassign element " + str(billingProfileId) + " for BillingProfile on AdAccount"

		try:
			# get the AdAccount from db
			adAccount = self.get( adAccountId ).first()	
			
			# assign to None for unassignment
			adAccount.billingProfile = None			

			#save it
			adAccount.save()

			# reload and return the appropriate version					
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDsp( self, adAccountId, dspId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.DSPDelegate import DSPDelegate

		errMsg = "Failed to assign element " + str(dspId) + " for Dsp on AdAccount"

		try:
			# get the AdAccount from db
			adAccount = self.get( adAccountId ).first()	
			
			# get the DSP from db
			dSP = DSPDelegate().get(dspId).first();
			
			# assign the Dsp		
			adAccount.dsp = dSP
			
			#save it
			adAccount.save()

			# reload and return the appropriate version					
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except DSP.DoesNotExist:
			raise ProcessingError(errMsg + " : DSP with id " + str(dspId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDsp( self, adAccountId ):
		errMsg = "Failed to unassign element " + str(dspId) + " for Dsp on AdAccount"

		try:
			# get the AdAccount from db
			adAccount = self.get( adAccountId ).first()	
			
			# assign to None for unassignment
			adAccount.dSP = None			

			#save it
			adAccount.save()

			# reload and return the appropriate version					
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Exception:
			return None;
		
	def addUsers( self, adAccountId, usersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to add elements " + str(usersIds) + " for Users on AdAccount"

		try:
			# get the AdAccount
			adAccount = self.get( adAccountId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				adAccount.users.add(user)
				
			# save it		
			adAccount.save()
			
			# reload and return the appropriate version
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUsers( self, adAccountId, usersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to remove elements " + str(usersIds) + " for Users on AdAccount"

		try:
			# get the AdAccount
			adAccount = self.get( adAccountId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				adAccount.users.remove(user)
				
			# save it		
			adAccount.save()
			
			# reload and return the appropriate version
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, adAccountId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on AdAccount"

		try:
			# get the AdAccount
			adAccount = self.get( adAccountId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				adAccount.campaigns.add(campaign)
				
			# save it		
			adAccount.save()
			
			# reload and return the appropriate version
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, adAccountId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on AdAccount"

		try:
			# get the AdAccount
			adAccount = self.get( adAccountId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				adAccount.campaigns.remove(campaign)
				
			# save it		
			adAccount.save()
			
			# reload and return the appropriate version
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPerformanceMetrics( self, adAccountId, performanceMetricsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PerformanceMetricDelegate import PerformanceMetricDelegate

		errMsg = "Failed to add elements " + str(performanceMetricsIds) + " for PerformanceMetrics on AdAccount"

		try:
			# get the AdAccount
			adAccount = self.get( adAccountId ).first()
				
			# split on a comma with no spaces
			idList = performanceMetricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PerformanceMetric		
				performanceMetric = PerformanceMetricDelegate().get(id).first();	
				# add the PerformanceMetric
				adAccount.performanceMetrics.add(performanceMetric)
				
			# save it		
			adAccount.save()
			
			# reload and return the appropriate version
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePerformanceMetrics( self, adAccountId, performanceMetricsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.PerformanceMetricDelegate import PerformanceMetricDelegate

		errMsg = "Failed to remove elements " + str(performanceMetricsIds) + " for PerformanceMetrics on AdAccount"

		try:
			# get the AdAccount
			adAccount = self.get( adAccountId ).first()
				
			# split on a comma with no spaces
			idList = performanceMetricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PerformanceMetric		
				performanceMetric = PerformanceMetricDelegate().get(id).first();	
				# add the PerformanceMetric
				adAccount.performanceMetrics.remove(performanceMetric)
				
			# save it		
			adAccount.save()
			
			# reload and return the appropriate version
			return self.get( adAccountId );
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount with id " + str(adAccountId) + " does not exist.")
		except PerformanceMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : PerformanceMetric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
