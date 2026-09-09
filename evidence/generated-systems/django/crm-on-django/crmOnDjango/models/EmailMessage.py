from django.db import models
from crmOnDjango.models.EmailDirection import EmailDirection
from crmOnDjango.models.EmailStatus import EmailStatus

#======================================================================
# 
# Encapsulates data for model EmailMessage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmailMessage Declaration
#======================================================================
class EmailMessage (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	subject = models.CharField(max_length=200, null=True)
	body = models.CharField(max_length=200, null=True)
	sentAt = models.CharField(max_length=64, null=True)
	messageId = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	contact = models.ForeignKey('Contact', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lead = models.ForeignKey('Lead', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	case = models.ForeignKey('Case_', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	opportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	direction = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EmailDirection])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EmailStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.subject
		str = str + self.body
		str = str + self.sentAt
		str = str + self.messageId
		str = str + self.direction
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "EmailMessage";
    
	def objectType(self):
		return "EmailMessage";
