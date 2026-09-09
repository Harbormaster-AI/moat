import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Quote import Quote
from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

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
		quote.validityStart = datetime.datetime.now()
		quote.validityEnd = datetime.datetime.now()
		quote.totalAmount = "default totalAmount field value"
		quote.discountPercent = "default discountPercent field value"
		quote.taxAmount = "default taxAmount field value"
		quote.shippingAmount = "default shippingAmount field value"
		quote.status = "default status field value"
		
		delegate = QuoteDelegate()
		responseObj = delegate.create(quote)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


