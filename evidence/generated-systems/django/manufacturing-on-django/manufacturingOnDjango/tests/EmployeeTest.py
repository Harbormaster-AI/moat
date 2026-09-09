import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Employee import Employee
from manufacturingOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

 #======================================================================
# 
# Encapsulates data for model Employee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmployeeTest Declaration
#======================================================================
class EmployeeTest (TestCase) :
	def test_crud(self) :
		employee = Employee()
		employee.firstName = "default firstName field value"
		employee.lastName = "default lastName field value"
		employee.role = "default role field value"
		employee.skillLevel = "default skillLevel field value"
		
		delegate = EmployeeDelegate()
		responseObj = delegate.create(employee)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


