import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Order import Order
from crmOnDjango.delegates.OrderDelegate import OrderDelegate

 #======================================================================
# 
# Encapsulates data for model Order
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderTest Declaration
#======================================================================
class OrderTest (TestCase) :
	def test_crud(self) :
		order = Order()
		order.orderNumber = "default orderNumber field value"
		order.orderDate = datetime.datetime.now()
		order.totalAmount = "default totalAmount field value"
		order.taxAmount = "default taxAmount field value"
		order.shippingAmount = "default shippingAmount field value"
		order.status = "default status field value"
		
		delegate = OrderDelegate()
		responseObj = delegate.create(order)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


