import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.WorkShift import WorkShift
from hrOnDjango.delegates.WorkShiftDelegate import WorkShiftDelegate

 #======================================================================
# 
# Encapsulates data for model WorkShift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkShiftTest Declaration
#======================================================================
class WorkShiftTest (TestCase) :
	def test_crud(self) :
		workShift = WorkShift()
		workShift.startTime = "default startTime field value"
		workShift.endTime = "default endTime field value"
		workShift.breakMinutes = 22
		workShift.dayOfWeek = "default dayOfWeek field value"
		
		delegate = WorkShiftDelegate()
		responseObj = delegate.create(workShift)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


