import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.BenefitPlan import BenefitPlan
from hrOnDjango.delegates.BenefitPlanDelegate import BenefitPlanDelegate

 #======================================================================
# 
# Encapsulates data for model BenefitPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitPlanTest Declaration
#======================================================================
class BenefitPlanTest (TestCase) :
	def test_crud(self) :
		benefitPlan = BenefitPlan()
		benefitPlan.name = "default name field value"
		benefitPlan.providerName = "default providerName field value"
		benefitPlan.employeeContributionRate = "default employeeContributionRate field value"
		benefitPlan.employerContributionRate = "default employerContributionRate field value"
		benefitPlan.eligibilityRules = "default eligibilityRules field value"
		benefitPlan.benefitType = "default benefitType field value"
		
		delegate = BenefitPlanDelegate()
		responseObj = delegate.create(benefitPlan)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


