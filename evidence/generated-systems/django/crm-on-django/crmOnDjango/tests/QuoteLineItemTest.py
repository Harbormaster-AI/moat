import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.QuoteLineItem import QuoteLineItem
from crmOnDjango.delegates.QuoteLineItemDelegate import QuoteLineItemDelegate

 #======================================================================
# 
# Encapsulates data for model QuoteLineItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QuoteLineItemTest Declaration
#======================================================================
class QuoteLineItemTest (TestCase) :
	def test_crud(self) :
		quoteLineItem = QuoteLineItem()
		quoteLineItem.quantity = "default quantity field value"
		quoteLineItem.unitPrice = "default unitPrice field value"
		quoteLineItem.discountAmount = "default discountAmount field value"
		quoteLineItem.taxAmount = "default taxAmount field value"
		quoteLineItem.totalAmount = "default totalAmount field value"
		
		delegate = QuoteLineItemDelegate()
		responseObj = delegate.create(quoteLineItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


