import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AdmissionService } from '../../../services/Admission.service';
import { SubBaseComponent } from '../../Admission/sub.base.component';


@Component({
    selector: 'app-edit-admission',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAdmissionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Admission';

    admissionForm: FormGroup;
    admission: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AdmissionService,
        private fb: FormBuilder
) {
        super(http);
        this.admissionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  admitDateTime: ['', Validators.required],
      bed: ['', Validators.required],
      Encounter: ['', ],
      Facility: ['', ],
      AdmissionType: ['', ]
        });
    }

    
    updateAdmission(admitDateTime, bed, Encounter, Facility, AdmissionType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAdmission(admitDateTime, bed, Encounter, Facility, AdmissionType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAdmission']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAdmission(params['id']).subscribe(res => {
                this.admission = res;
            });
        });
    }
}