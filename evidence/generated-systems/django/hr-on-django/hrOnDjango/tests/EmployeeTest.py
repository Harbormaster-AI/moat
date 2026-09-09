import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Employee import Employee
from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

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
		employee.employeeNumber = "default employeeNumber field value"
		employee.name = "default name field value"
		employee.workEmail = "default workEmail field value"
		employee.workPhone = "default workPhone field value"
		employee.dateOfHire = datetime.datetime.now()
		employee.nationalId = "default nationalId field value"
		employee.status = "default status field value"
		
		delegate = EmployeeDelegate()
		responseObj = delegate.create(employee)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


