import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Supplier import Supplier
from manufacturingOnDjango.delegates.SupplierDelegate import SupplierDelegate

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
		supplier.supplierCode = "default supplierCode field value"
		supplier.address = "default address field value"
		supplier.supplierTier = "default supplierTier field value"
		supplier.paymentTerms = "default paymentTerms field value"
		
		delegate = SupplierDelegate()
		responseObj = delegate.create(supplier)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


