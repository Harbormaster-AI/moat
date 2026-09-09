import datetime

from django.test import TestCase
from django.utils import timezone
from advertisingOnDjango.models.User import User
from advertisingOnDjango.delegates.UserDelegate import UserDelegate

 #======================================================================
# 
# Encapsulates data for model User
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UserTest Declaration
#======================================================================
class UserTest (TestCase) :
	def test_crud(self) :
		user = User()
		user.firstName = "default firstName field value"
		user.lastName = "default lastName field value"
		user.email = "default email field value"
		user.role = "default role field value"
		
		delegate = UserDelegate()
		responseObj = delegate.create(user)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


