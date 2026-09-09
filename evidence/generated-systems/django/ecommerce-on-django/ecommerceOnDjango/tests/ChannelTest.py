import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

 #======================================================================
# 
# Encapsulates data for model Channel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChannelTest Declaration
#======================================================================
class ChannelTest (TestCase) :
	def test_crud(self) :
		channel = Channel()
		channel.name = "default name field value"
		channel.channelCode = "default channelCode field value"
		channel.locale = "default locale field value"
		channel.domain = "default domain field value"
		channel.asActive = False
		channel.defaultCurrency = "default defaultCurrency field value"
		channel.channelType = "default channelType field value"
		
		delegate = ChannelDelegate()
		responseObj = delegate.create(channel)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


