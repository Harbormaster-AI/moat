import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

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
		supplier.supplierType = "default supplierType field value"
		supplier.approvalStatus = "default approvalStatus field value"
		
		delegate = SupplierDelegate()
		responseObj = delegate.create(supplier)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


