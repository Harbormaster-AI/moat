import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.PayrollCalendar import PayrollCalendar
from hrOnDjango.delegates.PayrollCalendarDelegate import PayrollCalendarDelegate

 #======================================================================
# 
# Encapsulates data for model PayrollCalendar
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollCalendarTest Declaration
#======================================================================
class PayrollCalendarTest (TestCase) :
	def test_crud(self) :
		payrollCalendar = PayrollCalendar()
		payrollCalendar.name = "default name field value"
		payrollCalendar.country = "default country field value"
		payrollCalendar.payFrequency = "default payFrequency field value"
		
		delegate = PayrollCalendarDelegate()
		responseObj = delegate.create(payrollCalendar)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


