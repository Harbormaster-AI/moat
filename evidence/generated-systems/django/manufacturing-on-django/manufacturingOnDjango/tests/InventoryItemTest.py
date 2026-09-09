import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.InventoryItem import InventoryItem
from manufacturingOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

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
		inventoryItem.quantityOnHand = "default quantityOnHand field value"
		inventoryItem.quantityReserved = "default quantityReserved field value"
		inventoryItem.lotNumber = "default lotNumber field value"
		inventoryItem.serialNumber = "default serialNumber field value"
		
		delegate = InventoryItemDelegate()
		responseObj = delegate.create(inventoryItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


