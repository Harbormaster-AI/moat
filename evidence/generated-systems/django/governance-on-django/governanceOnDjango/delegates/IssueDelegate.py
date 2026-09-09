from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Issue import Issue
from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.models.AuditFinding import AuditFinding
from governanceOnDjango.models.CorrectiveAction import CorrectiveAction
from governanceOnDjango.models.Control import Control
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Issue
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IssueDelegate Declaration
#======================================================================
class IssueDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, issueId ):
		try:	
			issue = Issue.objects.filter(id=issueId)
			return issue.first();
		except Issue.DoesNotExist:
			raise ProcessingError("Issue with id " + str(issueId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, issue):
		for model in serializers.deserialize("json", issue):
			model.save()
			return model;

	def create(self, issue):
		issue.save()
		return issue;

	def saveFromJson(self, issue):
		for model in serializers.deserialize("json", issue):
			model.save()
			return issue;
	
	def save(self, issue):
		issue.save()
		return issue;
	
	def delete(self, issueId ):
		errMsg = "Failed to delete Issue from db using id " + str(issueId)
		
		try:
			issue = Issue.objects.get(id=issueId)
			issue.delete()
			return True
		except Issue.DoesNotExist:
			raise ProcessingError("Issue with id " + str(issueId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Issue.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Issue from db")
		except Exception:
			return None;
		
	def assignRisk( self, issueId, riskId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to assign element " + str(riskId) + " for Risk on Issue"

		try:
			# get the Issue from db
			issue = self.get( issueId ).first()	
			
			# get the Risk from db
			risk = RiskDelegate().get(riskId).first();
			
			# assign the Risk		
			issue.risk = risk
			
			#save it
			issue.save()

			# reload and return the appropriate version					
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRisk( self, issueId ):
		errMsg = "Failed to unassign element " + str(riskId) + " for Risk on Issue"

		try:
			# get the Issue from db
			issue = self.get( issueId ).first()	
			
			# assign to None for unassignment
			issue.risk = None			

			#save it
			issue.save()

			# reload and return the appropriate version					
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except Exception:
			return None;
		
	def assignFinding( self, issueId, findingId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to assign element " + str(findingId) + " for Finding on Issue"

		try:
			# get the Issue from db
			issue = self.get( issueId ).first()	
			
			# get the AuditFinding from db
			auditFinding = AuditFindingDelegate().get(findingId).first();
			
			# assign the Finding		
			issue.finding = auditFinding
			
			#save it
			issue.save()

			# reload and return the appropriate version					
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(findingId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFinding( self, issueId ):
		errMsg = "Failed to unassign element " + str(findingId) + " for Finding on Issue"

		try:
			# get the Issue from db
			issue = self.get( issueId ).first()	
			
			# assign to None for unassignment
			issue.auditFinding = None			

			#save it
			issue.save()

			# reload and return the appropriate version					
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except Exception:
			return None;
		
	def assignControl( self, issueId, controlId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to assign element " + str(controlId) + " for Control on Issue"

		try:
			# get the Issue from db
			issue = self.get( issueId ).first()	
			
			# get the Control from db
			control = ControlDelegate().get(controlId).first();
			
			# assign the Control		
			issue.control = control
			
			#save it
			issue.save()

			# reload and return the appropriate version					
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignControl( self, issueId ):
		errMsg = "Failed to unassign element " + str(controlId) + " for Control on Issue"

		try:
			# get the Issue from db
			issue = self.get( issueId ).first()	
			
			# assign to None for unassignment
			issue.control = None			

			#save it
			issue.save()

			# reload and return the appropriate version					
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except Exception:
			return None;
		
	def addCorrectiveActions( self, issueId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to add elements " + str(correctiveActionsIds) + " for CorrectiveActions on Issue"

		try:
			# get the Issue
			issue = self.get( issueId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				issue.correctiveActions.add(correctiveAction)
				
			# save it		
			issue.save()
			
			# reload and return the appropriate version
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCorrectiveActions( self, issueId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to remove elements " + str(correctiveActionsIds) + " for CorrectiveActions on Issue"

		try:
			# get the Issue
			issue = self.get( issueId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				issue.correctiveActions.remove(correctiveAction)
				
			# save it		
			issue.save()
			
			# reload and return the appropriate version
			return self.get( issueId );
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue with id " + str(issueId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
