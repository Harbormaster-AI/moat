import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.InsuranceProduct import InsuranceProduct
from insuranceOnDjango.delegates.InsuranceProductDelegate import InsuranceProductDelegate

 #======================================================================
# 
# Encapsulates data for model InsuranceProduct
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuranceProductTest Declaration
#======================================================================
class InsuranceProductTest (TestCase) :
	def test_crud(self) :
		insuranceProduct = InsuranceProduct()
		insuranceProduct.name = "default name field value"
		insuranceProduct.productCode = "default productCode field value"
		insuranceProduct.lineOfBusiness = "default lineOfBusiness field value"
		
		delegate = InsuranceProductDelegate()
		responseObj = delegate.create(insuranceProduct)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


