import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.PaymentOrder import PaymentOrder
from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

 #======================================================================
# 
# Encapsulates data for model PaymentOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentOrderTest Declaration
#======================================================================
class PaymentOrderTest (TestCase) :
	def test_crud(self) :
		paymentOrder = PaymentOrder()
		paymentOrder.orderReference = "default orderReference field value"
		paymentOrder.requestedExecutionDate = datetime.datetime.now()
		paymentOrder.purpose = "default purpose field value"
		paymentOrder.paymentMethod = "default paymentMethod field value"
		paymentOrder.status = "default status field value"
		paymentOrder.priority = "default priority field value"
		
		delegate = PaymentOrderDelegate()
		responseObj = delegate.create(paymentOrder)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


