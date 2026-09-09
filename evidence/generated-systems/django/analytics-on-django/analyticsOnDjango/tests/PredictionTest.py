import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Prediction import Prediction
from analyticsOnDjango.delegates.PredictionDelegate import PredictionDelegate

 #======================================================================
# 
# Encapsulates data for model Prediction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PredictionTest Declaration
#======================================================================
class PredictionTest (TestCase) :
	def test_crud(self) :
		prediction = Prediction()
		prediction.referenceKey = "default referenceKey field value"
		prediction.predictedAt = datetime.datetime.now()
		prediction.score = "default score field value"
		
		delegate = PredictionDelegate()
		responseObj = delegate.create(prediction)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


