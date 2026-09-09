import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Measure import Measure
from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

 #======================================================================
# 
# Encapsulates data for model Measure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MeasureTest Declaration
#======================================================================
class MeasureTest (TestCase) :
	def test_crud(self) :
		measure = Measure()
		measure.name = "default name field value"
		measure.format = "default format field value"
		measure.aggregation = "default aggregation field value"
		
		delegate = MeasureDelegate()
		responseObj = delegate.create(measure)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


