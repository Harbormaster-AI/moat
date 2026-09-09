from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.RiskAssessment import RiskAssessment
from governanceOnDjango.models.Issue import Issue
from governanceOnDjango.models.AuditFinding import AuditFinding
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Risk
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskDelegate Declaration
#======================================================================
class RiskDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, riskId ):
		try:	
			risk = Risk.objects.filter(id=riskId)
			return risk.first();
		except Risk.DoesNotExist:
			raise ProcessingError("Risk with id " + str(riskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, risk):
		for model in serializers.deserialize("json", risk):
			model.save()
			return model;

	def create(self, risk):
		risk.save()
		return risk;

	def saveFromJson(self, risk):
		for model in serializers.deserialize("json", risk):
			model.save()
			return risk;
	
	def save(self, risk):
		risk.save()
		return risk;
	
	def delete(self, riskId ):
		errMsg = "Failed to delete Risk from db using id " + str(riskId)
		
		try:
			risk = Risk.objects.get(id=riskId)
			risk.delete()
			return True
		except Risk.DoesNotExist:
			raise ProcessingError("Risk with id " + str(riskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Risk.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Risk from db")
		except Exception:
			return None;
		
	def assignOrganization( self, riskId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Risk"

		try:
			# get the Risk from db
			risk = self.get( riskId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			risk.organization = organization
			
			#save it
			risk.save()

			# reload and return the appropriate version					
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, riskId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Risk"

		try:
			# get the Risk from db
			risk = self.get( riskId ).first()	
			
			# assign to None for unassignment
			risk.organization = None			

			#save it
			risk.save()

			# reload and return the appropriate version					
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Exception:
			return None;
		
	def addControls( self, riskId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to add elements " + str(controlsIds) + " for Controls on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				risk.controls.add(control)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControls( self, riskId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to remove elements " + str(controlsIds) + " for Controls on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				risk.controls.remove(control)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAssessments( self, riskId, assessmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskAssessmentDelegate import RiskAssessmentDelegate

		errMsg = "Failed to add elements " + str(assessmentsIds) + " for Assessments on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = assessmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RiskAssessment		
				riskAssessment = RiskAssessmentDelegate().get(id).first();	
				# add the RiskAssessment
				risk.assessments.add(riskAssessment)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except RiskAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : RiskAssessment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAssessments( self, riskId, assessmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskAssessmentDelegate import RiskAssessmentDelegate

		errMsg = "Failed to remove elements " + str(assessmentsIds) + " for Assessments on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = assessmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RiskAssessment		
				riskAssessment = RiskAssessmentDelegate().get(id).first();	
				# add the RiskAssessment
				risk.assessments.remove(riskAssessment)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except RiskAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : RiskAssessment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addIssues( self, riskId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to add elements " + str(issuesIds) + " for Issues on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				risk.issues.add(issue)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeIssues( self, riskId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to remove elements " + str(issuesIds) + " for Issues on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				risk.issues.remove(issue)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFindings( self, riskId, findingsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to add elements " + str(findingsIds) + " for Findings on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = findingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AuditFinding		
				auditFinding = AuditFindingDelegate().get(id).first();	
				# add the AuditFinding
				risk.findings.add(auditFinding)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFindings( self, riskId, findingsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to remove elements " + str(findingsIds) + " for Findings on Risk"

		try:
			# get the Risk
			risk = self.get( riskId ).first()
				
			# split on a comma with no spaces
			idList = findingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AuditFinding		
				auditFinding = AuditFindingDelegate().get(id).first();	
				# add the AuditFinding
				risk.findings.remove(auditFinding)
				
			# save it		
			risk.save()
			
			# reload and return the appropriate version
			return self.get( riskId );
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
