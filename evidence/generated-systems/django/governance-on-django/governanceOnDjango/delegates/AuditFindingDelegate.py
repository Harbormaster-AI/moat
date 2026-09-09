from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.AuditFinding import AuditFinding
from governanceOnDjango.models.AuditEngagement import AuditEngagement
from governanceOnDjango.models.AuditWorkpaper import AuditWorkpaper
from governanceOnDjango.models.CorrectiveAction import CorrectiveAction
from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Issue import Issue
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AuditFinding
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditFindingDelegate Declaration
#======================================================================
class AuditFindingDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, auditFindingId ):
		try:	
			auditFinding = AuditFinding.objects.filter(id=auditFindingId)
			return auditFinding.first();
		except AuditFinding.DoesNotExist:
			raise ProcessingError("AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, auditFinding):
		for model in serializers.deserialize("json", auditFinding):
			model.save()
			return model;

	def create(self, auditFinding):
		auditFinding.save()
		return auditFinding;

	def saveFromJson(self, auditFinding):
		for model in serializers.deserialize("json", auditFinding):
			model.save()
			return auditFinding;
	
	def save(self, auditFinding):
		auditFinding.save()
		return auditFinding;
	
	def delete(self, auditFindingId ):
		errMsg = "Failed to delete AuditFinding from db using id " + str(auditFindingId)
		
		try:
			auditFinding = AuditFinding.objects.get(id=auditFindingId)
			auditFinding.delete()
			return True
		except AuditFinding.DoesNotExist:
			raise ProcessingError("AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AuditFinding.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AuditFinding from db")
		except Exception:
			return None;
		
	def assignEngagement( self, auditFindingId, engagementId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

		errMsg = "Failed to assign element " + str(engagementId) + " for Engagement on AuditFinding"

		try:
			# get the AuditFinding from db
			auditFinding = self.get( auditFindingId ).first()	
			
			# get the AuditEngagement from db
			auditEngagement = AuditEngagementDelegate().get(engagementId).first();
			
			# assign the Engagement		
			auditFinding.engagement = auditEngagement
			
			#save it
			auditFinding.save()

			# reload and return the appropriate version					
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(engagementId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEngagement( self, auditFindingId ):
		errMsg = "Failed to unassign element " + str(engagementId) + " for Engagement on AuditFinding"

		try:
			# get the AuditFinding from db
			auditFinding = self.get( auditFindingId ).first()	
			
			# assign to None for unassignment
			auditFinding.auditEngagement = None			

			#save it
			auditFinding.save()

			# reload and return the appropriate version					
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkpaper( self, auditFindingId, workpaperId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditWorkpaperDelegate import AuditWorkpaperDelegate

		errMsg = "Failed to assign element " + str(workpaperId) + " for Workpaper on AuditFinding"

		try:
			# get the AuditFinding from db
			auditFinding = self.get( auditFindingId ).first()	
			
			# get the AuditWorkpaper from db
			auditWorkpaper = AuditWorkpaperDelegate().get(workpaperId).first();
			
			# assign the Workpaper		
			auditFinding.workpaper = auditWorkpaper
			
			#save it
			auditFinding.save()

			# reload and return the appropriate version					
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(workpaperId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkpaper( self, auditFindingId ):
		errMsg = "Failed to unassign element " + str(workpaperId) + " for Workpaper on AuditFinding"

		try:
			# get the AuditFinding from db
			auditFinding = self.get( auditFindingId ).first()	
			
			# assign to None for unassignment
			auditFinding.auditWorkpaper = None			

			#save it
			auditFinding.save()

			# reload and return the appropriate version					
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Exception:
			return None;
		
	def addCorrectiveActions( self, auditFindingId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to add elements " + str(correctiveActionsIds) + " for CorrectiveActions on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				auditFinding.correctiveActions.add(correctiveAction)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCorrectiveActions( self, auditFindingId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to remove elements " + str(correctiveActionsIds) + " for CorrectiveActions on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				auditFinding.correctiveActions.remove(correctiveAction)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRelatedRisks( self, auditFindingId, relatedRisksIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to add elements " + str(relatedRisksIds) + " for RelatedRisks on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = relatedRisksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Risk		
				risk = RiskDelegate().get(id).first();	
				# add the Risk
				auditFinding.relatedRisks.add(risk)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRelatedRisks( self, auditFindingId, relatedRisksIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to remove elements " + str(relatedRisksIds) + " for RelatedRisks on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = relatedRisksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Risk		
				risk = RiskDelegate().get(id).first();	
				# add the Risk
				auditFinding.relatedRisks.remove(risk)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRelatedControls( self, auditFindingId, relatedControlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to add elements " + str(relatedControlsIds) + " for RelatedControls on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = relatedControlsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				auditFinding.relatedControls.add(control)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRelatedControls( self, auditFindingId, relatedControlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to remove elements " + str(relatedControlsIds) + " for RelatedControls on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = relatedControlsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				auditFinding.relatedControls.remove(control)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addIssues( self, auditFindingId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to add elements " + str(issuesIds) + " for Issues on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				auditFinding.issues.add(issue)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeIssues( self, auditFindingId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to remove elements " + str(issuesIds) + " for Issues on AuditFinding"

		try:
			# get the AuditFinding
			auditFinding = self.get( auditFindingId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				auditFinding.issues.remove(issue)
				
			# save it		
			auditFinding.save()
			
			# reload and return the appropriate version
			return self.get( auditFindingId );
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding with id " + str(auditFindingId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
