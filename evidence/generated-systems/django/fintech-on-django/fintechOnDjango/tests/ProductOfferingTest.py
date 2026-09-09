import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.ProductOffering import ProductOffering
from fintechOnDjango.delegates.ProductOfferingDelegate import ProductOfferingDelegate

 #======================================================================
# 
# Encapsulates data for model ProductOffering
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductOfferingTest Declaration
#======================================================================
class ProductOfferingTest (TestCase) :
	def test_crud(self) :
		productOffering = ProductOffering()
		productOffering.name = "default name field value"
		productOffering.productCode = "default productCode field value"
		productOffering.category = "default category field value"
		
		delegate = ProductOfferingDelegate()
		responseObj = delegate.create(productOffering)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


