from django.urls import path
from hrOnDjango.views import CertificationView

urlpatterns = [
    path('', CertificationView.index, name='index'),
	path('create', CertificationView.get, name='create'),
	path('get/<int:certificationId>/', CertificationView.get, name='get'),
	path('save', CertificationView.save, name='save'),
	path('getAll', CertificationView.getAll, name='getAll'),
	path('delete/<int:certificationId>/', CertificationView.delete, name='delete'),
	path('assignEmployee/<int:certificationId>/<int:EmployeeId>/', CertificationView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:certificationId>/', CertificationView.unassignEmployee, name='unassignEmployee'),
	path('assignCourse/<int:certificationId>/<int:CourseId>/', CertificationView.assignCourse, name='assignCourse'),
	path('unassignCourse/<int:certificationId>/', CertificationView.unassignCourse, name='unassignCourse'),
]
