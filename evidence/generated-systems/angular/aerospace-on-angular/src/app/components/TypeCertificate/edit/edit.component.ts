import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TypeCertificateService } from '../../../services/TypeCertificate.service';
import { SubBaseComponent } from '../../TypeCertificate/sub.base.component';


@Component({
    selector: 'app-edit-typeCertificate',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTypeCertificateComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TypeCertificate';

    typeCertificateForm: FormGroup;
    typeCertificate: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TypeCertificateService,
        private fb: FormBuilder
) {
        super(http);
        this.typeCertificateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  certificateNumber: ['', Validators.required],
      authority: ['', Validators.required],
      Program: ['', ]
        });
    }

    
    updateTypeCertificate(certificateNumber, authority, Program): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTypeCertificate(certificateNumber, authority, Program, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTypeCertificate']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTypeCertificate(params['id']).subscribe(res => {
                this.typeCertificate = res;
            });
        });
    }
}