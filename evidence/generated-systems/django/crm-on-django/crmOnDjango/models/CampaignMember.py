from django.db import models
from crmOnDjango.models.CampaignMemberStatus import CampaignMemberStatus
from crmOnDjango.models.CampaignMemberType import CampaignMemberType

#======================================================================
# 
# Encapsulates data for model CampaignMember
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignMember Declaration
#======================================================================
class CampaignMember (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	responded = models.BooleanField(null=True)
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lead = models.ForeignKey('Lead', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	contact = models.ForeignKey('Contact', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CampaignMemberStatus])
	memberType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CampaignMemberType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.responded
		str = str + self.status
		str = str + self.memberType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CampaignMember";
    
	def objectType(self):
		return "CampaignMember";
