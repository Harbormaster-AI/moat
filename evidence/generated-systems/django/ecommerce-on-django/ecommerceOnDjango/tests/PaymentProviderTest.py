import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.PaymentProvider import PaymentProvider
from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

 #======================================================================
# 
# Encapsulates data for model PaymentProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProviderTest Declaration
#======================================================================
class PaymentProviderTest (TestCase) :
	def test_crud(self) :
		paymentProvider = PaymentProvider()
		paymentProvider.name = "default name field value"
		paymentProvider.enabled = False
		paymentProvider.merchantAccountId = "default merchantAccountId field value"
		paymentProvider.providerType = "default providerType field value"
		
		delegate = PaymentProviderDelegate()
		responseObj = delegate.create(paymentProvider)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


