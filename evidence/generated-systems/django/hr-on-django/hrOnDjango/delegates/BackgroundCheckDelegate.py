from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.BackgroundCheck import BackgroundCheck
from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.models.JobRequisition import JobRequisition
from hrOnDjango.models.Document import Document
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BackgroundCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BackgroundCheckDelegate Declaration
#======================================================================
class BackgroundCheckDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, backgroundCheckId ):
		try:	
			backgroundCheck = BackgroundCheck.objects.filter(id=backgroundCheckId)
			return backgroundCheck.first();
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError("BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, backgroundCheck):
		for model in serializers.deserialize("json", backgroundCheck):
			model.save()
			return model;

	def create(self, backgroundCheck):
		backgroundCheck.save()
		return backgroundCheck;

	def saveFromJson(self, backgroundCheck):
		for model in serializers.deserialize("json", backgroundCheck):
			model.save()
			return backgroundCheck;
	
	def save(self, backgroundCheck):
		backgroundCheck.save()
		return backgroundCheck;
	
	def delete(self, backgroundCheckId ):
		errMsg = "Failed to delete BackgroundCheck from db using id " + str(backgroundCheckId)
		
		try:
			backgroundCheck = BackgroundCheck.objects.get(id=backgroundCheckId)
			backgroundCheck.delete()
			return True
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError("BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BackgroundCheck.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BackgroundCheck from db")
		except Exception:
			return None;
		
	def assignCandidate( self, backgroundCheckId, candidateId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

		errMsg = "Failed to assign element " + str(candidateId) + " for Candidate on BackgroundCheck"

		try:
			# get the BackgroundCheck from db
			backgroundCheck = self.get( backgroundCheckId ).first()	
			
			# get the Candidate from db
			candidate = CandidateDelegate().get(candidateId).first();
			
			# assign the Candidate		
			backgroundCheck.candidate = candidate
			
			#save it
			backgroundCheck.save()

			# reload and return the appropriate version					
			return self.get( backgroundCheckId );
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCandidate( self, backgroundCheckId ):
		errMsg = "Failed to unassign element " + str(candidateId) + " for Candidate on BackgroundCheck"

		try:
			# get the BackgroundCheck from db
			backgroundCheck = self.get( backgroundCheckId ).first()	
			
			# assign to None for unassignment
			backgroundCheck.candidate = None			

			#save it
			backgroundCheck.save()

			# reload and return the appropriate version					
			return self.get( backgroundCheckId );
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRequisition( self, backgroundCheckId, requisitionId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobRequisitionDelegate import JobRequisitionDelegate

		errMsg = "Failed to assign element " + str(requisitionId) + " for Requisition on BackgroundCheck"

		try:
			# get the BackgroundCheck from db
			backgroundCheck = self.get( backgroundCheckId ).first()	
			
			# get the JobRequisition from db
			jobRequisition = JobRequisitionDelegate().get(requisitionId).first();
			
			# assign the Requisition		
			backgroundCheck.requisition = jobRequisition
			
			#save it
			backgroundCheck.save()

			# reload and return the appropriate version					
			return self.get( backgroundCheckId );
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except JobRequisition.DoesNotExist:
			raise ProcessingError(errMsg + " : JobRequisition with id " + str(requisitionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRequisition( self, backgroundCheckId ):
		errMsg = "Failed to unassign element " + str(requisitionId) + " for Requisition on BackgroundCheck"

		try:
			# get the BackgroundCheck from db
			backgroundCheck = self.get( backgroundCheckId ).first()	
			
			# assign to None for unassignment
			backgroundCheck.jobRequisition = None			

			#save it
			backgroundCheck.save()

			# reload and return the appropriate version					
			return self.get( backgroundCheckId );
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except Exception:
			return None;
		
	def assignReport( self, backgroundCheckId, reportId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DocumentDelegate import DocumentDelegate

		errMsg = "Failed to assign element " + str(reportId) + " for Report on BackgroundCheck"

		try:
			# get the BackgroundCheck from db
			backgroundCheck = self.get( backgroundCheckId ).first()	
			
			# get the Document from db
			document = DocumentDelegate().get(reportId).first();
			
			# assign the Report		
			backgroundCheck.report = document
			
			#save it
			backgroundCheck.save()

			# reload and return the appropriate version					
			return self.get( backgroundCheckId );
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(reportId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignReport( self, backgroundCheckId ):
		errMsg = "Failed to unassign element " + str(reportId) + " for Report on BackgroundCheck"

		try:
			# get the BackgroundCheck from db
			backgroundCheck = self.get( backgroundCheckId ).first()	
			
			# assign to None for unassignment
			backgroundCheck.document = None			

			#save it
			backgroundCheck.save()

			# reload and return the appropriate version					
			return self.get( backgroundCheckId );
		except BackgroundCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : BackgroundCheck with id " + str(backgroundCheckId) + " does not exist.")
		except Exception:
			return None;
		
