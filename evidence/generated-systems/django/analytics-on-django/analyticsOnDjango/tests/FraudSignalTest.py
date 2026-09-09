import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.FraudSignal import FraudSignal
from analyticsOnDjango.delegates.FraudSignalDelegate import FraudSignalDelegate

 #======================================================================
# 
# Encapsulates data for model FraudSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudSignalTest Declaration
#======================================================================
class FraudSignalTest (TestCase) :
	def test_crud(self) :
		fraudSignal = FraudSignal()
		fraudSignal.name = "default name field value"
		fraudSignal.ruleLogic = "default ruleLogic field value"
		fraudSignal.signalType = "default signalType field value"
		
		delegate = FraudSignalDelegate()
		responseObj = delegate.create(fraudSignal)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


