import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.PriceBookEntry import PriceBookEntry
from crmOnDjango.delegates.PriceBookEntryDelegate import PriceBookEntryDelegate

 #======================================================================
# 
# Encapsulates data for model PriceBookEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PriceBookEntryTest Declaration
#======================================================================
class PriceBookEntryTest (TestCase) :
	def test_crud(self) :
		priceBookEntry = PriceBookEntry()
		priceBookEntry.unitPrice = "default unitPrice field value"
		priceBookEntry.effectiveDate = datetime.datetime.now()
		priceBookEntry.expirationDate = datetime.datetime.now()
		priceBookEntry.asActive = False
		
		delegate = PriceBookEntryDelegate()
		responseObj = delegate.create(priceBookEntry)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


