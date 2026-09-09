import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.PriceBook import PriceBook
from crmOnDjango.delegates.PriceBookDelegate import PriceBookDelegate

 #======================================================================
# 
# Encapsulates data for model PriceBook
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBookTest Declaration
#======================================================================
class PriceBookTest (TestCase) :
	def test_crud(self) :
		priceBook = PriceBook()
		priceBook.name = "default name field value"
		priceBook.asActive = False
		priceBook.description = "default description field value"
		
		delegate = PriceBookDelegate()
		responseObj = delegate.create(priceBook)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


