import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Supplier import Supplier
from ecommerceOnDjango.delegates.SupplierDelegate import SupplierDelegate

 #======================================================================
# 
# Encapsulates data for model Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierTest Declaration
#======================================================================
class SupplierTest (TestCase) :
	def test_crud(self) :
		supplier = Supplier()
		supplier.name = "default name field value"
		supplier.contactEmail = "default contactEmail field value"
		supplier.website = "default website field value"
		supplier.status = "default status field value"
		
		delegate = SupplierDelegate()
		responseObj = delegate.create(supplier)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


