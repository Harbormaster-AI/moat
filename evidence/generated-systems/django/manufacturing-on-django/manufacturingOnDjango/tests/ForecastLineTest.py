import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.ForecastLine import ForecastLine
from manufacturingOnDjango.delegates.ForecastLineDelegate import ForecastLineDelegate

 #======================================================================
# 
# Encapsulates data for model ForecastLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastLineTest Declaration
#======================================================================
class ForecastLineTest (TestCase) :
	def test_crud(self) :
		forecastLine = ForecastLine()
		forecastLine.period = datetime.datetime.now()
		forecastLine.quantity = "default quantity field value"
		forecastLine.confidence = "default confidence field value"
		
		delegate = ForecastLineDelegate()
		responseObj = delegate.create(forecastLine)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


