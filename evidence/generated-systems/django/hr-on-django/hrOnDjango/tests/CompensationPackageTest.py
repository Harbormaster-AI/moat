import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.CompensationPackage import CompensationPackage
from hrOnDjango.delegates.CompensationPackageDelegate import CompensationPackageDelegate

 #======================================================================
# 
# Encapsulates data for model CompensationPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompensationPackageTest Declaration
#======================================================================
class CompensationPackageTest (TestCase) :
	def test_crud(self) :
		compensationPackage = CompensationPackage()
		compensationPackage.effectiveFrom = datetime.datetime.now()
		compensationPackage.effectiveTo = datetime.datetime.now()
		compensationPackage.currency = "default currency field value"
		
		delegate = CompensationPackageDelegate()
		responseObj = delegate.create(compensationPackage)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


