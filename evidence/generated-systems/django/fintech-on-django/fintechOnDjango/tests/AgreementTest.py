import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Agreement import Agreement
from fintechOnDjango.delegates.AgreementDelegate import AgreementDelegate

 #======================================================================
# 
# Encapsulates data for model Agreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgreementTest Declaration
#======================================================================
class AgreementTest (TestCase) :
	def test_crud(self) :
		agreement = Agreement()
		agreement.agreementNumber = "default agreementNumber field value"
		agreement.effectiveDate = datetime.datetime.now()
		agreement.agreementType = "default agreementType field value"
		agreement.status = "default status field value"
		
		delegate = AgreementDelegate()
		responseObj = delegate.create(agreement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


