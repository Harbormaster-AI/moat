import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.QualityCheck import QualityCheck
from analyticsOnDjango.delegates.QualityCheckDelegate import QualityCheckDelegate

 #======================================================================
# 
# Encapsulates data for model QualityCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityCheckTest Declaration
#======================================================================
class QualityCheckTest (TestCase) :
	def test_crud(self) :
		qualityCheck = QualityCheck()
		qualityCheck.checkedAt = datetime.datetime.now()
		qualityCheck.observedValue = "default observedValue field value"
		qualityCheck.sampleSize = 22
		qualityCheck.status = "default status field value"
		
		delegate = QualityCheckDelegate()
		responseObj = delegate.create(qualityCheck)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


