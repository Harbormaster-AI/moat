import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Forecast import Forecast
from manufacturingOnDjango.delegates.ForecastDelegate import ForecastDelegate

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
		forecast.forecastNumber = "default forecastNumber field value"
		forecast.forecastHorizonStart = datetime.datetime.now()
		forecast.forecastHorizonEnd = datetime.datetime.now()
		forecast.method = "default method field value"
		
		delegate = ForecastDelegate()
		responseObj = delegate.create(forecast)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


