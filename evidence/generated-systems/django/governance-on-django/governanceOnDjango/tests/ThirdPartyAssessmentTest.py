import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.ThirdPartyAssessment import ThirdPartyAssessment
from governanceOnDjango.delegates.ThirdPartyAssessmentDelegate import ThirdPartyAssessmentDelegate

 #======================================================================
# 
# Encapsulates data for model ThirdPartyAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyAssessmentTest Declaration
#======================================================================
class ThirdPartyAssessmentTest (TestCase) :
	def test_crud(self) :
		thirdPartyAssessment = ThirdPartyAssessment()
		thirdPartyAssessment.assessmentDate = datetime.datetime.now()
		thirdPartyAssessment.assessor = "default assessor field value"
		thirdPartyAssessment.assessmentType = "default assessmentType field value"
		thirdPartyAssessment.result = "default result field value"
		
		delegate = ThirdPartyAssessmentDelegate()
		responseObj = delegate.create(thirdPartyAssessment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


