import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.OrderLine import OrderLine
from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

 #======================================================================
# 
# Encapsulates data for model OrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderLineTest Declaration
#======================================================================
class OrderLineTest (TestCase) :
	def test_crud(self) :
		orderLine = OrderLine()
		orderLine.quantity = 22
		orderLine.unitPrice = "default unitPrice field value"
		orderLine.totalPrice = "default totalPrice field value"
		orderLine.taxRate = "default taxRate field value"
		orderLine.lineStatus = "default lineStatus field value"
		
		delegate = OrderLineDelegate()
		responseObj = delegate.create(orderLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


