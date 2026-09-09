import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

 #======================================================================
# 
# Encapsulates data for model Merchant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MerchantTest Declaration
#======================================================================
class MerchantTest (TestCase) :
	def test_crud(self) :
		merchant = Merchant()
		merchant.name = "default name field value"
		merchant.legalName = "default legalName field value"
		merchant.website = "default website field value"
		merchant.defaultCurrency = "default defaultCurrency field value"
		merchant.defaultLocale = "default defaultLocale field value"
		merchant.supportEmail = "default supportEmail field value"
		
		delegate = MerchantDelegate()
		responseObj = delegate.create(merchant)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


