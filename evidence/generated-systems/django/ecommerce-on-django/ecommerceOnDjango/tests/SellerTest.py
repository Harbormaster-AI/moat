import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Seller import Seller
from ecommerceOnDjango.delegates.SellerDelegate import SellerDelegate

 #======================================================================
# 
# Encapsulates data for model Seller
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SellerTest Declaration
#======================================================================
class SellerTest (TestCase) :
	def test_crud(self) :
		seller = Seller()
		seller.name = "default name field value"
		seller.sellerCode = "default sellerCode field value"
		seller.contactEmail = "default contactEmail field value"
		seller.status = "default status field value"
		
		delegate = SellerDelegate()
		responseObj = delegate.create(seller)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


