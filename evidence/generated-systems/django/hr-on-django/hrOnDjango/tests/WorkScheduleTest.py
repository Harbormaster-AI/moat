import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.WorkSchedule import WorkSchedule
from hrOnDjango.delegates.WorkScheduleDelegate import WorkScheduleDelegate

 #======================================================================
# 
# Encapsulates data for model WorkSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkScheduleTest Declaration
#======================================================================
class WorkScheduleTest (TestCase) :
	def test_crud(self) :
		workSchedule = WorkSchedule()
		workSchedule.name = "default name field value"
		workSchedule.standardHoursPerWeek = "default standardHoursPerWeek field value"
		workSchedule.scheduleType = "default scheduleType field value"
		
		delegate = WorkScheduleDelegate()
		responseObj = delegate.create(workSchedule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


