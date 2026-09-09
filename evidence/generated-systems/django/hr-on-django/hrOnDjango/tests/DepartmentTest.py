import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Department import Department
from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

 #======================================================================
# 
# Encapsulates data for model Department
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DepartmentTest Declaration
#======================================================================
class DepartmentTest (TestCase) :
	def test_crud(self) :
		department = Department()
		department.name = "default name field value"
		department.code = "default code field value"
		
		delegate = DepartmentDelegate()
		responseObj = delegate.create(department)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


