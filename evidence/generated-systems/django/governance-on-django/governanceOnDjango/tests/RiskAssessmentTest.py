import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.RiskAssessment import RiskAssessment
from governanceOnDjango.delegates.RiskAssessmentDelegate import RiskAssessmentDelegate

 #======================================================================
# 
# Encapsulates data for model RiskAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskAssessmentTest Declaration
#======================================================================
class RiskAssessmentTest (TestCase) :
	def test_crud(self) :
		riskAssessment = RiskAssessment()
		riskAssessment.assessmentDate = datetime.datetime.now()
		riskAssessment.assessor = "default assessor field value"
		riskAssessment.summary = "default summary field value"
		riskAssessment.assessmentType = "default assessmentType field value"
		
		delegate = RiskAssessmentDelegate()
		responseObj = delegate.create(riskAssessment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


