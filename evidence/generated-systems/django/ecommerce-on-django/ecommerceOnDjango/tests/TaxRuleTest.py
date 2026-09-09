import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.TaxRule import TaxRule
from ecommerceOnDjango.delegates.TaxRuleDelegate import TaxRuleDelegate

 #======================================================================
# 
# Encapsulates data for model TaxRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxRuleTest Declaration
#======================================================================
class TaxRuleTest (TestCase) :
	def test_crud(self) :
		taxRule = TaxRule()
		taxRule.name = "default name field value"
		taxRule.country = "default country field value"
		taxRule.region = "default region field value"
		taxRule.rate = "default rate field value"
		taxRule.taxInclusive = False
		taxRule.taxClass = "default taxClass field value"
		
		delegate = TaxRuleDelegate()
		responseObj = delegate.create(taxRule)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


