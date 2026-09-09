import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.FeeSchedule import FeeSchedule
from fintechOnDjango.delegates.FeeScheduleDelegate import FeeScheduleDelegate

 #======================================================================
# 
# Encapsulates data for model FeeSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeeScheduleTest Declaration
#======================================================================
class FeeScheduleTest (TestCase) :
	def test_crud(self) :
		feeSchedule = FeeSchedule()
		feeSchedule.name = "default name field value"
		feeSchedule.amount = "default amount field value"
		feeSchedule.percentage = "default percentage field value"
		feeSchedule.minimum = "default minimum field value"
		feeSchedule.maximum = "default maximum field value"
		feeSchedule.feeType = "default feeType field value"
		feeSchedule.calculationMethod = "default calculationMethod field value"
		
		delegate = FeeScheduleDelegate()
		responseObj = delegate.create(feeSchedule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


