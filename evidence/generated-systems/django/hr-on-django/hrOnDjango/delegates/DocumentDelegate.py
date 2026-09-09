from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Document import Document
from hrOnDjango.models.Candidate import Candidate
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Document
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DocumentDelegate Declaration
#======================================================================
class DocumentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, documentId ):
		try:	
			document = Document.objects.filter(id=documentId)
			return document.first();
		except Document.DoesNotExist:
			raise ProcessingError("Document with id " + str(documentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, document):
		for model in serializers.deserialize("json", document):
			model.save()
			return model;

	def create(self, document):
		document.save()
		return document;

	def saveFromJson(self, document):
		for model in serializers.deserialize("json", document):
			model.save()
			return document;
	
	def save(self, document):
		document.save()
		return document;
	
	def delete(self, documentId ):
		errMsg = "Failed to delete Document from db using id " + str(documentId)
		
		try:
			document = Document.objects.get(id=documentId)
			document.delete()
			return True
		except Document.DoesNotExist:
			raise ProcessingError("Document with id " + str(documentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Document.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Document from db")
		except Exception:
			return None;
		
	def assignCandidate( self, documentId, candidateId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CandidateDelegate import CandidateDelegate

		errMsg = "Failed to assign element " + str(candidateId) + " for Candidate on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# get the Candidate from db
			candidate = CandidateDelegate().get(candidateId).first();
			
			# assign the Candidate		
			document.candidate = candidate
			
			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Candidate.DoesNotExist:
			raise ProcessingError(errMsg + " : Candidate with id " + str(candidateId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCandidate( self, documentId ):
		errMsg = "Failed to unassign element " + str(candidateId) + " for Candidate on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# assign to None for unassignment
			document.candidate = None			

			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, documentId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			document.employee = employee
			
			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, documentId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on Document"

		try:
			# get the Document from db
			document = self.get( documentId ).first()	
			
			# assign to None for unassignment
			document.employee = None			

			#save it
			document.save()

			# reload and return the appropriate version					
			return self.get( documentId );
		except Document.DoesNotExist:
			raise ProcessingError(errMsg + " : Document with id " + str(documentId) + " does not exist.")
		except Exception:
			return None;
		
