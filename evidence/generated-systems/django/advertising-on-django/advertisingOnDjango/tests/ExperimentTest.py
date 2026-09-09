import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Experiment import Experiment
from advertisingOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

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
		experiment.hypothesis = "default hypothesis field value"
		experiment.startDate = datetime.datetime.now()
		experiment.endDate = datetime.datetime.now()
		experiment.status = "default status field value"
		
		delegate = ExperimentDelegate()
		responseObj = delegate.create(experiment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


