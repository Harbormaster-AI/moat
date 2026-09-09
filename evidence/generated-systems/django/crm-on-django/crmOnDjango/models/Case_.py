from django.db import models
from crmOnDjango.models.CaseStatus import CaseStatus
from crmOnDjango.models.CasePriority import CasePriority
from crmOnDjango.models.CaseOrigin import CaseOrigin
from crmOnDjango.models.CaseSeverity import CaseSeverity

#======================================================================
# 
# Encapsulates data for model Case_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Case_ Declaration
#======================================================================
class Case_ (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	caseNumber = models.CharField(max_length=200, null=True)
	subject = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	slaDue = models.CharField(max_length=64, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	account = models.ForeignKey('Account', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	contact = models.ForeignKey('Contact', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	owner = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	team = models.ForeignKey('Team', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	activities = models.ManyToManyField('Activity',  blank=True, related_name='+')
	caseComments = models.ManyToManyField('Note',  blank=True, related_name='+')
	emails = models.ManyToManyField('EmailMessage',  blank=True, related_name='+')
	relatedOpportunities = models.ManyToManyField('Opportunity',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CaseStatus])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CasePriority])
	origin = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CaseOrigin])
	severity = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CaseSeverity])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.caseNumber
		str = str + self.subject
		str = str + self.description
		str = str + self.slaDue
		str = str + self.status
		str = str + self.priority
		str = str + self.origin
		str = str + self.severity
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Case_";
    
	def objectType(self):
		return "Case_";
