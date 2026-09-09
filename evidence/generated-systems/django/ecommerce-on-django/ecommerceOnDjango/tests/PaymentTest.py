import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Payment import Payment
from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

 #======================================================================
# 
# Encapsulates data for model Payment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentTest Declaration
#======================================================================
class PaymentTest (TestCase) :
	def test_crud(self) :
		payment = Payment()
		payment.paymentNumber = "default paymentNumber field value"
		payment.amount = "default amount field value"
		payment.transactionId = "default transactionId field value"
		payment.authorizedAt = datetime.datetime.now()
		payment.capturedAt = datetime.datetime.now()
		payment.status = "default status field value"
		payment.paymentMethod = "default paymentMethod field value"
		
		delegate = PaymentDelegate()
		responseObj = delegate.create(payment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


