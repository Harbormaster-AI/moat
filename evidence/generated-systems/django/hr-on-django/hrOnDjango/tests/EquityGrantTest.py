import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.EquityGrant import EquityGrant
from hrOnDjango.delegates.EquityGrantDelegate import EquityGrantDelegate

 #======================================================================
# 
# Encapsulates data for model EquityGrant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EquityGrantTest Declaration
#======================================================================
class EquityGrantTest (TestCase) :
	def test_crud(self) :
		equityGrant = EquityGrant()
		equityGrant.grantId = "default grantId field value"
		equityGrant.grantedUnits = 22
		equityGrant.vestingStart = datetime.datetime.now()
		equityGrant.grantType = "default grantType field value"
		
		delegate = EquityGrantDelegate()
		responseObj = delegate.create(equityGrant)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


