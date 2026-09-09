import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.CartItem import CartItem
from ecommerceOnDjango.delegates.CartItemDelegate import CartItemDelegate

 #======================================================================
# 
# Encapsulates data for model CartItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartItemTest Declaration
#======================================================================
class CartItemTest (TestCase) :
	def test_crud(self) :
		cartItem = CartItem()
		cartItem.quantity = 22
		cartItem.unitPrice = "default unitPrice field value"
		cartItem.totalPrice = "default totalPrice field value"
		
		delegate = CartItemDelegate()
		responseObj = delegate.create(cartItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


