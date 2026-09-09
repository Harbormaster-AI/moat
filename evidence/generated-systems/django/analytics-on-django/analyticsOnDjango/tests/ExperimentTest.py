import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Experiment import Experiment
from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

 #======================================================================
# 
# Encapsulates data for model Experiment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentTest Declaration
#======================================================================
class ExperimentTest (TestCase) :
	def test_crud(self) :
		experiment = Experiment()
		experiment.name = "default name field value"
		experiment.objective = "default objective field value"
		experiment.status = "default status field value"
		
		delegate = ExperimentDelegate()
		responseObj = delegate.create(experiment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


