import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.EmploymentAssignment import EmploymentAssignment
from hrOnDjango.delegates.EmploymentAssignmentDelegate import EmploymentAssignmentDelegate

 #======================================================================
# 
# Encapsulates data for model EmploymentAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentAssignmentTest Declaration
#======================================================================
class EmploymentAssignmentTest (TestCase) :
	def test_crud(self) :
		employmentAssignment = EmploymentAssignment()
		employmentAssignment.startDate = datetime.datetime.now()
		employmentAssignment.endDate = datetime.datetime.now()
		employmentAssignment.primary = False
		employmentAssignment.assignmentType = "default assignmentType field value"
		employmentAssignment.status = "default status field value"
		
		delegate = EmploymentAssignmentDelegate()
		responseObj = delegate.create(employmentAssignment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


