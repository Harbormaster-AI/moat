import datetime

from django.test import TestCase
from django.utils import timezone
from healthcareOnDjango.models.Appointment import Appointment
from healthcareOnDjango.delegates.AppointmentDelegate import AppointmentDelegate

 #======================================================================
# 
# Encapsulates data for model Appointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppointmentTest Declaration
#======================================================================
class AppointmentTest (TestCase) :
	def test_crud(self) :
		appointment = Appointment()
		appointment.appointmentDate = "default appointmentDate field value"
		appointment.reason = "default reason field value"
		appointment.status = "default status field value"
		appointment.priority = "default priority field value"
		
		delegate = AppointmentDelegate()
		responseObj = delegate.create(appointment)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


