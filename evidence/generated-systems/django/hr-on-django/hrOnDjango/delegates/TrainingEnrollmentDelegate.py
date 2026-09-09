from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.TrainingEnrollment import TrainingEnrollment
from hrOnDjango.models.TrainingCourse import TrainingCourse
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TrainingEnrollment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingEnrollmentDelegate Declaration
#======================================================================
class TrainingEnrollmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, trainingEnrollmentId ):
		try:	
			trainingEnrollment = TrainingEnrollment.objects.filter(id=trainingEnrollmentId)
			return trainingEnrollment.first();
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError("TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, trainingEnrollment):
		for model in serializers.deserialize("json", trainingEnrollment):
			model.save()
			return model;

	def create(self, trainingEnrollment):
		trainingEnrollment.save()
		return trainingEnrollment;

	def saveFromJson(self, trainingEnrollment):
		for model in serializers.deserialize("json", trainingEnrollment):
			model.save()
			return trainingEnrollment;
	
	def save(self, trainingEnrollment):
		trainingEnrollment.save()
		return trainingEnrollment;
	
	def delete(self, trainingEnrollmentId ):
		errMsg = "Failed to delete TrainingEnrollment from db using id " + str(trainingEnrollmentId)
		
		try:
			trainingEnrollment = TrainingEnrollment.objects.get(id=trainingEnrollmentId)
			trainingEnrollment.delete()
			return True
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError("TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TrainingEnrollment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TrainingEnrollment from db")
		except Exception:
			return None;
		
	def assignCourse( self, trainingEnrollmentId, courseId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.TrainingCourseDelegate import TrainingCourseDelegate

		errMsg = "Failed to assign element " + str(courseId) + " for Course on TrainingEnrollment"

		try:
			# get the TrainingEnrollment from db
			trainingEnrollment = self.get( trainingEnrollmentId ).first()	
			
			# get the TrainingCourse from db
			trainingCourse = TrainingCourseDelegate().get(courseId).first();
			
			# assign the Course		
			trainingEnrollment.course = trainingCourse
			
			#save it
			trainingEnrollment.save()

			# reload and return the appropriate version					
			return self.get( trainingEnrollmentId );
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except TrainingCourse.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingCourse with id " + str(courseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCourse( self, trainingEnrollmentId ):
		errMsg = "Failed to unassign element " + str(courseId) + " for Course on TrainingEnrollment"

		try:
			# get the TrainingEnrollment from db
			trainingEnrollment = self.get( trainingEnrollmentId ).first()	
			
			# assign to None for unassignment
			trainingEnrollment.trainingCourse = None			

			#save it
			trainingEnrollment.save()

			# reload and return the appropriate version					
			return self.get( trainingEnrollmentId );
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, trainingEnrollmentId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on TrainingEnrollment"

		try:
			# get the TrainingEnrollment from db
			trainingEnrollment = self.get( trainingEnrollmentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			trainingEnrollment.employee = employee
			
			#save it
			trainingEnrollment.save()

			# reload and return the appropriate version					
			return self.get( trainingEnrollmentId );
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, trainingEnrollmentId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on TrainingEnrollment"

		try:
			# get the TrainingEnrollment from db
			trainingEnrollment = self.get( trainingEnrollmentId ).first()	
			
			# assign to None for unassignment
			trainingEnrollment.employee = None			

			#save it
			trainingEnrollment.save()

			# reload and return the appropriate version					
			return self.get( trainingEnrollmentId );
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInstructor( self, trainingEnrollmentId, instructorId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(instructorId) + " for Instructor on TrainingEnrollment"

		try:
			# get the TrainingEnrollment from db
			trainingEnrollment = self.get( trainingEnrollmentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(instructorId).first();
			
			# assign the Instructor		
			trainingEnrollment.instructor = employee
			
			#save it
			trainingEnrollment.save()

			# reload and return the appropriate version					
			return self.get( trainingEnrollmentId );
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(instructorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInstructor( self, trainingEnrollmentId ):
		errMsg = "Failed to unassign element " + str(instructorId) + " for Instructor on TrainingEnrollment"

		try:
			# get the TrainingEnrollment from db
			trainingEnrollment = self.get( trainingEnrollmentId ).first()	
			
			# assign to None for unassignment
			trainingEnrollment.employee = None			

			#save it
			trainingEnrollment.save()

			# reload and return the appropriate version					
			return self.get( trainingEnrollmentId );
		except TrainingEnrollment.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingEnrollment with id " + str(trainingEnrollmentId) + " does not exist.")
		except Exception:
			return None;
		
