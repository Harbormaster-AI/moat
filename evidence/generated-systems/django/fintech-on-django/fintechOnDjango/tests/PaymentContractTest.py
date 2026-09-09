import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.PaymentContract import PaymentContract
from fintechOnDjango.delegates.PaymentContractDelegate import PaymentContractDelegate

 #======================================================================
# 
# Encapsulates data for model PaymentContract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentContractTest Declaration
#======================================================================
class PaymentContractTest (TestCase) :
	def test_crud(self) :
		paymentContract = PaymentContract()
		paymentContract.contractNumber = "default contractNumber field value"
		paymentContract.pricingPlanCode = "default pricingPlanCode field value"
		paymentContract.status = "default status field value"
		
		delegate = PaymentContractDelegate()
		responseObj = delegate.create(paymentContract)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


