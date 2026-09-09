import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

 #======================================================================
# 
# Encapsulates data for model Metric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MetricTest Declaration
#======================================================================
class MetricTest (TestCase) :
	def test_crud(self) :
		metric = Metric()
		metric.name = "default name field value"
		metric.expression = "default expression field value"
		metric.unit = "default unit field value"
		metric.metricType = "default metricType field value"
		
		delegate = MetricDelegate()
		responseObj = delegate.create(metric)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


