import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Cart import Cart
from ecommerceOnDjango.delegates.CartDelegate import CartDelegate

 #======================================================================
# 
# Encapsulates data for model Cart
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartTest Declaration
#======================================================================
class CartTest (TestCase) :
	def test_crud(self) :
		cart = Cart()
		cart.cartNumber = "default cartNumber field value"
		cart.createdAt = datetime.datetime.now()
		cart.currency = "default currency field value"
		cart.shippingAddress = "default shippingAddress field value"
		cart.billingAddress = "default billingAddress field value"
		cart.status = "default status field value"
		
		delegate = CartDelegate()
		responseObj = delegate.create(cart)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


