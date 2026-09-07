import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TypeCertificateService } from '../../../services/TypeCertificate.service';
import { TypeCertificate } from '../../../models/TypeCertificate';
import { SubBaseComponent } from '../../TypeCertificate/sub.base.component';

@Component({
    selector: 'app-create-typeCertificate',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTypeCertificateComponent extends SubBaseComponent implements OnInit {

    title = 'Add TypeCertificate';

    typeCertificateForm: FormGroup;
    typeCertificate: TypeCertificate;

    constructor( http: HttpClient,
        private typeCertificateService: TypeCertificateService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTypeCertificate(certificateNumber, authority, Program): void {
        this.typeCertificateService
        .addTypeCertificate(certificateNumber, authority, Program)
            .subscribe(() => {
                this.router.navigate(['/indexTypeCertificate']);
            });
    }

    ngOnInit(): void {
    }
}