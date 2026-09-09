import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.CampaignMember import CampaignMember
from crmOnDjango.delegates.CampaignMemberDelegate import CampaignMemberDelegate

 #======================================================================
# 
# Encapsulates data for model CampaignMember
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignMemberTest Declaration
#======================================================================
class CampaignMemberTest (TestCase) :
	def test_crud(self) :
		campaignMember = CampaignMember()
		campaignMember.responded = False
		campaignMember.status = "default status field value"
		campaignMember.memberType = "default memberType field value"
		
		delegate = CampaignMemberDelegate()
		responseObj = delegate.create(campaignMember)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


