import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Timesheet import Timesheet
from hrOnDjango.delegates.TimesheetDelegate import TimesheetDelegate

 #======================================================================
# 
# Encapsulates data for model Timesheet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimesheetTest Declaration
#======================================================================
class TimesheetTest (TestCase) :
	def test_crud(self) :
		timesheet = Timesheet()
		timesheet.periodStart = datetime.datetime.now()
		timesheet.periodEnd = datetime.datetime.now()
		timesheet.submissionDate = datetime.datetime.now()
		timesheet.status = "default status field value"
		
		delegate = TimesheetDelegate()
		responseObj = delegate.create(timesheet)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


