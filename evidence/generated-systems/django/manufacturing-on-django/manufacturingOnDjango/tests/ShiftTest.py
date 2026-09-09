import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Shift import Shift
from manufacturingOnDjango.delegates.ShiftDelegate import ShiftDelegate

 #======================================================================
# 
# Encapsulates data for model Shift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftTest Declaration
#======================================================================
class ShiftTest (TestCase) :
	def test_crud(self) :
		shift = Shift()
		shift.shiftName = "default shiftName field value"
		shift.startTime = "default startTime field value"
		shift.endTime = "default endTime field value"
		shift.shiftType = "default shiftType field value"
		
		delegate = ShiftDelegate()
		responseObj = delegate.create(shift)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


