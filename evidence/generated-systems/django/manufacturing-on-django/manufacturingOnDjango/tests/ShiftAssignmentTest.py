import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.ShiftAssignment import ShiftAssignment
from manufacturingOnDjango.delegates.ShiftAssignmentDelegate import ShiftAssignmentDelegate

 #======================================================================
# 
# Encapsulates data for model ShiftAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftAssignmentTest Declaration
#======================================================================
class ShiftAssignmentTest (TestCase) :
	def test_crud(self) :
		shiftAssignment = ShiftAssignment()
		shiftAssignment.assignmentDate = datetime.datetime.now()
		
		delegate = ShiftAssignmentDelegate()
		responseObj = delegate.create(shiftAssignment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


