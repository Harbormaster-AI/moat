import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Pharmacy import Pharmacy
from healthcareOnDjango.delegates.PharmacyDelegate import PharmacyDelegate

 #======================================================================
# 
# Encapsulates data for model Pharmacy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PharmacyTest Declaration
#======================================================================
class PharmacyTest (TestCase) :
	def test_crud(self) :
		pharmacy = Pharmacy()
		pharmacy.name = "default name field value"
		
		delegate = PharmacyDelegate()
		responseObj = delegate.create(pharmacy)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


