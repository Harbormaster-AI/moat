from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Advertiser import Advertiser
from advertisingOnDjango.models.Agency import Agency
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.models.BillingProfile import BillingProfile
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.models.TrackingPixel import TrackingPixel
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Advertiser
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdvertiserDelegate Declaration
#======================================================================
class AdvertiserDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, advertiserId ):
		try:	
			advertiser = Advertiser.objects.filter(id=advertiserId)
			return advertiser.first();
		except Advertiser.DoesNotExist:
			raise ProcessingError("Advertiser with id " + str(advertiserId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, advertiser):
		for model in serializers.deserialize("json", advertiser):
			model.save()
			return model;

	def create(self, advertiser):
		advertiser.save()
		return advertiser;

	def saveFromJson(self, advertiser):
		for model in serializers.deserialize("json", advertiser):
			model.save()
			return advertiser;
	
	def save(self, advertiser):
		advertiser.save()
		return advertiser;
	
	def delete(self, advertiserId ):
		errMsg = "Failed to delete Advertiser from db using id " + str(advertiserId)
		
		try:
			advertiser = Advertiser.objects.get(id=advertiserId)
			advertiser.delete()
			return True
		except Advertiser.DoesNotExist:
			raise ProcessingError("Advertiser with id " + str(advertiserId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Advertiser.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Advertiser from db")
		except Exception:
			return None;
		
	def assignAgency( self, advertiserId, agencyId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AgencyDelegate import AgencyDelegate

		errMsg = "Failed to assign element " + str(agencyId) + " for Agency on Advertiser"

		try:
			# get the Advertiser from db
			advertiser = self.get( advertiserId ).first()	
			
			# get the Agency from db
			agency = AgencyDelegate().get(agencyId).first();
			
			# assign the Agency		
			advertiser.agency = agency
			
			#save it
			advertiser.save()

			# reload and return the appropriate version					
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAgency( self, advertiserId ):
		errMsg = "Failed to unassign element " + str(agencyId) + " for Agency on Advertiser"

		try:
			# get the Advertiser from db
			advertiser = self.get( advertiserId ).first()	
			
			# assign to None for unassignment
			advertiser.agency = None			

			#save it
			advertiser.save()

			# reload and return the appropriate version					
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Exception:
			return None;
		
	def addAdAccounts( self, advertiserId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to add elements " + str(adAccountsIds) + " for AdAccounts on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				advertiser.adAccounts.add(adAccount)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAdAccounts( self, advertiserId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to remove elements " + str(adAccountsIds) + " for AdAccounts on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				advertiser.adAccounts.remove(adAccount)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBillingProfiles( self, advertiserId, billingProfilesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.BillingProfileDelegate import BillingProfileDelegate

		errMsg = "Failed to add elements " + str(billingProfilesIds) + " for BillingProfiles on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = billingProfilesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BillingProfile		
				billingProfile = BillingProfileDelegate().get(id).first();	
				# add the BillingProfile
				advertiser.billingProfiles.add(billingProfile)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBillingProfiles( self, advertiserId, billingProfilesIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.BillingProfileDelegate import BillingProfileDelegate

		errMsg = "Failed to remove elements " + str(billingProfilesIds) + " for BillingProfiles on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = billingProfilesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BillingProfile		
				billingProfile = BillingProfileDelegate().get(id).first();	
				# add the BillingProfile
				advertiser.billingProfiles.remove(billingProfile)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except BillingProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingProfile does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, advertiserId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				advertiser.campaigns.add(campaign)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, advertiserId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				advertiser.campaigns.remove(campaign)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTrackingPixels( self, advertiserId, trackingPixelsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TrackingPixelDelegate import TrackingPixelDelegate

		errMsg = "Failed to add elements " + str(trackingPixelsIds) + " for TrackingPixels on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = trackingPixelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TrackingPixel		
				trackingPixel = TrackingPixelDelegate().get(id).first();	
				# add the TrackingPixel
				advertiser.trackingPixels.add(trackingPixel)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrackingPixels( self, advertiserId, trackingPixelsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TrackingPixelDelegate import TrackingPixelDelegate

		errMsg = "Failed to remove elements " + str(trackingPixelsIds) + " for TrackingPixels on Advertiser"

		try:
			# get the Advertiser
			advertiser = self.get( advertiserId ).first()
				
			# split on a comma with no spaces
			idList = trackingPixelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TrackingPixel		
				trackingPixel = TrackingPixelDelegate().get(id).first();	
				# add the TrackingPixel
				advertiser.trackingPixels.remove(trackingPixel)
				
			# save it		
			advertiser.save()
			
			# reload and return the appropriate version
			return self.get( advertiserId );
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser with id " + str(advertiserId) + " does not exist.")
		except TrackingPixel.DoesNotExist:
			raise ProcessingError(errMsg + " : TrackingPixel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
