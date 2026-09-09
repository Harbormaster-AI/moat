import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Deal import Deal
from advertisingOnDjango.delegates.DealDelegate import DealDelegate

 #======================================================================
# 
# Encapsulates data for model Deal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DealTest Declaration
#======================================================================
class DealTest (TestCase) :
	def test_crud(self) :
		deal = Deal()
		deal.floorPrice = "default floorPrice field value"
		deal.dealType = "default dealType field value"
		
		delegate = DealDelegate()
		responseObj = delegate.create(deal)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


