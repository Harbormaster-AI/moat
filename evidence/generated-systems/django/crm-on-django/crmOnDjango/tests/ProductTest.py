import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Product import Product
from crmOnDjango.delegates.ProductDelegate import ProductDelegate

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
		product.sku = "default sku field value"
		product.name = "default name field value"
		product.asActive = False
		product.standardPrice = "default standardPrice field value"
		product.description = "default description field value"
		product.productType = "default productType field value"
		product.uom = "default uom field value"
		
		delegate = ProductDelegate()
		responseObj = delegate.create(product)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


