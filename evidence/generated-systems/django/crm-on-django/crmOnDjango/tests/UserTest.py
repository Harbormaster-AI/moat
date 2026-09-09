import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.User import User
from crmOnDjango.delegates.UserDelegate import UserDelegate

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
		user.username = "default username field value"
		user.fullName = "default fullName field value"
		user.email = "default email field value"
		user.locale = "default locale field value"
		user.role = "default role field value"
		user.status = "default status field value"
		
		delegate = UserDelegate()
		responseObj = delegate.create(user)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


