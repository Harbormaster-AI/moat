import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Brand import Brand
from ecommerceOnDjango.delegates.BrandDelegate import BrandDelegate

 #======================================================================
# 
# Encapsulates data for model Brand
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandTest Declaration
#======================================================================
class BrandTest (TestCase) :
	def test_crud(self) :
		brand = Brand()
		brand.name = "default name field value"
		brand.description = "default description field value"
		brand.website = "default website field value"
		
		delegate = BrandDelegate()
		responseObj = delegate.create(brand)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


