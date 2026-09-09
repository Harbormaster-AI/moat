import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.EmailMessage import EmailMessage
from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

 #======================================================================
# 
# Encapsulates data for model EmailMessage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmailMessageTest Declaration
#======================================================================
class EmailMessageTest (TestCase) :
	def test_crud(self) :
		emailMessage = EmailMessage()
		emailMessage.subject = "default subject field value"
		emailMessage.body = "default body field value"
		emailMessage.sentAt = "default sentAt field value"
		emailMessage.messageId = "default messageId field value"
		emailMessage.direction = "default direction field value"
		emailMessage.status = "default status field value"
		
		delegate = EmailMessageDelegate()
		responseObj = delegate.create(emailMessage)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


