import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Quote import Quote
from insuranceOnDjango.delegates.QuoteDelegate import QuoteDelegate

 #======================================================================
# 
# Encapsulates data for model Quote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteTest Declaration
#======================================================================
class QuoteTest (TestCase) :
	def test_crud(self) :
		quote = Quote()
		quote.quoteNumber = "default quoteNumber field value"
		quote.totalPremium = "default totalPremium field value"
		quote.ratingDate = datetime.datetime.now()
		quote.asBound = False
		
		delegate = QuoteDelegate()
		responseObj = delegate.create(quote)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


