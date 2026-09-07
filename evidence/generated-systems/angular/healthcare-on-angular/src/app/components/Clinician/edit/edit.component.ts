import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ClinicianService } from '../../../services/Clinician.service';
import { SubBaseComponent } from '../../Clinician/sub.base.component';


@Component({
    selector: 'app-edit-clinician',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditClinicianComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Clinician';

    clinicianForm: FormGroup;
    clinician: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ClinicianService,
        private fb: FormBuilder
) {
        super(http);
        this.clinicianForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      licenseNumber: ['', Validators.required],
      CareTeams: ['', ],
      Appointments: ['', ],
      Encounters: ['', ],
      Procedures: ['', ],
      ImagingReports: ['', ],
      ClinicianType: ['', ],
      Specialty: ['', ]
        });
    }

    
    updateClinician(firstName, lastName, licenseNumber, CareTeams, Appointments, Encounters, Procedures, ImagingReports, ClinicianType, Specialty): void {
        this.route.params.subscribe((params) => {

                        this.service.updateClinician(firstName, lastName, licenseNumber, CareTeams, Appointments, Encounters, Procedures, ImagingReports, ClinicianType, Specialty, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexClinician']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getClinician(params['id']).subscribe(res => {
                this.clinician = res;
            });
        });
    }
}