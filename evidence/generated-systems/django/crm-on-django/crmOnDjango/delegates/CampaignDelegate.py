from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.CampaignMember import CampaignMember
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Lead import Lead
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.Team import Team
from crmOnDjango.models.Activity import Activity
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Campaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignDelegate Declaration
#======================================================================
class CampaignDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, campaignId ):
		try:	
			campaign = Campaign.objects.filter(id=campaignId)
			return campaign.first();
		except Campaign.DoesNotExist:
			raise ProcessingError("Campaign with id " + str(campaignId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, campaign):
		for model in serializers.deserialize("json", campaign):
			model.save()
			return model;

	def create(self, campaign):
		campaign.save()
		return campaign;

	def saveFromJson(self, campaign):
		for model in serializers.deserialize("json", campaign):
			model.save()
			return campaign;
	
	def save(self, campaign):
		campaign.save()
		return campaign;
	
	def delete(self, campaignId ):
		errMsg = "Failed to delete Campaign from db using id " + str(campaignId)
		
		try:
			campaign = Campaign.objects.get(id=campaignId)
			campaign.delete()
			return True
		except Campaign.DoesNotExist:
			raise ProcessingError("Campaign with id " + str(campaignId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Campaign.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Campaign from db")
		except Exception:
			return None;
		
	def assignOrganization( self, campaignId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			campaign.organization = organization
			
			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, campaignId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# assign to None for unassignment
			campaign.organization = None			

			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
		
	def assignParentCampaign( self, campaignId, parentCampaignId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(parentCampaignId) + " for ParentCampaign on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(parentCampaignId).first();
			
			# assign the ParentCampaign		
			campaign.parentCampaign = campaign
			
			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(parentCampaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignParentCampaign( self, campaignId ):
		errMsg = "Failed to unassign element " + str(parentCampaignId) + " for ParentCampaign on Campaign"

		try:
			# get the Campaign from db
			campaign = self.get( campaignId ).first()	
			
			# assign to None for unassignment
			campaign.campaign = None			

			#save it
			campaign.save()

			# reload and return the appropriate version					
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
		
	def addChildCampaigns( self, campaignId, childCampaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(childCampaignsIds) + " for ChildCampaigns on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = childCampaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				campaign.childCampaigns.add(campaign)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChildCampaigns( self, campaignId, childCampaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(childCampaignsIds) + " for ChildCampaigns on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = childCampaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				campaign.childCampaigns.remove(campaign)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMembers( self, campaignId, membersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignMemberDelegate import CampaignMemberDelegate

		errMsg = "Failed to add elements " + str(membersIds) + " for Members on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = membersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CampaignMember		
				campaignMember = CampaignMemberDelegate().get(id).first();	
				# add the CampaignMember
				campaign.members.add(campaignMember)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMembers( self, campaignId, membersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignMemberDelegate import CampaignMemberDelegate

		errMsg = "Failed to remove elements " + str(membersIds) + " for Members on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = membersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CampaignMember		
				campaignMember = CampaignMemberDelegate().get(id).first();	
				# add the CampaignMember
				campaign.members.remove(campaignMember)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOpportunities( self, campaignId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to add elements " + str(opportunitiesIds) + " for Opportunities on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				campaign.opportunities.add(opportunity)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOpportunities( self, campaignId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to remove elements " + str(opportunitiesIds) + " for Opportunities on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				campaign.opportunities.remove(opportunity)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAccounts( self, campaignId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to add elements " + str(accountsIds) + " for Accounts on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				campaign.accounts.add(account)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAccounts( self, campaignId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to remove elements " + str(accountsIds) + " for Accounts on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				campaign.accounts.remove(account)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLeads( self, campaignId, leadsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to add elements " + str(leadsIds) + " for Leads on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = leadsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Lead		
				lead = LeadDelegate().get(id).first();	
				# add the Lead
				campaign.leads.add(lead)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLeads( self, campaignId, leadsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to remove elements " + str(leadsIds) + " for Leads on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = leadsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Lead		
				lead = LeadDelegate().get(id).first();	
				# add the Lead
				campaign.leads.remove(lead)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContacts( self, campaignId, contactsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to add elements " + str(contactsIds) + " for Contacts on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = contactsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contact		
				contact = ContactDelegate().get(id).first();	
				# add the Contact
				campaign.contacts.add(contact)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContacts( self, campaignId, contactsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to remove elements " + str(contactsIds) + " for Contacts on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = contactsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contact		
				contact = ContactDelegate().get(id).first();	
				# add the Contact
				campaign.contacts.remove(contact)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTeams( self, campaignId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to add elements " + str(teamsIds) + " for Teams on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				campaign.teams.add(team)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTeams( self, campaignId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to remove elements " + str(teamsIds) + " for Teams on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				campaign.teams.remove(team)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addActivities( self, campaignId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to add elements " + str(activitiesIds) + " for Activities on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				campaign.activities.add(activity)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeActivities( self, campaignId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to remove elements " + str(activitiesIds) + " for Activities on Campaign"

		try:
			# get the Campaign
			campaign = self.get( campaignId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				campaign.activities.remove(activity)
				
			# save it		
			campaign.save()
			
			# reload and return the appropriate version
			return self.get( campaignId );
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
