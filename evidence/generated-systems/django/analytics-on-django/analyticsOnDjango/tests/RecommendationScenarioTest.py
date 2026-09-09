import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.RecommendationScenario import RecommendationScenario
from analyticsOnDjango.delegates.RecommendationScenarioDelegate import RecommendationScenarioDelegate

 #======================================================================
# 
# Encapsulates data for model RecommendationScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecommendationScenarioTest Declaration
#======================================================================
class RecommendationScenarioTest (TestCase) :
	def test_crud(self) :
		recommendationScenario = RecommendationScenario()
		recommendationScenario.name = "default name field value"
		recommendationScenario.objective = "default objective field value"
		recommendationScenario.recommendationType = "default recommendationType field value"
		
		delegate = RecommendationScenarioDelegate()
		responseObj = delegate.create(recommendationScenario)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


