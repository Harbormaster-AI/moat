import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.MedicalSupplier import MedicalSupplier
from healthcareOnDjango.delegates.MedicalSupplierDelegate import MedicalSupplierDelegate

 #======================================================================
# 
# Encapsulates data for model MedicalSupplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalSupplierTest Declaration
#======================================================================
class MedicalSupplierTest (TestCase) :
	def test_crud(self) :
		medicalSupplier = MedicalSupplier()
		medicalSupplier.name = "default name field value"
		medicalSupplier.website = "default website field value"
		medicalSupplier.supplierTier = "default supplierTier field value"
		
		delegate = MedicalSupplierDelegate()
		responseObj = delegate.create(medicalSupplier)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


