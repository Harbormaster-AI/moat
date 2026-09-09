import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.Warehouse import Warehouse
from aerospaceOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

 #======================================================================
# 
# Encapsulates data for model Warehouse
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WarehouseTest Declaration
#======================================================================
class WarehouseTest (TestCase) :
	def test_crud(self) :
		warehouse = Warehouse()
		warehouse.name = "default name field value"
		
		delegate = WarehouseDelegate()
		responseObj = delegate.create(warehouse)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


