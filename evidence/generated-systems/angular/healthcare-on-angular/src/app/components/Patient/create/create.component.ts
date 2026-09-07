import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PatientService } from '../../../services/Patient.service';
import { Patient } from '../../../models/Patient';
import { SubBaseComponent } from '../../Patient/sub.base.component';

@Component({
    selector: 'app-create-patient',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePatientComponent extends SubBaseComponent implements OnInit {

    title = 'Add Patient';

    patientForm: FormGroup;
    patient: Patient;

    constructor( http: HttpClient,
        private patientService: PatientService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.patientForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      mrn: ['', Validators.required],
      dateOfBirth: ['', Validators.required],
      address: ['', Validators.required],
      primaryLanguage: ['', Validators.required],
      Appointments: ['', ],
      Encounters: ['', ],
      CarePlans: ['', ],
      Allergies: ['', ],
      Conditions: ['', ],
      MedicationOrders: ['', ],
      LabOrders: ['', ],
      ImagingOrders: ['', ],
      Coverages: ['', ],
      Claims: ['', ],
      Devices: ['', ],
      Observations: ['', ],
      SexAtBirth: ['', ],
      BloodType: ['', ]
        });
    }

    
    addPatient(firstName, lastName, mrn, dateOfBirth, address, primaryLanguage, Appointments, Encounters, CarePlans, Allergies, Conditions, MedicationOrders, LabOrders, ImagingOrders, Coverages, Claims, Devices, Observations, SexAtBirth, BloodType): void {
        this.patientService
        .addPatient(firstName, lastName, mrn, dateOfBirth, address, primaryLanguage, Appointments, Encounters, CarePlans, Allergies, Conditions, MedicationOrders, LabOrders, ImagingOrders, Coverages, Claims, Devices, Observations, SexAtBirth, BloodType)
            .subscribe(() => {
                this.router.navigate(['/indexPatient']);
            });
    }

    ngOnInit(): void {
    }
}