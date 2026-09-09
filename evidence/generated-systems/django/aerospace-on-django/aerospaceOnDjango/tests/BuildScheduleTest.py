import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.BuildSchedule import BuildSchedule
from aerospaceOnDjango.delegates.BuildScheduleDelegate import BuildScheduleDelegate

 #======================================================================
# 
# Encapsulates data for model BuildSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BuildScheduleTest Declaration
#======================================================================
class BuildScheduleTest (TestCase) :
	def test_crud(self) :
		buildSchedule = BuildSchedule()
		buildSchedule.scheduleNumber = "default scheduleNumber field value"
		buildSchedule.status = "default status field value"
		
		delegate = BuildScheduleDelegate()
		responseObj = delegate.create(buildSchedule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


