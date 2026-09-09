import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Rate import Rate
from advertisingOnDjango.delegates.RateDelegate import RateDelegate

 #======================================================================
# 
# Encapsulates data for model Rate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RateTest Declaration
#======================================================================
class RateTest (TestCase) :
	def test_crud(self) :
		rate = Rate()
		rate.unitPrice = "default unitPrice field value"
		rate.adFormat = "default adFormat field value"
		rate.pricingModel = "default pricingModel field value"
		
		delegate = RateDelegate()
		responseObj = delegate.create(rate)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


