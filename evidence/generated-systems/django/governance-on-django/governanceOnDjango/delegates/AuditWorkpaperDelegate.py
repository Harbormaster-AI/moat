from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.AuditWorkpaper import AuditWorkpaper
from governanceOnDjango.models.AuditEngagement import AuditEngagement
from governanceOnDjango.models.Evidence import Evidence
from governanceOnDjango.models.AuditFinding import AuditFinding
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AuditWorkpaper
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuditWorkpaperDelegate Declaration
#======================================================================
class AuditWorkpaperDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, auditWorkpaperId ):
		try:	
			auditWorkpaper = AuditWorkpaper.objects.filter(id=auditWorkpaperId)
			return auditWorkpaper.first();
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError("AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, auditWorkpaper):
		for model in serializers.deserialize("json", auditWorkpaper):
			model.save()
			return model;

	def create(self, auditWorkpaper):
		auditWorkpaper.save()
		return auditWorkpaper;

	def saveFromJson(self, auditWorkpaper):
		for model in serializers.deserialize("json", auditWorkpaper):
			model.save()
			return auditWorkpaper;
	
	def save(self, auditWorkpaper):
		auditWorkpaper.save()
		return auditWorkpaper;
	
	def delete(self, auditWorkpaperId ):
		errMsg = "Failed to delete AuditWorkpaper from db using id " + str(auditWorkpaperId)
		
		try:
			auditWorkpaper = AuditWorkpaper.objects.get(id=auditWorkpaperId)
			auditWorkpaper.delete()
			return True
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError("AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AuditWorkpaper.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AuditWorkpaper from db")
		except Exception:
			return None;
		
	def assignEngagement( self, auditWorkpaperId, engagementId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

		errMsg = "Failed to assign element " + str(engagementId) + " for Engagement on AuditWorkpaper"

		try:
			# get the AuditWorkpaper from db
			auditWorkpaper = self.get( auditWorkpaperId ).first()	
			
			# get the AuditEngagement from db
			auditEngagement = AuditEngagementDelegate().get(engagementId).first();
			
			# assign the Engagement		
			auditWorkpaper.engagement = auditEngagement
			
			#save it
			auditWorkpaper.save()

			# reload and return the appropriate version					
			return self.get( auditWorkpaperId );
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(engagementId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEngagement( self, auditWorkpaperId ):
		errMsg = "Failed to unassign element " + str(engagementId) + " for Engagement on AuditWorkpaper"

		try:
			# get the AuditWorkpaper from db
			auditWorkpaper = self.get( auditWorkpaperId ).first()	
			
			# assign to None for unassignment
			auditWorkpaper.auditEngagement = None			

			#save it
			auditWorkpaper.save()

			# reload and return the appropriate version					
			return self.get( auditWorkpaperId );
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except Exception:
			return None;
		
	def addEvidence( self, auditWorkpaperId, evidenceIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

		errMsg = "Failed to add elements " + str(evidenceIds) + " for Evidence on AuditWorkpaper"

		try:
			# get the AuditWorkpaper
			auditWorkpaper = self.get( auditWorkpaperId ).first()
				
			# split on a comma with no spaces
			idList = evidenceIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Evidence		
				evidence = EvidenceDelegate().get(id).first();	
				# add the Evidence
				auditWorkpaper.evidence.add(evidence)
				
			# save it		
			auditWorkpaper.save()
			
			# reload and return the appropriate version
			return self.get( auditWorkpaperId );
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEvidence( self, auditWorkpaperId, evidenceIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

		errMsg = "Failed to remove elements " + str(evidenceIds) + " for Evidence on AuditWorkpaper"

		try:
			# get the AuditWorkpaper
			auditWorkpaper = self.get( auditWorkpaperId ).first()
				
			# split on a comma with no spaces
			idList = evidenceIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Evidence		
				evidence = EvidenceDelegate().get(id).first();	
				# add the Evidence
				auditWorkpaper.evidence.remove(evidence)
				
			# save it		
			auditWorkpaper.save()
			
			# reload and return the appropriate version
			return self.get( auditWorkpaperId );
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFindings( self, auditWorkpaperId, findingsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to add elements " + str(findingsIds) + " for Findings on AuditWorkpaper"

		try:
			# get the AuditWorkpaper
			auditWorkpaper = self.get( auditWorkpaperId ).first()
				
			# split on a comma with no spaces
			idList = findingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AuditFinding		
				auditFinding = AuditFindingDelegate().get(id).first();	
				# add the AuditFinding
				auditWorkpaper.findings.add(auditFinding)
				
			# save it		
			auditWorkpaper.save()
			
			# reload and return the appropriate version
			return self.get( auditWorkpaperId );
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFindings( self, auditWorkpaperId, findingsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditFindingDelegate import AuditFindingDelegate

		errMsg = "Failed to remove elements " + str(findingsIds) + " for Findings on AuditWorkpaper"

		try:
			# get the AuditWorkpaper
			auditWorkpaper = self.get( auditWorkpaperId ).first()
				
			# split on a comma with no spaces
			idList = findingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AuditFinding		
				auditFinding = AuditFindingDelegate().get(id).first();	
				# add the AuditFinding
				auditWorkpaper.findings.remove(auditFinding)
				
			# save it		
			auditWorkpaper.save()
			
			# reload and return the appropriate version
			return self.get( auditWorkpaperId );
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(auditWorkpaperId) + " does not exist.")
		except AuditFinding.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditFinding does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
