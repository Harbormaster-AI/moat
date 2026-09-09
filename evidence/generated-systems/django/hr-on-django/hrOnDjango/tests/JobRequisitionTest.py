import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.JobRequisition import JobRequisition
from hrOnDjango.delegates.JobRequisitionDelegate import JobRequisitionDelegate

 #======================================================================
# 
# Encapsulates data for model JobRequisition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobRequisitionTest Declaration
#======================================================================
class JobRequisitionTest (TestCase) :
	def test_crud(self) :
		jobRequisition = JobRequisition()
		jobRequisition.requisitionNumber = "default requisitionNumber field value"
		jobRequisition.title = "default title field value"
		jobRequisition.openings = 22
		jobRequisition.targetStartDate = datetime.datetime.now()
		jobRequisition.status = "default status field value"
		jobRequisition.priority = "default priority field value"
		
		delegate = JobRequisitionDelegate()
		responseObj = delegate.create(jobRequisition)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


