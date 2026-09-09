import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Beneficiary import Beneficiary
from insuranceOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

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
		beneficiary.share = "default share field value"
		beneficiary.relationship = "default relationship field value"
		
		delegate = BeneficiaryDelegate()
		responseObj = delegate.create(beneficiary)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


