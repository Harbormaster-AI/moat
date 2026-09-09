import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.ServiceBulletin import ServiceBulletin
from aerospaceOnDjango.delegates.ServiceBulletinDelegate import ServiceBulletinDelegate

 #======================================================================
# 
# Encapsulates data for model ServiceBulletin
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceBulletinTest Declaration
#======================================================================
class ServiceBulletinTest (TestCase) :
	def test_crud(self) :
		serviceBulletin = ServiceBulletin()
		serviceBulletin.bulletinNumber = "default bulletinNumber field value"
		serviceBulletin.category = "default category field value"
		
		delegate = ServiceBulletinDelegate()
		responseObj = delegate.create(serviceBulletin)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


