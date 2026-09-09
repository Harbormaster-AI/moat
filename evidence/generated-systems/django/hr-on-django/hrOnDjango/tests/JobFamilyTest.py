import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.JobFamily import JobFamily
from hrOnDjango.delegates.JobFamilyDelegate import JobFamilyDelegate

 #======================================================================
# 
# Encapsulates data for model JobFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobFamilyTest Declaration
#======================================================================
class JobFamilyTest (TestCase) :
	def test_crud(self) :
		jobFamily = JobFamily()
		jobFamily.name = "default name field value"
		jobFamily.description = "default description field value"
		
		delegate = JobFamilyDelegate()
		responseObj = delegate.create(jobFamily)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


