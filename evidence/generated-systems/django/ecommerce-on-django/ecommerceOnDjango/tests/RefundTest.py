import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Refund import Refund
from ecommerceOnDjango.delegates.RefundDelegate import RefundDelegate

 #======================================================================
# 
# Encapsulates data for model Refund
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RefundTest Declaration
#======================================================================
class RefundTest (TestCase) :
	def test_crud(self) :
		refund = Refund()
		refund.refundNumber = "default refundNumber field value"
		refund.amount = "default amount field value"
		refund.reason = "default reason field value"
		refund.createdAt = datetime.datetime.now()
		refund.status = "default status field value"
		
		delegate = RefundDelegate()
		responseObj = delegate.create(refund)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


