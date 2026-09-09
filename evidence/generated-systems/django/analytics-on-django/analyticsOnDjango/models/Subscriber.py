from django.db import models
from analyticsOnDjango.models.NotificationChannel import NotificationChannel

#======================================================================
# 
# Encapsulates data for model Subscriber
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Subscriber Declaration
#======================================================================
class Subscriber (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	address = models.CharField(max_length=200, null=True)
	alerts = models.ManyToManyField('Alert',  blank=True, related_name='+')
	channel = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in NotificationChannel])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.address
		str = str + self.channel
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Subscriber";
    
	def objectType(self):
		return "Subscriber";
