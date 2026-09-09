import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.ProductVariant import ProductVariant
from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

 #======================================================================
# 
# Encapsulates data for model ProductVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductVariantTest Declaration
#======================================================================
class ProductVariantTest (TestCase) :
	def test_crud(self) :
		productVariant = ProductVariant()
		productVariant.sku = "default sku field value"
		productVariant.barcode = "default barcode field value"
		productVariant.title = "default title field value"
		productVariant.weight = "default weight field value"
		productVariant.requiresShipping = False
		productVariant.weightUnit = "default weightUnit field value"
		
		delegate = ProductVariantDelegate()
		responseObj = delegate.create(productVariant)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


