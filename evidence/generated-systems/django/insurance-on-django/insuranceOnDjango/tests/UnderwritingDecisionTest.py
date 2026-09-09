import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.UnderwritingDecision import UnderwritingDecision
from insuranceOnDjango.delegates.UnderwritingDecisionDelegate import UnderwritingDecisionDelegate

 #======================================================================
# 
# Encapsulates data for model UnderwritingDecision
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwritingDecisionTest Declaration
#======================================================================
class UnderwritingDecisionTest (TestCase) :
	def test_crud(self) :
		underwritingDecision = UnderwritingDecision()
		underwritingDecision.notes = "default notes field value"
		underwritingDecision.decisionDate = datetime.datetime.now()
		underwritingDecision.decision = "default decision field value"
		
		delegate = UnderwritingDecisionDelegate()
		responseObj = delegate.create(underwritingDecision)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


