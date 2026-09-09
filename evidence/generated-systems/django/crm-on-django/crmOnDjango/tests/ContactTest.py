import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Contact import Contact
from crmOnDjango.delegates.ContactDelegate import ContactDelegate

 #======================================================================
# 
# Encapsulates data for model Contact
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContactTest Declaration
#======================================================================
class ContactTest (TestCase) :
	def test_crud(self) :
		contact = Contact()
		contact.firstName = "default firstName field value"
		contact.lastName = "default lastName field value"
		contact.title = "default title field value"
		contact.email = "default email field value"
		contact.phone = "default phone field value"
		contact.mobile = "default mobile field value"
		contact.mailingAddress = "default mailingAddress field value"
		contact.preferredContactMethod = "default preferredContactMethod field value"
		
		delegate = ContactDelegate()
		responseObj = delegate.create(contact)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


