import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PatientService } from '../../../services/Patient.service';
import { SubBaseComponent } from '../../Patient/sub.base.component';


@Component({
    selector: 'app-edit-patient',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPatientComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Patient';

    patientForm: FormGroup;
    patient: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PatientService,
        private fb: FormBuilder
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

    
    updatePatient(firstName, lastName, mrn, dateOfBirth, address, primaryLanguage, Appointments, Encounters, CarePlans, Allergies, Conditions, MedicationOrders, LabOrders, ImagingOrders, Coverages, Claims, Devices, Observations, SexAtBirth, BloodType): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePatient(firstName, lastName, mrn, dateOfBirth, address, primaryLanguage, Appointments, Encounters, CarePlans, Allergies, Conditions, MedicationOrders, LabOrders, ImagingOrders, Coverages, Claims, Devices, Observations, SexAtBirth, BloodType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPatient']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPatient(params['id']).subscribe(res => {
                this.patient = res;
            });
        });
    }
}