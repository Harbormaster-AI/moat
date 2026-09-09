import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.PurchaseAgreement import PurchaseAgreement
from aerospaceOnDjango.delegates.PurchaseAgreementDelegate import PurchaseAgreementDelegate

 #======================================================================
# 
# Encapsulates data for model PurchaseAgreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PurchaseAgreementTest Declaration
#======================================================================
class PurchaseAgreementTest (TestCase) :
	def test_crud(self) :
		purchaseAgreement = PurchaseAgreement()
		purchaseAgreement.agreementNumber = "default agreementNumber field value"
		purchaseAgreement.effectiveDate = datetime.datetime.now()
		
		delegate = PurchaseAgreementDelegate()
		responseObj = delegate.create(purchaseAgreement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


