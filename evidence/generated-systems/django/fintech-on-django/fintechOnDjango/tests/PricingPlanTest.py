import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.PricingPlan import PricingPlan
from fintechOnDjango.delegates.PricingPlanDelegate import PricingPlanDelegate

 #======================================================================
# 
# Encapsulates data for model PricingPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PricingPlanTest Declaration
#======================================================================
class PricingPlanTest (TestCase) :
	def test_crud(self) :
		pricingPlan = PricingPlan()
		pricingPlan.name = "default name field value"
		pricingPlan.planCode = "default planCode field value"
		pricingPlan.baseCurrency = "default baseCurrency field value"
		pricingPlan.status = "default status field value"
		
		delegate = PricingPlanDelegate()
		responseObj = delegate.create(pricingPlan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


