from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.CorrectiveAction import CorrectiveAction
from governanceOnDjango.models.AuditFinding import AuditFinding
from governanceOnDjango.models.Issue import Issue
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CorrectiveAction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CorrectiveActionDelegate Declaration
#======================================================================
class CorrectiveActionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, correctiveActionId ):
		try:	
			correctiveAction = CorrectiveAction.objects.filter(id=correctiveActionId)
			return correctiveAction.first();
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError("CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, correctiveAction):
		for model in serializers.deserialize("json", correctiveAction):
			model.save()
			return model;

	def create(self, correctiveAction):
		correctiveAction.save()
		return correctiveAction;

	def saveFromJson(self, correctiveAction):
		for model in serializers.deserialize("json", correctiveAction):
			model.save()
			return correctiveAction;
	
	def save(self, correctiveAction):
		correctiveAction.save()
		return correctiveAction;
	
	def delete(self, correctiveActionId ):
		errMsg = "Failed to delete CorrectiveAction from db using id " + str(correctiveActionId)
		
		try:
			correctiveAction = CorrectiveAction.objects.get(id=correctiveActionId)
			correctiveAction.delete()
			return True
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError("CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CorrectiveAction.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CorrectiveAction from db")
		except Exception:
			return None;
		
	def assignFinding( self, correctiveActionId, findingId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to assign element " + str(findingId) + " for Finding on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# get the AuditFinding from db
			auditFinding = AuditFindingDelegate().get(findingId).first();
			
			# assign the Finding		
			correctiveAction.finding = auditFinding
			
			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(findingId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFinding( self, correctiveActionId ):
		errMsg = "Failed to unassign element " + str(findingId) + " for Finding on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# assign to None for unassignment
			correctiveAction.auditFinding = None			

			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignIssue( self, correctiveActionId, issueId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to assign element " + str(issueId) + " for Issue on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# get the Issue from db
			issue = IssueDelegate().get(issueId).first();
			
			# assign the Issue		
			correctiveAction.issue = issue
			
			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignIssue( self, correctiveActionId ):
		errMsg = "Failed to unassign element " + str(issueId) + " for Issue on CorrectiveAction"

		try:
			# get the CorrectiveAction from db
			correctiveAction = self.get( correctiveActionId ).first()	
			
			# assign to None for unassignment
			correctiveAction.issue = None			

			#save it
			correctiveAction.save()

			# reload and return the appropriate version					
			return self.get( correctiveActionId );
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction with id " + str(correctiveActionId) + " does not exist.")
		except Exception:
			return None;
		
