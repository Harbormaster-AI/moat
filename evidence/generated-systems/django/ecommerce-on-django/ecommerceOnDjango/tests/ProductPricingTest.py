import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.ProductPricing import ProductPricing
from ecommerceOnDjango.delegates.ProductPricingDelegate import ProductPricingDelegate

 #======================================================================
# 
# Encapsulates data for model ProductPricing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductPricingTest Declaration
#======================================================================
class ProductPricingTest (TestCase) :
	def test_crud(self) :
		productPricing = ProductPricing()
		productPricing.listPrice = "default listPrice field value"
		productPricing.salePrice = "default salePrice field value"
		productPricing.validFrom = datetime.datetime.now()
		productPricing.validTo = datetime.datetime.now()
		
		delegate = ProductPricingDelegate()
		responseObj = delegate.create(productPricing)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


