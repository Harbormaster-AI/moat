import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

 #======================================================================
# 
# Encapsulates data for model Campaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignTest Declaration
#======================================================================
class CampaignTest (TestCase) :
	def test_crud(self) :
		campaign = Campaign()
		campaign.name = "default name field value"
		campaign.totalBudget = "default totalBudget field value"
		campaign.flight = "default flight field value"
		campaign.objective = "default objective field value"
		campaign.status = "default status field value"
		
		delegate = CampaignDelegate()
		responseObj = delegate.create(campaign)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


