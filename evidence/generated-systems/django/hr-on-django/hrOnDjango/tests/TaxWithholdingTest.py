import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.TaxWithholding import TaxWithholding
from hrOnDjango.delegates.TaxWithholdingDelegate import TaxWithholdingDelegate

 #======================================================================
# 
# Encapsulates data for model TaxWithholding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TaxWithholdingTest Declaration
#======================================================================
class TaxWithholdingTest (TestCase) :
	def test_crud(self) :
		taxWithholding = TaxWithholding()
		taxWithholding.taxId = "default taxId field value"
		taxWithholding.allowances = 22
		taxWithholding.additionalAmount = "default additionalAmount field value"
		taxWithholding.filingStatus = "default filingStatus field value"
		
		delegate = TaxWithholdingDelegate()
		responseObj = delegate.create(taxWithholding)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


