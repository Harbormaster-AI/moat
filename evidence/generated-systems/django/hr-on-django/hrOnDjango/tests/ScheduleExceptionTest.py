import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.ScheduleException import ScheduleException
from hrOnDjango.delegates.ScheduleExceptionDelegate import ScheduleExceptionDelegate

 #======================================================================
# 
# Encapsulates data for model ScheduleException
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScheduleExceptionTest Declaration
#======================================================================
class ScheduleExceptionTest (TestCase) :
	def test_crud(self) :
		scheduleException = ScheduleException()
		scheduleException.date = datetime.datetime.now()
		scheduleException.reason = "default reason field value"
		scheduleException.hours = "default hours field value"
		
		delegate = ScheduleExceptionDelegate()
		responseObj = delegate.create(scheduleException)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


