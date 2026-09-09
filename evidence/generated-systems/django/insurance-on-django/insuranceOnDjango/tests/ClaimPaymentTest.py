import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.ClaimPayment import ClaimPayment
from insuranceOnDjango.delegates.ClaimPaymentDelegate import ClaimPaymentDelegate

 #======================================================================
# 
# Encapsulates data for model ClaimPayment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimPaymentTest Declaration
#======================================================================
class ClaimPaymentTest (TestCase) :
	def test_crud(self) :
		claimPayment = ClaimPayment()
		claimPayment.paymentNumber = "default paymentNumber field value"
		claimPayment.amount = "default amount field value"
		claimPayment.paymentDate = datetime.datetime.now()
		claimPayment.payeeType = "default payeeType field value"
		claimPayment.method = "default method field value"
		claimPayment.status = "default status field value"
		
		delegate = ClaimPaymentDelegate()
		responseObj = delegate.create(claimPayment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


