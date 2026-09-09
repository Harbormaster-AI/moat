from django.db import models
from crmOnDjango.models.ActivityType import ActivityType
from crmOnDjango.models.ActivityStatus import ActivityStatus
from crmOnDjango.models.ActivityPriority import ActivityPriority

#======================================================================
# 
# Encapsulates data for model Activity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Activity Declaration
#======================================================================
class Activity (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	subject = models.CharField(max_length=200, null=True)
	dueDate = models.DateField(null=True)
	startAt = models.CharField(max_length=64, null=True)
	endAt = models.CharField(max_length=64, null=True)
	location = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	contact = models.ForeignKey('Contact', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lead = models.ForeignKey('Lead', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	opportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	case = models.ForeignKey('Case_', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	activityType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ActivityType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ActivityStatus])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ActivityPriority])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.subject
		str = str + self.dueDate
		str = str + self.startAt
		str = str + self.endAt
		str = str + self.location
		str = str + self.activityType
		str = str + self.status
		str = str + self.priority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Activity";
    
	def objectType(self):
		return "Activity";
