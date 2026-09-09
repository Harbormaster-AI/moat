import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.GeoRegion import GeoRegion
from advertisingOnDjango.delegates.GeoRegionDelegate import GeoRegionDelegate

 #======================================================================
# 
# Encapsulates data for model GeoRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GeoRegionTest Declaration
#======================================================================
class GeoRegionTest (TestCase) :
	def test_crud(self) :
		geoRegion = GeoRegion()
		geoRegion.code = "default code field value"
		geoRegion.name = "default name field value"
		geoRegion.regionType = "default regionType field value"
		
		delegate = GeoRegionDelegate()
		responseObj = delegate.create(geoRegion)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


