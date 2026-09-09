import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Location import Location
from hrOnDjango.delegates.LocationDelegate import LocationDelegate

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
		location.name = "default name field value"
		location.address = "default address field value"
		location.timezone = "default timezone field value"
		
		delegate = LocationDelegate()
		responseObj = delegate.create(location)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


