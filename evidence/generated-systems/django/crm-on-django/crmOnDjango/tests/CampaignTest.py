import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

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
		campaign.startDate = datetime.datetime.now()
		campaign.endDate = datetime.datetime.now()
		campaign.budget = "default budget field value"
		campaign.actualCost = "default actualCost field value"
		campaign.expectedRevenue = "default expectedRevenue field value"
		campaign.status = "default status field value"
		campaign.type = "default type field value"
		
		delegate = CampaignDelegate()
		responseObj = delegate.create(campaign)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


