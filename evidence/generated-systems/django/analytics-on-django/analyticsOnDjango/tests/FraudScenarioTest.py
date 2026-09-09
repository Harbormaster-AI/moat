import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.FraudScenario import FraudScenario
from analyticsOnDjango.delegates.FraudScenarioDelegate import FraudScenarioDelegate

 #======================================================================
# 
# Encapsulates data for model FraudScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudScenarioTest Declaration
#======================================================================
class FraudScenarioTest (TestCase) :
	def test_crud(self) :
		fraudScenario = FraudScenario()
		fraudScenario.name = "default name field value"
		fraudScenario.riskAppetite = "default riskAppetite field value"
		fraudScenario.detectionType = "default detectionType field value"
		
		delegate = FraudScenarioDelegate()
		responseObj = delegate.create(fraudScenario)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


