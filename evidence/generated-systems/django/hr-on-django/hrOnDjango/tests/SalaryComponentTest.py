import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.SalaryComponent import SalaryComponent
from hrOnDjango.delegates.SalaryComponentDelegate import SalaryComponentDelegate

 #======================================================================
# 
# Encapsulates data for model SalaryComponent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalaryComponentTest Declaration
#======================================================================
class SalaryComponentTest (TestCase) :
	def test_crud(self) :
		salaryComponent = SalaryComponent()
		salaryComponent.amount = "default amount field value"
		salaryComponent.recurring = False
		salaryComponent.componentType = "default componentType field value"
		
		delegate = SalaryComponentDelegate()
		responseObj = delegate.create(salaryComponent)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


