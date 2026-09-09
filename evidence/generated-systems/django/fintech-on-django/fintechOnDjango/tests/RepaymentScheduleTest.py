import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.RepaymentSchedule import RepaymentSchedule
from fintechOnDjango.delegates.RepaymentScheduleDelegate import RepaymentScheduleDelegate

 #======================================================================
# 
# Encapsulates data for model RepaymentSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RepaymentScheduleTest Declaration
#======================================================================
class RepaymentScheduleTest (TestCase) :
	def test_crud(self) :
		repaymentSchedule = RepaymentSchedule()
		repaymentSchedule.installmentNumber = 22
		repaymentSchedule.dueDate = datetime.datetime.now()
		repaymentSchedule.amountDue = "default amountDue field value"
		repaymentSchedule.principalDue = "default principalDue field value"
		repaymentSchedule.interestDue = "default interestDue field value"
		repaymentSchedule.status = "default status field value"
		
		delegate = RepaymentScheduleDelegate()
		responseObj = delegate.create(repaymentSchedule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


