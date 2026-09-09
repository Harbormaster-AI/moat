import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

 #======================================================================
# 
# Encapsulates data for model Risk
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskTest Declaration
#======================================================================
class RiskTest (TestCase) :
	def test_crud(self) :
		risk = Risk()
		risk.name = "default name field value"
		risk.description = "default description field value"
		risk.inherentRiskScore = 22
		risk.residualRiskScore = 22
		risk.category = "default category field value"
		risk.impact = "default impact field value"
		risk.likelihood = "default likelihood field value"
		risk.status = "default status field value"
		
		delegate = RiskDelegate()
		responseObj = delegate.create(risk)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


