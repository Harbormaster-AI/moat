import datetime

from django.test import TestCase
from django.utils import timezone
from inventoryOnDjango.models.StorageLocation import StorageLocation
from inventoryOnDjango.delegates.StorageLocationDelegate import StorageLocationDelegate

 #======================================================================
# 
# Encapsulates data for model StorageLocation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class StorageLocationTest Declaration
#======================================================================
class StorageLocationTest (TestCase) :
	def test_crud(self) :
		storageLocation = StorageLocation()
		storageLocation.code = "default code field value"
		storageLocation.temperatureControlled = False
		storageLocation.capacity = "default capacity field value"
		storageLocation.capacityUnit = "default capacityUnit field value"
		storageLocation.locationType = "default locationType field value"
		
		delegate = StorageLocationDelegate()
		responseObj = delegate.create(storageLocation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


