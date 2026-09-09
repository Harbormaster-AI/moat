import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.RunMetric import RunMetric
from analyticsOnDjango.delegates.RunMetricDelegate import RunMetricDelegate

 #======================================================================
# 
# Encapsulates data for model RunMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunMetricTest Declaration
#======================================================================
class RunMetricTest (TestCase) :
	def test_crud(self) :
		runMetric = RunMetric()
		runMetric.name = "default name field value"
		runMetric.value = "default value field value"
		
		delegate = RunMetricDelegate()
		responseObj = delegate.create(runMetric)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


