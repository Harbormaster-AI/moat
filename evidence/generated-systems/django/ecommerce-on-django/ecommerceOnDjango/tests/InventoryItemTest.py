import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.InventoryItem import InventoryItem
from ecommerceOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

 #======================================================================
# 
# Encapsulates data for model InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItemTest Declaration
#======================================================================
class InventoryItemTest (TestCase) :
	def test_crud(self) :
		inventoryItem = InventoryItem()
		inventoryItem.quantityOnHand = 22
		inventoryItem.quantityReserved = 22
		inventoryItem.safetyStock = 22
		inventoryItem.status = "default status field value"
		
		delegate = InventoryItemDelegate()
		responseObj = delegate.create(inventoryItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


