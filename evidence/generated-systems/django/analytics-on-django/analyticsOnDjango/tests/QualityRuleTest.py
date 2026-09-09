import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.QualityRule import QualityRule
from analyticsOnDjango.delegates.QualityRuleDelegate import QualityRuleDelegate

 #======================================================================
# 
# Encapsulates data for model QualityRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityRuleTest Declaration
#======================================================================
class QualityRuleTest (TestCase) :
	def test_crud(self) :
		qualityRule = QualityRule()
		qualityRule.name = "default name field value"
		qualityRule.threshold = "default threshold field value"
		qualityRule.targetField = "default targetField field value"
		qualityRule.dimension = "default dimension field value"
		qualityRule.operator = "default operator field value"
		
		delegate = QualityRuleDelegate()
		responseObj = delegate.create(qualityRule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


