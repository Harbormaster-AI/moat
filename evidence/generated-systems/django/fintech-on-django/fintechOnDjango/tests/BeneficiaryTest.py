import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.Beneficiary import Beneficiary
from fintechOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

 #======================================================================
# 
# Encapsulates data for model Beneficiary
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BeneficiaryTest Declaration
#======================================================================
class BeneficiaryTest (TestCase) :
	def test_crud(self) :
		beneficiary = Beneficiary()
		beneficiary.name = "default name field value"
		beneficiary.accountIdentifier = "default accountIdentifier field value"
		beneficiary.iban = "default iban field value"
		beneficiary.bic = "default bic field value"
		beneficiary.address = "default address field value"
		
		delegate = BeneficiaryDelegate()
		responseObj = delegate.create(beneficiary)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


