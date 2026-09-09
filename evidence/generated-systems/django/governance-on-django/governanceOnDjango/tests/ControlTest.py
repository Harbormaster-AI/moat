import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Control import Control
from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

 #======================================================================
# 
# Encapsulates data for model Control
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlTest Declaration
#======================================================================
class ControlTest (TestCase) :
	def test_crud(self) :
		control = Control()
		control.name = "default name field value"
		control.objective = "default objective field value"
		control.ownerDepartment = "default ownerDepartment field value"
		control.controlType = "default controlType field value"
		control.frequency = "default frequency field value"
		control.status = "default status field value"
		
		delegate = ControlDelegate()
		responseObj = delegate.create(control)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


