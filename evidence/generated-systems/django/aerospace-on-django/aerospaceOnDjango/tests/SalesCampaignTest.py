import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.SalesCampaign import SalesCampaign
from aerospaceOnDjango.delegates.SalesCampaignDelegate import SalesCampaignDelegate

 #======================================================================
# 
# Encapsulates data for model SalesCampaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesCampaignTest Declaration
#======================================================================
class SalesCampaignTest (TestCase) :
	def test_crud(self) :
		salesCampaign = SalesCampaign()
		salesCampaign.campaignCode = "default campaignCode field value"
		salesCampaign.status = "default status field value"
		
		delegate = SalesCampaignDelegate()
		responseObj = delegate.create(salesCampaign)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


