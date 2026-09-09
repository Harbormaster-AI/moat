import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

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
		order.placedDate = datetime.datetime.now()
		order.subtotal = "default subtotal field value"
		order.discountTotal = "default discountTotal field value"
		order.shippingTotal = "default shippingTotal field value"
		order.taxTotal = "default taxTotal field value"
		order.grandTotal = "default grandTotal field value"
		order.shippingAddress = "default shippingAddress field value"
		order.billingAddress = "default billingAddress field value"
		order.status = "default status field value"
		
		delegate = OrderDelegate()
		responseObj = delegate.create(order)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


