import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.ReturnRequest import ReturnRequest
from ecommerceOnDjango.delegates.ReturnRequestDelegate import ReturnRequestDelegate

 #======================================================================
# 
# Encapsulates data for model ReturnRequest
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReturnRequestTest Declaration
#======================================================================
class ReturnRequestTest (TestCase) :
	def test_crud(self) :
		returnRequest = ReturnRequest()
		returnRequest.returnNumber = "default returnNumber field value"
		returnRequest.createdAt = datetime.datetime.now()
		returnRequest.refundAmount = "default refundAmount field value"
		returnRequest.status = "default status field value"
		
		delegate = ReturnRequestDelegate()
		responseObj = delegate.create(returnRequest)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


