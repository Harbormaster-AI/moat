import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

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
		merchant.mcc = "default mcc field value"
		merchant.url = "default url field value"
		merchant.country = "default country field value"
		merchant.settlementCurrency = "default settlementCurrency field value"
		
		delegate = MerchantDelegate()
		responseObj = delegate.create(merchant)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


