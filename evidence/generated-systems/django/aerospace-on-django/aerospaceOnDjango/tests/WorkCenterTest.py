import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.WorkCenter import WorkCenter
from aerospaceOnDjango.delegates.WorkCenterDelegate import WorkCenterDelegate

 #======================================================================
# 
# Encapsulates data for model WorkCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkCenterTest Declaration
#======================================================================
class WorkCenterTest (TestCase) :
	def test_crud(self) :
		workCenter = WorkCenter()
		workCenter.name = "default name field value"
		workCenter.capability = "default capability field value"
		
		delegate = WorkCenterDelegate()
		responseObj = delegate.create(workCenter)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


