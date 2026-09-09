import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Person import Person
from governanceOnDjango.delegates.PersonDelegate import PersonDelegate

 #======================================================================
# 
# Encapsulates data for model Person
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PersonTest Declaration
#======================================================================
class PersonTest (TestCase) :
	def test_crud(self) :
		person = Person()
		person.firstName = "default firstName field value"
		person.lastName = "default lastName field value"
		person.email = "default email field value"
		person.department = "default department field value"
		
		delegate = PersonDelegate()
		responseObj = delegate.create(person)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


