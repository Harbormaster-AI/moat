import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.ExchangeRate import ExchangeRate
from fintechOnDjango.delegates.ExchangeRateDelegate import ExchangeRateDelegate

 #======================================================================
# 
# Encapsulates data for model ExchangeRate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExchangeRateTest Declaration
#======================================================================
class ExchangeRateTest (TestCase) :
	def test_crud(self) :
		exchangeRate = ExchangeRate()
		exchangeRate.baseCurrency = "default baseCurrency field value"
		exchangeRate.quoteCurrency = "default quoteCurrency field value"
		exchangeRate.rate = "default rate field value"
		exchangeRate.asOf = "default asOf field value"
		exchangeRate.source = "default source field value"
		
		delegate = ExchangeRateDelegate()
		responseObj = delegate.create(exchangeRate)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


