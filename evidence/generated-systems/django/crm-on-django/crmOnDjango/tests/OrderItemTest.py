import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.OrderItem import OrderItem
from crmOnDjango.delegates.OrderItemDelegate import OrderItemDelegate

 #======================================================================
# 
# Encapsulates data for model OrderItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderItemTest Declaration
#======================================================================
class OrderItemTest (TestCase) :
	def test_crud(self) :
		orderItem = OrderItem()
		orderItem.quantity = "default quantity field value"
		orderItem.unitPrice = "default unitPrice field value"
		orderItem.discountAmount = "default discountAmount field value"
		orderItem.taxAmount = "default taxAmount field value"
		orderItem.totalAmount = "default totalAmount field value"
		
		delegate = OrderItemDelegate()
		responseObj = delegate.create(orderItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


