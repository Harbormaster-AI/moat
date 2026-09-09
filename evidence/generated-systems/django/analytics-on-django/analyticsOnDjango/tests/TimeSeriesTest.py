import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.TimeSeries import TimeSeries
from analyticsOnDjango.delegates.TimeSeriesDelegate import TimeSeriesDelegate

 #======================================================================
# 
# Encapsulates data for model TimeSeries
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeSeriesTest Declaration
#======================================================================
class TimeSeriesTest (TestCase) :
	def test_crud(self) :
		timeSeries = TimeSeries()
		timeSeries.name = "default name field value"
		timeSeries.timezone = "default timezone field value"
		timeSeries.granularity = "default granularity field value"
		
		delegate = TimeSeriesDelegate()
		responseObj = delegate.create(timeSeries)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


