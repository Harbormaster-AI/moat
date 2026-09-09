import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.InsurancePayer import InsurancePayer
from healthcareOnDjango.delegates.InsurancePayerDelegate import InsurancePayerDelegate

 #======================================================================
# 
# Encapsulates data for model InsurancePayer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePayerTest Declaration
#======================================================================
class InsurancePayerTest (TestCase) :
	def test_crud(self) :
		insurancePayer = InsurancePayer()
		insurancePayer.name = "default name field value"
		insurancePayer.website = "default website field value"
		insurancePayer.payerType = "default payerType field value"
		
		delegate = InsurancePayerDelegate()
		responseObj = delegate.create(insurancePayer)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


