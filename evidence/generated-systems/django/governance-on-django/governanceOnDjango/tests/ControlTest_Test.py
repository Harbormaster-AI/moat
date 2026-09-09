import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.ControlTest_ import ControlTest_
from governanceOnDjango.delegates.ControlTest_Delegate import ControlTest_Delegate

 #======================================================================
# 
# Encapsulates data for model ControlTest_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlTest_Test Declaration
#======================================================================
class ControlTest_Test (TestCase) :
	def test_crud(self) :
		controlTest_ = ControlTest_()
		controlTest_.name = "default name field value"
		controlTest_.testPeriodStart = datetime.datetime.now()
		controlTest_.testPeriodEnd = datetime.datetime.now()
		controlTest_.sampleSize = 22
		controlTest_.testType = "default testType field value"
		controlTest_.effectiveness = "default effectiveness field value"
		controlTest_.status = "default status field value"
		
		delegate = ControlTest_Delegate()
		responseObj = delegate.create(controlTest_)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


