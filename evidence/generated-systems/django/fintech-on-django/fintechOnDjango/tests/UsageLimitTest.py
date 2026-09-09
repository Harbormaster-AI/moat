import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.UsageLimit import UsageLimit
from fintechOnDjango.delegates.UsageLimitDelegate import UsageLimitDelegate

 #======================================================================
# 
# Encapsulates data for model UsageLimit
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UsageLimitTest Declaration
#======================================================================
class UsageLimitTest (TestCase) :
	def test_crud(self) :
		usageLimit = UsageLimit()
		usageLimit.name = "default name field value"
		usageLimit.amount = "default amount field value"
		usageLimit.count = 22
		usageLimit.scope = "default scope field value"
		usageLimit.period = "default period field value"
		
		delegate = UsageLimitDelegate()
		responseObj = delegate.create(usageLimit)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


