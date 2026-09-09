import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.PerformanceCycle import PerformanceCycle
from hrOnDjango.delegates.PerformanceCycleDelegate import PerformanceCycleDelegate

 #======================================================================
# 
# Encapsulates data for model PerformanceCycle
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PerformanceCycleTest Declaration
#======================================================================
class PerformanceCycleTest (TestCase) :
	def test_crud(self) :
		performanceCycle = PerformanceCycle()
		performanceCycle.name = "default name field value"
		performanceCycle.startDate = datetime.datetime.now()
		performanceCycle.endDate = datetime.datetime.now()
		performanceCycle.status = "default status field value"
		
		delegate = PerformanceCycleDelegate()
		responseObj = delegate.create(performanceCycle)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


