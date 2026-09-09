import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.RetentionSchedule import RetentionSchedule
from governanceOnDjango.delegates.RetentionScheduleDelegate import RetentionScheduleDelegate

 #======================================================================
# 
# Encapsulates data for model RetentionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RetentionScheduleTest Declaration
#======================================================================
class RetentionScheduleTest (TestCase) :
	def test_crud(self) :
		retentionSchedule = RetentionSchedule()
		retentionSchedule.name = "default name field value"
		retentionSchedule.retentionPeriodMonths = 22
		retentionSchedule.retentionTrigger = "default retentionTrigger field value"
		retentionSchedule.dispositionAction = "default dispositionAction field value"
		retentionSchedule.status = "default status field value"
		
		delegate = RetentionScheduleDelegate()
		responseObj = delegate.create(retentionSchedule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


