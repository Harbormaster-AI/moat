import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.SalesRegion import SalesRegion
from aerospaceOnDjango.delegates.SalesRegionDelegate import SalesRegionDelegate

 #======================================================================
# 
# Encapsulates data for model SalesRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesRegionTest Declaration
#======================================================================
class SalesRegionTest (TestCase) :
	def test_crud(self) :
		salesRegion = SalesRegion()
		salesRegion.name = "default name field value"
		salesRegion.regionCode = "default regionCode field value"
		
		delegate = SalesRegionDelegate()
		responseObj = delegate.create(salesRegion)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


