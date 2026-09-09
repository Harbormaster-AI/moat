import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.RoleAssignment import RoleAssignment
from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

 #======================================================================
# 
# Encapsulates data for model RoleAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoleAssignmentTest Declaration
#======================================================================
class RoleAssignmentTest (TestCase) :
	def test_crud(self) :
		roleAssignment = RoleAssignment()
		roleAssignment.effectiveFrom = datetime.datetime.now()
		roleAssignment.effectiveTo = datetime.datetime.now()
		
		delegate = RoleAssignmentDelegate()
		responseObj = delegate.create(roleAssignment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


