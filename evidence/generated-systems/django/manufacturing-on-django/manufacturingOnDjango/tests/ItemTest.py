import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

 #======================================================================
# 
# Encapsulates data for model Item
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ItemTest Declaration
#======================================================================
class ItemTest (TestCase) :
	def test_crud(self) :
		item = Item()
		item.itemNumber = "default itemNumber field value"
		item.name = "default name field value"
		item.standardCost = "default standardCost field value"
		item.weight = "default weight field value"
		item.asSerialControlled = False
		item.itemType = "default itemType field value"
		item.procurementType = "default procurementType field value"
		item.unitOfMeasure = "default unitOfMeasure field value"
		item.lifecycleStatus = "default lifecycleStatus field value"
		
		delegate = ItemDelegate()
		responseObj = delegate.create(item)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


