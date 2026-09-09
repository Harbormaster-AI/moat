import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.TimeEntry import TimeEntry
from hrOnDjango.delegates.TimeEntryDelegate import TimeEntryDelegate

 #======================================================================
# 
# Encapsulates data for model TimeEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeEntryTest Declaration
#======================================================================
class TimeEntryTest (TestCase) :
	def test_crud(self) :
		timeEntry = TimeEntry()
		timeEntry.entryDate = datetime.datetime.now()
		timeEntry.hoursWorked = "default hoursWorked field value"
		timeEntry.entryType = "default entryType field value"
		
		delegate = TimeEntryDelegate()
		responseObj = delegate.create(timeEntry)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


