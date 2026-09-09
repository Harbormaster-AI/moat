import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Location import Location
from manufacturingOnDjango.delegates.LocationDelegate import LocationDelegate

 #======================================================================
# 
# Encapsulates data for model Location
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LocationTest Declaration
#======================================================================
class LocationTest (TestCase) :
	def test_crud(self) :
		location = Location()
		location.locationCode = "default locationCode field value"
		location.description = "default description field value"
		location.locationType = "default locationType field value"
		
		delegate = LocationDelegate()
		responseObj = delegate.create(location)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


