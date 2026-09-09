import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.Report import Report
from advertisingOnDjango.delegates.ReportDelegate import ReportDelegate

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
		report.reportName = "default reportName field value"
		report.generatedAt = "default generatedAt field value"
		report.fileUrl = "default fileUrl field value"
		report.reportType = "default reportType field value"
		
		delegate = ReportDelegate()
		responseObj = delegate.create(report)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


