from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.OpportunityStageHistory import OpportunityStageHistory
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.User import User
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model OpportunityStageHistory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityStageHistoryDelegate Declaration
#======================================================================
class OpportunityStageHistoryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, opportunityStageHistoryId ):
		try:	
			opportunityStageHistory = OpportunityStageHistory.objects.filter(id=opportunityStageHistoryId)
			return opportunityStageHistory.first();
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError("OpportunityStageHistory with id " + str(opportunityStageHistoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, opportunityStageHistory):
		for model in serializers.deserialize("json", opportunityStageHistory):
			model.save()
			return model;

	def create(self, opportunityStageHistory):
		opportunityStageHistory.save()
		return opportunityStageHistory;

	def saveFromJson(self, opportunityStageHistory):
		for model in serializers.deserialize("json", opportunityStageHistory):
			model.save()
			return opportunityStageHistory;
	
	def save(self, opportunityStageHistory):
		opportunityStageHistory.save()
		return opportunityStageHistory;
	
	def delete(self, opportunityStageHistoryId ):
		errMsg = "Failed to delete OpportunityStageHistory from db using id " + str(opportunityStageHistoryId)
		
		try:
			opportunityStageHistory = OpportunityStageHistory.objects.get(id=opportunityStageHistoryId)
			opportunityStageHistory.delete()
			return True
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError("OpportunityStageHistory with id " + str(opportunityStageHistoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = OpportunityStageHistory.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all OpportunityStageHistory from db")
		except Exception:
			return None;
		
	def assignOpportunity( self, opportunityStageHistoryId, opportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(opportunityId) + " for Opportunity on OpportunityStageHistory"

		try:
			# get the OpportunityStageHistory from db
			opportunityStageHistory = self.get( opportunityStageHistoryId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(opportunityId).first();
			
			# assign the Opportunity		
			opportunityStageHistory.opportunity = opportunity
			
			#save it
			opportunityStageHistory.save()

			# reload and return the appropriate version					
			return self.get( opportunityStageHistoryId );
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityStageHistory with id " + str(opportunityStageHistoryId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunity( self, opportunityStageHistoryId ):
		errMsg = "Failed to unassign element " + str(opportunityId) + " for Opportunity on OpportunityStageHistory"

		try:
			# get the OpportunityStageHistory from db
			opportunityStageHistory = self.get( opportunityStageHistoryId ).first()	
			
			# assign to None for unassignment
			opportunityStageHistory.opportunity = None			

			#save it
			opportunityStageHistory.save()

			# reload and return the appropriate version					
			return self.get( opportunityStageHistoryId );
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityStageHistory with id " + str(opportunityStageHistoryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignChangedBy( self, opportunityStageHistoryId, changedById ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(changedById) + " for ChangedBy on OpportunityStageHistory"

		try:
			# get the OpportunityStageHistory from db
			opportunityStageHistory = self.get( opportunityStageHistoryId ).first()	
			
			# get the User from db
			user = UserDelegate().get(changedById).first();
			
			# assign the ChangedBy		
			opportunityStageHistory.changedBy = user
			
			#save it
			opportunityStageHistory.save()

			# reload and return the appropriate version					
			return self.get( opportunityStageHistoryId );
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityStageHistory with id " + str(opportunityStageHistoryId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(changedById) + " does not exist.")
		except Exception:
			return None;
				
	def unassignChangedBy( self, opportunityStageHistoryId ):
		errMsg = "Failed to unassign element " + str(changedById) + " for ChangedBy on OpportunityStageHistory"

		try:
			# get the OpportunityStageHistory from db
			opportunityStageHistory = self.get( opportunityStageHistoryId ).first()	
			
			# assign to None for unassignment
			opportunityStageHistory.user = None			

			#save it
			opportunityStageHistory.save()

			# reload and return the appropriate version					
			return self.get( opportunityStageHistoryId );
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityStageHistory with id " + str(opportunityStageHistoryId) + " does not exist.")
		except Exception:
			return None;
		
