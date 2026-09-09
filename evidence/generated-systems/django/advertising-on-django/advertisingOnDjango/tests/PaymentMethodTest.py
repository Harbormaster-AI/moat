import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.PaymentMethod import PaymentMethod
from advertisingOnDjango.delegates.PaymentMethodDelegate import PaymentMethodDelegate

 #======================================================================
# 
# Encapsulates data for model PaymentMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethodTest Declaration
#======================================================================
class PaymentMethodTest (TestCase) :
	def test_crud(self) :
		paymentMethod = PaymentMethod()
		paymentMethod.last4 = "default last4 field value"
		paymentMethod.cardholderName = "default cardholderName field value"
		paymentMethod.billingAddress = "default billingAddress field value"
		paymentMethod.methodType = "default methodType field value"
		
		delegate = PaymentMethodDelegate()
		responseObj = delegate.create(paymentMethod)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


