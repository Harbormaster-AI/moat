import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Wishlist import Wishlist
from ecommerceOnDjango.delegates.WishlistDelegate import WishlistDelegate

 #======================================================================
# 
# Encapsulates data for model Wishlist
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WishlistTest Declaration
#======================================================================
class WishlistTest (TestCase) :
	def test_crud(self) :
		wishlist = Wishlist()
		wishlist.name = "default name field value"
		wishlist.asPublic = False
		wishlist.createdAt = datetime.datetime.now()
		
		delegate = WishlistDelegate()
		responseObj = delegate.create(wishlist)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


