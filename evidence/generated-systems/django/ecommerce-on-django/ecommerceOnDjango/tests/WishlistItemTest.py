import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.WishlistItem import WishlistItem
from ecommerceOnDjango.delegates.WishlistItemDelegate import WishlistItemDelegate

 #======================================================================
# 
# Encapsulates data for model WishlistItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WishlistItemTest Declaration
#======================================================================
class WishlistItemTest (TestCase) :
	def test_crud(self) :
		wishlistItem = WishlistItem()
		wishlistItem.addedDate = datetime.datetime.now()
		
		delegate = WishlistItemDelegate()
		responseObj = delegate.create(wishlistItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


