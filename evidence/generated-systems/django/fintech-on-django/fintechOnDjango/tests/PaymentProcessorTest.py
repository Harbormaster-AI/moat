import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.PaymentProcessor import PaymentProcessor
from fintechOnDjango.delegates.PaymentProcessorDelegate import PaymentProcessorDelegate

 #======================================================================
# 
# Encapsulates data for model PaymentProcessor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProcessorTest Declaration
#======================================================================
class PaymentProcessorTest (TestCase) :
	def test_crud(self) :
		paymentProcessor = PaymentProcessor()
		paymentProcessor.name = "default name field value"
		paymentProcessor.processorCode = "default processorCode field value"
		paymentProcessor.networkSupport = "default networkSupport field value"
		
		delegate = PaymentProcessorDelegate()
		responseObj = delegate.create(paymentProcessor)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


