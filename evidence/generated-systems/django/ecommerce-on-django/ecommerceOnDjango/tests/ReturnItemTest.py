import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.ReturnItem import ReturnItem
from ecommerceOnDjango.delegates.ReturnItemDelegate import ReturnItemDelegate

 #======================================================================
# 
# Encapsulates data for model ReturnItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnItemTest Declaration
#======================================================================
class ReturnItemTest (TestCase) :
	def test_crud(self) :
		returnItem = ReturnItem()
		returnItem.quantity = 22
		returnItem.reason = "default reason field value"
		returnItem.condition = "default condition field value"
		
		delegate = ReturnItemDelegate()
		responseObj = delegate.create(returnItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


