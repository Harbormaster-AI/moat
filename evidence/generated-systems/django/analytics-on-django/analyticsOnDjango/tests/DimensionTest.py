import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Dimension import Dimension
from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

 #======================================================================
# 
# Encapsulates data for model Dimension
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DimensionTest Declaration
#======================================================================
class DimensionTest (TestCase) :
	def test_crud(self) :
		dimension = Dimension()
		dimension.name = "default name field value"
		dimension.typeTime = False
		dimension.dimensionType = "default dimensionType field value"
		
		delegate = DimensionDelegate()
		responseObj = delegate.create(dimension)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


