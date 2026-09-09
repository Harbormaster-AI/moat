import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Distributor import Distributor
from insuranceOnDjango.delegates.DistributorDelegate import DistributorDelegate

 #======================================================================
# 
# Encapsulates data for model Distributor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DistributorTest Declaration
#======================================================================
class DistributorTest (TestCase) :
	def test_crud(self) :
		distributor = Distributor()
		distributor.name = "default name field value"
		distributor.licenseNumber = "default licenseNumber field value"
		distributor.region = "default region field value"
		distributor.distributorType = "default distributorType field value"
		
		delegate = DistributorDelegate()
		responseObj = delegate.create(distributor)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


