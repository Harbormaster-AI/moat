import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.RiskAssessment import RiskAssessment
from fintechOnDjango.delegates.RiskAssessmentDelegate import RiskAssessmentDelegate

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
		riskAssessment.score = "default score field value"
		riskAssessment.assessedAt = "default assessedAt field value"
		riskAssessment.modelVersion = "default modelVersion field value"
		riskAssessment.notes = "default notes field value"
		riskAssessment.decision = "default decision field value"
		
		delegate = RiskAssessmentDelegate()
		responseObj = delegate.create(riskAssessment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


