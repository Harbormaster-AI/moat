import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.TrainingRun import TrainingRun
from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

 #======================================================================
# 
# Encapsulates data for model TrainingRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingRunTest Declaration
#======================================================================
class TrainingRunTest (TestCase) :
	def test_crud(self) :
		trainingRun = TrainingRun()
		trainingRun.runLabel = "default runLabel field value"
		trainingRun.startedAt = datetime.datetime.now()
		trainingRun.completedAt = datetime.datetime.now()
		trainingRun.status = "default status field value"
		
		delegate = TrainingRunDelegate()
		responseObj = delegate.create(trainingRun)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


