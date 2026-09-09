import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.EvaluationMetric import EvaluationMetric
from analyticsOnDjango.delegates.EvaluationMetricDelegate import EvaluationMetricDelegate

 #======================================================================
# 
# Encapsulates data for model EvaluationMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EvaluationMetricTest Declaration
#======================================================================
class EvaluationMetricTest (TestCase) :
	def test_crud(self) :
		evaluationMetric = EvaluationMetric()
		evaluationMetric.name = "default name field value"
		evaluationMetric.value = "default value field value"
		
		delegate = EvaluationMetricDelegate()
		responseObj = delegate.create(evaluationMetric)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


