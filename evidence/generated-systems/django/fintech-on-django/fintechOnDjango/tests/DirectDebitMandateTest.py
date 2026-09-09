import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.DirectDebitMandate import DirectDebitMandate
from fintechOnDjango.delegates.DirectDebitMandateDelegate import DirectDebitMandateDelegate

 #======================================================================
# 
# Encapsulates data for model DirectDebitMandate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DirectDebitMandateTest Declaration
#======================================================================
class DirectDebitMandateTest (TestCase) :
	def test_crud(self) :
		directDebitMandate = DirectDebitMandate()
		directDebitMandate.mandateId = "default mandateId field value"
		directDebitMandate.signedAt = "default signedAt field value"
		directDebitMandate.scheme = "default scheme field value"
		directDebitMandate.status = "default status field value"
		
		delegate = DirectDebitMandateDelegate()
		responseObj = delegate.create(directDebitMandate)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


