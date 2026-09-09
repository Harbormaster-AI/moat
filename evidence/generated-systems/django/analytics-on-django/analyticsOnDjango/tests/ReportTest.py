import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

 #======================================================================
# 
# Encapsulates data for model Report
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReportTest Declaration
#======================================================================
class ReportTest (TestCase) :
	def test_crud(self) :
		report = Report()
		report.title = "default title field value"
		report.audience = "default audience field value"
		report.status = "default status field value"
		
		delegate = ReportDelegate()
		responseObj = delegate.create(report)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


