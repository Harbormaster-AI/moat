import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.FXQuote import FXQuote
from fintechOnDjango.delegates.FXQuoteDelegate import FXQuoteDelegate

 #======================================================================
# 
# Encapsulates data for model FXQuote
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FXQuoteTest Declaration
#======================================================================
class FXQuoteTest (TestCase) :
	def test_crud(self) :
		fXQuote = FXQuote()
		fXQuote.baseCurrency = "default baseCurrency field value"
		fXQuote.quoteCurrency = "default quoteCurrency field value"
		fXQuote.rate = "default rate field value"
		fXQuote.quotedAt = "default quotedAt field value"
		fXQuote.expiresAt = "default expiresAt field value"
		fXQuote.priceType = "default priceType field value"
		
		delegate = FXQuoteDelegate()
		responseObj = delegate.create(fXQuote)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


