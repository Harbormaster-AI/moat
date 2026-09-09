import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

 #======================================================================
# 
# Encapsulates data for model Opportunity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityTest Declaration
#======================================================================
class OpportunityTest (TestCase) :
	def test_crud(self) :
		opportunity = Opportunity()
		opportunity.name = "default name field value"
		opportunity.amount = "default amount field value"
		opportunity.closeDate = datetime.datetime.now()
		opportunity.probability = "default probability field value"
		opportunity.description = "default description field value"
		opportunity.stage = "default stage field value"
		opportunity.type = "default type field value"
		opportunity.forecastCategory = "default forecastCategory field value"
		
		delegate = OpportunityDelegate()
		responseObj = delegate.create(opportunity)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


