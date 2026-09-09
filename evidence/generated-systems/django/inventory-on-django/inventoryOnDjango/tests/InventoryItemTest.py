import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.InventoryItem import InventoryItem
from inventoryOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

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
		inventoryItem.quantityAvailable = "default quantityAvailable field value"
		inventoryItem.quantityReserved = "default quantityReserved field value"
		inventoryItem.unitCost = "default unitCost field value"
		inventoryItem.lastUpdated = datetime.datetime.now()
		inventoryItem.stockStatus = "default stockStatus field value"
		
		delegate = InventoryItemDelegate()
		responseObj = delegate.create(inventoryItem)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


