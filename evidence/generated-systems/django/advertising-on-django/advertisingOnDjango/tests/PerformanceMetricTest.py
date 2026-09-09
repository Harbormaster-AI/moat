import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.PerformanceMetric import PerformanceMetric
from advertisingOnDjango.delegates.PerformanceMetricDelegate import PerformanceMetricDelegate

 #======================================================================
# 
# Encapsulates data for model PerformanceMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceMetricTest Declaration
#======================================================================
class PerformanceMetricTest (TestCase) :
	def test_crud(self) :
		performanceMetric = PerformanceMetric()
		performanceMetric.date = datetime.datetime.now()
		performanceMetric.value = "default value field value"
		performanceMetric.metricType = "default metricType field value"
		
		delegate = PerformanceMetricDelegate()
		responseObj = delegate.create(performanceMetric)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


