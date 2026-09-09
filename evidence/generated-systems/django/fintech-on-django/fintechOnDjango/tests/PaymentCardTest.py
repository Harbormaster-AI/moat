import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.PaymentCard import PaymentCard
from fintechOnDjango.delegates.PaymentCardDelegate import PaymentCardDelegate

 #======================================================================
# 
# Encapsulates data for model PaymentCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentCardTest Declaration
#======================================================================
class PaymentCardTest (TestCase) :
	def test_crud(self) :
		paymentCard = PaymentCard()
		paymentCard.cardToken = "default cardToken field value"
		paymentCard.maskedPan = "default maskedPan field value"
		paymentCard.expiryMonth = 22
		paymentCard.expiryYear = 22
		paymentCard.cardholderName = "default cardholderName field value"
		paymentCard.scheme = "default scheme field value"
		paymentCard.status = "default status field value"
		
		delegate = PaymentCardDelegate()
		responseObj = delegate.create(paymentCard)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


