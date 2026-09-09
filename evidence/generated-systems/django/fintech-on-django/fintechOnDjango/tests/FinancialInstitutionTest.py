import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

 #======================================================================
# 
# Encapsulates data for model FinancialInstitution
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FinancialInstitutionTest Declaration
#======================================================================
class FinancialInstitutionTest (TestCase) :
	def test_crud(self) :
		financialInstitution = FinancialInstitution()
		financialInstitution.name = "default name field value"
		financialInstitution.legalName = "default legalName field value"
		financialInstitution.countryOfIncorporation = "default countryOfIncorporation field value"
		financialInstitution.bic = "default bic field value"
		financialInstitution.website = "default website field value"
		
		delegate = FinancialInstitutionDelegate()
		responseObj = delegate.create(financialInstitution)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


