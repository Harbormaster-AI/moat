import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.ImagingReport import ImagingReport
from healthcareOnDjango.delegates.ImagingReportDelegate import ImagingReportDelegate

 #======================================================================
# 
# Encapsulates data for model ImagingReport
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingReportTest Declaration
#======================================================================
class ImagingReportTest (TestCase) :
	def test_crud(self) :
		imagingReport = ImagingReport()
		imagingReport.reportNumber = "default reportNumber field value"
		imagingReport.impression = "default impression field value"
		imagingReport.reportedDate = "default reportedDate field value"
		imagingReport.status = "default status field value"
		
		delegate = ImagingReportDelegate()
		responseObj = delegate.create(imagingReport)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


