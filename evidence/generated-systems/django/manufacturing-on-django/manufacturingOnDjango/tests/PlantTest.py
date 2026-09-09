import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

 #======================================================================
# 
# Encapsulates data for model Plant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlantTest Declaration
#======================================================================
class PlantTest (TestCase) :
	def test_crud(self) :
		plant = Plant()
		plant.name = "default name field value"
		plant.plantCode = "default plantCode field value"
		plant.address = "default address field value"
		plant.timeZone = "default timeZone field value"
		
		delegate = PlantDelegate()
		responseObj = delegate.create(plant)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


