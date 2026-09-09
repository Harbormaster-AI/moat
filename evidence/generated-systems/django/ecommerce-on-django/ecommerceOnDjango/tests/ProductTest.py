import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

 #======================================================================
# 
# Encapsulates data for model Product
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductTest Declaration
#======================================================================
class ProductTest (TestCase) :
	def test_crud(self) :
		product = Product()
		product.name = "default name field value"
		product.slug = "default slug field value"
		product.asActive = False
		product.productType = "default productType field value"
		product.defaultTaxClass = "default defaultTaxClass field value"
		
		delegate = ProductDelegate()
		responseObj = delegate.create(product)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


