import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.RateCard import RateCard
from advertisingOnDjango.delegates.RateCardDelegate import RateCardDelegate

 #======================================================================
# 
# Encapsulates data for model RateCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RateCardTest Declaration
#======================================================================
class RateCardTest (TestCase) :
	def test_crud(self) :
		rateCard = RateCard()
		rateCard.name = "default name field value"
		rateCard.effectiveDate = datetime.datetime.now()
		rateCard.currency = "default currency field value"
		
		delegate = RateCardDelegate()
		responseObj = delegate.create(rateCard)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


