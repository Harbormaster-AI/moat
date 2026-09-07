import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CertificationService } from '../../../services/Certification.service';
import { SubBaseComponent } from '../../Certification/sub.base.component';


@Component({
    selector: 'app-edit-certification',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCertificationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Certification';

    certificationForm: FormGroup;
    certification: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CertificationService,
        private fb: FormBuilder
) {
        super(http);
        this.certificationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      issuer: ['', Validators.required],
      validFrom: ['', Validators.required],
      validTo: ['', Validators.required],
      credentialId: ['', Validators.required],
      Employee: ['', ],
      Course: ['', ]
        });
    }

    
    updateCertification(name, issuer, validFrom, validTo, credentialId, Employee, Course): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCertification(name, issuer, validFrom, validTo, credentialId, Employee, Course, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCertification']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCertification(params['id']).subscribe(res => {
                this.certification = res;
            });
        });
    }
}