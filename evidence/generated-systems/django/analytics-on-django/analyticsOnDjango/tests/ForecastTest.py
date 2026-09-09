import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Forecast import Forecast
from analyticsOnDjango.delegates.ForecastDelegate import ForecastDelegate

 #======================================================================
# 
# Encapsulates data for model Forecast
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastTest Declaration
#======================================================================
class ForecastTest (TestCase) :
	def test_crud(self) :
		forecast = Forecast()
		forecast.name = "default name field value"
		forecast.horizon = 22
		forecast.granularity = "default granularity field value"
		
		delegate = ForecastDelegate()
		responseObj = delegate.create(forecast)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


