import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.InventorySource import InventorySource
from advertisingOnDjango.delegates.InventorySourceDelegate import InventorySourceDelegate

 #======================================================================
# 
# Encapsulates data for model InventorySource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventorySourceTest Declaration
#======================================================================
class InventorySourceTest (TestCase) :
	def test_crud(self) :
		inventorySource = InventorySource()
		inventorySource.name = "default name field value"
		inventorySource.domain = "default domain field value"
		inventorySource.channel = "default channel field value"
		inventorySource.primaryFormat = "default primaryFormat field value"
		
		delegate = InventorySourceDelegate()
		responseObj = delegate.create(inventorySource)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


