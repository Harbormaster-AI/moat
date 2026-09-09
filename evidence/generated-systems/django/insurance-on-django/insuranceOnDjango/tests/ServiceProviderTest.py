import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.ServiceProvider import ServiceProvider
from insuranceOnDjango.delegates.ServiceProviderDelegate import ServiceProviderDelegate

 #======================================================================
# 
# Encapsulates data for model ServiceProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceProviderTest Declaration
#======================================================================
class ServiceProviderTest (TestCase) :
	def test_crud(self) :
		serviceProvider = ServiceProvider()
		serviceProvider.name = "default name field value"
		serviceProvider.taxId = "default taxId field value"
		serviceProvider.providerType = "default providerType field value"
		serviceProvider.networkStatus = "default networkStatus field value"
		
		delegate = ServiceProviderDelegate()
		responseObj = delegate.create(serviceProvider)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


