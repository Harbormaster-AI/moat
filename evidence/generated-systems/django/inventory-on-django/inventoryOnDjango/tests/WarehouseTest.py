import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.Warehouse import Warehouse
from inventoryOnDjango.delegates.WarehouseDelegate import WarehouseDelegate

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
		warehouse.code = "default code field value"
		warehouse.address = "default address field value"
		warehouse.timeZone = "default timeZone field value"
		warehouse.allowsOverAllocation = False
		
		delegate = WarehouseDelegate()
		responseObj = delegate.create(warehouse)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


