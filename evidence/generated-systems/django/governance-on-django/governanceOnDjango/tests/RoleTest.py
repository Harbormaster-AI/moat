import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Role import Role
from governanceOnDjango.delegates.RoleDelegate import RoleDelegate

 #======================================================================
# 
# Encapsulates data for model Role
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoleTest Declaration
#======================================================================
class RoleTest (TestCase) :
	def test_crud(self) :
		role = Role()
		role.name = "default name field value"
		role.responsibility = "default responsibility field value"
		
		delegate = RoleDelegate()
		responseObj = delegate.create(role)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


