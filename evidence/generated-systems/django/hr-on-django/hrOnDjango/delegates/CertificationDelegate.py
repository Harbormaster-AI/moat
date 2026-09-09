from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Certification import Certification
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.TrainingCourse import TrainingCourse
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Certification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CertificationDelegate Declaration
#======================================================================
class CertificationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, certificationId ):
		try:	
			certification = Certification.objects.filter(id=certificationId)
			return certification.first();
		except Certification.DoesNotExist:
			raise ProcessingError("Certification with id " + str(certificationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, certification):
		for model in serializers.deserialize("json", certification):
			model.save()
			return model;

	def create(self, certification):
		certification.save()
		return certification;

	def saveFromJson(self, certification):
		for model in serializers.deserialize("json", certification):
			model.save()
			return certification;
	
	def save(self, certification):
		certification.save()
		return certification;
	
	def delete(self, certificationId ):
		errMsg = "Failed to delete Certification from db using id " + str(certificationId)
		
		try:
			certification = Certification.objects.get(id=certificationId)
			certification.delete()
			return True
		except Certification.DoesNotExist:
			raise ProcessingError("Certification with id " + str(certificationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Certification.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Certification from db")
		except Exception:
			return None;
		
	def assignEmployee( self, certificationId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on Certification"

		try:
			# get the Certification from db
			certification = self.get( certificationId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			certification.employee = employee
			
			#save it
			certification.save()

			# reload and return the appropriate version					
			return self.get( certificationId );
		except Certification.DoesNotExist:
			raise ProcessingError(errMsg + " : Certification with id " + str(certificationId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, certificationId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on Certification"

		try:
			# get the Certification from db
			certification = self.get( certificationId ).first()	
			
			# assign to None for unassignment
			certification.employee = None			

			#save it
			certification.save()

			# reload and return the appropriate version					
			return self.get( certificationId );
		except Certification.DoesNotExist:
			raise ProcessingError(errMsg + " : Certification with id " + str(certificationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCourse( self, certificationId, courseId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

		errMsg = "Failed to assign element " + str(courseId) + " for Course on Certification"

		try:
			# get the Certification from db
			certification = self.get( certificationId ).first()	
			
			# get the TrainingCourse from db
			trainingCourse = TrainingCourseDelegate().get(courseId).first();
			
			# assign the Course		
			certification.course = trainingCourse
			
			#save it
			certification.save()

			# reload and return the appropriate version					
			return self.get( certificationId );
		except Certification.DoesNotExist:
			raise ProcessingError(errMsg + " : Certification with id " + str(certificationId) + " does not exist.")
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(courseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCourse( self, certificationId ):
		errMsg = "Failed to unassign element " + str(courseId) + " for Course on Certification"

		try:
			# get the Certification from db
			certification = self.get( certificationId ).first()	
			
			# assign to None for unassignment
			certification.trainingCourse = None			

			#save it
			certification.save()

			# reload and return the appropriate version					
			return self.get( certificationId );
		except Certification.DoesNotExist:
			raise ProcessingError(errMsg + " : Certification with id " + str(certificationId) + " does not exist.")
		except Exception:
			return None;
		
