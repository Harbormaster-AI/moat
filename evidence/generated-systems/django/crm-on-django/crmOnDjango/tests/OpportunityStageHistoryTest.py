import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.OpportunityStageHistory import OpportunityStageHistory
from crmOnDjango.delegates.OpportunityStageHistoryDelegate import OpportunityStageHistoryDelegate

 #======================================================================
# 
# Encapsulates data for model OpportunityStageHistory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityStageHistoryTest Declaration
#======================================================================
class OpportunityStageHistoryTest (TestCase) :
	def test_crud(self) :
		opportunityStageHistory = OpportunityStageHistory()
		opportunityStageHistory.changedAt = "default changedAt field value"
		opportunityStageHistory.comment = "default comment field value"
		opportunityStageHistory.fromStage = "default fromStage field value"
		opportunityStageHistory.toStage = "default toStage field value"
		
		delegate = OpportunityStageHistoryDelegate()
		responseObj = delegate.create(opportunityStageHistory)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


